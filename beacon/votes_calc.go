package beacon

import (
	"math/big"

	"go.uber.org/zap"
)

func calcVotes(logger *zap.Logger, theta *big.Float, s *state) (allVotes, proposalList) {
	_ = "STUB: not implemented"
	return *new(allVotes), *new(proposalList)
}

func votingThreshold(theta *big.Float, epochWeight uint64) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func tallyUndecided(votes *allVotes, undecided proposalList, coinFlip bool) {
	_ = "STUB: not implemented"
	return
}
