package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Namespace is the basic namespace where all metrics are defined under.
	Namespace = "spacemesh"
)

// NewCounter creates a Counter metrics under the global namespace returns nop if metrics are disabled.
func NewCounter(name, subsystem, help string, labels []string) *prometheus.CounterVec {
	_ = "STUB: not implemented"
	return nil
}

// NewGauge creates a Gauge metrics under the global namespace returns nop if metrics are disabled.
func NewGauge(name, subsystem, help string, labels []string) *prometheus.GaugeVec {
	_ = "STUB: not implemented"
	return nil
}

// NewHistogram creates a Histogram metrics under the global namespace returns nop if metrics are disabled.
func NewHistogram(name, subsystem, help string, labels []string) *prometheus.HistogramVec {
	_ = "STUB: not implemented"
	return nil
}

// NewSimpleHistogram returns a histogram without labels.
func NewSimpleHistogram(name, subsystem, help string, buckets []float64) prometheus.Histogram {
	_ = "STUB: not implemented"
	return *new(prometheus.Histogram)
}

// NewHistogramWithBuckets creates a Histogram metrics with custom buckets.
func NewHistogramWithBuckets(
	name, subsystem, help string,
	labels []string,
	buckets []float64,
) *prometheus.HistogramVec {
	_ = "STUB: not implemented"
	return nil
}

// receivedMessagesLatency measures the time a message was received relative to
// the time it should be sent in the protocol. Metrics are labeled by protocol
// and sign. Sign is either "pos" or "neg" and is chosen depending on the sign
// of the observed latency. Negative latencies occur when a message is received
// before the time it should be sent, this could happen in a network where
// nodes' clocks are not in sync.
var receivedMessagesLatency = NewHistogramWithBuckets(
	"message_latency_seconds",
	"",
	"Observed latency for message",
	[]string{"protocol", "type", "sign"},
	prometheus.ExponentialBuckets(0.1, 2, 12),
)

// ReportMessageLatency records the latency for the given protocol and message
// type, if the protocol consists only of a single message then the protocol
// should be provided as the message type.
func ReportMessageLatency(protocol, msgType string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// If the observation is negative make it positive.

// NewSimpleCounter creates a simple prometheus counter.
func NewSimpleCounter(subsystem, name, help string) prometheus.Counter {
	_ = "STUB: not implemented"
	return *new(prometheus.Counter)
}

func NewCounterOpts(subsystem, name, help string) prometheus.CounterOpts {
	_ = "STUB: not implemented"
	return *new(prometheus.CounterOpts)
}
