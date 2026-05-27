package nipost

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

func encodeNipostChallenge(nodeID types.NodeID, ch *types.NIPostChallenge) func(stmt *sql.Statement) {
	_ = "STUB: not implemented"
	return nil
}

func AddChallenge(db sql.Executor, nodeID types.NodeID, ch *types.NIPostChallenge) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateChallenge(db sql.Executor, nodeID types.NodeID, ch *types.NIPostChallenge) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveChallenge(db sql.Executor, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func Challenge(db sql.Executor, nodeID types.NodeID) (*types.NIPostChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdatePoetProofRef(
	db sql.Executor,
	nodeID types.NodeID,
	ref types.PoetProofRef,
	membership *types.MerkleProof,
) error {
	_ = "STUB: not implemented"
	return nil
}

func PoetProofRef(db sql.Executor, nodeID types.NodeID) (types.PoetProofRef, *types.MerkleProof, error) {
	_ = "STUB: not implemented"
	return *new(types.PoetProofRef), nil, nil
}
