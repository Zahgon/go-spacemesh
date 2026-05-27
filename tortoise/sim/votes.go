package sim

import (
	"math/rand"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// Voting contains blocks voting.
type Voting = types.Votes

// VotesGenerator allows to replace default votes generator.
// TODO(dshulyak) what is the best way to encapsulate all configuration that is required to generate votes?
type VotesGenerator func(rng *rand.Rand, layers []*types.Layer, i int) Voting

// PerfectVoting selects base ballot from previous layer and supports all blocks from previous layer.
// Used by default.
func PerfectVoting(rng *rand.Rand, layers []*types.Layer, _ int) Voting {
	_ = "STUB: not implemented"
	return *new(Voting)
}

// ConsistentVoting selects same base ballot for ballot at a specific index.
func ConsistentVoting(rng *rand.Rand, layers []*types.Layer, i int) Voting {
	_ = "STUB: not implemented"
	return *new(Voting)
}

// VaryingVoting votes using first generator for ballots before mid, and with second generator after mid.
func VaryingVoting(mid int, first, second VotesGenerator) VotesGenerator {
	_ = "STUB: not implemented"
	return *new(VotesGenerator)
}
