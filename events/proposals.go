package events

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
)

// ProposalStatus is a type for proposal status.
type ProposalStatus int

func (s ProposalStatus) String() string { _ = "STUB: not implemented"; return "" }

const (
	// ProposalCreated is a status of the proposal that was created.
	ProposalCreated ProposalStatus = iota
	// ProposalIncluded is a status of the proposal when it is included into the block.
	ProposalIncluded
)

// EventProposal includes proposal and proposal status.
type EventProposal struct {
	Status   ProposalStatus
	Proposal *types.Proposal
}

// ReportProposal reports a proposal.
func ReportProposal(status ProposalStatus, proposal *types.Proposal) {
	_ = "STUB: not implemented"
	return
}

// TODO(dshulyak) check why it can error. i think it is reasonable to panic here

// SubscribeProposals subscribes to the proposals.
func SubscribeProposals() Subscription { _ = "STUB: not implemented"; return *new(Subscription) }
