package multipeer

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// syncRange represents a range of keys to be synchronized, along with the status
// information.
type syncRange struct {
	X, Y            rangesync.KeyBytes
	LastSyncStarted time.Time
	Done            bool
	NumSyncers      int
	Index           int
}

// syncQueue is a priority queue for syncRanges.
type syncQueue []*syncRange

// Len implements heap.Interface.
func (sq syncQueue) Len() int {
	_ = "STUB: not implemented"

	// Less implements heap.Interface.
	return 0
}

func (sq syncQueue) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// We want Pop to give us syncRange for which which sync has started the
	// earliest. Items which are not being synced are considered "most earliest"
	return false
}

// Swap implements heap.Interface.
func (sq syncQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Push implements heap.Interface.
func (sq *syncQueue) Push(i any) { _ = "STUB: not implemented"; return }

// Pop implements heap.Interface.
func (sq *syncQueue) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// avoid memory leak
// not in the queue anymore

func newSyncQueue(numPeers, keyLen, maxDepth int) syncQueue {
	_ = "STUB: not implemented"
	return *new(syncQueue)
}

func (sq *syncQueue) empty() bool { _ = "STUB: not implemented"; return false }

func (sq *syncQueue) PopRange() *syncRange { _ = "STUB: not implemented"; return nil }

func (sq *syncQueue) PushRange(sr *syncRange) { _ = "STUB: not implemented"; return }

func (sq *syncQueue) Update(sr *syncRange, lastSyncStarted time.Time) {
	_ = "STUB: not implemented"
	return
}
