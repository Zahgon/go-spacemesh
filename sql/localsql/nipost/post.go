package nipost

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type Post struct {
	Nonce     uint32
	Indices   []byte
	Pow       uint64
	Challenge []byte

	NumUnits      uint32
	CommitmentATX types.ATXID
	VRFNonce      types.VRFPostIndex
}

func AddPost(db sql.Executor, nodeID types.NodeID, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func RemovePost(db sql.Executor, nodeID types.NodeID) error { _ = "STUB: not implemented"; return nil }

func GetPost(db sql.Executor, nodeID types.NodeID) (*Post, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
