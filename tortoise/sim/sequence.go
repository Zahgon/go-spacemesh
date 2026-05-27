package sim

import "github.com/spacemeshos/go-spacemesh/common/types"

// Sequence of layers with same configuration.
type Sequence struct {
	Length int
	Opts   []NextOpt
}

// WithSequence creates Sequence object.
func WithSequence(lth int, opts ...NextOpt) Sequence {
	_ = "STUB: not implemented"
	return *new(Sequence)
}

// GenLayers produces sequence of layers using all configurators.
func GenLayers(g *Generator, seqs ...Sequence) []types.LayerID {
	_ = "STUB: not implemented"
	return nil
}
