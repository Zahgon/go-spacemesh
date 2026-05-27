package sim

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
)

// DefaultNumBlocks is a number of blocks in a layer by default.
const DefaultNumBlocks = 5

// NextOpt is for configuring layer generator.
type NextOpt func(*nextConf)

func nextConfDefaults() nextConf { _ = "STUB: not implemented"; return *new(nextConf) }

type nextConf struct {
	Reorder          uint32
	FailHare         bool
	EmptyHare        bool
	HareOutputIndex  int
	Coinflip         bool
	LayerSize        int
	NumBlocks        int
	BlockTickHeights []uint64
	VoteGen          VotesGenerator
}

// WithNextReorder configures when reordered layer should be returned.
// Examples:
// Next() Next(WithNextReorder(1)) Next()
// 1      3                        2
// Next() Next(WithNextReorder(2)) Next() Next()
// 1      3                        4      2
//
// So the Next layer with WithNextReorder will be delayed exactly by `delay` value.
func WithNextReorder(delay uint32) NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithoutHareOutput will prevent from saving hare output.
func WithoutHareOutput() NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithEmptyHareOutput will save an empty vector for hare output.
func WithEmptyHareOutput() NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithLayerSizeOverwrite overwrite expected layer size.
func WithLayerSizeOverwrite(size int) NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithVoteGenerator declares vote generator for a layer.
func WithVoteGenerator(gen VotesGenerator) NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithCoin is to setup weak coin for voting. By default coin will support blocks.
func WithCoin(coin bool) NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithBlockTickHeights updates height of the blocks.
func WithBlockTickHeights(heights ...uint64) NextOpt {
	_ = "STUB: not implemented"
	return *new(NextOpt)
}

// WithNumBlocks sets number of the generated blocks.
func WithNumBlocks(num int) NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// WithHareOutputIndex sets the index of the block that will be stored as a hare output.
func WithHareOutputIndex(i int) NextOpt { _ = "STUB: not implemented"; return *new(NextOpt) }

// Next generates the next layer.
func (g *Generator) Next(opts ...NextOpt) types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

// TODO(dshulyak) we are not reordering already reordered layer

// Add(1) to account for generated layer at the end

func (g *Generator) genBeacon() { _ = "STUB: not implemented"; return }

func (g *Generator) genTXIDs(n int) []types.TransactionID { _ = "STUB: not implemented"; return nil }

func (g *Generator) genLayer(cfg nextConf) types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}
