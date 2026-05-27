package fptree

import (
	"testing"
)

// checkNode checks that the tree node at the given index is correct and also recursively
// checks its children.
func checkNode(t *testing.T, ft *FPTree, idx nodeIndex, depth int) {
	_ = "STUB: not implemented"
	return
}

// CheckTree checks that the tree has correct structure.
func CheckTree(t *testing.T, ft *FPTree) { _ = "STUB: not implemented"; return }

// analyzeTreeNodeRefs checks that the reference counts in the node pool are correct.
func analyzeTreeNodeRefs(t *testing.T, np *nodePool, trees ...*FPTree) {
	_ = "STUB: not implemented"
	return
}

// AnalyzeTreeNodeRefs checks that the reference counts are correct for the given trees in
// their respective node pools.
func AnalyzeTreeNodeRefs(t *testing.T, trees ...*FPTree) {
	_ = "STUB: not implemented"

	// group trees by node pool they use
	return
}
