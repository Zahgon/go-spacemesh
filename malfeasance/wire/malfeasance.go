package wire

import (
	"time"

	"github.com/spacemeshos/go-scale"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/activation/wire"
	"github.com/spacemeshos/go-spacemesh/common/types"
)

//go:generate scalegen -types MalfeasanceProof,MalfeasanceGossip,AtxProof,BallotProof,HareProof,AtxProofMsg,BallotProofMsg,HareProofMsg,HareMetadata,InvalidPostIndexProof,InvalidPrevATXProof

const (
	MultipleATXs byte = iota + 1
	MultipleBallots
	HareEquivocation
	InvalidPostIndex
	InvalidPrevATX
)

type MalfeasanceProof struct {
	// for network upgrade
	Layer types.LayerID
	Proof Proof

	received time.Time
}

func (mp *MalfeasanceProof) Received() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (mp *MalfeasanceProof) SetReceived(received time.Time) { _ = "STUB: not implemented"; return }

func (mp *MalfeasanceProof) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Proof struct {
	// MultipleATXs | MultipleBallots | HareEquivocation | InvalidPostIndex
	Type byte
	// AtxProof | BallotProof | HareProof | InvalidPostIndexProof
	Data ProofData
}

type ProofData interface {
	scale.Type
}

func (e *Proof) EncodeScale(enc *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"

	// not compact, as scale spec uses "full" uint8 for enums
	return 0, nil
}

func (e *Proof) DecodeScale(dec *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type MalfeasanceGossip struct {
	MalfeasanceProof
	Eligibility *types.HareEligibilityGossip // deprecated - to be removed in the next version
}

func (mg *MalfeasanceGossip) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type AtxProof struct {
	Messages [2]AtxProofMsg
}

func (ap *AtxProof) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type BallotProof struct {
	Messages [2]BallotProofMsg
}

func (bp *BallotProof) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type HareProof struct {
	Messages [2]HareProofMsg
}

func (hp *HareProof) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (hp *HareProof) ToMalfeasanceProof() *MalfeasanceProof { _ = "STUB: not implemented"; return nil }

type AtxProofMsg struct {
	InnerMsg types.ATXMetadata

	SmesherID types.NodeID
	Signature types.EdSignature
}

// SignedBytes returns the actual data being signed in a AtxProofMsg.
func (m *AtxProofMsg) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

type InvalidPostIndexProof struct {
	Atx wire.ActivationTxV1

	// Which index in POST is invalid
	InvalidIdx uint32
}

func (p *InvalidPostIndexProof) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type BallotProofMsg struct {
	InnerMsg types.BallotMetadata

	SmesherID types.NodeID
	Signature types.EdSignature
}

// SignedBytes returns the actual data being signed in a BallotProofMsg.
func (m *BallotProofMsg) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

type HareMetadata struct {
	Layer types.LayerID
	// the round counter (K)
	Round uint32
	// hash of hare.Message.InnerMessage
	MsgHash types.Hash32
}

func (hm *HareMetadata) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Equivocation detects if two messages form an equivocation, based on their HareMetadata.
// It returns true if the two messages are from the same layer and round, but have different hashes.
func (hm *HareMetadata) Equivocation(other *HareMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

func (hm HareMetadata) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

type HareProofMsg struct {
	InnerMsg HareMetadata

	SmesherID types.NodeID
	Signature types.EdSignature
}

// SignedBytes returns the actual data being signed in a HareProofMsg.
func (m *HareProofMsg) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

// InvalidPrevAtxProof is a proof that a smesher published an ATX with an old previous ATX ID.
// The proof contains two ATXs that reference the same previous ATX.
type InvalidPrevATXProof struct {
	Atx1 wire.ActivationTxV1
	Atx2 wire.ActivationTxV1
}

func MalfeasanceInfo(smesher types.NodeID, mp *MalfeasanceProof) string {
	_ = "STUB: not implemented"
	return ""
}
