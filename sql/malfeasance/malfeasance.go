package malfeasance

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
	"github.com/spacemeshos/go-spacemesh/sql/marriage"
)

func AddProof(
	db sql.Executor,
	nodeID types.NodeID,
	marriageID *marriage.ID,
	proof []byte,
	domain int,
	received time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func SetMalicious(db sql.Executor, nodeID types.NodeID, marriageID marriage.ID, received time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func IsMalicious(db sql.Executor, nodeID types.NodeID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func Count(db sql.Executor) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func IterateOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(types.NodeID, []byte, int, time.Time) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// NodeIDProof returns the malfeasance proof and its domain for the given node ID. Returns sql.ErrNotFound if no proof
// for the given node ID exists. To return a proof for a marriage set use MarriageProof instead.
func NodeIDProof(db sql.Executor, nodeID types.NodeID) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// MarriageProof returns the malfeasance proof for the marriage set. Returns sql.ErrNotFound if no proof for the given
// marriage ID exists. To return a proof for a node ID use NodeIDProof instead.
func MarriageProof(db sql.Executor, marriageID marriage.ID) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
