package fixture

import (
	"math/rand"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// NewRewardsGenerator with some random parameters.
func NewRewardsGenerator() *RewardsGenerator { _ = "STUB: not implemented"; return nil }

// RewardsGenerator generates random rewards.
// Rewards are not syntactically or contextually valid. This is for testing databases and APIs.
type RewardsGenerator struct {
	rng *rand.Rand

	UniqueCoinbase bool
	Addrs          []types.Address
	Layers         []types.LayerID
}

// WithSeed update randomness source.
func (g *RewardsGenerator) WithSeed(seed int64) *RewardsGenerator {
	_ = "STUB: not implemented"
	return nil
}

// WithAddresses update addresses.
func (g *RewardsGenerator) WithAddresses(n int) *RewardsGenerator {
	_ = "STUB: not implemented"
	return nil
}

// WithLayers updates layers.
func (g *RewardsGenerator) WithLayers(start, n int) *RewardsGenerator {
	_ = "STUB: not implemented"
	return nil
}

// WithUniqueCoinbase will generate every reward with different address.
func (g *RewardsGenerator) WithUniqueCoinbase() *RewardsGenerator {
	_ = "STUB: not implemented"
	return nil
}

// Next generates Reward.
func (g *RewardsGenerator) Next() *types.Reward { _ = "STUB: not implemented"; return nil }
