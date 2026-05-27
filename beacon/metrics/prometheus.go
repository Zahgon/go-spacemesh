package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spacemeshos/fixed"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/metrics"
)

const (
	subsystem   = "beacons"
	labelEpoch  = "epoch"
	labelBeacon = "beacon"
)

// BeaconStats hold metadata for each beacon value.
type BeaconStats struct {
	Epoch      types.EpochID
	Beacon     string
	Weight     fixed.Fixed
	WeightUnit int
}

// GatherCB returns stats for the observed and calculated beacons.
type GatherCB func() ([]*BeaconStats, *BeaconStats)

// BeaconMetricsCollector is a prometheus Collector for beacon metrics.
type BeaconMetricsCollector struct {
	gather GatherCB
	logger *zap.Logger

	registry               *prometheus.Registry
	observedBeaconCount    *prometheus.Desc
	observedBeaconWeight   *prometheus.Desc
	calculatedBeaconWeight *prometheus.Desc
}

var nameCalculatedWeight = prometheus.BuildFQName(metrics.Namespace, subsystem, "beacon_calculated_weight")

func MetricNameCalculatedWeight() string { _ = "STUB: not implemented"; return "" }

// NewBeaconMetricsCollector creates a prometheus Collector for beacons.
func NewBeaconMetricsCollector(cb GatherCB, logger *zap.Logger) *BeaconMetricsCollector {
	_ = "STUB: not implemented"
	return nil
}

// Start registers the Collector with specified prometheus registry and starts the metrics collection.
func (bmc *BeaconMetricsCollector) Start(registry *prometheus.Registry) {
	_ = "STUB: not implemented"
	return
}

// use Register instead of MustRegister because during app test, multiple instances
// will register the same set of metrics with the default registry and panic

// Stop unregisters the Collector with specified prometheus registry and stops the metrics collection.
func (bmc *BeaconMetricsCollector) Stop() { _ = "STUB: not implemented"; return }

// Describe implements Collector.
func (bmc *BeaconMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

// Collect implements Collector.
func (bmc *BeaconMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

// export the calculated beacon for the target epoch for ease of monitoring along with the observed beacons

var NumMaliciousProps = metrics.NewCounter(
	"malicious_proposals",
	subsystem,
	"number of malicious proposals",
	[]string{},
).WithLabelValues()
