package mesh

import (
	"context"
	"errors"
	"sync"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var (
	ErrLayerNotInOrder = errors.New("layers not applied in order")
	ErrLayerApplied    = errors.New("layer already applied")
)

type Executor struct {
	logger   *zap.Logger
	db       sql.Executor
	atxsdata *atxsdata.Data
	vm       vmState
	cs       conservativeState

	mu sync.Mutex
}

func NewExecutor(db sql.Executor, atxsdata *atxsdata.Data, vm vmState, cs conservativeState, lg *zap.Logger) *Executor {
	_ = "STUB: not implemented"
	return nil
}

// Revert reverts the VM state and conservative cache to the given layer.
func (e *Executor) Revert(ctx context.Context, revertTo types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// ExecuteOptimistic executes the specified transactions and returns a block that contains
// only successfully executed transactions.
func (e *Executor) ExecuteOptimistic(
	ctx context.Context,
	lid types.LayerID,
	tickHeight uint64,
	rewards []types.AnyReward,
	tids []types.TransactionID,
) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Execute transactions in the specified block and update the conservative cache.
func (e *Executor) Execute(ctx context.Context, lid types.LayerID, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) convertRewards(lid types.LayerID, rewards []types.AnyReward) ([]types.CoinbaseReward, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Executor) executeEmpty(ctx context.Context, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) checkOrder(lid types.LayerID) error { _ = "STUB: not implemented"; return nil }

func updateResults(bid types.BlockID, executed []types.TransactionWithResult) {
	_ = "STUB: not implemented"
	return
}

// getExecutableTxs retrieves a list of txs filtering transaction that were previously executed.
func (e *Executor) getExecutableTxs(ids []types.TransactionID) ([]types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
