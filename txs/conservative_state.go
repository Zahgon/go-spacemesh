package txs

import (
	"context"
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

// CSConfig is the config for the conservative state/cache.
type CSConfig struct {
	BlockGasLimit     uint64
	NumTXsPerProposal int
}

func defaultCSConfig() CSConfig { _ = "STUB: not implemented"; return *new(CSConfig) }

// ConservativeStateOpt for configuring conservative state.
type ConservativeStateOpt func(cs *ConservativeState)

// WithCSConfig defines the config used for the conservative state.
func WithCSConfig(cfg CSConfig) ConservativeStateOpt {
	_ = "STUB: not implemented"
	return *new(ConservativeStateOpt)
}

// WithLogger defines logger for conservative state.
func WithLogger(logger *zap.Logger) ConservativeStateOpt {
	_ = "STUB: not implemented"
	return *new(ConservativeStateOpt)
}

// ConservativeState provides the conservative version of the VM state by taking into accounts of
// nonce and balances for pending transactions in un-applied blocks and mempool.
type ConservativeState struct {
	vmState

	logger *zap.Logger
	cfg    CSConfig
	db     sql.StateDatabase
	cache  *Cache
}

// NewConservativeState returns a ConservativeState.
func NewConservativeState(state vmState, db sql.StateDatabase, opts ...ConservativeStateOpt) *ConservativeState {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ConservativeState) getState(addr types.Address) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// SelectProposalTXs picks a specific number of random txs for miner to pack in a proposal.
func (cs *ConservativeState) SelectProposalTXs(lid types.LayerID, numEligibility int) []types.TransactionID {
	_ = "STUB: not implemented"
	return nil
}

func getProposalTXs(
	logger *zap.Logger,
	numTXs int,
	predictedBlock []*NanoTX,
	byAddrAndNonce map[types.Address][]*NanoTX,
) []types.TransactionID {
	_ = "STUB: not implemented"
	return nil
}

// randomly select transactions from the predicted block.

// Validation initializes validation request.
func (cs *ConservativeState) Validation(raw types.RawTx) system.ValidationRequest {
	_ = "STUB: not implemented"
	return *new(system.ValidationRequest)
}

// AddToCache adds the provided transaction to the conservative cache.
func (cs *ConservativeState) AddToCache(ctx context.Context, tx *types.Transaction, received time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// RevertCache reverts the conservative cache to the given layer.
func (cs *ConservativeState) RevertCache(revertTo types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ConservativeState) UpdateCache(
	ctx context.Context,
	lid types.LayerID,
	bid types.BlockID,
	results []types.TransactionWithResult,
	ineffective []types.Transaction,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetProjection returns the projected nonce and balance for an account, including
// pending transactions that are paced in proposals/blocks but not yet applied to the state.
func (cs *ConservativeState) GetProjection(addr types.Address) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// LinkTXsWithProposal associates the transactions to a proposal.
func (cs *ConservativeState) LinkTXsWithProposal(
	lid types.LayerID,
	pid types.ProposalID,
	tids []types.TransactionID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// LinkTXsWithBlock associates the transactions to a block.
func (cs *ConservativeState) LinkTXsWithBlock(lid types.LayerID, bid types.BlockID, tids []types.TransactionID) error {
	_ = "STUB: not implemented"
	return nil
}

// AddToDB adds a transaction to the database.
func (cs *ConservativeState) AddToDB(tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// HasTx returns true if transaction exists in the database.
func (cs *ConservativeState) HasTx(tid types.TransactionID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetMeshHash gets the aggregated layer hash at the specified layer.
func (cs *ConservativeState) GetMeshHash(lid types.LayerID) (types.Hash32, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// GetMeshTransaction retrieves a tx by its id.
func (cs *ConservativeState) GetMeshTransaction(tid types.TransactionID) (*types.MeshTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMeshTransactions retrieves a list of txs by their id's.
func (cs *ConservativeState) GetMeshTransactions(
	ids []types.TransactionID,
) ([]*types.MeshTransaction, map[types.TransactionID]struct{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTransactionsByAddress retrieves txs for a single address in between layers [from, to].
// Guarantees that transaction will appear exactly once, even if origin and recipient is the same,
// and in insertion order.
func (cs *ConservativeState) GetTransactionsByAddress(
	from, to types.LayerID,
	address types.Address,
) ([]*types.MeshTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ShuffleWithNonceOrder perform a Fisher-Yates shuffle on the transactions.
// Note that after shuffling, the original list of transactions are no longer in nonce order
// within the same principal. We simply check which principal occupies the spot after
// the shuffle and retrieve their transactions in nonce order.
func ShuffleWithNonceOrder(
	logger *zap.Logger,
	rng *rand.Rand,
	numTXs int,
	ntxs []*NanoTX,
	byAddrAndNonce map[types.Address][]*NanoTX,
) []types.TransactionID {
	_ = "STUB: not implemented"
	return nil
}

// if a spot is taken by a principal, we add its TX for the next eligible nonce

func (cs *ConservativeState) HasEvicted(tid types.TransactionID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
