package fixture

import (
	"math/rand"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// NewTransactionResultGenerator with some random parameters.
func NewTransactionResultGenerator() *TransactionResultGenerator {
	_ = "STUB: not implemented"
	return nil
}

// TransactionResultGenerator generates random transaction with results.
// Transactions are not syntactically or contextually valid. This is for testing databases and APIs.
type TransactionResultGenerator struct {
	rng *rand.Rand

	Blocks []types.BlockID
	Addrs  []types.Address
	Layers []types.LayerID
}

// WithSeed update randomness source.
func (g *TransactionResultGenerator) WithSeed(seed int64) *TransactionResultGenerator {
	_ = "STUB: not implemented"
	return nil
}

// WithBlocks update blocks ids.
func (g *TransactionResultGenerator) WithBlocks(n int) *TransactionResultGenerator {
	_ = "STUB: not implemented"
	return nil
}

// WithAddresses update addresses.
func (g *TransactionResultGenerator) WithAddresses(n int) *TransactionResultGenerator {
	_ = "STUB: not implemented"
	return nil
}

// WithLayers updates layers.
func (g *TransactionResultGenerator) WithLayers(start, n int) *TransactionResultGenerator {
	_ = "STUB: not implemented"
	return nil
}

// Next generates TransactionWithResult.
func (g *TransactionResultGenerator) Next() *types.TransactionWithResult {
	_ = "STUB: not implemented"
	return nil
}
