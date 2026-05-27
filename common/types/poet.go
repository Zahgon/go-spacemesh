package types

import (
	"net/url"
	"time"

	poetShared "github.com/spacemeshos/poet/shared"
	"go.uber.org/zap/zapcore"
)

//go:generate scalegen -types PoetProof,PoetProofMessage

type PoetServer struct {
	Address string    `mapstructure:"address" json:"address"`
	Pubkey  Base64Enc `mapstructure:"pubkey"  json:"pubkey"`
}

func ByteToPoetProofRef(b []byte) (ref PoetProofRef) {
	_ = "STUB: not implemented"
	return *new(PoetProofRef)
}

type PoetProofRef Hash32

func (r *PoetProofRef) String() string { _ = "STUB: not implemented"; return "" }

// EmptyPoetProofRef is an empty PoET proof reference.
var EmptyPoetProofRef = PoetProofRef{}

// PoetProof is the full PoET service proof of elapsed time.
// It includes the number of leaves produced and the actual PoET Merkle proof.
type PoetProof struct {
	poetShared.MerkleProof
	LeafCount uint64
}

func (p *PoetProof) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// PoetProofMessage is the envelope which includes the PoetProof, service ID, round ID and signature.
type PoetProofMessage struct {
	PoetProof
	PoetServiceID []byte `scale:"max=32"` // public key of the PoET service
	RoundID       string `scale:"max=32"`
	// The input to Poet's POSW.
	// It's the root of a merkle tree built from all of the members
	// that are included in the proof.
	Statement Hash32
	Signature EdSignature
}

func (p *PoetProofMessage) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Ref returns the reference to the PoET proof message. It's the blake3 sum of the entire proof.
func (p *PoetProof) Ref() (PoetProofRef, error) {
	_ = "STUB: not implemented"
	return *new(PoetProofRef), nil
}

// PoetRound includes the PoET's round ID.
type PoetRound struct {
	ID  string `scale:"max=32"`
	End time.Time
}

type PoetInfo struct {
	ServicePubkey []byte
	PhaseShift    time.Duration
	CycleGap      time.Duration
	Certifier     *CertifierInfo
}

type CertifierInfo struct {
	Url    *url.URL
	Pubkey []byte
}
