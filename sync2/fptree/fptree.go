package fptree

import (
	"errors"
	"io"

	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
	"github.com/spacemeshos/go-spacemesh/sync2/sqlstore"
)

var errEasySplitFailed = errors.New("easy split failed")

const (
	FingerprintSize = rangesync.FingerprintSize
	// sizeHintCoef is used to calculate the number of pool entries to preallocate for
	// an FPTree based on the expected number of items which this tree may contain.
	sizeHintCoef = 2.1
)

// FPResult represents the result of a range fingerprint query against FPTree, as returned
// by FingerprintInterval.
type FPResult struct {
	// Range fingerprint
	FP rangesync.Fingerprint
	// Number of items in the range
	Count uint32
	// Interval type: -1 for normal, 0 for the whole set, 1 for wrapped around ("inverse")
	IType int
	// Items in the range
	Items rangesync.SeqResult
	// The item following the range
	Next rangesync.KeyBytes
}

// SplitResult represents the result of a split operation.
type SplitResult struct {
	// The two parts of the interval
	Part0, Part1 FPResult
	// Middle point value
	Middle rangesync.KeyBytes
}

// aggContext is the context used for aggregation operations.
type aggContext struct {
	// nodePool used by the tree
	np *nodePool
	// Bounds of the interval being aggregated
	x, y rangesync.KeyBytes
	// The current fingerprint of the items aggregated so far, since the beginning or
	// after the split ("easy split")
	fp rangesync.Fingerprint
	// The fingerprint of the items aggregated in the first part of the split
	fp0 rangesync.Fingerprint
	// Number of items aggregated so far, since the beginning or after the split
	// ("easy split")
	count uint32
	// Number of items aggregated in the first part of the split
	count0 uint32
	// Interval type: -1 for normal, 0 for the whole set, 1 for wrapped around ("inverse")
	itype int
	// Maximum remaining number of items to aggregate.
	limit int
	// The number of items aggregated so far.
	total uint32
	// The resulting item sequence.
	items rangesync.SeqResult
	// The item immediately following the aggregated items.
	next rangesync.KeyBytes
	// The prefix corresponding to the last aggregated node.
	lastPrefix *prefix
	// The prefix corresponding to the last aggregated node in the first part of the split.
	lastPrefix0 *prefix
	// Whether the aggregation is being done for an "easy split" (split operation
	// without querying the underlying IDStore).
	easySplit bool
}

// prefixAtOrAfterX verifies that the any key with the prefix p is at or after x.
// It can be used for the whole interval in case of a normal interval.
// With inverse intervals, it should only be used when processing the [x, max) part of the
// interval.
func (ac *aggContext) prefixAtOrAfterX(p prefix) bool { _ = "STUB: not implemented"; return false }

// prefixBelowY verifies that the any key with the prefix p is below y.
// It can be used for the whole interval in case of a normal interval.
// With inverse intervals, it should only be used when processing the [0, y) part of the
// interval.
func (ac *aggContext) prefixBelowY(p prefix) bool { _ = "STUB: not implemented"; return false }

// If p.idAfter(b) is true, this means there's wraparound and
// b is zero whereas all the possible keys beginning with prefix p
// are non-zero. In this case, there can be no key y such that
// all the keys beginning with prefix p are below y.

// fingerprintAtOrAfterX verifies that the specified fingerprint, which should be derived
// from a single key, is at or after x bound of the interval.
func (ac *aggContext) fingerprintAtOrAfterX(fp rangesync.Fingerprint) bool {
	_ = "STUB: not implemented"
	return false
}

// fingerprintBelowY verifies that the specified fingerprint, which should be derived from a
// single key, is below y bound of the interval.
func (ac *aggContext) fingerprintBelowY(fp rangesync.Fingerprint) bool {
	_ = "STUB: not implemented"
	return false
}

// 1 after max key derived from the fingerprint

// nodeAtOrAfterX verifies that the node with the given index is at or after x bound of the
// interval.
func (ac *aggContext) nodeAtOrAfterX(idx nodeIndex, p prefix) bool {
	_ = "STUB: not implemented"
	return false
}

// nodeBelowY verifies that the node with the given index is below y bound of the interval.
func (ac *aggContext) nodeBelowY(idx nodeIndex, p prefix) bool {
	_ = "STUB: not implemented"
	return false
}

// pruneX returns true if the specified node can be pruned during left-aggregation because
// all of its keys are below the x bound of the interval.
func (ac *aggContext) pruneX(idx nodeIndex, p prefix) bool { _ = "STUB: not implemented"; return false }

// idAfter derived from the prefix is at or below y => prune

// node has count > 1, so we can't use its fingerprint or value to
// determine if it's at or after X

// 1 after max key derived from the fingerprint

// pruneY returns true if the specified node can be pruned during right-aggregation
// because all of its keys are at or after the y bound of the interval.
func (ac *aggContext) pruneY(idx nodeIndex, p prefix) bool { _ = "STUB: not implemented"; return false }

// min ID derived from the prefix is at or after y => prune

// node has count > 1, so we can't use its fingerprint or value to
// determine if it's below y

// switchToSecondPart switches aggregation to the second part of the "easy split".
func (ac *aggContext) switchToSecondPart() { _ = "STUB: not implemented"; return }

// maybeIncludeNode returns tries to include the full contents of the specified node in
// the aggregation and returns if it succeeded, based on the remaining limit and the number
// of items in the node.
// It also handles "easy split" happening at the node.
func (ac *aggContext) maybeIncludeNode(idx nodeIndex, p prefix) bool {
	_ = "STUB: not implemented"
	return false
}

// We're doing a split and this node is over the limit, but the first part
// is still empty so we include this node in the first part and
// then switch to the second part

// We're doing a split and this node is over the limit, so store count and
// fingerprint for the first part and include the current node in the
// second part

// We're doing a split and this node is exactly at the limit, or it was
// above the limit but first part was still empty, so store count and
// fingerprint for the first part which includes the current node and zero
// out count and fingerprint for the second part

// FPTree is a binary tree data structure designed to perform range fingerprint queries
// efficiently.
// FPTree can work on its own, with fingerprint query complexity being O(log n).
// It can also be backed by an IDStore with a depth limit the binary tree, in which
// case the query efficiency degrades with the number of items growing.
// O(log n) query efficiency can be retained in this case for queries which
// have the number of non-zero bits, starting from the high bit, below maxDepth.
// FPTree does not do any special balancing and relies on the IDs added on it being
// uniformly distributed, which is the case for the IDs based on cryptographic hashes.
type FPTree struct {
	trace
	idStore  sqlstore.IDStore
	np       *nodePool
	root     nodeIndex
	keyLen   int
	maxDepth int
}

var _ sqlstore.IDStore = &FPTree{}

// NewFPTreeWithValues creates an FPTree which also stores the items themselves and does
// not make use of a backing IDStore.
// sizeHint specifies the approximate expected number of items.
// keyLen specifies the number of bytes in keys used.
func NewFPTreeWithValues(sizeHint, keyLen int) *FPTree { _ = "STUB: not implemented"; return nil }

// NewFPTree creates an FPTree of limited depth backed by an IDStore.
// sizeHint specifies the approximate expected number of items.
// keyLen specifies the number of bytes in keys used.
func NewFPTree(sizeHint int, idStore sqlstore.IDStore, keyLen, maxDepth int) *FPTree {
	_ = "STUB: not implemented"
	return nil
}

// traverse traverses the subtree rooted in idx in order and calls the given function for
// each item.
func (ft *FPTree) traverse(idx nodeIndex, yield func(rangesync.KeyBytes) bool) (res bool) {
	_ = "STUB: not implemented"
	return false
}

// traverseFrom traverses the subtree rooted in idx in order and calls the given function for
// each item starting from the given key.
func (ft *FPTree) traverseFrom(
	idx nodeIndex,
	p prefix,
	from rangesync.KeyBytes,
	yield func(rangesync.KeyBytes) bool,
) (res bool) {
	_ = "STUB: not implemented"
	return false
}

func (ft *FPTree) all() rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// All returns all the items currently in the tree (including those in the IDStore).
// The sequence in SeqResult is either empty or infinite.
// Implements sqlstore.All.
func (ft *FPTree) All() rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// From returns all the items in the tree that are greater than or equal to the given key.
// The sequence in SeqResult is either empty or infinite.
// Implements sqlstore.IDStore.
func (ft *FPTree) From(from rangesync.KeyBytes, sizeHint int) rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

func (ft *FPTree) from(from rangesync.KeyBytes, sizeHint int) rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// Release releases resources used by the tree.
// Implements sqlstore.IDStore.
func (ft *FPTree) Release() { _ = "STUB: not implemented"; return }

// Clear removes all items from the tree.
// It should only be used with trees that were created using NewFPtreeWithValues.
func (ft *FPTree) Clear() { _ = "STUB: not implemented"; return }

// if we have an idStore, it can't be cleared and thus the tree can't be
// cleared either

// Clone makes a copy of the tree.
// The copy operation is thread-safe and has complexity of O(1).
func (ft *FPTree) Clone() sqlstore.IDStore {
	_ = "STUB: not implemented"
	return *new(sqlstore.IDStore)
}

// pushLeafDown pushes a leaf node down the tree when the node's path matches that of the
// new to be added, splitting it if necessary.
func (ft *FPTree) pushLeafDown(
	idx nodeIndex,
	replace bool,
	singleFP, prevFP rangesync.Fingerprint,
	depth int,
	curCount uint32,
	value, prevValue rangesync.KeyBytes,
) (newIdx nodeIndex) {
	_ = "STUB: not implemented"
	return *new(nodeIndex)
}

// Once we stumble upon a node with refCount > 1, we no longer can replace nodes
// as they're also referenced by another tree.

// TODO: in the proper radix tree, these 1-child nodes should never be
// created, accumulating the prefix instead

// addValue adds a value to the subtree rooted in idx.
func (ft *FPTree) addValue(
	idx nodeIndex,
	replace bool,
	fp rangesync.Fingerprint,
	depth int,
	value rangesync.KeyBytes,
) (newIdx nodeIndex) {
	_ = "STUB: not implemented"
	return *new(nodeIndex)
}

// Once we stumble upon a node with refCount > 1, we no longer can replace nodes
// as they're also referenced by another tree.

// we're at a leaf node, need to push down the old fingerprint, or,
// if we've reached the max depth, just update the current node

// the original node is not being replaced, so the reused left
// node has acquired another reference

// the original node is not being replaced, so the reused right
// node has acquired another reference

// AddStoredKey adds a key to the tree, assuming that either the tree doesn't have an
// IDStore ar the IDStore already contains the key.
func (ft *FPTree) AddStoredKey(k rangesync.KeyBytes) { _ = "STUB: not implemented"; return }

// RegisterKey registers a key in the tree.
// If the tree has an IDStore, the key is also registered with the IDStore.
func (ft *FPTree) RegisterKey(k rangesync.KeyBytes) error { _ = "STUB: not implemented"; return nil }

// storeValues returns true if the tree stores the values (has no IDStore).
func (ft *FPTree) storeValues() bool { _ = "STUB: not implemented"; return false }

// CheckKey returns true if the tree contains or may contain the given key.
// If this function returns false, the tree definitely doesn't contain the key.
// If this function returns true and the tree stores the values, the key is definitely
// contained in the tree.
// If this function returns true and the tree doesn't store the values, the key may be
// contained in the tree.
func (ft *FPTree) CheckKey(k rangesync.KeyBytes) bool { _ = "STUB: not implemented"; return false }

// We're unlikely to be able to find a node with the full prefix, but if we can
// find a leaf node with matching partial prefix, that's good enough except
// that we also need to check the node's fingerprint.

// followPrefix follows the bit prefix p from the node idx.
func (ft *FPTree) followPrefix(from nodeIndex, p, followed prefix) (idx nodeIndex, rp prefix, found bool) {
	_ = "STUB: not implemented"
	return *new(nodeIndex), *new(prefix), false
}

// aggregateEdge aggregates an edge of the interval, which can be bounded by x, y, both x
// and y or none of x and y, have a common prefix and optionally bounded by a limit of N of
// aggregated items.
// It returns a boolean indicating whether the limit or the right edge (y) was reached and
// an error, if any.
func (ft *FPTree) aggregateEdge(
	x, y rangesync.KeyBytes,
	idx nodeIndex,
	p prefix,
	ac *aggContext,
) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// easySplit means we should not be querying the database,
// so we'll have to retry using slower strategy

// aggregateUpToLimit aggregates the subtree rooted in idx up to the limit of N of nodes.
func (ft *FPTree) aggregateUpToLimit(idx nodeIndex, p prefix, ac *aggContext) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// node is fully included

// reached the limit on this node, do not need to continue after
// done with it

// left node is fully included, after which
// we need to stop somewhere in the right subtree

// we must stop somewhere in the left subtree,
// and the right subtree is irrelevant unless
// easySplit is being done and we must restart
// after the limit is exhausted

// aggregateLeft aggregates the subtree that covers the left subtree of the LCA in case of
// normal intervals, and the subtree that covers [x, MAX] part for the inverse (wrapped
// around) intervals.
func (ft *FPTree) aggregateLeft(
	idx nodeIndex,
	k rangesync.KeyBytes,
	p prefix,
	ac *aggContext,
) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// for ac.limit == 0, it's important that we still visit the node
// so that we can get the item immediately following the included items

// we've not reached X yet so we should not stop, thus true

// aggregateRight aggregates the subtree that covers the right subtree of the LCA in case
// of normal intervals, and the subtree that covers [0, y) part for the inverse (wrapped
// around) intervals.
func (ft *FPTree) aggregateRight(
	idx nodeIndex,
	k rangesync.KeyBytes,
	p prefix,
	ac *aggContext,
) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// aggregateXX aggregates intervals of form [x, x) which denotes the whole set.
func (ft *FPTree) aggregateXX(ac *aggContext) (err error) {
	_ = "STUB: not implemented"
	// [x, x) interval which denotes the whole set unless
	// the limit is specified, in which case we need to start aggregating
	// with x and wrap around if necessary
	return nil
}

// We need to aggregate up to ac.limit number of items starting
// from x and wrapping around if necessary

// aggregateSimple aggregates simple (normal) intervals of form [x, y) where x < y.
func (ft *FPTree) aggregateSimple(ac *aggContext) (err error) {
	_ = "STUB: not implemented"
	// "proper" interval: [x, lca); (lca, y)
	return nil
}

// leaf 1-node with value that could not be included should be skipped

// aggregateInverse aggregates inverse intervals of form [x, y) where x > y.
func (ft *FPTree) aggregateInverse(ac *aggContext) (err error) {
	_ = "STUB: not implemented"
	// inverse interval: [min, y); [x, max]
	return nil
}

// First, we handle [x, max] part
// For this, we process the subtree rooted in the LCA of 0x000000... (all 0s) and x

// nothing to do

// node is fully included

// the node is below X

// Then we handle [min, y) part.
// For this, we process the subtree rooted in the LCA of y and 0xffffff... (all 1s)

// nothing to do

// node is fully included

// the node is at or after Y

// aggregateInterval aggregates an interval, updating the aggContext accordingly.
func (ft *FPTree) aggregateInterval(ac *aggContext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// startFromPrefix returns a SeqResult which begins with the first item that has the
// specified prefix.
func (ft *FPTree) startFromPrefix(ac *aggContext, p prefix) rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// nextFromPrefix return the first item that has the prefix p.
func (ft *FPTree) nextFromPrefix(ac *aggContext, p prefix) (rangesync.KeyBytes, error) {
	_ = "STUB: not implemented"
	return *new(rangesync.KeyBytes), nil
}

// FingerprintInterval performs a range fingerprint query with specified bounds and limit.
func (ft *FPTree) FingerprintInterval(x, y rangesync.KeyBytes, limit int) (fpr FPResult, err error) {
	_ = "STUB: not implemented"
	return *new(FPResult), nil
}

func (ft *FPTree) fingerprintInterval(
	x, y rangesync.KeyBytes,
	limit int,
	needNext bool,
) (fpr FPResult, err error) {
	_ = "STUB: not implemented"
	return *new(FPResult), nil
}

// The next item is only needed for splitting in case if easy split is not
// feasible, and it's better to avoid getting it as that may incur
// database access

// We apply limit after we have retrieved the next item

// FingerprintAll returns FingerprintResult for all items in the tree.
// Unlike FingerprintInterval, it is guaranteed not to query the underlying idStore, so it
// will never cause database access.
func (ft *FPTree) FingerprintAll() FPResult { _ = "STUB: not implemented"; return *new(FPResult) }

// easySplit splits an interval in two parts trying to do it in such way that the first
// part has close to limit items while not making any idStore queries so that the database
// is not accessed. If the split can't be done, which includes the situation where one of
// the sides has 0 items, easySplit returns errEasySplitFailed error.
// easySplit never fails for a tree with values.
func (ft *FPTree) easySplit(x, y rangesync.KeyBytes, limit int) (sr SplitResult, err error) {
	_ = "STUB: not implemented"
	return *new(SplitResult), nil
}

// need to get some items on both sides for the easy split to succeed

// It should not be possible to have ac.lastPrefix0 == nil or ac.lastPrefix == nil
// if both ac.count0 and ac.count are non-zero, b/c of how
// aggContext.maybeIncludeNode works

// ac.start / ac.end are only set in aggregateEdge which fails with
// errEasySplitFailed if easySplit is enabled, so we can ignore them here

// Next is only used during splitting itself, and thus not included

// Next is only used during splitting itself, and thus not included

// Split splits an interval in two parts.
func (ft *FPTree) Split(x, y rangesync.KeyBytes, limit int) (sr SplitResult, err error) {
	_ = "STUB: not implemented"
	return *new(SplitResult), nil
}

// dumpNode prints the node structure to the writer.
func (ft *FPTree) dumpNode(w io.Writer, idx nodeIndex, indent, dir string) {
	_ = "STUB: not implemented"
	return
}

// Dump prints the tree structure to the writer.
func (ft *FPTree) Dump(w io.Writer) { _ = "STUB: not implemented"; return }

// DumpToString returns the tree structure as a string.
func (ft *FPTree) DumpToString() string { _ = "STUB: not implemented"; return "" }

// Count returns the number of items in the tree.
func (ft *FPTree) Count() int { _ = "STUB: not implemented"; return 0 }

// EnableTrace enables or disables tracing for the tree.
func (ft *FPTree) EnableTrace(enable bool) { _ = "STUB: not implemented"; return }
