package types

import (
	"github.com/spacemeshos/go-scale"
)

// BytesToNodeID is a helper to copy buffer into a NodeID.
func BytesToNodeID(buf []byte) (id NodeID) { _ = "STUB: not implemented"; return *new(NodeID) }

// NodeID contains a miner's public key.
type NodeID Hash32

const (
	// NodeIDSize in bytes.
	NodeIDSize = Hash32Length
)

// String returns a string representation of the NodeID, for logging purposes.
// It implements the Stringer interface.
func (id NodeID) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the byte representation of the Edwards public key.
func (id NodeID) Bytes() []byte {
	_ = "STUB: not implemented"

	// ShortString returns a the first 3 hex-encoded bytes of the ID, for logging purposes.
	return nil
}

func (id NodeID) ShortString() string { _ = "STUB: not implemented"; return "" }

// EmptyNodeID is a canonical empty NodeID.
var EmptyNodeID NodeID

// EncodeScale implements scale codec interface.
func (id *NodeID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (id *NodeID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (id NodeID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *NodeID) UnmarshalText(buf []byte) error { _ = "STUB: not implemented"; return nil }

// NodeIDsToHashes turns a list of NodeID into their Hash32 representation.
func NodeIDsToHashes(ids []NodeID) []Hash32 { _ = "STUB: not implemented"; return nil }
