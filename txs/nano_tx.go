package txs

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// NanoTX represents minimal info about a transaction for the conservative cache/mempool.
type NanoTX struct {
	types.TxHeader
	ID types.TransactionID

	Received time.Time

	Block types.BlockID
	Layer types.LayerID
}

// NewNanoTX converts a NanoTX instance from a MeshTransaction.
func NewNanoTX(mtx *types.MeshTransaction) *NanoTX { _ = "STUB: not implemented"; return nil }

// MaxSpending returns the maximal amount a transaction can spend.
func (n *NanoTX) MaxSpending() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *NanoTX) combinedHash(blockSeed []byte) []byte { _ = "STUB: not implemented"; return nil }

// Better returns true if this transaction takes priority than `other`.
// When the block seed is non-empty, this tx is being considered for a block.
// The block seed then is used to tie-break (deterministically) transactions for
// the same account/nonce.
func (n *NanoTX) Better(other *NanoTX, blockSeed []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// UpdateLayerMaybe updates the layer of a transaction if it's lower than the current value.
func (n *NanoTX) UpdateLayerMaybe(lid types.LayerID, bid types.BlockID) {
	_ = "STUB: not implemented"
	return
}

// UpdateLayer updates the layer of a transaction.
func (n *NanoTX) UpdateLayer(lid types.LayerID, bid types.BlockID) {
	_ = "STUB: not implemented"
	return
}
