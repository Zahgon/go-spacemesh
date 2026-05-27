package beacons

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// Get gets a beacon for a given epoch.
func Get(db sql.Executor, epoch types.EpochID) (beacon types.Beacon, err error) {
	_ = "STUB: not implemented"
	return *new(types.Beacon), nil
}

// Add adds a beacon for a given epoch.
func Add(db sql.Executor, epoch types.EpochID, beacon types.Beacon) error {
	_ = "STUB: not implemented"
	return nil
}

func Set(db sql.Executor, epoch types.EpochID, beacon types.Beacon) error {
	_ = "STUB: not implemented"
	return nil
}
