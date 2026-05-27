package wire

import (
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

//go:generate scalegen

type ActivationTxV1 struct {
	InnerActivationTxV1

	SmesherID types.NodeID
	Signature types.EdSignature

	id   types.ATXID
	blob []byte
}

// InnerActivationTxV1 is the set of all fields of an ATX that are used to calculate the ATX ID. It contains all fields
// besides the Signature and the Smesher ID.
type InnerActivationTxV1 struct {
	NIPostChallengeV1

	Coinbase types.Address
	NumUnits uint32

	NIPost   *NIPostV1
	NodeID   *types.NodeID // only present in initial ATX to make hash of the first InnerActivationTxV2 unique
	VRFNonce *uint64       // only present when the nonce changed (or initial ATX)
}

type PostV1 struct {
	Nonce   uint32
	Indices []byte `scale:"max=800"` // up to K2=100
	Pow     uint64
}

type PostRoot types.Hash32

func (p *PostV1) Root() (result PostRoot) { _ = "STUB: not implemented"; return *new(PostRoot) }

type MerkleProofV1 struct {
	// Nodes on path from leaf to root (not including leaf)
	Nodes     []types.Hash32 `scale:"max=32"`
	LeafIndex uint64
}

type NIPostV1 struct {
	// Membership proves that the challenge for the PoET, which is
	// constructed from fields in the activation transaction,
	// is a member of the poet's proof.
	// Proof.Root must match the Poet's POSW statement.
	Membership MerkleProofV1

	// Post is the proof that the prover data is still stored (or was recomputed) at
	// the time he learned the challenge constructed from the PoET.
	Post *PostV1

	// PostMetadata is the Post metadata, associated with the proof.
	// The proof should be verified upon the metadata during the syntactic validation,
	// while the metadata should be verified during the contextual validation.
	PostMetadata *PostMetadataV1
}

type PostMetadataV1 struct {
	Challenge     []byte `scale:"max=32"`
	LabelsPerUnit uint64
}

type ATXMetadataV1 struct {
	Publish types.EpochID
	MsgHash types.Hash32
}

func (atx *ActivationTxV1) ID() types.ATXID { _ = "STUB: not implemented"; return *new(types.ATXID) }

func (atx *ActivationTxV1) SetID(id types.ATXID) { _ = "STUB: not implemented"; return }

func (atx *ActivationTxV1) Sign(signer *signing.EdSigner) { _ = "STUB: not implemented"; return }

func (atx *ActivationTxV1) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

func (atx *ActivationTxV1) Blob() types.AtxBlob {
	_ = "STUB: not implemented"
	return *new(types.AtxBlob)
}

func DecodeAtxV1(blob []byte) (*ActivationTxV1, error) { _ = "STUB: not implemented"; return nil, nil }

func (atx *ActivationTxV1) HashInnerBytes() (result types.Hash32) {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

func PostToWireV1(p *types.Post) *PostV1 { _ = "STUB: not implemented"; return nil }

func NiPostToWireV1(n *types.NIPost) *NIPostV1 { _ = "STUB: not implemented"; return nil }

func NIPostChallengeToWireV1(c *types.NIPostChallenge) *NIPostChallengeV1 {
	_ = "STUB: not implemented"
	return nil
}

func ActivationTxFromWireV1(atx *ActivationTxV1) *types.ActivationTx {
	_ = "STUB: not implemented"
	return nil
}

func NiPostFromWireV1(nipost *NIPostV1) *types.NIPost { _ = "STUB: not implemented"; return nil }

func PostFromWireV1(post *PostV1) *types.Post { _ = "STUB: not implemented"; return nil }

func (p *PostV1) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (nipost *NIPostV1) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (atx *ActivationTxV1) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
