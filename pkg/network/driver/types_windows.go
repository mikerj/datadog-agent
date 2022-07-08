// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo.exe -godefs -- -fsigned-char C:\src\agent.isabelle\pkg\network\driver\types.go

package driver

const Signature = 0xddfd00000012

const (
	GetStatsIOCTL             = 0x122004
	SetFlowFilterIOCTL        = 0x122010
	SetDataFilterIOCTL        = 0x12200c
	GetFlowsIOCTL             = 0x122014
	SetMaxOpenFlowsIOCTL      = 0x122024
	SetMaxClosedFlowsIOCTL    = 0x122028
	FlushPendingHttpTxnsIOCTL = 0x122020
	GetHttpTransactionsIOCTL  = 0x122030
	EnableHttpIOCTL           = 0x122034
)

type FilterAddress struct {
	Af         uint64
	V4_address [4]uint8
	V4_padding [4]uint8
	V6_address [16]uint8
	Mask       uint64
}

type FilterDefinition struct {
	FilterVersion  uint64
	Size           uint64
	FilterLayer    uint64
	Af             uint64
	LocalAddress   FilterAddress
	RemoteAddress  FilterAddress
	LocalPort      uint64
	RemotePort     uint64
	Protocol       uint64
	Direction      uint64
	InterfaceIndex uint64
}

const FilterDefinitionSize = 0x98

type FilterPacketHeader struct {
	FilterVersion    uint64
	Sz               uint64
	SkippedSinceLast uint64
	FilterId         uint64
	Direction        uint64
	PktSize          uint64
	Af               uint64
	OwnerPid         uint64
	Timestamp        uint64
}

const FilterPacketHeaderSize = 0x48

type HandleStats struct {
	Read_calls             int64
	Read_calls_outstanding int64
	Read_calls_completed   int64
	Read_calls_cancelled   int64
	Write_calls            int64
	Write_bytes            int64
	Ioctl_calls            int64
}
type FlowStats struct {
	Packets_observed                        int64
	Packets_processed                       int64
	Open_flows                              int64
	Total_flows                             int64
	Num_flow_collisions                     int64
	Num_flow_structures                     int64
	Num_flow_closed_structures              int64
	Peak_num_flow_structures                int64
	Num_flow_alloc_skipped_max_exceeded     int64
	Num_flow_closed_dropped_max_exceeded    int64
	Open_table_adds                         int64
	Open_table_removes                      int64
	Closed_table_adds                       int64
	Closed_table_removes                    int64
	Num_flows_no_handle                     int64
	Peak_num_flows_no_handle                int64
	Num_flows_missed_max_no_handle_exceeded int64
	Num_packets_after_flow_closed           int64
}
type TransportStats struct {
	Packets_processed    int64
	Read_packets_skipped int64
	Packets_reported     int64
}
type HttpStats struct {
	Packets_processed             int64
	Num_flow_collisions           int64
	Num_flows_missed_max_exceeded int64
	Read_batch_skipped            int64
	Batches_reported              int64
}
type Stats struct {
	Handle_stats    HandleStats
	Flow_stats      FlowStats
	Transport_stats TransportStats
	Http_stats      HttpStats
}
type DriverStats struct {
	FilterVersion uint64
	Total         Stats
	Handle        Stats
}

const DriverStatsSize = 0x218

type PerFlowData struct {
	FlowHandle         uint64
	ProcessId          uint64
	AddressFamily      uint16
	Protocol           uint16
	Flags              uint32
	LocalAddress       [16]uint8
	RemoteAddress      [16]uint8
	PacketsOut         uint64
	MonotonicSentBytes uint64
	TransportBytesOut  uint64
	PacketsIn          uint64
	MonotonicRecvBytes uint64
	TransportBytesIn   uint64
	Timestamp          uint64
	LocalPort          uint16
	RemotePort         uint16
	U                  [32]byte
}
type TCPFlowData struct {
	IRTT            uint64
	SRTT            uint64
	RttVariance     uint64
	RetransmitCount uint64
}
type UDPFlowData struct {
	Reserved uint64
}

const PerFlowDataSize = 0x94

const (
	FlowDirectionMask     = 0x300
	FlowDirectionBits     = 0x8
	FlowDirectionInbound  = 0x1
	FlowDirectionOutbound = 0x2

	FlowClosedMask         = 0x10
	TCPFlowEstablishedMask = 0x20
)

const (
	DirectionInbound  = 0x0
	DirectionOutbound = 0x1
)

const (
	LayerTransport = 0x1
)

type HttpTransactionType struct {
	RequestStarted     uint64
	ResponseLastSeen   uint64
	Tup                ConnTupleType
	RequestMethod      uint32
	ResponseStatusCode uint16
	RequestFragment    [25]uint8
	Pad_cgo_0          [1]byte
}
type ConnTupleType struct {
	CliAddr [16]uint8
	SrvAddr [16]uint8
	CliPort uint16
	SrvPort uint16
	Family  uint16
}
type HttpMethodType uint32

const (
	HttpBatchSize           = 0xf
	HttpBufferSize          = 0x19
	HttpTransactionTypeSize = 0x58
)
