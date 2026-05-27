package fptree

import (
	"sync"

	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// nodeIndex represents an index of a node in the node pool.
type nodeIndex uint32

const (
	// noIndex represents an invalid node index.
	noIndex = ^nodeIndex(0)
	// leafFlag is a flag that indicates that a node is a leaf node.
	leafFlag = uint32(1 << 31)
)

// node represents an fpTree node.
type node struct {
	// Fingerprint
	fp rangesync.Fingerprint
	// Item count
	c uint32
	// Left child, noIndex if not present.
	l nodeIndex
	// Right child, noIndex if not present.
	r nodeIndex
}

// nodePool represents a pool of tree nodes.
// The pool is shared between the orignal tree and its clones.
type nodePool struct {
	mtx     sync.RWMutex
	rcPool  rcPool[node, uint32]
	leafMap map[uint32]rangesync.KeyBytes
}

// init pre-allocates the node pool with n nodes.
func (np *nodePool) init(n int) {
	np.rcPool.init(n)
}

// lockWrite locks the node pool for writing.
// There can only be one writer at a time.
// This blocks until all other reader and writer locks are released.
func (np *nodePool) lockWrite() {
	_ = "STUB: not implemented"

	// unlockWrite unlocks the node pool for writing.
	return
}

func (np *nodePool) unlockWrite() {
	_ = "STUB: not implemented"

	// lockRead locks the node pool for reading.
	// There can be multiple reader locks held at a time.
	// This blocks until the writer lock is released, if it's held.
	return
}

func (np *nodePool) lockRead() {
	_ = "STUB: not implemented"

	// unlockRead unlocks the node pool for reading.
	return
}

func (np *nodePool) unlockRead() {
	_ = "STUB: not implemented"

	// add adds a new node to the pool.
	return
}

func (np *nodePool) add(
	fp rangesync.Fingerprint,
	c uint32,
	left, right nodeIndex,
	v rangesync.KeyBytes,
	replaceIdx nodeIndex,
) nodeIndex {
	_ = "STUB: not implemented"
	return *new(nodeIndex)
}

// value returns the value of the node at the given index.
func (np *nodePool) value(idx nodeIndex) rangesync.KeyBytes {
	_ = "STUB: not implemented"
	return *new(rangesync.KeyBytes)
}

// left returns the left child of the node at the given index.
func (np *nodePool) left(idx nodeIndex) nodeIndex {
	_ = "STUB: not implemented"
	return *new(nodeIndex)
}

// right returns the right child of the node at the given index.
func (np *nodePool) right(idx nodeIndex) nodeIndex {
	_ = "STUB: not implemented"
	return *new(nodeIndex)
}

// leaf returns true if this is a leaf node.
func (np *nodePool) leaf(idx nodeIndex) bool { _ = "STUB: not implemented"; return false }

// count returns number of set items to which the node at the given index corresponds.
func (np *nodePool) count(idx nodeIndex) uint32 { _ = "STUB: not implemented"; return 0 }

// info returns the count, fingerprint, and leaf flag of the node at the given index.
func (np *nodePool) info(idx nodeIndex) (count uint32, fp rangesync.Fingerprint, leaf bool) {
	_ = "STUB: not implemented"
	return 0, *new(rangesync.Fingerprint), false
}

// releaseOne releases the node at the given index, returning it to the pool.
func (np *nodePool) releaseOne(idx nodeIndex) bool { _ = "STUB: not implemented"; return false }

// release releases the node at the given index, returning it to the pool, and recursively
// releases its children.
func (np *nodePool) release(idx nodeIndex) bool { _ = "STUB: not implemented"; return false }

// ref adds a reference to the given node.
func (np *nodePool) ref(idx nodeIndex) { _ = "STUB: not implemented"; return }

// refCount returns the reference count for the node at the given index.
func (np *nodePool) refCount(idx nodeIndex) uint32 { _ = "STUB: not implemented"; return 0 }

// nodeCount returns the number of nodes in the pool.
func (np *nodePool) nodeCount() int { _ = "STUB: not implemented"; return 0 }
