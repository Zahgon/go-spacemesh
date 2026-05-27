package metrics

import (
	"time"
)

// StartPushingMetrics begins pushing metrics to the url specified by the --metrics-push flag
// with period specified by the --metrics-push-period flag.
func StartPushingMetrics(
	url, username, password string,
	headers map[string]string,
	period time.Duration,
	nodeID, networkID string,
) {
	_ = "STUB: not implemented"
	return
}
