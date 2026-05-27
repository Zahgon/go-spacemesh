package types

import (
	"github.com/spacemeshos/go-scale"
	"go.uber.org/zap/zapcore"
)

const (
	// BallotIDSize in bytes.
	// FIXME(dshulyak) why do we cast to hash32 when returning bytes?
	BallotIDSize = Hash32Length
)

//go:generate scalegen

// BallotID is a 20-byte blake3 sum of the serialized ballot used to identify a Ballot.
type BallotID Hash20

// EmptyBallotID is a canonical empty BallotID.
var EmptyBallotID = BallotID{}

// EncodeScale implements scale codec interface.
func (id *BallotID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (id *BallotID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (id *BallotID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *BallotID) UnmarshalText(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Ballot contains the smeshers signed vote on the mesh history.
type Ballot struct {
	InnerBallot
	// smeshers signature on InnerBallot
	Signature EdSignature
	// the public key of the smesher that produced this ballot.
	SmesherID NodeID
	// Votes field is not signed.
	Votes Votes
	// the proof of the smeshers eligibility to vote and propose block content in this epoch.
	// Eligibilities must be produced in the ascending order.
	// the proofs are vrf signatures and need not be included in the ballot's signature.
	//
	// The number of eligibility proofs depends on the smeshers weight and the total weight of the network.
	// For epoch 16 the largest smesher had 1 337 SUs and the total weight of the network was ~20.2 Mio SUs.
	// This means that the largest smesher received 1 337 / 26 481 043 SU = 0.005% of all eligibility slots for the
	// epoch.
	// There are 4032 layers in an epoch and 50 eligibility slots per layer, so the largest smesher received
	// 0.005% * 4032 * 50 = ~10 eligibility slots.
	//
	// Assuming the largest smesher won't control more than 10% of space in the network, we can assume that the
	// highest number of eligibilities in a single ballot will be below 25000. (10% of 4032 * 50 = 20160)
	EligibilityProofs []VotingEligibility `scale:"max=25000"`
	// from the smesher's view, the set of ATXs eligible to vote and propose block content in this epoch
	// only present in smesher's first ballot of the epoch
	// this field isn't actually used any more (replaced by InnerBallot.EpochData)
	// TODO (mafa): remove this field in Ballot v2
	ActiveSet []ATXID `scale:"max=1"`

	// the following fields are kept private and from being serialized
	ballotID BallotID
	// malicious is set to true if smesher that produced this ballot is known to be malicious.
	malicious bool
}

func (b Ballot) Equal(other Ballot) bool { _ = "STUB: not implemented"; return false }

// BallotMetadata is the signed part of Ballot.
type BallotMetadata struct {
	Layer   LayerID // the layer ID in which this ballot is eligible for. this will be validated via EligibilityProof
	MsgHash Hash32  // Hash of InnerBallot (returned by HashInnerBytes)
}

func (m *BallotMetadata) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// InnerBallot contains all info about a smeshers votes on the mesh history. This structure is
// serialized and signed to produce the signature in Ballot.
type InnerBallot struct {
	Layer LayerID
	// the smeshers ATX in the epoch this ballot is cast.
	AtxID ATXID
	// OpinionHash is a aggregated opinion on all previous layers.
	// It is included into transferred data explicitly, so that signature
	// can be verified before decoding votes.
	OpinionHash Hash32

	// the first Ballot the smesher cast in the epoch. this Ballot is a special Ballot that contains information
	// that cannot be changed mid-epoch.
	RefBallot BallotID
	EpochData *EpochData
}

// Votes is for encoding local votes to send over the wire.
//
// A smesher creates votes in the following steps:
// - select a Ballot in the past as a base Ballot
// - calculate the opinion difference on history between the smesher and the base Ballot
// - encode the opinion difference in 3 list:
//   - ForDiff
//     contains blocks we support while the base ballot did not support (i.e. voted against)
//     for blocks we support in layers later than the base ballot, we also add them to this list
//   - AgainstDiff
//     contains blocks we vote against while the base ballot explicitly supported
//   - NeutralDiff
//     contains layers we vote neutral while the base ballot explicitly supported or voted against
//
// example:
// layer | unified content block
// -----------------------------------------------------------------------------------------------
//
//	N   | UCB_A (genesis)
//
// -----------------------------------------------------------------------------------------------
//
//	N+1  | UCB_B base:UCB_A, for:[UCB_A], against:[], neutral:[]
//
// -----------------------------------------------------------------------------------------------
//
//	N+2  | UCB_C base:UCB_B, for:[UCB_B], against:[], neutral:[]
//
// -----------------------------------------------------------------------------------------------
//
//	(hare hasn't terminated for N+2)
//	N+3  | UCB_D base:UCB_B, for:[UCB_B], against:[], neutral:[N+2]
//
// -----------------------------------------------------------------------------------------------
//
//	(hare succeeded for N+2 but failed for N+3)
//	N+4  | UCB_E base:UCB_C, for:[UCB_C], against:[], neutral:[]
//
// -----------------------------------------------------------------------------------------------
// NOTE on neutral votes: a base block is by default neutral on all blocks and layers that come after it, so
// there's no need to explicitly add neutral votes for more recent layers.
//
// TODO: maybe collapse Support and Against into a single list.
//
//	see https://github.com/spacemeshos/go-spacemesh/issues/2369.
type Votes struct {
	// Base ballot.
	Base BallotID `json:"base"`
	// Support block id at a particular layer and height.
	// sliding vote window size is 10k layers, vote for one block per layer
	Support []Vote `json:"support,omitempty" scale:"max=10000"`
	// Against previously supported block.
	// sliding vote window size is 10k layers, vote for one block per layer
	Against []Vote `json:"against,omitempty" scale:"max=10000"`
	// Abstain on layers until they are terminated.
	// sliding vote window size is 10k layers, vote to abstain on any layer
	Abstain []LayerID `json:"abstain,omitempty" scale:"max=10000"`
}

// MarshalLogObject implements logging interface.
func (v *Votes) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type BlockHeader struct {
	ID      BlockID `json:"id"`
	LayerID LayerID `json:"lid"`
	Height  uint64  `json:"height"`
}

// MarshalLogObject implements logging interface.
func (header *BlockHeader) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Vote additionally carries layer id and height
// in order for the tortoise to count votes without downloading block body.
type Vote = BlockHeader

// Opinion is a tuple from opinion hash and votes that decode to opinion hash.
type Opinion struct {
	Hash  Hash32 `json:"hash"`
	Votes `       json:",inline"`
}

// MarshalLogObject implements logging interface.
func (o *Opinion) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// EpochData contains information that cannot be changed mid-epoch.
type EpochData struct {
	ActiveSetHash Hash32
	// the beacon value the smesher recorded for this epoch
	Beacon Beacon
	// total number of ballots the smesher is eligible in this epoch.
	EligibilityCount uint32
}

// Initialize calculates and sets the Ballot's cached ballotID and smesherID.
// This should be called once all the other fields of the Ballot are set.
func (b *Ballot) Initialize() error { _ = "STUB: not implemented"; return nil }

// SignedBytes returns the serialization of the BallotMetadata for signing.
func (b *Ballot) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

// HashInnerBytes returns the hash of the InnerBallot.
func (b *Ballot) HashInnerBytes() []byte { _ = "STUB: not implemented"; return nil }

// SetID from stored data.
func (b *Ballot) SetID(id BallotID) {
	_ = "STUB: not implemented"

	// ID returns the BallotID.
	return
}

func (b *Ballot) ID() BallotID {
	_ = "STUB: not implemented"

	// SetMalicious sets ballot as malicious.
	return *new(BallotID)
}

func (b *Ballot) SetMalicious() {
	_ = "STUB: not implemented"

	// IsMalicious returns true if ballot is malicious.
	return
}

func (b *Ballot) IsMalicious() bool {
	_ = "STUB: not implemented"

	// MarshalLogObject implements logging encoder for Ballot.
	return false
}

func (b *Ballot) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Ballot) ToTortoiseData() *BallotTortoiseData { _ = "STUB: not implemented"; return nil }

// ToBallotIDs turns a list of Ballot into a list of BallotID.
func ToBallotIDs(ballots []*Ballot) []BallotID { _ = "STUB: not implemented"; return nil }

// String returns a short prefix of the hex representation of the ID.
func (id BallotID) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the BallotID as a byte slice.
func (id BallotID) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// AsHash32 returns a Hash32 whose first 20 bytes are the bytes of this BallotID, it is right-padded with zeros.
func (id BallotID) AsHash32() Hash32 { _ = "STUB: not implemented"; return *new(Hash32) }

// Compare returns true if other (the given BallotID) is less than this BallotID, by lexicographic comparison.
func (id BallotID) Compare(other BallotID) bool { _ = "STUB: not implemented"; return false }

// BallotIDsToHashes turns a list of BallotID into their Hash32 representation.
func BallotIDsToHashes(ids []BallotID) []Hash32 { _ = "STUB: not implemented"; return nil }

// NewExistingBallot creates ballot from stored data.
func NewExistingBallot(id BallotID, sig EdSignature, nodeId NodeID, layer LayerID) Ballot {
	_ = "STUB: not implemented"
	return *new(Ballot)
}
