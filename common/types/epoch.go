package types

import (
	"github.com/spacemeshos/go-scale"
)

// EpochID is the running epoch number. It's zero-based, so the genesis epoch has EpochID == 0.
type EpochID uint32

func (e EpochID) Uint32() uint32 {
	_ = "STUB: not implemented"

	// EncodeScale implements scale codec interface.
	return 0
}

func (e EpochID) EncodeScale(enc *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (e *EpochID) DecodeScale(dec *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FirstLayer returns the layer ID of the first layer in the epoch.
func (e EpochID) FirstLayer() LayerID { _ = "STUB: not implemented"; return *new(LayerID) }

// Add Epochs to the EpochID. Panics on wraparound.
func (e EpochID) Add(epochs uint32) EpochID { _ = "STUB: not implemented"; return *new(EpochID) }

// String returns string representation of the epoch id numeric value.
func (e EpochID) String() string { _ = "STUB: not implemented"; return "" }
