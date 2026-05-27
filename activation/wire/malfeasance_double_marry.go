package wire

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate scalegen

// ProofDoubleMarry is a proof that two distinct ATXs contain a marriage certificate signed by the same identity.
//
// We are proving the following:
//  1. The ATXs have different IDs.
//  2. Both ATXs have a valid signature.
//  3. Both ATXs contain a marriage certificate created by the same identity.
//  4. Both marriage certificates have valid signatures.
type ProofDoubleMarry struct {
	// NodeID is the node ID that married twice.
	NodeID types.NodeID

	// ATXID1 is the ID of the ATX being proven to have the marriage certificate of interest.
	ATXID1 types.ATXID
	// SmesherID1 is the ID of the smesher that published ATX1.
	SmesherID1 types.NodeID
	// Signature1 is the signature of the ATXID by the smesher.
	Signature1 types.EdSignature
	// Proof1 is the proof that the marriage certificate is contained in the ATX1.
	Proof1 MarryProof

	// ATXID2 is the ID of the ATX being proven to have the marriage certificate of interest.
	ATXID2 types.ATXID
	// SmesherID2 is the ID of the smesher that published ATX2.
	SmesherID2 types.NodeID
	// Signature2 is the signature of the ATXID by the smesher.
	Signature2 types.EdSignature
	// Proof2 is the proof that the marriage certificate is contained in the ATX2.
	Proof2 MarryProof
}

func (p ProofDoubleMarry) AllowNoRefATXs() bool { _ = "STUB: not implemented"; return false }

func (p ProofDoubleMarry) TypeName() string { _ = "STUB: not implemented"; return "" }

func (p ProofDoubleMarry) Type() ProofType { _ = "STUB: not implemented"; return *new(ProofType) }

func (p ProofDoubleMarry) Info() map[string]string { _ = "STUB: not implemented"; return nil }

var _ Proof = &ProofDoubleMarry{}

func NewDoubleMarryProof(db sql.Executor, atx1, atx2 *ActivationTxV2, nodeID types.NodeID) (*ProofDoubleMarry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p ProofDoubleMarry) Valid(_ context.Context, malValidator MalfeasanceValidator) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}
