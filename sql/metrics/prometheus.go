package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	enabledDBStat = "ENABLE_DBSTAT_VTAB"
	subsystem     = "database" // subsystem shared by all metrics exposed by this package.
)

// DBMetricsCollector collects metrics from db.
type DBMetricsCollector struct {
	logger        *zap.Logger
	checkInterval time.Duration
	db            sql.StateDatabase
	tablesList    map[string]struct{}
	eg            errgroup.Group
	cancel        context.CancelFunc

	tableSize *prometheus.GaugeVec
	indexSize *prometheus.GaugeVec
	totalSize *prometheus.GaugeVec
}

// NewDBMetricsCollector creates new DBMetricsCollector.
func NewDBMetricsCollector(
	ctx context.Context,
	db sql.StateDatabase,
	logger *zap.Logger,
	checkInterval time.Duration,
) *DBMetricsCollector {
	_ = "STUB: not implemented"
	return nil
}

// Close closes DBMetricsCollector.
func (d *DBMetricsCollector) Close() { _ = "STUB: not implemented"; return }

// CollectMetrics collects metrics from db.
func (d *DBMetricsCollector) CollectMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *DBMetricsCollector) collect() error { _ = "STUB: not implemented"; return nil }

func (d *DBMetricsCollector) checkCompiledWithDBStat() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// getListOfTables returns list of tables in db. Need to separate size indexes and tables.
func (d *DBMetricsCollector) getListOfTables() (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
