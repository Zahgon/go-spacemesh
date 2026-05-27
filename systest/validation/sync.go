package validation

import (
	"context"
	"time"

	"github.com/spacemeshos/go-spacemesh/systest/cluster"
)

// Periodic runs validation once in a period, starting immediately.
func Periodic(ctx context.Context, period time.Duration, f Validation) error {
	_ = "STUB: not implemented"
	return nil
}

type Validation func(context.Context) error

func isSynced(ctx context.Context, node *cluster.NodeClient) bool {
	_ = "STUB: not implemented"
	return false
}

func Sync(c *cluster.Cluster, tolerate int) Validation {
	_ = "STUB: not implemented"
	return *new(Validation)
}

type SyncValidation struct {
	failures []int
	tolerate int
}

func (s *SyncValidation) OnData(id int, synced bool) error { _ = "STUB: not implemented"; return nil }
