package wire

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate scalegen

// ProofInvalidPost is a proof that a merged ATX with an invalid Post was published by a smesher.
//
// We are proofing the following:
//  1. The ATX has a valid signature.
//  2. If NodeID is different from SmesherID, we prove that NodeID and SmesherID are married.
//  3. The commitment ATX of NodeID used for the invalid PoST based on their initial ATX.
//  4. The provided Post is invalid for the given NodeID.
type ProofInvalidPost struct {
	// ATXID is the ID of the ATX containing the invalid PoST.
	ATXID types.ATXID
	// SmesherID is the ID of the smesher that published the ATX.
	SmesherID types.NodeID
	// Signature is the signature of the ATXID by the smesher.
	Signature types.EdSignature

	// NodeID is the node ID that created the invalid PoST.
	NodeID types.NodeID

	// MarriageProof is the proof that NodeID and SmesherID are married. It is nil if NodeID == SmesherID.
	MarriageProof *MarriageProof
	// InvalidPostProof is the proof for the invalid PoST of the ATX. It contains the PoST and the merkle proofs to
	// verify the PoST.
	InvalidPostProof InvalidPostProof
}

func (p ProofInvalidPost) AllowNoRefATXs() bool { _ = "STUB: not implemented"; return false }

func (p ProofInvalidPost) TypeName() string { _ = "STUB: not implemented"; return "" }

func (p ProofInvalidPost) Type() ProofType { _ = "STUB: not implemented"; return *new(ProofType) }

func (p ProofInvalidPost) Info() map[string]string { _ = "STUB: not implemented"; return nil }

var _ Proof = &ProofInvalidPost{}

func NewInvalidPostProof(
	db sql.Executor,
	atx *ActivationTxV2,
	commitmentATX types.ATXID,
	nodeID types.NodeID,
	nipostIndex int,
	invalidPostIndex uint32,
	validPostIndex uint32,
) (*ProofInvalidPost, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p ProofInvalidPost) Valid(ctx context.Context, malValidator MalfeasanceValidator) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

// InvalidPostProof is a proof for an invalid PoST in an ATX. It contains the PoST and the merkle proofs to verify the
// PoST.
//
// It contains both a valid and an invalid PoST index. This is required to proof that the commitment ATX was used to
// initialize the data for the invalid PoST. If a PoST contains no valid indices, then the ATX is syntactically invalid.
type InvalidPostProof struct {
	// NIPostsRoot and its proof that it is contained in the ATX.
	NIPostsRoot      NIPostsRoot
	NIPostsRootProof NIPostsRootProof `scale:"max=32"`

	// NIPostRoot and its proof that it is contained at the given index in the NIPostsRoot.
	NIPostRoot      NIPostRoot
	NIPostRootProof NIPostRootProof `scale:"max=32"`
	NIPostIndex     uint16

	// Challenge and its proof that it is contained in the NIPostRoot.
	Challenge      types.Hash32
	ChallengeProof ChallengeProof `scale:"max=32"`

	// SubPostsRoot and its proof that it is contained in the NIPostRoot.
	SubPostsRoot      SubPostsRoot
	SubPostsRootProof SubPostsRootProof `scale:"max=32"`

	// SubPostRoot and its proof that is contained at the given index in the SubPostsRoot.
	SubPostRoot      SubPostRoot
	SubPostRootProof SubPostRootProof `scale:"max=32"`
	SubPostRootIndex uint16

	// MarriageIndexProof is the proof that the MarriageIndex (CertificateIndex from NodeIDMarryProof) is contained in
	// the SubPostRoot.
	MarriageIndexProof MarriageIndexProof `scale:"max=32"`

	// Post is the invalid PoST and its proof that it is contained in the SubPostRoot.
	Post      PostV1
	PostProof PostRootProof `scale:"max=32"`

	// NumUnits and its proof that it is contained in the SubPostRoot.
	NumUnits      uint32
	NumUnitsProof NumUnitsProof `scale:"max=32"`

	// CommitmentATX is the ATX that was used to initialize data for the invalid PoST.
	CommitmentATX types.ATXID

	// InvalidPostIndex is the index of the leaf that was identified to be invalid.
	InvalidPostIndex uint32

	// ValidPostIndex is the index of a leaf that was identified to be valid.
	ValidPostIndex uint32
}

func createInvalidPostProof(
	atx *ActivationTxV2,
	commitmentATX types.ATXID,
	nipostIndex,
	postIndex int,
	invalidPostIndex uint32,
	validPostIndex uint32,
) (InvalidPostProof, error) {
	_ = "STUB: not implemented"
	return *new(InvalidPostProof), nil
}

// Valid returns no error if the proof is valid. It verifies that the signature is valid, that the merkle proofs are
// and that the provided post is invalid.
func (p InvalidPostProof) Valid(
	ctx context.Context,
	malValidator MalfeasanceValidator,
	atxID types.ATXID,
	nodeID types.NodeID,
	marriageIndex *uint32,
) error {
	_ = "STUB: not implemented"
	return nil
}
