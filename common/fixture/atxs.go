package fixture

import (
	"math/rand"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// NewAtxsGenerator with some random parameters.
func NewAtxsGenerator() *AtxsGenerator { _ = "STUB: not implemented"; return nil }

// AtxsGenerator generates random activations.
// Activations are not syntactically or contextually valid. This is for testing databases and APIs.
type AtxsGenerator struct {
	rng *rand.Rand

	Epochs []types.EpochID
	Addrs  []types.Address
}

// WithSeed update randomness source.
func (g *AtxsGenerator) WithSeed(seed int64) *AtxsGenerator { _ = "STUB: not implemented"; return nil }

// WithEpochs update epochs ids.
func (g *AtxsGenerator) WithEpochs(start, n int) *AtxsGenerator {
	_ = "STUB: not implemented"
	return nil
}

// Next generates ActivationTx.
func (g *AtxsGenerator) Next() *types.ActivationTx { _ = "STUB: not implemented"; return nil }
