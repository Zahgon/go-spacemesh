package tortoise

import (
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

func newFullTortoise(config Config, state *state) *full { _ = "STUB: not implemented"; return nil }

type full struct {
	Config
	*state

	// counted weights up to this layer.
	//
	// counting votes is what makes full tortoise expensive during rerun.
	// we want to wait until verifying can't make progress before they are counted.
	// storing them in current version is cheap.
	counted types.LayerID
	// delayed ballots by the layer when they are safe to count
	delayed map[types.LayerID][]*ballotInfo
}

func (f *full) countBallot(logger *zap.Logger, ballot *ballotInfo) {
	_ = "STUB: not implemented"
	return
}

func (f *full) countForLateBlock(block *blockInfo) { _ = "STUB: not implemented"; return }

func (f *full) countDelayed(logger *zap.Logger, lid types.LayerID) {
	_ = "STUB: not implemented"
	return
}

func (f *full) countVotes(logger *zap.Logger) { _ = "STUB: not implemented"; return }

func (f *full) verify(logger *zap.Logger, lid types.LayerID) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (f *full) shouldBeDelayed(logger *zap.Logger, ballot *ballotInfo) bool {
	_ = "STUB: not implemented"
	return false
}
