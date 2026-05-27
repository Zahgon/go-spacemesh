package server

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/spacemeshos/go-spacemesh/metrics"
)

const (
	namespace  = "server"
	protoLabel = "protocol"
)

var (
	targetQueue = metrics.NewGauge(
		"target_queue",
		namespace,
		"target size of the queue",
		[]string{protoLabel},
	)
	queue = metrics.NewGauge(
		"queue",
		namespace,
		"actual size of the queue",
		[]string{protoLabel},
	)
	targetRps = metrics.NewGauge(
		"rps",
		namespace,
		"target requests per second",
		[]string{protoLabel},
	)
	requests = metrics.NewCounter(
		"requests",
		namespace,
		"requests counter",
		[]string{protoLabel, "state"},
	)
	clientRequests = metrics.NewCounter(
		"client_requests",
		namespace,
		"client request counter",
		[]string{protoLabel, "state"})
	clientLatency = metrics.NewHistogramWithBuckets(
		"client_latency_seconds",
		namespace,
		"latency since initiating a request",
		[]string{protoLabel, "result"},
		prometheus.ExponentialBuckets(0.01, 2, 20),
	)
	serverLatency = metrics.NewHistogramWithBuckets(
		"server_latency_seconds",
		namespace,
		"latency since accepting new stream",
		[]string{protoLabel},
		prometheus.ExponentialBuckets(0.01, 2, 20),
	)
	inQueueLatency = metrics.NewHistogramWithBuckets(
		"in_queue_latency_seconds",
		namespace,
		"latency spent in queue",
		[]string{protoLabel},
		prometheus.ExponentialBuckets(0.01, 2, 20),
	)
)

func newTracker(protocol string) *tracker { _ = "STUB: not implemented"; return nil }

type tracker struct {
	targetQueue                         prometheus.Gauge
	queue                               prometheus.Gauge
	targetRps                           prometheus.Gauge
	completed                           prometheus.Counter
	failed                              prometheus.Counter
	accepted                            prometheus.Counter
	dropped                             prometheus.Counter
	clientSucceeded                     prometheus.Counter
	clientFailed                        prometheus.Counter
	clientServerError                   prometheus.Counter
	inQueueLatency                      prometheus.Observer
	serverLatency                       prometheus.Observer
	clientLatency, clientLatencyFailure prometheus.Observer
}
