package atxsync

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type EpochSyncState = map[types.ATXID]*IDSyncState

func fromDatabase(tries int) *IDSyncState { _ = "STUB: not implemented"; return nil }

// IDSyncState tracks how many times we tried to sync an ATX.
type IDSyncState struct {
	Tries          int
	persisted      bool
	persistedTries int
}

// TriesPersisted returns true if the number of tries is equal to the number of persisted tries.
func (s *IDSyncState) TriesPersisted() bool { _ = "STUB: not implemented"; return false }

// SetPersisted sets the number of tries to the number of persisted tries.
func (s *IDSyncState) setPersisted() { _ = "STUB: not implemented"; return }

func GetSyncState(db sql.Executor, epoch types.EpochID) (EpochSyncState, error) {
	_ = "STUB: not implemented"
	return *new(EpochSyncState), nil
}

func SaveSyncState(db sql.Executor, epoch types.EpochID, states EpochSyncState, max int) error {
	_ = "STUB: not implemented"
	return nil
}

func SaveRequest(db sql.Executor, epoch types.EpochID, timestamp time.Time, total, downloaded int64) error {
	_ = "STUB: not implemented"
	return nil
}

func GetRequest(db sql.Executor, epoch types.EpochID) (time.Time, int64, int64, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), 0, 0, nil
}

func Clear(db sql.Executor) error { _ = "STUB: not implemented"; return nil }
