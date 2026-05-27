package beacon

const (
	up   = uint(1)
	down = uint(0)
)

func encodeVotes(currentRound allVotes, firstRound proposalList) []byte {
	_ = "STUB: not implemented"
	return nil
}

// no need to set invalid votes as big.Int will have unset bits
// return the default value 0

func decodeVotes(votesBitVector []byte, firstRound proposalList) allVotes {
	_ = "STUB: not implemented"
	return *new(allVotes)
}
