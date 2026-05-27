package txs

import (
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

const (
	MinTXGas = uint64(1) // gas required for the most basic transaction
)

type item struct {
	*NanoTX

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type priorityQueue []*item

// Len implements head.Interface.
func (pq priorityQueue) Len() int {
	_ = "STUB: not implemented"

	// Less implements head.Interface.
	return 0
}

func (pq priorityQueue) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// We want Pop to give us the highest, not lowest, fee, so we use greater than here.
	return false
}

// if fees are equal, we want the older tx first

// if fees and timestamps are equal, we want the tx with the lower ID first

// Swap implements head.Interface.
func (pq priorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Push implements head.Interface.
func (pq *priorityQueue) Push(i any) { _ = "STUB: not implemented"; return }

// Pop implements head.Interface.
func (pq *priorityQueue) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// avoid memory leak
// for safety

// update modifies the fee and value of an item in the queue.
func (pq *priorityQueue) update(it *item, ntx *NanoTX) { _ = "STUB: not implemented"; return }

// mempoolIterator holds the best transaction from the conservative state mempool.
// Not thread-safe.
type mempoolIterator struct {
	logger       *zap.Logger
	gasRemaining uint64
	pq           priorityQueue
	txs          map[types.Address][]*NanoTX
}

// newMempoolIterator builds and returns a mempoolIterator.
func newMempoolIterator(logger *zap.Logger, cs conStateCache, gasLimit uint64) *mempoolIterator {
	_ = "STUB: not implemented"
	return nil
}

func (mi *mempoolIterator) buildPQ() { _ = "STUB: not implemented"; return }

func (mi *mempoolIterator) getNext(addr types.Address) *NanoTX {
	_ = "STUB: not implemented"
	return nil
}

func (mi *mempoolIterator) pop() *NanoTX { _ = "STUB: not implemented"; return nil }

// the first item in priority queue is always the item to be popped with the heap

// remove all txs for this principal since we cannot fulfill the lowest nonce for this principal

// updating the item (for the same address) in the heap is less expensive than a pop followed by a push.

// PopAll returns all the transaction in the mempoolIterator.
func (mi *mempoolIterator) PopAll() ([]*NanoTX, map[types.Address][]*NanoTX) {
	_ = "STUB: not implemented"
	return nil, nil
}
