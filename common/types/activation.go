package types

import (
	"time"

	"github.com/spacemeshos/go-scale"
	"github.com/spacemeshos/post/shared"
	"go.uber.org/zap/zapcore"
)

//go:generate scalegen -types ATXMetadata,MerkleProof,EpochActiveSet

// BytesToATXID is a helper to copy buffer into a ATXID.
func BytesToATXID(buf []byte) (id ATXID) { _ = "STUB: not implemented"; return *new(ATXID) }

type Validity int

const (
	Unknown Validity = iota
	Valid
	Invalid
)

// ATXID is a 32 byte hash used to identify an activation transaction.
type ATXID Hash32

const (
	// ATXIDSize in bytes.
	ATXIDSize = Hash32Length
)

// String implements stringer interface.
func (t ATXID) String() string { _ = "STUB: not implemented"; return "" }

// ShortString returns the first few characters of the ID, for logging purposes.
func (t ATXID) ShortString() string { _ = "STUB: not implemented"; return "" }

// Hash32 returns the ATXID as a Hash32.
func (t ATXID) Hash32() Hash32 {
	_ = "STUB: not implemented"

	// Bytes returns the ATXID as a byte slice.
	return *new(Hash32)
}

func (t ATXID) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// EncodeScale implements scale codec interface.
func (t *ATXID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (t *ATXID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t ATXID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *ATXID) UnmarshalText(buf []byte) error { _ = "STUB: not implemented"; return nil }

// EmptyATXID is a canonical empty ATXID.
var EmptyATXID = ATXID{}

type ATXIDs []ATXID

// impl zap's ArrayMarshaler interface.
func (ids ATXIDs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// NIPostChallenge is the set of fields that's serialized, hashed and submitted to the PoET service
// to be included in the PoET membership proof.
type NIPostChallenge struct {
	PublishEpoch EpochID
	// Sequence number counts the number of ancestors of the ATX. It sequentially increases for each ATX in the chain.
	// Two ATXs with the same sequence number from the same miner can be used as the proof of malfeasance against
	// that miner.
	Sequence uint64
	// the previous ATX's ID (for all but the first in the sequence)
	PrevATXID      ATXID
	PositioningATX ATXID

	// CommitmentATX is the ATX used in the commitment for initializing the PoST of the node.
	CommitmentATX *ATXID
	InitialPost   *Post
}

func (c *NIPostChallenge) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// TargetEpoch returns the target epoch of the NIPostChallenge. This is the epoch in which the miner is eligible
// to participate thanks to the ATX.
func (challenge *NIPostChallenge) TargetEpoch() EpochID {
	_ = "STUB: not implemented"
	return *new(EpochID)
}

// ATXMetadata is the data of ActivationTx that is signed.
// It is also used for Malfeasance proofs.
type ATXMetadata struct {
	PublishEpoch EpochID
	MsgHash      Hash32 // Hash of InnerActivationTx (returned by HashInnerBytes)
}

func (m *ATXMetadata) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type AtxVersion uint

const (
	AtxV1   AtxVersion = 1
	AtxV2   AtxVersion = 2
	AtxVMAX AtxVersion = AtxV2
)

func (v AtxVersion) String() string { _ = "STUB: not implemented"; return "" }

type AtxBlob struct {
	Blob    []byte
	Version AtxVersion
}

// ActivationTx is a full, signed activation transaction. It includes (or references) everything a miner needs to prove
// they are eligible to actively participate in the Spacemesh protocol in the next epoch.
type ActivationTx struct {
	PublishEpoch EpochID
	// Sequence number counts the number of ancestors of the ATX. It sequentially increases for each ATX in the chain.
	// Two ATXs with the same sequence number from the same miner can be used as the proof of malfeasance against
	// that miner.
	Sequence uint64

	// CommitmentATX is the ATX used in the commitment for initializing the PoST of the node.
	CommitmentATX *ATXID
	// The marriage ATX, used in merged ATXs only.
	MarriageATX    *ATXID
	Coinbase       Address
	NumUnits       uint32 // the minimum number of space units in this and the previous ATX
	BaseTickHeight uint64
	TickCount      uint64
	VRFNonce       VRFPostIndex
	SmesherID      NodeID
	// Weight of the ATX. The total weight of the epoch is expected to fit in a uint64.
	// The total ATX weight is sum(NumUnits * TickCount) for identity it holds.
	// Space Units sizes are chosen such that NumUnits for all ATXs in an epoch is expected to be < 10^6.
	// PoETs should produce ~10k ticks at genesis, but are expected due to technological advances
	// to produce more over time. A uint64 should be large enough to hold the total weight of an epoch,
	// for at least the first few years.
	Weight uint64

	golden   bool
	id       ATXID     // non-exported cache of the ATXID
	received time.Time // time received by node, gossiped or synced
	validity Validity  // whether the chain is fully verified and OK
}

// TargetEpoch returns the target epoch of the ATX. This is the epoch in which the miner is eligible
// to participate thanks to the ATX.
func (atx *ActivationTx) TargetEpoch() EpochID { _ = "STUB: not implemented"; return *new(EpochID) }

// Golden returns true if atx is from a checkpoint snapshot.
// A golden ATX is not verifiable, and is only allowed to be prev atx or positioning atx.
func (atx *ActivationTx) Golden() bool {
	_ = "STUB: not implemented"

	// SetGolden set atx to golden.
	return false
}

func (atx *ActivationTx) SetGolden() {
	_ = "STUB: not implemented"

	// TickHeight returns a sum of base tick height and tick count.
	return
}

func (atx *ActivationTx) TickHeight() uint64 { _ = "STUB: not implemented"; return 0 }

// MarshalLogObject implements logging interface.
func (atx *ActivationTx) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// ShortString returns the first 5 characters of the ID, for logging purposes.
func (atx *ActivationTx) ShortString() string { _ = "STUB: not implemented"; return "" }

// ID returns the ATX's ID.
func (atx *ActivationTx) ID() ATXID {
	_ = "STUB: not implemented"

	// SetID sets the ATXID in this ATX's cache.
	return *new(ATXID)
}

func (atx *ActivationTx) SetID(id ATXID) { _ = "STUB: not implemented"; return }

func (atx *ActivationTx) SetReceived(received time.Time) { _ = "STUB: not implemented"; return }

func (atx *ActivationTx) Received() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (atx *ActivationTx) Validity() Validity { _ = "STUB: not implemented"; return *new(Validity) }

func (atx *ActivationTx) SetValidity(validity Validity) { _ = "STUB: not implemented"; return }

// Merkle proof proving that a given leaf is included in the root of merkle tree.
type MerkleProof struct {
	// Nodes on path from leaf to root (not including leaf)
	Nodes     []Hash32 `scale:"max=32"`
	LeafIndex uint64
}

type MultiMerkleProof struct {
	// Nodes on path from leaf to root (not including leaf)
	Nodes       []Hash32 `scale:"max=32"`
	LeafIndices []uint64
}

// NIPost is Non-Interactive Proof of Space-Time.
// Given an id, a space parameter S, a duration D and a challenge C,
// it can convince a verifier that (1) the prover expended S * D space-time
// after learning the challenge C. (2) the prover did not know the NIPost until D time
// after the prover learned C.
type NIPost struct {
	// Membership proves that the challenge for the PoET, which is
	// constructed from fields in the activation transaction,
	// is a member of the poet's proof.
	// Proof.Root must match the Poet's POSW statement.
	Membership MerkleProof

	// Post is the proof that the prover data is still stored (or was recomputed) at
	// the time he learned the challenge constructed from the PoET.
	Post *Post

	// PostMetadata is the Post metadata, associated with the proof.
	// The proof should be verified upon the metadata during the syntactic validation,
	// while the metadata should be verified during the contextual validation.
	PostMetadata *PostMetadata
}

// VRFPostIndex is the nonce generated using Pow during post initialization. It is used as a mitigation for
// grinding of identities for VRF eligibility.
type VRFPostIndex uint64

// Post is an alias to postShared.Proof.
type Post shared.Proof

func (p *Post) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// PostMetadata is similar postShared.ProofMetadata, but without the fields which can be derived elsewhere
// in a given ATX (eg. NodeID, NumUnits).
type PostMetadata struct {
	Challenge     []byte
	LabelsPerUnit uint64
}

func (m *PostMetadata) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// ToATXIDs returns a slice of ATXID corresponding to the given activation tx.
func ToATXIDs(atxs []*ActivationTx) []ATXID { _ = "STUB: not implemented"; return nil }

// ATXIDsToHashes turns a list of ATXID into their Hash32 representation.
func ATXIDsToHashes(ids []ATXID) []Hash32 { _ = "STUB: not implemented"; return nil }

type EpochActiveSet struct {
	Epoch EpochID
	Set   []ATXID `scale:"max=8000000"` // to be in line with `EpochData` in fetch/wire_types.go
}

var MaxEpochActiveSetSize = scale.MustGetMaxElements[EpochActiveSet]("Set")
