package tests

import (
	"context"
	"errors"
)

var errMetricNotFound = errors.New("metric not found")

func fetchCounterMetric(
	ctx context.Context,
	url string,
	metricName string,
	labelFilters map[string]string,
) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check if the metric has the specified labels
