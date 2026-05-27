package sim

// Frac is a shortcut for creating Fraction object.
func Frac(nominator, denominator int) Fraction { _ = "STUB: not implemented"; return *new(Fraction) }

// Fraction of something.
type Fraction struct {
	Nominator, Denominator int
}

func (f Fraction) String() string { _ = "STUB: not implemented"; return "" }

// SplitOpt is for configuring partition.
type SplitOpt func(*splitConf)

// WithPartitions configures number of miners in each partition
// relative to original Generator.
func WithPartitions(parts ...Fraction) SplitOpt { _ = "STUB: not implemented"; return *new(SplitOpt) }

func splitConfDefaults() splitConf { _ = "STUB: not implemented"; return *new(splitConf) }

type splitConf struct {
	Partitions []Fraction
	// TODO(dshulyak) support both transient and full partitions
	//
	// transient partitions can be configured without splitting state, in transient partitions
	// - votes are split
	// - hare will not reach consensus (or will reach only in one of the partition)
	// - but atxs and beacons remain the same
	//
	// for now transient can be configured by options to Next
}

// Split generator into multiple partitions.
// First generator will use original tortoise state (mesh, activations).
func (g *Generator) Split(opts ...SplitOpt) []*Generator {
	_ = "STUB: not implemented"
	// Split everything that would be changed after long network partition
	// - activations
	// - beacons
	// - number of miners in each partition
	//
	//	things that should remain the same:
	//
	// - blocks, contextually valid blocks and hare output before partition
	// - activations before partition
	// - beacons before partition
	return nil
}

func (g *Generator) mergeKeys(other *Generator) { _ = "STUB: not implemented"; return }

func (g *Generator) mergeActivations(other *Generator) { _ = "STUB: not implemented"; return }

func (g *Generator) mergeLayers(other *Generator) { _ = "STUB: not implemented"; return }

// Merge other Generator state into this Generator state.
func (g *Generator) Merge(other *Generator) { _ = "STUB: not implemented"; return }
