package tortoise

import (
	"math/big"

	"github.com/spacemeshos/fixed"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
)

type (
	weight = fixed.Fixed

	verifyingInfo struct {
		// goodUncounted is a weight that doesn't vote for this layer
		//
		// for example ballot created in layer 10 doesn't vote for layers
		// 10 and above, therefore its weight needs to be added to goodUncounted
		// for layers 10 and above
		goodUncounted   weight
		referenceHeight uint64
	}

	epochInfo struct {
		// weight is a sum of all atxs
		weight weight
		// median height from atxs
		height uint64
		beacon *types.Beacon
	}

	state struct {
		// last received layer
		// TODO should be last layer according to the clock
		// https://github.com/spacemeshos/go-spacemesh/issues/2921
		last types.LayerID
		// localThreshold is updated together with the last layer.
		localThreshold weight

		// last verified layer
		verified types.LayerID
		// last processed layer
		processed types.LayerID
		// last evicted layer
		evicted types.LayerID

		atxsdata *atxsdata.Data
		epochs   map[types.EpochID]*epochInfo
		layers   layerSlice
		// ballots should not be referenced by other ballots
		// each ballot stores references (votes) for X previous layers
		// those X layers may reference another set of ballots that will
		// reference recursively more layers with another set of ballots
		ballots map[types.LayerID][]*ballotInfo

		// to efficiently find base and reference ballots
		ballotRefs map[types.BallotID]*ballotInfo

		// malnodes is a collection with all nodes that equivocated in history.
		// each node id is 32 bytes. 100 000 of such nodes is only about ~3MB
		malnodes map[types.NodeID]struct{}
	}
)

func newState(atxdata *atxsdata.Data) *state { _ = "STUB: not implemented"; return nil }

func (s *state) globalThreshold(cfg Config, target types.LayerID) weight {
	_ = "STUB: not implemented"
	return *new(weight)
}

func (s *state) expectedWeight(cfg Config, target types.LayerID) weight {
	_ = "STUB: not implemented"
	return *new(weight)
}

func (s *state) layer(lid types.LayerID) *layerInfo { _ = "STUB: not implemented"; return nil }

func (s *state) epoch(eid types.EpochID) *epochInfo { _ = "STUB: not implemented"; return nil }

func (s *state) addBallot(ballot *ballotInfo) { _ = "STUB: not implemented"; return }

func (s *state) addBlock(block *blockInfo) { _ = "STUB: not implemented"; return }

func (s *state) getBlock(header types.Vote) *blockInfo { _ = "STUB: not implemented"; return nil }

func (s *state) findRefHeightBelow(lid types.LayerID) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *state) updateRefHeight(layer *layerInfo, block *blockInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *state) isMalfeasant(id types.NodeID) bool { _ = "STUB: not implemented"; return false }

func (s *state) markMalfeasant(id types.NodeID) bool { _ = "STUB: not implemented"; return false }

type layerInfo struct {
	lid            types.LayerID
	empty          weight
	hareTerminated bool
	blocks         []*blockInfo
	verifying      verifyingInfo
	coinflip       sign

	// unique opinions recorded from the ballots in this layer.
	// ballot votes an opinion and encodes sidecar
	opinions map[types.Hash32]votes

	opinion types.Hash32
	// a pointer to the value stored on the previous layerInfo object
	// it is stored as a pointer so that when previous layerInfo is evicted
	// we still have access in case we need to recompute opinion for this layer
	prevOpinion *types.Hash32
}

func (l *layerInfo) computeOpinion(hdist uint32, last types.LayerID) {
	_ = "STUB: not implemented"
	return
}

type (
	conditions struct {
		// set after comparing with local beacon
		badBeacon bool
	}

	referenceInfo struct {
		smesher         types.NodeID
		atxid           types.ATXID
		expectedBallots uint32
		beacon          types.Beacon
		weight          *big.Rat
		height          uint64
	}

	ballotInfo struct {
		id         types.BallotID
		layer      types.LayerID
		malicious  bool
		weight     weight
		reference  *referenceInfo
		votes      votes
		conditions conditions
	}
)

func (b *ballotInfo) opinion() types.Hash32 { _ = "STUB: not implemented"; return *new(types.Hash32) }

func (b *ballotInfo) overwriteOpinion(opinion types.Hash32) { _ = "STUB: not implemented"; return }

type votes struct {
	tail *layerVote
}

func (v *votes) append(lv *layerVote) { _ = "STUB: not implemented"; return }

func (v *votes) update(from types.LayerID, diff map[types.LayerID]map[types.BlockID]headerWithSign) (votes, error) {
	_ = "STUB: not implemented"
	return *new(votes), nil
}

// cutBefore cuts all pointers to votes before the specified layer.
func (v *votes) cutBefore(lid types.LayerID) { _ = "STUB: not implemented"; return }

func (v *votes) opinion() types.Hash32 { _ = "STUB: not implemented"; return *new(types.Hash32) }

type layerVote struct {
	lid       types.LayerID
	opinion   types.Hash32
	vote      sign
	supported []*blockInfo

	prev *layerVote
}

func (l *layerVote) getVote(binfo *blockInfo) sign { _ = "STUB: not implemented"; return *new(sign) }

func (l *layerVote) copy() *layerVote { _ = "STUB: not implemented"; return nil }

func (l *layerVote) append(lv *layerVote) *layerVote { _ = "STUB: not implemented"; return nil }

func (l *layerVote) update(
	from types.LayerID,
	diff map[types.LayerID]map[types.BlockID]headerWithSign,
) (*layerVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *layerVote) sortSupported() { _ = "STUB: not implemented"; return }

func (l *layerVote) computeOpinion() { _ = "STUB: not implemented"; return }

func sortBlocks(blocks []*blockInfo) { _ = "STUB: not implemented"; return }

func newBlockInfo(header types.Vote) *blockInfo { _ = "STUB: not implemented"; return nil }

type blockInfo struct {
	id     types.BlockID
	layer  types.LayerID
	height uint64
	hare   sign
	margin weight

	validity sign

	data bool // set to true if block is available locally
}

func (b *blockInfo) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *blockInfo) header() types.Vote { _ = "STUB: not implemented"; return *new(types.Vote) }

type headerWithSign struct {
	header types.Vote
	sign   sign
}

func decodeVotes(_, blid types.LayerID, base *ballotInfo, exceptions types.Votes) (votes, types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(votes), *new(types.LayerID), nil
}

// FIXME(dshulyak) this needs to be ignored when recovering from disk
// if from <= evicted {
// 	return votes{}, 0, fmt.Errorf("votes for a block in the layer (%d) outside the window (evicted %d)",
// 	 from, evicted,
// 	)
// }

// inherit opinion from the base ballot by copying votes

// add new opinions after the base layer

type layerSlice struct {
	data []*layerInfo
}

func (s *layerSlice) get(offset, index types.LayerID) *layerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *layerSlice) pop() { _ = "STUB: not implemented"; return }
