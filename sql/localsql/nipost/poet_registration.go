package nipost

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type PoETRegistration struct {
	ChallengeHash types.Hash32
	Address       string
	RoundID       string
	RoundEnd      time.Time
}

func AddPoetRegistration(
	db sql.Executor,
	nodeID types.NodeID,
	registration PoETRegistration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ClearPoetRegistrations(db sql.Executor, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func PoetRegistrations(db sql.Executor, nodeID types.NodeID) ([]PoETRegistration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
