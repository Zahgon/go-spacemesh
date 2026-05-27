package opinionhash

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hash"
)

var abstainSentinel = []byte{0}

// New returns new instance of the opinion hasher.
func New() *OpinionHasher { _ = "STUB: not implemented"; return nil }

// OpinionHasher is a utility for computing opinion hash.
type OpinionHasher struct {
	enc *scale.Encoder
	h   hash.Hash
}

// WritePrevious aggregated hash.
func (h *OpinionHasher) WritePrevious(hash types.Hash32) { _ = "STUB: not implemented"; return }

// WriteAbstain writes abstain sentinel as an opinion.
func (h *OpinionHasher) WriteAbstain() { _ = "STUB: not implemented"; return }

// WriteSupport writes id and height of the block.
func (h *OpinionHasher) WriteSupport(id types.BlockID, height uint64) {
	_ = "STUB: not implemented"
	return
}

// Sum appends hash sum to dst.
func (h *OpinionHasher) Sum(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

// Hash instantiates 32bytes and write hash.Sum to it.
func (h *OpinionHasher) Hash() (rst types.Hash32) {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

// Reset opinion hasher state.
func (h *OpinionHasher) Reset() { _ = "STUB: not implemented"; return }
