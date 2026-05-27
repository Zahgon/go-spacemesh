package wire

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate scalegen

// ProofDoubleMerge is a proof that two distinct ATXs published in the same epoch
// contain the same marriage ATX.
//
// We are proving the following:
//  1. The ATXs have different IDs.
//  2. Both ATXs have a valid signature.
//  3. Both ATXs contain the same marriage ATX.
//  4. Both ATXs were published in the same epoch.
//  5. Signers of both ATXs are married - to prevent banning others by
//     publishing an ATX with the same marriage ATX.
type ProofDoubleMerge struct {
	// PublishEpoch and its proof that it is contained in the ATX.
	PublishEpoch types.EpochID

	// MarriageATXID is the ID of the marriage ATX.
	MarriageATX types.ATXID
	// MarriageATXSmesherID is the ID of the smesher that published the marriage ATX.
	MarriageATXSmesherID types.NodeID

	// ATXID1 is the ID of the ATX being proven.
	ATXID1 types.ATXID
	// SmesherID1 is the ID of the smesher that published the ATX.
	SmesherID1 types.NodeID
	// Signature1 is the signature of the ATXID by the smesher.
	Signature1 types.EdSignature
	// PublishEpochProof1 is the proof that the publish epoch is contained in the ATX.
	PublishEpochProof1 PublishEpochProof `scale:"max=32"`
	// MarriageATXProof1 is the proof that MarriageATX is contained in the ATX.
	MarriageATXProof1 MarriageATXProof `scale:"max=32"`
	// SmesherID1MarryProof is the proof that they married in MarriageATX.
	SmesherID1MarryProof MarryProof

	// ATXID2 is the ID of the ATX being proven.
	ATXID2 types.ATXID
	// SmesherID is the ID of the smesher that published the ATX.
	SmesherID2 types.NodeID
	// Signature2 is the signature of the ATXID by the smesher.
	Signature2 types.EdSignature
	// PublishEpochProof2 is the proof that the publish epoch is contained in the ATX.
	PublishEpochProof2 PublishEpochProof `scale:"max=32"`
	// MarriageATXProof1 is the proof that MarriageATX is contained in the ATX.
	MarriageATXProof2 MarriageATXProof `scale:"max=32"`
	// SmesherID1MarryProof is the proof that they married in MarriageATX.
	SmesherID2MarryProof MarryProof
}

func (p ProofDoubleMerge) AllowNoRefATXs() bool { _ = "STUB: not implemented"; return false }

func (p ProofDoubleMerge) TypeName() string { _ = "STUB: not implemented"; return "" }

func (p ProofDoubleMerge) Type() ProofType { _ = "STUB: not implemented"; return *new(ProofType) }

func (p ProofDoubleMerge) Info() map[string]string { _ = "STUB: not implemented"; return nil }

var _ Proof = &ProofDoubleMerge{}

func NewDoubleMergeProof(db sql.Executor, atx1, atx2 *ActivationTxV2) (*ProofDoubleMerge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ProofDoubleMerge) Valid(_ context.Context, malValidator MalfeasanceValidator) (types.NodeID, error) {
	_ = "STUB: not implemented"
	// 1. The ATXs have different IDs.
	return *new(types.NodeID), nil
}

// 2. Both ATXs have a valid signature.

// 3. and 4. publish epoch is contained in the ATXs

// 5. signers are married

// 6. smeshers have published valid ATXs before
