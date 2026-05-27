package tortoise

import (
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

func newVerifying(config Config, state *state) *verifying { _ = "STUB: not implemented"; return nil }

type verifying struct {
	Config
	*state

	// total weight of good ballots from verified + 1 up to last processed
	totalGoodWeight weight
}

// reset all weight that can vote on a voted layer.
func (v *verifying) resetWeights(voted types.LayerID) { _ = "STUB: not implemented"; return }

func (v *verifying) countBallot(logger *zap.Logger, ballot *ballotInfo) {
	_ = "STUB: not implemented"
	return
}

func (v *verifying) countVotes(logger *zap.Logger, ballots []*ballotInfo) {
	_ = "STUB: not implemented"
	return
}

func (v *verifying) verify(logger *zap.Logger, lid types.LayerID) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// GreaterThan(zero) returns true even if value with negative sign
