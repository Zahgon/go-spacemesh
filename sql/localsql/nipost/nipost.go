package nipost

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type NIPostState struct {
	*types.NIPost

	NumUnits uint32
	VRFNonce types.VRFPostIndex
}

func AddNIPost(db sql.Executor, nodeID types.NodeID, nipost *NIPostState) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveNIPost(db sql.Executor, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func NIPost(db sql.Executor, nodeID types.NodeID) (*NIPostState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
