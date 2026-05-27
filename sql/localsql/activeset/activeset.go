package activeset

import (
	"math"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type Kind uint8

const (
	Tortoise Kind = iota
	Hare
)

// disable verification for max number of atxs as it is performed before calling functions in this module.
const maxAtxs = math.MaxUint32

// Add adds an activeset with weight to the database.
// It is expected to return error if activeset is not unique per epoch.
func Add(
	db sql.Executor,
	kind Kind,
	epoch types.EpochID,
	id types.Hash32,
	weight uint64,
	set []types.ATXID,
) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(db sql.Executor, kind Kind, epoch types.EpochID) (types.Hash32, uint64, []types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), 0, nil, nil
}
