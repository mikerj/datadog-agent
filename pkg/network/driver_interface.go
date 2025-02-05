// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build windows && npm
// +build windows,npm

package network

import (
	"fmt"
	"math"
	"sync"
	"unsafe"

	"github.com/DataDog/datadog-agent/pkg/network/config"
	"github.com/DataDog/datadog-agent/pkg/network/driver"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"go.uber.org/atomic"
	"golang.org/x/sys/windows"
)

// DriverExpvar is the name of a top-level driver expvar returned from GetStats
type DriverExpvar string

const (
	totalFlowStats  DriverExpvar = "driver_total_flow_stats"
	flowHandleStats              = "driver_flow_handle_stats"
	flowStats                    = "flows"
	driverStats                  = "driver"
)

const (
	// set default max open & closed flows for windows.  See note in setParams(),
	// these are only sort-of honored for now
	defaultMaxOpenFlows   = uint64(32767)
	defaultMaxClosedFlows = uint64(32767)

	// starting number of entries usermode flow buffer can contain
	defaultFlowEntries      = 50
	defaultDriverBufferSize = defaultFlowEntries * driver.PerFlowDataSize
)

// DriverExpvarNames is a list of all the DriverExpvar names returned from GetStats
var DriverExpvarNames = []DriverExpvar{totalFlowStats, flowHandleStats, flowStats, driverStats}

// DriverInterface holds all necessary information for interacting with the windows driver
type DriverInterface struct {
	totalFlows     *atomic.Int64
	closedFlows    *atomic.Int64
	openFlows      *atomic.Int64
	moreDataErrors *atomic.Int64
	bufferSize     *atomic.Int64

	maxOpenFlows   uint64
	maxClosedFlows uint64

	driverFlowHandle  *driver.Handle
	driverStatsHandle *driver.Handle

	enableMonotonicCounts bool

	bufferLock sync.Mutex
	readBuffer []uint8

	cfg *config.Config
}

// NewDriverInterface returns a DriverInterface struct for interacting with the driver
func NewDriverInterface(cfg *config.Config) (*DriverInterface, error) {
	dc := &DriverInterface{
		totalFlows:     atomic.NewInt64(0),
		closedFlows:    atomic.NewInt64(0),
		openFlows:      atomic.NewInt64(0),
		moreDataErrors: atomic.NewInt64(0),
		bufferSize:     atomic.NewInt64(defaultDriverBufferSize),

		cfg:                   cfg,
		enableMonotonicCounts: cfg.EnableMonotonicCount,
		readBuffer:            make([]byte, defaultDriverBufferSize),
		maxOpenFlows:          uint64(cfg.MaxTrackedConnections),
		maxClosedFlows:        uint64(cfg.MaxClosedConnectionsBuffered),
	}

	err := dc.setupFlowHandle()
	if err != nil {
		return nil, fmt.Errorf("error creating driver flow handle: %w", err)
	}

	err = dc.setupStatsHandle()
	if err != nil {
		return nil, fmt.Errorf("Error creating stats handle: %w", err)
	}

	return dc, nil
}

// Close shuts down the driver interface
func (di *DriverInterface) Close() error {
	if err := di.driverFlowHandle.Close(); err != nil {
		return fmt.Errorf("error closing flow file handle: %w", err)
	}
	if err := di.driverStatsHandle.Close(); err != nil {
		return fmt.Errorf("error closing stat file handle: %w", err)
	}
	return nil
}

// setupFlowHandle generates a windows Driver Handle, and creates a DriverHandle struct to pull flows from the driver
// by setting the necessary filters
func (di *DriverInterface) setupFlowHandle() error {
	dh, err := driver.NewHandle(0, driver.FlowHandle)
	if err != nil {
		return err
	}
	di.driverFlowHandle = dh

	filters, err := di.createFlowHandleFilters()
	if err != nil {
		return err
	}

	// Create and set flow filters for each interface
	err = di.driverFlowHandle.SetFlowFilters(filters)
	if err != nil {
		return err
	}

	// Set up the maximum amount of connections
	err = di.setFlowParams()
	if err != nil {
		return err
	}
	return nil
}

// setupStatsHandle generates a windows Driver Handle, and creates a DriverHandle struct
func (di *DriverInterface) setupStatsHandle() error {
	dh, err := driver.NewHandle(0, driver.StatsHandle)
	if err != nil {
		return err
	}

	di.driverStatsHandle = dh
	return nil
}

// GetStats returns statistics for the driver interface used by the windows tracer
func (di *DriverInterface) GetStats() (map[DriverExpvar]interface{}, error) {
	handleStats, err := di.driverFlowHandle.GetStatsForHandle()
	if err != nil {
		return nil, err
	}

	totalDriverStats, err := di.driverStatsHandle.GetStatsForHandle()
	if err != nil {
		return nil, err
	}
	totalFlows := di.totalFlows.Load()
	openFlows := di.openFlows.Swap(0)
	closedFlows := di.closedFlows.Swap(0)
	moreDataErrors := di.moreDataErrors.Swap(0)
	bufferSize := di.bufferSize.Load()

	return map[DriverExpvar]interface{}{
		totalFlowStats:  totalDriverStats,
		flowHandleStats: handleStats,
		flowStats: map[string]int64{
			"total":  totalFlows,
			"open":   openFlows,
			"closed": closedFlows,
		},
		driverStats: map[string]int64{
			"more_data_errors": moreDataErrors,
			"buffer_size":      bufferSize,
		},
	}, nil
}

// GetConnectionStats will read all flows from the driver and convert them into ConnectionStats.
// It returns the count of connections added to the active and closed buffers, respectively.
func (di *DriverInterface) GetConnectionStats(activeBuf *ConnectionBuffer, closedBuf *ConnectionBuffer, filter func(*ConnectionStats) bool) (int, int, error) {
	di.bufferLock.Lock()
	defer di.bufferLock.Unlock()

	startActive, startClosed := activeBuf.Len(), closedBuf.Len()

	var bytesRead uint32
	var totalBytesRead uint32
	// keep reading while driver says there is more data available
	for err := error(windows.ERROR_MORE_DATA); err == windows.ERROR_MORE_DATA; {
		err = windows.ReadFile(di.driverFlowHandle.Handle, di.readBuffer, &bytesRead, nil)
		if err != nil {
			if err == windows.ERROR_NO_MORE_ITEMS {
				break
			}
			if err != windows.ERROR_MORE_DATA {
				return 0, 0, fmt.Errorf("ReadFile: %w", err)
			}
			di.moreDataErrors.Inc()
		}
		totalBytesRead += bytesRead

		var buf []byte
		for bytesUsed := uint32(0); bytesUsed < bytesRead; bytesUsed += driver.PerFlowDataSize {
			buf = di.readBuffer[bytesUsed:]
			pfd := (*driver.PerFlowData)(unsafe.Pointer(&(buf[0])))

			if isFlowClosed(pfd.Flags) {
				c := closedBuf.Next()
				FlowToConnStat(c, pfd, di.enableMonotonicCounts)
				if !filter(c) {
					closedBuf.Reclaim(1)
					continue
				}
			} else {
				c := activeBuf.Next()
				FlowToConnStat(c, pfd, di.enableMonotonicCounts)
				if !filter(c) {
					activeBuf.Reclaim(1)
					continue
				}
			}
		}

		di.readBuffer = resizeDriverBuffer(int(totalBytesRead), di.readBuffer)
		di.bufferSize.Store(int64(len(di.readBuffer)))
	}

	activeCount := activeBuf.Len() - startActive
	closedCount := closedBuf.Len() - startClosed
	di.openFlows.Add(int64(activeCount))
	di.closedFlows.Add(int64(closedCount))
	di.totalFlows.Add(int64(activeCount + closedCount))

	return activeCount, closedCount, nil
}

func resizeDriverBuffer(compareSize int, buffer []uint8) []uint8 {
	// Explicitly setting len to 0 causes the ReadFile syscall to break, so allocate buffer with cap = len
	if compareSize >= cap(buffer)*2 {
		return make([]uint8, cap(buffer)*2)
	} else if compareSize <= cap(buffer)/2 {
		// Take the max of buffer/2 and compareSize to limit future array resizes
		return make([]uint8, int(math.Max(float64(cap(buffer)/2), float64(compareSize))))
	}
	return buffer
}

func minUint64(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

// setParams passes any configuration values from the config file down
// to the driver.
func (di *DriverInterface) setFlowParams() error {
	// set up the maximum flows

	// temporary setup.  Will set the maximum flows to the sum of the configured
	// max_tracked_connections and max_closed_connections_buffered, setting a
	// (hard_coded) maximum.  This will be updated to actually honor the separate
	// config values when the driver is updated to track them separately.

	// this makes it so that the config can clamp down, but can never make it
	// larger than the coded defaults above.
	maxFlows := minUint64(defaultMaxOpenFlows+defaultMaxClosedFlows, di.maxOpenFlows+di.maxClosedFlows)
	log.Debugf("Setting max flows in driver to %v", maxFlows)
	err := windows.DeviceIoControl(di.driverFlowHandle.Handle,
		driver.SetMaxFlowsIOCTL,
		(*byte)(unsafe.Pointer(&maxFlows)),
		uint32(unsafe.Sizeof(maxFlows)),
		nil,
		uint32(0), nil, nil)
	if err != nil {
		log.Warnf("Failed to set max number of flows to %v %v", maxFlows, err)
	}
	return err
}

func (di *DriverInterface) createFlowHandleFilters() ([]driver.FilterDefinition, error) {
	var filters []driver.FilterDefinition
	log.Debugf("Creating filters for all interfaces")
	if di.cfg.CollectTCPConns {
		filters = append(filters, driver.FilterDefinition{
			FilterVersion:  driver.Signature,
			Size:           driver.FilterDefinitionSize,
			Direction:      driver.DirectionOutbound,
			FilterLayer:    driver.LayerTransport,
			InterfaceIndex: uint64(0),
			Af:             windows.AF_INET,
			Protocol:       windows.IPPROTO_TCP,
		}, driver.FilterDefinition{
			FilterVersion:  driver.Signature,
			Size:           driver.FilterDefinitionSize,
			Direction:      driver.DirectionInbound,
			FilterLayer:    driver.LayerTransport,
			InterfaceIndex: uint64(0),
			Af:             windows.AF_INET,
			Protocol:       windows.IPPROTO_TCP,
		})
		if di.cfg.CollectIPv6Conns {
			filters = append(filters, driver.FilterDefinition{
				FilterVersion:  driver.Signature,
				Size:           driver.FilterDefinitionSize,
				Direction:      driver.DirectionOutbound,
				FilterLayer:    driver.LayerTransport,
				InterfaceIndex: uint64(0),
				Af:             windows.AF_INET6,
				Protocol:       windows.IPPROTO_TCP,
			}, driver.FilterDefinition{
				FilterVersion:  driver.Signature,
				Size:           driver.FilterDefinitionSize,
				Direction:      driver.DirectionInbound,
				FilterLayer:    driver.LayerTransport,
				InterfaceIndex: uint64(0),
				Af:             windows.AF_INET6,
				Protocol:       windows.IPPROTO_TCP,
			})
		}
	}

	if di.cfg.CollectUDPConns {
		filters = append(filters, driver.FilterDefinition{
			FilterVersion:  driver.Signature,
			Size:           driver.FilterDefinitionSize,
			Direction:      driver.DirectionOutbound,
			FilterLayer:    driver.LayerTransport,
			InterfaceIndex: uint64(0),
			Af:             windows.AF_INET,
			Protocol:       windows.IPPROTO_UDP,
		}, driver.FilterDefinition{
			FilterVersion:  driver.Signature,
			Size:           driver.FilterDefinitionSize,
			Direction:      driver.DirectionInbound,
			FilterLayer:    driver.LayerTransport,
			InterfaceIndex: uint64(0),
			Af:             windows.AF_INET,
			Protocol:       windows.IPPROTO_UDP,
		})
		if di.cfg.CollectIPv6Conns {
			filters = append(filters, driver.FilterDefinition{
				FilterVersion:  driver.Signature,
				Size:           driver.FilterDefinitionSize,
				Direction:      driver.DirectionOutbound,
				FilterLayer:    driver.LayerTransport,
				InterfaceIndex: uint64(0),
				Af:             windows.AF_INET6,
				Protocol:       windows.IPPROTO_UDP,
			}, driver.FilterDefinition{
				FilterVersion:  driver.Signature,
				Size:           driver.FilterDefinitionSize,
				Direction:      driver.DirectionInbound,
				FilterLayer:    driver.LayerTransport,
				InterfaceIndex: uint64(0),
				Af:             windows.AF_INET6,
				Protocol:       windows.IPPROTO_UDP,
			})
		}
	}

	return filters, nil
}
