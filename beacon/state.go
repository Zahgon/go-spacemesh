package beacon

import (
	"math/big"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type minerInfo struct {
	atxid     types.ATXID
	malicious bool
}

// state does the data management for epoch specific data for the protocol.
// Not thread-safe. It relies on ProtocolDriver's thread-safety mechanism.
type state struct {
	logger      *zap.Logger
	active      map[types.NodeID]participant
	epochWeight uint64
	// the original proposals as received, bucketed by validity.
	incomingProposals proposals
	// minerPublicKey -> list of proposal.
	// this list is used in encoding/decoding votes for each miner in all subsequent voting rounds.
	firstRoundIncomingVotes map[types.NodeID]proposalList
	// TODO(nkryuchkov): For every round excluding first round consider having a vector of opinions.
	votesMargin               map[Proposal]*big.Int
	hasProposed               map[types.NodeID]struct{}
	hasVoted                  map[types.NodeID]*votesTracker
	proposalPhaseFinishedTime time.Time
	proposalChecker           eligibilityChecker
	minerAtxs                 map[types.NodeID]*minerInfo
}

func newState(
	logger *zap.Logger,
	active map[types.NodeID]participant,
	epochWeight uint64,
	miners map[types.NodeID]*minerInfo,
	checker eligibilityChecker,
) *state {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) addValidProposal(proposal Proposal) { _ = "STUB: not implemented"; return }

func (s *state) addPotentiallyValidProposal(proposal Proposal) { _ = "STUB: not implemented"; return }

func (s *state) setMinerFirstRoundVote(nodeID types.NodeID, voteList []Proposal) {
	_ = "STUB: not implemented"
	return
}

func (s *state) getMinerFirstRoundVote(nodeID types.NodeID) (proposalList, error) {
	_ = "STUB: not implemented"
	return *new(proposalList), nil
}

func (s *state) addVote(proposal Proposal, vote uint, voteWeight *big.Int) {
	_ = "STUB: not implemented"
	return
}

// voteMargin is updated during the proposal phase.
// ignore votes on proposals not in the original proposals.

func (s *state) registerProposed(nodeID types.NodeID) error { _ = "STUB: not implemented"; return nil }

func (s *state) registerVoted(nodeID types.NodeID, round types.RoundID) error {
	_ = "STUB: not implemented"
	return nil
}
