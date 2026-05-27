package types

import (
	"github.com/spacemeshos/go-scale"
	"go.uber.org/zap/zapcore"
)

const (
	// ProposalIDSize in bytes.
	// FIXME(dshulyak) why do we cast to hash32 when returning bytes?
	// Probably required for fetching by hash between peers.
	ProposalIDSize = Hash32Length
)

//go:generate scalegen

// ProposalID is a 20-byte blake3 sum of the serialized ballot used to identify a Proposal.
type ProposalID Hash20

// EmptyProposalID is a canonical empty ProposalID.
var EmptyProposalID = ProposalID{}

type CompactProposalID [4]byte

// EncodeScale implements scale codec interface.
func (id *CompactProposalID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (id *CompactProposalID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EncodeScale implements scale codec interface.
func (id *ProposalID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (id *ProposalID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Proposal contains the smesher's signed content proposal for a given layer and vote on the mesh history.
// Proposal is ephemeral and will be discarded after the unified content block is created. The Ballot within
// the Proposal will remain in the mesh.
type Proposal struct {
	// the content proposal for a given layer and the votes on the mesh history
	InnerProposal
	// the smesher's signature on the InnerProposal
	Signature EdSignature

	// the following fields are kept private and from being serialized
	proposalID ProposalID
	// beacon that this proposal references either via its ballot's EpochData or RefBallot
	beacon Beacon
}

func (p Proposal) Equal(other Proposal) bool { _ = "STUB: not implemented"; return false }

// InnerProposal contains a smesher's content proposal for layer and its votes on the mesh history.
// This structure is serialized and signed to produce the signature in Proposal.
type InnerProposal struct {
	// smesher's votes on the mesh history
	Ballot
	// smesher's content proposal for a layer
	TxIDs []TransactionID `scale:"max=100000"`
	// aggregated hash up to the layer before this proposal.
	MeshHash Hash32
	// TODO add this when a state commitment mechanism is implemented.
	// state root up to the layer before this proposal.
	// note: this is needed in addition to mesh hash to detect bug in SVM
	// StateHash Hash32
}

// Initialize calculates and sets the Proposal's cached proposalID.
// This should be called once all the other fields of the Proposal are set.
func (p *Proposal) Initialize() error { _ = "STUB: not implemented"; return nil }

func (p *Proposal) MustInitialize() { _ = "STUB: not implemented"; return }

// SignedBytes returns the serialization of the InnerProposal.
func (p *Proposal) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

// HashInnerProposal returns the hash of the InnerProposal.
func (p *Proposal) HashInnerProposal() []byte { _ = "STUB: not implemented"; return nil }

// ID returns the ProposalID.
func (p *Proposal) ID() ProposalID {
	_ = "STUB: not implemented"
	return *

	// SetID set the ProposalID.
	new(ProposalID)
}

func (p *Proposal) SetID(pid ProposalID) { _ = "STUB: not implemented"; return }

func (p *Proposal) Beacon() Beacon { _ = "STUB: not implemented"; return *new(Beacon) }

func (p *Proposal) SetBeacon(beacon Beacon) {
	_ = "STUB: not implemented"

	// MarshalLogObject implements logging interface.
	return
}

func (p *Proposal) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns a short prefix of the hex representation of the ID.
func (id ProposalID) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the ProposalID as a byte slice.
func (id ProposalID) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// AsHash32 returns a Hash32 whose first 20 bytes are the bytes of this ProposalID, it is right-padded with zeros.
func (id ProposalID) AsHash32() Hash32 { _ = "STUB: not implemented"; return *new(Hash32) }

// Compare returns true if other (the given ProposalID) is less than this ProposalID, by lexicographic comparison.
func (id ProposalID) Compare(other ProposalID) bool { _ = "STUB: not implemented"; return false }

// ToProposalIDs returns a slice of ProposalID corresponding to the given proposals.
func ToProposalIDs(proposals []*Proposal) []ProposalID { _ = "STUB: not implemented"; return nil }

// SortProposalIDs sorts a list of ProposalID in lexicographic order, in-place.
func SortProposalIDs(ids []ProposalID) []ProposalID { _ = "STUB: not implemented"; return nil }

// ProposalIDsToHashes turns a list of ProposalID into their Hash32 representation.
func ProposalIDsToHashes(ids []ProposalID) []Hash32 { _ = "STUB: not implemented"; return nil }
