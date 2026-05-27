package tortoise

import (
	"go.uber.org/zap"
)

const (
	support sign = 1
	against sign = -1
	abstain sign = 0

	neutral = abstain
)

type sign int8

func (a sign) String() string { _ = "STUB: not implemented"; return "" }

type voteReason string

func (v voteReason) String() string { _ = "STUB: not implemented"; return "" }

const (
	reasonHareOutput     voteReason = "hare"
	reasonValidity       voteReason = "validity"
	reasonLocalThreshold voteReason = "local_threshold"
	reasonCoinflip       voteReason = "coinflip"
	reasonMissingData    voteReason = "missing data"
)

func verifyLayer(logger *zap.Logger, blocks []*blockInfo, getDecision func(*blockInfo) sign) (bool, bool) {
	_ = "STUB: not implemented"
	// order blocks by height in ascending order
	// if there is a support before any abstain
	// and a previous height is lower than the current one
	// the layer is verified
	//
	// it will modify original slice
	return false, false
}

// all blocks with the same height should be finalized

func zapBlocks(blocks []*blockInfo) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }
