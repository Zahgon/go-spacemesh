package tortoise

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// Recover tortoise state from database.
func Recover(
	ctx context.Context,
	db sql.Executor,
	atxdata *atxsdata.Data,
	current types.LayerID,
	opts ...Opt,
) (*Tortoise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we want to emulate the same condition as during genesis with one difference.
// genesis starts with zero opinion (aggregated hash) - see computeOpinion method.
// but in this case first processed layer should use non-zero opinion of the previous layer.

// we start tallying votes from the first layer of the epoch to guarantee that we load reference ballots.
// reference ballots track beacon and eligibilities

// tortoise will need reference to previous layer

// need to load the golden atxs after a checkpoint recovery

// load activations from future epochs that are not yet referenced by the ballots

// recoverEpoch expects target epoch

// find topmost layer that was already applied with same result
// and reset pending so that result for that layer is not returned

func recoverEpoch(target types.EpochID, trtl *Tortoise, db sql.Executor, atxdata *atxsdata.Data) error {
	_ = "STUB: not implemented"
	return nil
}

type ballotFunc func(*types.BallotTortoiseData)

func RecoverLayer(
	trtl *Tortoise,
	db sql.Executor,
	atxdata *atxsdata.Data,
	lid types.LayerID,
	onBallot ballotFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

// tortoise votes according to the hare only within hdist (protocol parameter).
// also node is free to prune certificates outside hdist to minimize space usage.

// NOTE(dshulyak) it is done in two steps so that if ballot from the same layer was used
// as reference or base ballot we will be able to decode it.
// it might be possible to invalidate such ballots, but until then this is required
