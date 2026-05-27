package transactions

import (
	"context"
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

// Add transaction to the database or update the header if it wasn't set originally.
func Add(db sql.Executor, tx *types.Transaction, received time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// AddToProposal associates a transaction with a proposal.
func AddToProposal(db sql.Executor, tid types.TransactionID, lid types.LayerID, pid types.ProposalID) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteProposalTxsBefore(db sql.Executor, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// HasProposalTX returns true if the given transaction is included in the given proposal.
func HasProposalTX(db sql.Executor, pid types.ProposalID, tid types.TransactionID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddToBlock associates a transaction with a block.
func AddToBlock(db sql.Executor, tid types.TransactionID, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// HasBlockTX returns true if the given transaction is included in the given block.
func HasBlockTX(db sql.Executor, bid types.BlockID, tid types.TransactionID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetAppliedLayer returns layer when transaction was applied.
func GetAppliedLayer(db sql.Executor, tid types.TransactionID) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// UndoLayers unset all transactions to `statePending` from `from` layer to the max layer with applied transactions.
func UndoLayers(tx sql.Transaction, from types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// tx, header, layer, block, timestamp.
func decodeTransaction(id types.TransactionID, stmt *sql.Statement) (*types.MeshTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get gets a transaction from database.
// Layer and Block fields are set if transaction was applied.
// If transaction is included, but not applied check references in proposals and blocks.
func Get(db sql.Executor, id types.TransactionID) (tx *types.MeshTransaction, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to the transactions
// with specified ids. For non-existent transactions, the corresponding items
// are set to -1.
func GetBlobSizes(db sql.Executor, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads transaction as an encoded blob, ready to be sent over the wire.
func LoadBlob(ctx context.Context, db sql.Executor, id []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

// Has returns true if transaction is stored in the database.
func Has(db sql.Executor, id types.TransactionID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func HasEvicted(db sql.Executor, id types.TransactionID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetByAddress finds all transactions for an address.
func GetByAddress(db sql.Executor, from, to types.LayerID, address types.Address) ([]*types.MeshTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddressesWithPendingTransactions returns list of addresses with pending transactions.
// Query is expensive, meant to be used only on startup.
func AddressesWithPendingTransactions(db sql.Executor) ([]types.AddressNonce, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAcctPendingFromNonce get all pending transactions with nonce after `from` for the given address.
func GetAcctPendingFromNonce(db sql.Executor, address types.Address, from uint64) ([]*types.MeshTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAcctPendingToNonce get all pending transactions with nonce before `to` for the given address.
func GetAcctPendingToNonce(db sql.Executor, address types.Address, to uint64) ([]types.TransactionID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetEvicted(db sql.Executor, id types.TransactionID) error {
	_ = "STUB: not implemented"
	return nil
}

func Delete(db sql.Executor, id types.TransactionID) error { _ = "STUB: not implemented"; return nil }

func PruneEvicted(db sql.Executor, before time.Time) error { _ = "STUB: not implemented"; return nil }

// query MUST ensure that this order of fields tx, header, layer, block, timestamp, id.
func queryPending(
	db sql.Executor,
	query string,
	encoder func(*sql.Statement),
	errStr string,
) (rst []*types.MeshTransaction, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddResult adds result for the transaction.
func AddResult(tx sql.Transaction, id types.TransactionID, rst *types.TransactionResult) error {
	_ = "STUB: not implemented"
	return nil
}

// TransactionInProposal returns lowest layer of the proposal where tx is included after the specified layer.
func TransactionInProposal(db sql.Executor, id types.TransactionID, after types.LayerID) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// TransactionInBlock returns lowest layer and id of the block where tx is included after the specified layer.
func TransactionInBlock(
	db sql.Executor,
	id types.TransactionID,
	after types.LayerID,
) (types.BlockID, types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.BlockID), *new(types.LayerID), nil
}

func IterateTransactionsOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(tx *types.MeshTransaction, result *types.TransactionResult) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
