// Package types defines the types used by go-spacemesh consensus algorithms and structs
package types

import (
	"github.com/spacemeshos/go-scale"
)

var (
	layersPerEpoch uint32
	// effectiveGenesis marks when actual proposals would start being created in the network. It takes into account
	// the first genesis epoch and the following epoch in which ATXs are published.
	effectiveGenesis uint32

	// EmptyLayerHash is the layer hash for an empty layer.
	EmptyLayerHash = EmptyHash32
)

// SetLayersPerEpoch sets global parameter of layers per epoch, all conversions from layer to epoch use this param.
func SetLayersPerEpoch(layers uint32) { _ = "STUB: not implemented"; return }

func SetEffectiveGenesis(layer uint32) { _ = "STUB: not implemented"; return }

// GetLayersPerEpoch returns number of layers per epoch.
func GetLayersPerEpoch() uint32 { _ = "STUB: not implemented"; return 0 }

// FirstEffectiveGenesis returns the first effective genesis layer.
func FirstEffectiveGenesis() LayerID { _ = "STUB: not implemented"; return *new(LayerID) }

// GetEffectiveGenesis returns the last layer of genesis.
// This value can change after a checkpoint recovery.
func GetEffectiveGenesis() LayerID { _ = "STUB: not implemented"; return *new(LayerID) }

// LayerID is representing a layer number. Zero value is safe to use, and means 0.
// Internally it is a simple wrapper over uint32 and should be considered immutable
// the same way as any integer.
type LayerID uint32

// EncodeScale implements scale codec interface.
func (l LayerID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (l *LayerID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetEpoch returns the epoch number of this LayerID.
func (l LayerID) GetEpoch() EpochID { _ = "STUB: not implemented"; return *new(EpochID) }

// Add layers to the layer. Panics on wraparound.
func (l LayerID) Add(layers uint32) LayerID { _ = "STUB: not implemented"; return *new(LayerID) }

// Sub layers from the layer. Panics on wraparound.
func (l LayerID) Sub(layers uint32) LayerID { _ = "STUB: not implemented"; return *new(LayerID) }

// OrdinalInEpoch returns layer ordinal in epoch.
func (l LayerID) OrdinalInEpoch() uint32 { _ = "STUB: not implemented"; return 0 }

// FirstInEpoch returns whether this LayerID is first in epoch.
func (l LayerID) FirstInEpoch() bool { _ = "STUB: not implemented"; return false }

// Mul layer by the layers. Panics on wraparound.
func (l LayerID) Mul(layers uint32) LayerID { _ = "STUB: not implemented"; return *new(LayerID) }

// Uint32 returns the LayerID as a uint32.
func (l LayerID) Uint32() uint32 {
	_ = "STUB: not implemented"

	// Before returns true if this layer is lower than the other.
	return 0
}

func (l LayerID) Before(other LayerID) bool {
	_ = "STUB: not implemented"

	// After returns true if this layer is higher than the other.
	return false
}

func (l LayerID) After(other LayerID) bool {
	_ = "STUB: not implemented"

	// Difference returns the difference between current and other layer.
	return false
}

func (l LayerID) Difference(other LayerID) uint32 { _ = "STUB: not implemented"; return 0 }

// String returns string representation of the layer id numeric value.
func (l LayerID) String() string { _ = "STUB: not implemented"; return "" }

// Layer contains a list of proposals and their corresponding LayerID.
type Layer struct {
	index   LayerID
	ballots []*Ballot
	blocks  []*Block
}

// Index returns the layer's ID.
func (l *Layer) Index() LayerID {
	_ = "STUB: not implemented"

	// Blocks returns the list of Block in this layer.
	return *new(LayerID)
}

func (l *Layer) Blocks() []*Block {
	_ = "STUB: not implemented"

	// BlocksIDs returns the list of IDs for blocks in this layer.
	return nil
}

func (l *Layer) BlocksIDs() []BlockID { _ = "STUB: not implemented"; return nil }

// Ballots returns the list of ballots in this layer.
func (l *Layer) Ballots() []*Ballot {
	_ = "STUB: not implemented"

	// BallotIDs returns the list of IDs for ballots in this layer.
	return nil
}

func (l *Layer) BallotIDs() []BallotID { _ = "STUB: not implemented"; return nil }

// AddBallot adds a ballot to this layer. Panics if the ballot's index doesn't match the layer.
func (l *Layer) AddBallot(b *Ballot) { _ = "STUB: not implemented"; return }

// AddBlock adds a block to this layer. Panics if the block's index doesn't match the layer.
func (l *Layer) AddBlock(b *Block) { _ = "STUB: not implemented"; return }

// SetBallots sets the list of ballots for the layer without validation.
func (l *Layer) SetBallots(ballots []*Ballot) { _ = "STUB: not implemented"; return }

// SetBlocks sets the list of blocks for the layer without validation.
func (l *Layer) SetBlocks(blocks []*Block) {
	_ = "STUB: not implemented"

	// NewExistingLayer returns a new layer with the given list of blocks without validation.
	return
}

func NewExistingLayer(idx LayerID, ballots []*Ballot, blocks []*Block) *Layer {
	_ = "STUB: not implemented"
	return nil
}

// NewLayer returns a layer with no proposals.
func NewLayer(layerIndex LayerID) *Layer { _ = "STUB: not implemented"; return nil }
