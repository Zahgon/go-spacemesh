package rangesync

import (
	"sync"
	"time"
)

// stringToFP converts a string to a Fingerprint.
// It is only used in tests.
func stringToFP(s string) Fingerprint { _ = "STUB: not implemented"; return *new(Fingerprint) }

// gtePos returns the position of the first item in the given list that is greater than or
// equal to the given item. If no such item is found, it returns the length of the list.
func gtePos(all []KeyBytes, item KeyBytes) int { _ = "STUB: not implemented"; return 0 }

// naiveRange returns the items in the range [x, y) from the given list of items,
// supporting wraparound (x > y) and full (x == y) ranges.
// It also returns the first and last item falling within the range.
func naiveRange(
	all []KeyBytes,
	x, y KeyBytes,
	stopCount int,
) (
	items []KeyBytes,
	startID, endID KeyBytes,
) {
	_ = "STUB: not implemented"
	return nil, *new(KeyBytes), *new(KeyBytes)
}

func naiveFPFunc(items []KeyBytes) Fingerprint { _ = "STUB: not implemented"; return *new(Fingerprint) }

func realFPFunc(items []KeyBytes) Fingerprint { _ = "STUB: not implemented"; return *new(Fingerprint) }

// DumbSet is a simple OrderedSet implementation that doesn't include any optimizations.
// It is intended to be only used in tests.
type DumbSet struct {
	copyMtx           sync.Mutex
	keys              []KeyBytes
	received          map[string]int
	added             map[string]bool
	allowMultiReceive bool
	FPFunc            func(items []KeyBytes) Fingerprint
}

var _ OrderedSet = &DumbSet{}

// SetAllowMultiReceive sets whether the set allows receiving the same item multiple times.
func (ds *DumbSet) SetAllowMultiReceive(allow bool) { _ = "STUB: not implemented"; return }

// AddUnchecked adds an item to the set without registerting the item for checks
// as in case of Add and Receive.
func (ds *DumbSet) AddUnchecked(id KeyBytes) { _ = "STUB: not implemented"; return }

// already present

// AddReceived adds all the received items to the set.
func (ds *DumbSet) AddReceived() { _ = "STUB: not implemented"; return }

// DumbSet's Received implementation should never return an error

// Add implements the OrderedSet.
func (ds *DumbSet) Add(id KeyBytes) error { _ = "STUB: not implemented"; return nil }

// Add is invoked during recent sync.
// If the item was already received, this means it was received as part of the
// recent sync, and thus may be received again as the algorithm does not guarantee
// that items already in the set are not received from the remote peer, only that
// no item is received twice (with exception of recent sync).

// Receive implements the OrderedSet.
func (ds *DumbSet) Receive(id KeyBytes) error { _ = "STUB: not implemented"; return nil }

// Received implements the OrderedSet.
func (ds *DumbSet) Received() SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }

// seq returns an endless sequence as a SeqResult starting from the given index.
func (ds *DumbSet) seq(n int) SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }

// make the sequence reusable

// seqFor returns an endless sequence as a SeqResult starting from the given key, or the
// lowest key greater than the given key if the key is not present in the set.
func (ds *DumbSet) seqFor(s KeyBytes) SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }

func (ds *DumbSet) getRangeInfo(
	x, y KeyBytes,
	count int,
) (r RangeInfo, end KeyBytes, err error) {
	_ = "STUB: not implemented"
	return *new(RangeInfo), *new(KeyBytes), nil
}

// RangeInfo implements OrderedSet.
func (ds *DumbSet) RangeInfo(x, y KeyBytes) (RangeInfo, error) {
	_ = "STUB: not implemented"
	return *new(RangeInfo), nil
}

// SplitRange implements OrderedSet.
func (ds *DumbSet) SplitRange(x, y KeyBytes, count int) (SplitInfo, error) {
	_ = "STUB: not implemented"
	return *new(SplitInfo), nil
}

// SetInfo implements OrderedSet.
func (ds *DumbSet) SetInfo() (RangeInfo, error) {
	_ = "STUB: not implemented"
	return *new(RangeInfo), nil
}

// WithCopy implements OrderedSet.
func (ds *DumbSet) WithCopy(toCall func(OrderedSet) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Recent implements OrderedSet.
func (ds *DumbSet) Recent(since time.Time) (SeqResult, int) {
	_ = "STUB: not implemented"
	return *new(SeqResult), 0
}

// Loaded implements OrderedSet.
func (ds *DumbSet) Loaded() bool {
	_ = "STUB: not implemented"

	// EnsureLoaded implements OrderedSet.
	return false
}

func (ds *DumbSet) EnsureLoaded() error {
	_ = "STUB: not implemented"

	// Advance implements OrderedSet.
	return nil
}

func (ds *DumbSet) Advance() error {
	_ = "STUB: not implemented"

	// Has implements OrderedSet.
	return nil
}

func (ds *DumbSet) Has(k KeyBytes) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Release implements OrderedSet.
func (ds *DumbSet) Release() { _ = "STUB: not implemented"; return }
