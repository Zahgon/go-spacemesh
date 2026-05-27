package tortoise

import (
	"container/list"
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
)

var (
	errBeaconUnavailable = errors.New("beacon unavailable")
	errDanglingBase      = errors.New("base ballot not in state")
	errEvictedBlocks     = errors.New("voted blocks were evicted")
	ErrBallotExists      = errors.New("tortoise: ballot exists")
)

type turtle struct {
	Config
	logger *zap.Logger

	*state

	// pending is a minimal layer where opinion has changed
	pending types.LayerID

	// a linked list with retriable ballots
	// the purpose is to add ballot to the state even
	// if beacon is not available locally, as tortoise
	// can't count ballots without knowing local beacon
	retriable list.List

	verifying *verifying

	isFull bool
	full   *full
}

// newTurtle creates a new verifying tortoise algorithm instance.
func newTurtle(logger *zap.Logger, config Config, atxdata *atxsdata.Data) *turtle {
	_ = "STUB: not implemented"
	return nil
}

func (t *turtle) evict() { _ = "STUB: not implemented"; return }

// EncodeVotes by choosing base ballot and explicit votes.
func (t *turtle) EncodeVotes(ctx context.Context, conf *encodeConf) (*types.Opinion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip them as they are candidates for pruning

// encode differences between selected base ballot and local votes.
func (t *turtle) encodeVotes(
	ctx context.Context,
	base *ballotInfo,
	start types.LayerID,
	current types.LayerID,
) (*types.Opinion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encode difference with local opinion between [start, base.layer)

// there is nothing to encode if hare didn't terminate
// and base ballot voted abstain on the previous layer

// ballot vote is consistent with local opinion, exception is not necessary

// encode votes after base ballot votes [base layer, last)

// getFullVote unlike getLocalVote will vote according to the counted votes on blocks that are
// outside of hdist. If opinion is undecided according to the votes it will use coinflip recorded
// in the current layer.
func (t *turtle) getFullVote(verified, current types.LayerID, block *blockInfo) (sign, voteReason, error) {
	_ = "STUB: not implemented"
	return *new(sign), *new(voteReason), nil
}

func (t *turtle) updateLast(last types.LayerID) { _ = "STUB: not implemented"; return }

func (t *turtle) tallyVotes(last types.LayerID) { _ = "STUB: not implemented"; return }

// NOTE(dshulyak) i need this when running verifying tortoise not from the genesis.

// terminate layer that falls out of the zdist window and wasn't terminated
// by any other component

func (t *turtle) switchModes() { _ = "STUB: not implemented"; return }

func (t *turtle) countBallot(ballot *ballotInfo) error { _ = "STUB: not implemented"; return nil }

func (t *turtle) verifyLayers() { _ = "STUB: not implemented"; return }

// count all votes if next layer after verified is outside hdist

func (t *turtle) runVerifying() (verified, changed types.LayerID) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), *new(types.LayerID)
}

func (t *turtle) runFull() (verified, changed types.LayerID) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), *new(types.LayerID)
}

func (t *turtle) computeEpochHeight(lid types.LayerID) { _ = "STUB: not implemented"; return }

func (t *turtle) onBlock(header types.BlockHeader, data, valid bool) {
	_ = "STUB: not implemented"
	return
}

// update existing state without calling t.addBlock

func (t *turtle) addBlock(binfo *blockInfo) { _ = "STUB: not implemented"; return }

func (t *turtle) onHareOutput(lid types.LayerID, bid types.BlockID) {
	_ = "STUB: not implemented"
	return
}

// we do not compute opinion because opinion hashing recursive.
// so if we didn't receive layer in order this opinion will be wrong
// and we also need to copy previous layer opinion into layer.prevOpinion

func (t *turtle) onOpinionChange(lid types.LayerID, early bool) { _ = "STUB: not implemented"; return }

func (t *turtle) onAtx(target types.EpochID, id types.ATXID, atx *atxsdata.ATX) {
	_ = "STUB: not implemented"
	return
}

func (t *turtle) decodeBallot(ballot *types.BallotTortoiseData) (*ballotInfo, types.LayerID, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.LayerID), nil
}

func (t *turtle) storeBallot(ballot *ballotInfo, offset types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *turtle) onRecoveredBallot(ballot *types.BallotTortoiseData) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *turtle) onBallot(ballot *types.BallotTortoiseData) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *turtle) compareBeacons(bid types.BallotID, lid types.LayerID, beacon types.Beacon) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *turtle) retryLater(ballot *ballotInfo) { _ = "STUB: not implemented"; return }

func (t *turtle) drainRetriable() error { _ = "STUB: not implemented"; return nil }

// if beacon is still unavailable - exit and wait for the next call
// to drain this queue

func (t *turtle) checkDrained() error { _ = "STUB: not implemented"; return nil }

func withinDistance(dist uint32, lid, last types.LayerID) bool {
	_ = "STUB: not implemented"
	// layer + distance > last
	return false
}

func getLocalVote(config Config, verified, last types.LayerID, block *blockInfo) (sign, voteReason) {
	_ = "STUB: not implemented"
	return *new(sign), *new(voteReason)
}

// if layer was verified, but then global threshold became unreachable we will not
// update validity, but verified variable will be lowered
