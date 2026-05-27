package beacon

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

type category uint8

const (
	valid            category = 1
	potentiallyValid category = 2
	invalid          category = 3
)

var (
	errVRFNotVerified         = errors.New("proposal failed vrf verification")
	errAlreadyProposed        = errors.New("already proposed")
	errAlreadyVoted           = errors.New("already voted")
	errMinerNotActive         = errors.New("miner ATX not found in previous epoch")
	errProtocolNotRunning     = errors.New("beacon protocol not running")
	errEpochNotActive         = errors.New("epoch not active")
	errMalformedMessage       = fmt.Errorf("%w: malformed msg", pubsub.ErrValidationReject)
	errUntimelyMessage        = errors.New("untimely msg")
	errBeaconProtocolInactive = errors.New("beacon protocol inactive")
)

// HandleWeakCoinProposal handles weakcoin proposal from gossip.
func (pd *ProtocolDriver) HandleWeakCoinProposal(ctx context.Context, peer p2p.Peer, msg []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// HandleProposal handles beacon proposal from gossip.
func (pd *ProtocolDriver) HandleProposal(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) classifyProposal(
	logger *zap.Logger,
	m ProposalMessage,
	atxReceived, receivedTime time.Time,
	checker eligibilityChecker,
) category {
	_ = "STUB: not implemented"
	return *new(category)
}

// partition the proposals into three sets:
//	let W1 be the weight at δ before the end of the previous epoch, with eligibility threshold X1
//	let W2 be the weight at the end of the previous epoch, with eligibility threshold X2
//  - valid:
//	  * the proposer's ATX was received before the end of the previous epoch
//	  * the proposal is received before the proposal phase ends.
//	  * the proposal is lower than the threshold X2
//	- potentially valid:
//	  * the proposer's ATX was received within δ after the end of the previous epoch.
//	  * the proposal is received within δ after the proposal phase ends.
//	  * the proposal is lower than the threshold X1
//  - invalid
//    the proposal is neither valid nor potentially valid
//
// note that honest users cannot disagree on timing by more than δ, so if a proposal is timely for
// any honest user, it cannot be late for any honest user (and vice versa).

func (pd *ProtocolDriver) addProposal(m ProposalMessage, cat category) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) verifyProposalMessage(logger *zap.Logger, m ProposalMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(nkryuchkov): attach telemetry

// HandleFirstVotes handles beacon first votes from gossip.
func (pd *ProtocolDriver) HandleFirstVotes(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// don't accept more first vote after the round ends

func (pd *ProtocolDriver) verifyFirstVotes(ctx context.Context, m FirstVotingMessage) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (pd *ProtocolDriver) storeFirstVotes(m FirstVotingMessage, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// this is used for bit vector calculation

// HandleFollowingVotes handles beacon following votes from gossip.
func (pd *ProtocolDriver) HandleFollowingVotes(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// don't accept votes from future rounds

func (pd *ProtocolDriver) verifyFollowingVotes(ctx context.Context, m FollowingVotingMessage) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (pd *ProtocolDriver) storeFollowingVotes(m FollowingVotingMessage, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) getProposalPhaseFinishedTime(epoch types.EpochID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// if this epoch doesn't exist, it is finished. returns something
// always in the past but different from time.Time{}

func (pd *ProtocolDriver) addToVoteMargin(epoch types.EpochID, thisRoundVotes allVotes, voteWeight *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) currentEpoch() types.EpochID {
	_ = "STUB: not implemented"
	return *new(types.EpochID)
}

func (pd *ProtocolDriver) currentRound() types.RoundID {
	_ = "STUB: not implemented"
	return *new(types.RoundID)
}

func (pd *ProtocolDriver) isProposalTimely(p *ProposalMessage, receivedTime time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (pd *ProtocolDriver) isVoteTimely(m *FollowingVotingMessage, receivedTime time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (pd *ProtocolDriver) registerProposed(epoch types.EpochID, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) registerVoted(epoch types.EpochID, nodeID types.NodeID, round types.RoundID) error {
	_ = "STUB: not implemented"
	return nil
}

func newVotesTracker() *votesTracker { _ = "STUB: not implemented"; return nil }

type votesTracker struct {
	votes *big.Int
}

func (v *votesTracker) register(round types.RoundID) bool { _ = "STUB: not implemented"; return false }

func (v *votesTracker) voted(round types.RoundID) bool { _ = "STUB: not implemented"; return false }
