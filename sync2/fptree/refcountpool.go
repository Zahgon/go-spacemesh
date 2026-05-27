package fptree

import (
	"sync/atomic"
)

// freeBit is a bit that indicates that an entry is free.
const freeBit = 1 << 31

// freeListMask is a mask that extracts the free list index from a refCount.
const freeListMask = freeBit - 1

// poolEntry is an entry in the rcPool.
type poolEntry[T any, I ~uint32] struct {
	refCount uint32
	content  T
}

// rcPool is a reference-counted pool of items.
// The zero value is a valid, empty rcPool.
// Unlike sync.Pool, rcPool does not shrink, but uint32 indices can be used
// to reference items instead of larger 64-bit pointers, and the items
// can be shared between.
type rcPool[T any, I ~uint32] struct {
	entries []poolEntry[T, I]
	// freeList is 1-based so that rcPool doesn't need a constructor
	freeList   uint32
	allocCount atomic.Int64
}

// init pre-allocates the rcPool with n items.
func (rc *rcPool[T, I]) init(n int) {
	rc.entries = make([]poolEntry[T, I], 0, n)
	rc.freeList = 0
	rc.allocCount.Store(0)
}

// count returns the number of items in the rcPool.
func (rc *rcPool[T, I]) count() int { _ = "STUB: not implemented"; return 0 }

// item returns the item at the given index.
func (rc *rcPool[T, I]) item(idx I) T { _ = "STUB: not implemented"; return *new(T) }

// entry returns the pool entry at the given index.
func (rc *rcPool[T, I]) entry(idx I) *poolEntry[T, I] { _ = "STUB: not implemented"; return nil }

// replace replaces the item at the given index.
func (rc *rcPool[T, I]) replace(idx I, item T) { _ = "STUB: not implemented"; return }

// add adds an item to the rcPool and returns its index.
func (rc *rcPool[T, I]) add(item T) I { _ = "STUB: not implemented"; return *new(I) }

// release releases the item at the given index.
func (rc *rcPool[T, I]) release(idx I) bool { _ = "STUB: not implemented"; return false }

// ref adds a reference to the item at the given index.
func (rc *rcPool[T, I]) ref(idx I) { _ = "STUB: not implemented"; return }

// refCount returns the reference count for the item at the given index.
func (rc *rcPool[T, I]) refCount(idx I) uint32 { _ = "STUB: not implemented"; return 0 }
