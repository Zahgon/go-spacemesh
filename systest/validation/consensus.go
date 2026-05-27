package validation

import (
	"context"
	"sync"

	"github.com/spacemeshos/go-spacemesh/systest/cluster"
)

type ConsensusData struct {
	Consensus, State []byte
}

func getConsensusData(ctx context.Context, distance int, node *cluster.NodeClient) *ConsensusData {
	_ = "STUB: not implemented"
	return nil
}

// empty strings are always in consensus

// failMinority should increment number of failures for groups smaller than the largest one
// if there are several groups of the same size they all should be considered as failed.
func failMinority(failures []int, groups map[string][]int) { _ = "STUB: not implemented"; return }

// nolint:copyloopvar

func Consensus(c *cluster.Cluster, tolerate, distance int) Validation {
	_ = "STUB: not implemented"
	return *new(Validation)
}

func NewConsensusValidation(size, tolerate int) *ConsensusValidation {
	_ = "STUB: not implemented"
	return nil
}

type ConsensusValidation struct {
	failures []int
	tolerate int
}

func (c *ConsensusValidation) Next() *ConsensusValidationIteration {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConsensusValidation) Complete(iter *ConsensusValidationIteration) error {
	_ = "STUB: not implemented"
	return nil
}

type ConsensusValidationIteration struct {
	mu        sync.Mutex
	all       []*ConsensusData
	consensus map[string][]int
	state     map[string][]int
}

func (iter *ConsensusValidationIteration) OnData(id int, data *ConsensusData) {
	_ = "STUB: not implemented"
	return
}
