package certificates

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

func SetHareOutput(db sql.Executor, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

func SetHareOutputInvalid(db sql.Executor, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

func setHareOutput(db sql.Executor, lid types.LayerID, bid types.BlockID, valid bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHareOutput returns the block that's valid as hare output for the specified layer.
// If there are more than one valid blocks, return types.EmptyBlockID.
func GetHareOutput(db sql.Executor, lid types.LayerID) (types.BlockID, error) {
	_ = "STUB: not implemented"
	return *new(types.BlockID), nil
}

func FirstInEpoch(db sql.Executor, epoch types.EpochID) (types.BlockID, error) {
	_ = "STUB: not implemented"
	return *new(types.BlockID), nil
}

func Add(db sql.Executor, lid types.LayerID, cert *types.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

type CertValidity struct {
	Block types.BlockID
	Cert  *types.Certificate
	Valid bool
}

func Get(db sql.Executor, lid types.LayerID) ([]CertValidity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CertifiedBlock(db sql.Executor, lid types.LayerID) (types.BlockID, error) {
	_ = "STUB: not implemented"
	return *new(types.BlockID), nil
}

func SetValid(db sql.Executor, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

func SetInvalid(db sql.Executor, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteCertBefore(db sql.Executor, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}
