package hare4

import (
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
)

type Round uint8

var roundNames = [...]string{"preround", "hardlock", "softlock", "propose", "wait1", "wait2", "commit", "notify"}

func (r Round) String() string { _ = "STUB: not implemented"; return "" }

// NOTE(dshulyak) changes in order is a breaking change.
const (
	preround Round = iota
	hardlock
	softlock
	propose
	wait1
	wait2
	commit
	notify
)

//go:generate scalegen

type IterRound struct {
	Iter  uint8
	Round Round
}

// Delay returns number of network delays since specified iterround.
func (ir IterRound) Delay(since IterRound) uint32 { _ = "STUB: not implemented"; return 0 }

// we skip hardlock round in 0th iteration.

func (ir IterRound) Grade(since IterRound) grade { _ = "STUB: not implemented"; return *new(grade) }

func (ir IterRound) IsMessageRound() bool { _ = "STUB: not implemented"; return false }

func (ir IterRound) Absolute() uint32 { _ = "STUB: not implemented"; return 0 }

type Value struct {
	// Proposals is set in messages for preround and propose rounds.
	//
	// Worst case scenario is that a single smesher identity has > 99.97% of the total weight of the network.
	// In this case they will get all 50 available slots in all 4032 layers of the epoch.
	// Additionally every other identity on the network that successfully published an ATX will get 1 slot.
	//
	// If we expect 7.0 Mio ATXs that would be a total of 7.0 Mio + 50 * 4032 = 8 201 600 slots.
	// Since these are randomly distributed across the epoch, we can expect an average of n * p =
	// 8 201 600 / 4032 = 2034.1 eligibilities in a layer with a standard deviation of sqrt(n * p * (1 - p)) =
	// sqrt(8 201 600 * 1/4032 * 4031/4032) = 45.1
	//
	// This means that we can expect a maximum of 2034.1 + 6*45.1 = 2304.7 eligibilities in a layer with
	// > 99.9997% probability.
	Proposals []types.ProposalID `scale:"max=2350"`
	// Reference is set in messages for commit and notify rounds.
	Reference *types.Hash32
	// CompactProposals is the array of compacted proposals IDs which are represented as truncated
	// eligibility hashes.
	CompactProposals []types.CompactProposalID `scale:"max=2350"`
}

type Body struct {
	Layer types.LayerID
	IterRound
	Value       Value
	Eligibility types.HareEligibility
}

type Message struct {
	Body
	Sender    types.NodeID
	Signature types.EdSignature
}

func (m *Message) ToHash() types.Hash32 { _ = "STUB: not implemented"; return *new(types.Hash32) }

func (m *Message) ToMetadata() wire.HareMetadata {
	_ = "STUB: not implemented"
	return *new(wire.HareMetadata)
}

func (m *Message) ToMalfeasanceProof() wire.HareProofMsg {
	_ = "STUB: not implemented"
	return *new(wire.HareProofMsg)
}

func (m *Message) key() messageKey { _ = "STUB: not implemented"; return *new(messageKey) }

func (m *Message) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) Validate() error { _ = "STUB: not implemented"; return nil }

func (m *Message) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type CompactIdRequest struct {
	MsgId types.Hash32
}

type CompactIdResponse struct {
	Ids []types.ProposalID `scale:"max=2050"`
}
