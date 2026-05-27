package rangesync

type generator struct {
	nextFn func() (KeyBytes, bool)
	stop   func()
	k      KeyBytes
	error  SeqErrorFunc
	done   bool
}

func gen(sr SeqResult) *generator { _ = "STUB: not implemented"; return nil }

func (g *generator) next() (k KeyBytes, ok bool) {
	_ = "STUB: not implemented"
	return *new(KeyBytes), false
}

func (g *generator) peek() (k KeyBytes, ok bool) {
	_ = "STUB: not implemented"
	return *new(KeyBytes), false
}

type combinedSeq struct {
	gens    []*generator
	wrapped []*generator
}

// CombineSeqs combines multiple ordered sequences from SeqResults into one, returning the
// smallest current key among all iterators at each step.
// startingPoint is used to check if an iterator has wrapped around. If an iterator yields
// a value below startingPoint, it is considered to have wrapped around.
func CombineSeqs(startingPoint KeyBytes, srs ...SeqResult) SeqResult {
	_ = "STUB: not implemented"
	return *new(SeqResult)
}

// We clean up even if c.begin() returned an error so that we don't leak
// any pull iterators that are created before c.begin() failed

// In case if c.begin() succeeds, the error is reset. If yield
// calls SeqResult's Error function, it will get nil until the
// iteration is finished.

func (c *combinedSeq) begin(startingPoint KeyBytes, srs []SeqResult) error {
	_ = "STUB: not implemented"
	return nil
}

// all iterators wrapped around

func (c *combinedSeq) end() { _ = "STUB: not implemented"; return }

func (c *combinedSeq) aheadGen() (ahead *generator, aheadIdx int, err error) {
	_ = "STUB: not implemented"
	// remove any exhausted generators
	return nil, 0, nil
}

// if all the generators have wrapped around, move the wrapped generators

// If not all of the generators have wrapped around, then we
// already did a successful peek() on this generator above, so it
// should not be exhausted here.
// If all of the generators have wrapped around, then we have
// moved to the wrapped generators, but the generators may only
// end up in wrapped list after a successful peek(), too.
// So if we get here, then combinedSeq code is broken.

func (c *combinedSeq) iterate(yield func(KeyBytes) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// if this iterator is exhausted, it'll be removed by the
// next aheadGen call

// the iterator has wrapped around, move it to the wrapped
// list which will be used after all the iterators have
// wrapped around
