// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package info

// TraceWriterInfo represents statistics from the trace writer.
type TraceWriterInfo struct {
	Payloads int64
	Traces   int64
	Events   int64
	Spans    int64
	Errors   int64
	Retries  int64

	// Bytes specifies the actual number of bytes successfully written to the intake.
	Bytes int64

	// BytesUncompressed specifies the number of bytes that were written to
	// the intake before compression. This metric is used solely to verify the
	// accuracy of BytesEstimated (see below).
	BytesUncompressed int64

	// BytesEstimated specifies the estimated size of payloads before they were
	// marshalled and compressed. It is purely a heuristic used to estimate the
	// size of traces stored in memory in order to know when they have reached a
	// size big enough to be marshalled, compressed and flushed, in order to avoid
	// reaching the intakes' maximum payload size.
	BytesEstimated int64
}

// StatsWriterInfo represents statistics from the stats writer.
type StatsWriterInfo struct {
	Payloads       int64
	ClientPayloads int64
	StatsBuckets   int64
	StatsEntries   int64
	Errors         int64
	Retries        int64
	Splits         int64
	Bytes          int64
}

// UpdateTraceWriterInfo updates internal trace writer stats
func UpdateTraceWriterInfo(tws TraceWriterInfo) {
	infoMu.Lock()
	defer infoMu.Unlock()
	traceWriterInfo = tws
}

func publishTraceWriterInfo() interface{} {
	infoMu.RLock()
	defer infoMu.RUnlock()
	return traceWriterInfo
}

// UpdateStatsWriterInfo updates internal stats writer stats
func UpdateStatsWriterInfo(sws StatsWriterInfo) {
	infoMu.Lock()
	defer infoMu.Unlock()
	statsWriterInfo = sws
}

func publishStatsWriterInfo() interface{} {
	infoMu.RLock()
	defer infoMu.RUnlock()
	return statsWriterInfo
}
