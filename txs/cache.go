package txs

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	maxTXsPerAcct  = 100
	maxTXsPerNonce = 100
)

var (
	errBadNonce            = errors.New("bad nonce")
	errInsufficientBalance = errors.New("insufficient balance")
	errTooManyNonce        = errors.New("account has too many nonce pending")
	errLayerNotInOrder     = errors.New("layers not applied in order")
)

// a candidate for the mempool.
type candidate struct {
	// this is the best tx among all the txs with the same nonce
	best        *NanoTX
	postBalance uint64
}

func (s *candidate) id() types.TransactionID {
	_ = "STUB: not implemented"
	return *new(types.TransactionID)
}

func (s *candidate) layer() types.LayerID { _ = "STUB: not implemented"; return *new(types.LayerID) }

func (s *candidate) block() types.BlockID { _ = "STUB: not implemented"; return *new(types.BlockID) }

func (s *candidate) nonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *candidate) maxSpending() uint64 { _ = "STUB: not implemented"; return 0 }

type accountCache struct {
	addr         types.Address
	txsByNonce   *list.List
	startNonce   uint64
	startBalance uint64
	// moreInDB is used to indicate that an account has pending txs in db that need to be
	// reconsidered for mempool after a layer is applied.
	// - there are too many nonces for an account in the mempool. the extra (higher nonce) txs are in db only.
	// - txs deemed insufficient balance by the conservative state, but can be feasible after a layer applied
	//   (that may contain incoming funds for that account)
	// - a better tx arrived (higher fee) and made higher nonce txs infeasible due to insufficient balance
	//   deemed by conservative state.
	// TODO: evict accounts that only has DB-only txs
	// https://github.com/spacemeshos/go-spacemesh/issues/3668
	moreInDB bool

	cachedTXs map[types.TransactionID]*NanoTX // shared with the cache instance
}

func (ac *accountCache) nextNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (ac *accountCache) availBalance() uint64 { _ = "STUB: not implemented"; return 0 }

func (ac *accountCache) precheck(ntx *NanoTX) (*list.Element, *candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ac *accountCache) accept(logger *zap.Logger, ntx *NanoTX, blockSeed []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// insert at the first position

// existing nonce

// propagate the balance change

func (ac *accountCache) addBatch(logger *zap.Logger, nonce2TXs map[uint64][]*NanoTX, blockSeed []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func findBest(ntxs []*NanoTX, balance uint64, blockSeed []byte) *NanoTX {
	_ = "STUB: not implemented"
	return nil
}

// Adding a tx to the account cache. Possible outcomes:
//   - nonce is smaller than the next nonce in state: reject from cache
//   - too many txs present: reject from cache
//   - nonce already exists in the cache:
//     if it is better than the best candidate in that nonce group, swap
//   - nonce not present: add to cache.
func (ac *accountCache) add(logger *zap.Logger, tx *types.Transaction, received time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (ac *accountCache) addPendingFromNonce(
	logger *zap.Logger,
	db sql.StateDatabase,
	nonce uint64,
	applied types.LayerID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Find the first nonce without a layer.
// A nonce with a valid layer indicates that it's already packed in a proposal/block.
func (ac *accountCache) getMempool(logger *zap.Logger) []*NanoTX {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: this is the only point in time when we reconsider those previously rejected txs,
// because applying a layer changes the conservative balance in the cache.
func (ac *accountCache) resetAfterApply(
	logger *zap.Logger,
	db sql.StateDatabase,
	nextNonce, newBalance uint64,
	applied types.LayerID,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ac *accountCache) evictPendingNonce(db sql.StateDatabase) error {
	_ = "STUB: not implemented"
	return nil
}

func (ac *accountCache) shouldEvict() bool { _ = "STUB: not implemented"; return false }

type stateFunc func(types.Address) (uint64, uint64)

type Cache struct {
	logger *zap.Logger
	stateF stateFunc

	mu        sync.Mutex
	pending   map[types.Address]*accountCache
	cachedTXs map[types.TransactionID]*NanoTX // shared with accountCache instances
}

func NewCache(s stateFunc, logger *zap.Logger) *Cache { _ = "STUB: not implemented"; return nil }

func groupTXsByPrincipal(logger *zap.Logger, mtxs []*types.MeshTransaction) map[types.Address]map[uint64][]*NanoTX {
	_ = "STUB: not implemented"
	return nil
}

// buildFromScratch builds the cache from database.
func (c *Cache) buildFromScratch(db sql.StateDatabase) error { _ = "STUB: not implemented"; return nil }

// BuildFromTXs builds the cache from the provided transactions.
func (c *Cache) BuildFromTXs(rst []*types.MeshTransaction, blockSeed []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) createAcctIfNotPresent(addr types.Address) { _ = "STUB: not implemented"; return }

func (c *Cache) MoreInDB(addr types.Address) bool { _ = "STUB: not implemented"; return false }

func (c *Cache) cleanupAccounts(accounts ...types.Address) { _ = "STUB: not implemented"; return }

//   - errTooManyNonce: when a principal has way too many nonces, we don't want to blow up the memory. they should
//     be stored in db and retrieved after each earlier nonce is applied.
func acceptable(err error) bool { _ = "STUB: not implemented"; return false }

func (c *Cache) Add(ctx context.Context, db sql.StateDatabase, tx *types.Transaction, received time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// Get gets a transaction from the cache.
func (c *Cache) Get(tid types.TransactionID) *NanoTX { _ = "STUB: not implemented"; return nil }

// Has returns true if transaction exists in the cache.
func (c *Cache) Has(tid types.TransactionID) bool { _ = "STUB: not implemented"; return false }

func (c *Cache) has(tid types.TransactionID) bool { _ = "STUB: not implemented"; return false }

// LinkTXsWithProposal associates the transactions to a proposal.
func (c *Cache) LinkTXsWithProposal(
	db sql.StateDatabase,
	lid types.LayerID,
	pid types.ProposalID,
	tids []types.TransactionID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// LinkTXsWithBlock associates the transactions to a block.
func (c *Cache) LinkTXsWithBlock(
	db sql.StateDatabase,
	lid types.LayerID,
	bid types.BlockID,
	tids []types.TransactionID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// updateLayer associates the transactions to a layer and optionally a block.
// A transaction is tagged with a layer when it's included in a proposal/block.
// If a transaction is included in multiple proposals/blocks in different layers,
// the lowest layer is retained.
func (c *Cache) updateLayer(lid types.LayerID, bid types.BlockID, tids []types.TransactionID) {
	_ = "STUB: not implemented"
	return
}

// transaction is not considered best in its nonce group

func (c *Cache) applyEmptyLayer(db sql.StateDatabase, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyLayer retires the applied transactions from the cache and updates the balances.
func (c *Cache) ApplyLayer(
	ctx context.Context,
	db sql.StateDatabase,
	lid types.LayerID,
	bid types.BlockID,
	results []types.TransactionWithResult,
	ineffective []types.Transaction,
) error {
	_ = "STUB: not implemented"
	return nil
}

// commit results before reporting them
// TODO(dshulyak) save results in vm

func (c *Cache) RevertToLayer(db sql.StateDatabase, revertTo types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetProjection returns the projected nonce and balance for an account, including
// pending transactions that are paced in proposals/blocks but not yet applied to the state.
func (c *Cache) GetProjection(addr types.Address) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// GetMempool returns all the transactions that eligible for a proposal/block.
func (c *Cache) GetMempool() map[types.Address][]*NanoTX { _ = "STUB: not implemented"; return nil }

// checkApplyOrder returns an error if layers were not applied in order.
func checkApplyOrder(logger *zap.Logger, db sql.StateDatabase, toApply types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

func addToProposal(db sql.StateDatabase, lid types.LayerID, pid types.ProposalID, tids []types.TransactionID) error {
	_ = "STUB: not implemented"
	return nil
}

func addToBlock(db sql.StateDatabase, lid types.LayerID, bid types.BlockID, tids []types.TransactionID) error {
	_ = "STUB: not implemented"
	return nil
}

func undoLayers(db sql.StateDatabase, from types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

func getNextIncluded(
	db sql.Executor,
	id types.TransactionID,
	after types.LayerID,
) (types.LayerID, types.BlockID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), *new(types.BlockID), nil
}
