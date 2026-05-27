package wire

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate scalegen

// ProofInvalidPrevAtxV2 is a proof that two distinct ATXs reference the same previous ATX for one of the included
// identities.
//
// We are proving the following:
//  1. The ATXs have different IDs.
//  2. Both ATXs have a valid signature.
//  3. Both ATXs reference the same previous ATX for the same identity.
//  4. If the signer of one of the two ATXs is not the identity that referenced the same previous ATX, then the identity
//     that did is married to the signer via a valid marriage certificate in the referenced marriage ATX.
type ProofInvalidPrevAtxV2 struct {
	// NodeID is the node ID that referenced the same previous ATX twice.
	NodeID types.NodeID

	// PrevATXID is the ATX that was referenced twice.
	PrevATXID types.ATXID

	Proofs [2]InvalidPrevAtxProof
}

func (p ProofInvalidPrevAtxV2) AllowNoRefATXs() bool { _ = "STUB: not implemented"; return false }

func (p ProofInvalidPrevAtxV2) TypeName() string { _ = "STUB: not implemented"; return "" }

func (p ProofInvalidPrevAtxV2) Type() ProofType { _ = "STUB: not implemented"; return *new(ProofType) }

func (p ProofInvalidPrevAtxV2) Info() map[string]string { _ = "STUB: not implemented"; return nil }

var _ Proof = &ProofInvalidPrevAtxV2{}

func NewInvalidPrevAtxProofV2(
	db sql.Executor,
	atx1, atx2 *ActivationTxV2,
	nodeID types.NodeID,
) (*ProofInvalidPrevAtxV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createInvalidPrevAtxProof(
	atx *ActivationTxV2,
	prevATX types.ATXID,
	nipostIndex,
	postIndex int,
	marriageProof *MarriageProof,
) (InvalidPrevAtxProof, error) {
	_ = "STUB: not implemented"
	return *new(InvalidPrevAtxProof), nil
}

func (p ProofInvalidPrevAtxV2) Valid(_ context.Context, malValidator MalfeasanceValidator) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

// ProofInvalidPrevAtxV1 is a proof that two ATXs published by an identity reference the same previous ATX for an
// identity.
//
// We are proving the following:
//  1. Both ATXs have a valid signature.
//  2. Both ATXs reference the same previous ATX for the same identity.
//  3. If the signer of the ATXv2 is not the identity that referenced the same previous ATX, then the included marriage
//     proof is valid.
//  4. The ATXv1 has been signed by the identity that referenced the same previous ATX.
type ProofInvalidPrevAtxV1 struct {
	// NodeID is the node ID that referenced the same previous ATX twice.
	NodeID types.NodeID

	// PrevATXID is the ATX that was referenced twice.
	PrevATXID types.ATXID

	Proof InvalidPrevAtxProof
	ATXv1 ActivationTxV1
}

func (p ProofInvalidPrevAtxV1) AllowNoRefATXs() bool { _ = "STUB: not implemented"; return false }

func (p ProofInvalidPrevAtxV1) TypeName() string { _ = "STUB: not implemented"; return "" }

func (p ProofInvalidPrevAtxV1) Type() ProofType { _ = "STUB: not implemented"; return *new(ProofType) }

func (p ProofInvalidPrevAtxV1) Info() map[string]string { _ = "STUB: not implemented"; return nil }

var _ Proof = &ProofInvalidPrevAtxV1{}

func NewInvalidPrevAtxProofV1(
	db sql.Executor,
	atx1 *ActivationTxV2,
	atx2 *ActivationTxV1,
	nodeID types.NodeID,
) (*ProofInvalidPrevAtxV1, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p ProofInvalidPrevAtxV1) Valid(_ context.Context, malValidator MalfeasanceValidator) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

type InvalidPrevAtxProof struct {
	// ATXID is the ID of the ATX being proven.
	ATXID types.ATXID
	// SmesherID is the ID of the smesher that published the ATX.
	SmesherID types.NodeID
	// Signature is the signature of the ATXID by the smesher.
	Signature types.EdSignature

	// NIPostsRoot and its proof that it is contained in the ATX.
	NIPostsRoot      NIPostsRoot
	NIPostsRootProof NIPostsRootProof `scale:"max=32"`

	// NIPostRoot and its proof that it is contained at the given index in the NIPostsRoot.
	NIPostRoot      NIPostRoot
	NIPostRootProof NIPostRootProof `scale:"max=32"`
	NIPostIndex     uint16

	// SubPostsRoot and its proof that it is contained in the NIPostRoot.
	SubPostsRoot      SubPostsRoot
	SubPostsRootProof SubPostsRootProof `scale:"max=32"`

	// SubPostRoot and its proof that is contained at the given index in the SubPostsRoot.
	SubPostRoot      SubPostRoot
	SubPostRootProof SubPostRootProof `scale:"max=32"`
	SubPostRootIndex uint16

	// MarriageProof is the proof that NodeID and SmesherID are married. It is nil if NodeID == SmesherID.
	MarriageProof *MarriageProof

	// MarriageIndexProof is the proof that the MarriageIndex (CertificateIndex from NodeIDMarryProof) is contained in
	// the SubPostRoot.
	MarriageIndexProof MarriageIndexProof `scale:"max=32"`

	// PrevATXProof is the proof that the previous ATX is contained in the SubPostRoot.
	PrevATXProof PrevATXProof `scale:"max=32"`
}

func (p InvalidPrevAtxProof) Valid(prevATX types.ATXID, nodeID types.NodeID, malValidator MalfeasanceValidator) error {
	_ = "STUB: not implemented"
	return nil
}
