package prune

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/timesync"
)

type Opt func(*Pruner)

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func New(db sql.StateDatabase, safeDist uint32, activesetEpoch types.EpochID, opts ...Opt) *Pruner {
	_ = "STUB: not implemented"
	return nil
}

type Pruner struct {
	logger         *zap.Logger
	db             sql.StateDatabase
	safeDist       uint32
	activesetEpoch types.EpochID
}

func Run(ctx context.Context, p *Pruner, clock *timesync.NodeClock, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (p *Pruner) Prune(current types.LayerID) error { _ = "STUB: not implemented"; return nil }

// current - 1 as activesets will be fetched in hare eligibility oracle
// for example if we are in epoch 9, we want to prune 7 and below
// as activesets from 8 will be still be needed at the beginning of epoch 8
