package marriage

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

type ID int

type Info struct {
	ID            ID
	NodeID        types.NodeID
	ATX           types.ATXID
	MarriageIndex int
	Target        types.NodeID
	Signature     types.EdSignature
}

// NewID returns a new unique ID for a marriage. The ID is unique across all marriages.
// Use within a transaction to ensure that the returned ID is not used by another routine.
func NewID(db sql.Executor) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

// Add inserts a nodeID to the marriages table for the given MarriageID.
// If the nodeID already exists in the table, it is updated.
// Updates cannot change the MarriageID.
func Add(db sql.Executor, marriage Info) error { _ = "STUB: not implemented"; return nil }

func UpdateMarriageID(db sql.Executor, oldID, newID ID) error {
	_ = "STUB: not implemented"
	return nil
}

func FindIDByNodeID(db sql.Executor, nodeID types.NodeID) (ID, error) {
	_ = "STUB: not implemented"
	return *new(ID), nil
}

func FindByNodeID(db sql.Executor, nodeID types.NodeID) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func MarriageATXs(db sql.Executor, id ID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NodeIDsByID(db sql.Executor, id ID) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IterateOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(Info) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
