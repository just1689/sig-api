package trc

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func newId() string {
	return RandStringBytesMaskImprSrc(16)
}

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, (n+1)/2)
	if _, err := src.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)[:n]
}

func NewTrace(serviceName string) Span {
	var id = newId()
	var traceId = newId()
	result := Span{
		ID:     &id,
		Shared: false,
		LocalEndpoint: map[string]string{
			"serviceName": serviceName,
		},
		Timestamp: time.Now().Unix() * 1000 * 1000,
		TraceID:   &traceId,
	}
	return result
}
func NewSpan(traceID *string, serviceName, rpcName string, d time.Duration) Span {
	var id = newId()
	result := Span{
		Duration: d.Nanoseconds() / 1000,
		ID:       &id,
		Name:     rpcName,
		Shared:   true,
		LocalEndpoint: map[string]string{
			"serviceName": serviceName,
		},
		Timestamp: time.Now().Unix() * 1000 * 1000,
		TraceID:   traceID,
	}
	return result
}

//type Span struct {
//	TraceID string `json:"traceId"`
//	Name    string `json:"name"`
//
//	ParentID string `json:"parentId,omitempty"`
//	ID       string `json:"id"`
//	Duration int    `json:"duration,omitempty"`
//}

type Span struct {
	Duration int64   `json:"duration,omitempty"`
	ID       *string `json:"id"`

	// The logical operation this span represents in lowercase (e.g. rpc method).
	// Leave absent if unknown.
	//
	// As these are lookup labels, take care to ensure names are low cardinality.
	// For example, do not embed variables into the name.
	//
	Name string `json:"name,omitempty"`

	// The parent span ID or absent if this the root span in a trace.
	// Max Length: 16
	// Min Length: 16
	// Pattern: [a-z0-9]{16}
	ParentID string `json:"parentId,omitempty"`

	// True if we are contributing to a span started by another tracer (ex on a different host).
	Shared bool `json:"shared,omitempty"`

	LocalEndpoint map[string]string `json:"localEndpoint,omitempty"`

	// Tags give your span context for search, viewing and analysis.
	Tags Tags `json:"tags,omitempty"`

	// Epoch **microseconds** of the start of this span, possibly absent if incomplete.
	//
	// For example, 1502787600000000 corresponds to 2017-08-15 09:00 UTC
	//
	// This value should be set directly by instrumentation, using the most precise
	// value possible. For example, gettimeofday or multiplying epoch millis by 1000.
	//
	// There are three known edge-cases where this could be reported absent.
	//  * A span was allocated but never started (ex not yet received a timestamp)
	//  * The span's start event was lost
	//  * Data about a completed span (ex tags) were sent after the fact
	//
	Timestamp int64 `json:"timestamp,omitempty"`

	// Randomly generated, unique identifier for a trace, set on all spans within it.
	//
	// Encoded as 16 or 32 lowercase hex characters corresponding to 64 or 128 bits.
	// For example, a 128bit trace ID looks like 4e441824ec2b6a44ffdc9bb9a6453df3
	//
	// Required: true
	// Max Length: 32
	// Min Length: 16
	// Pattern: [a-z0-9]{16,32}
	TraceID *string `json:"traceId"`
}
type Tags map[string]string
