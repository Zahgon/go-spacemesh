package malsync

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/sql"
)

func LegacySyncState(db sql.Executor) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func SyncState(db sql.Executor) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func getSyncState(db sql.Executor, version int64) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func updateSyncState(db sql.Executor, version, ts int64) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateLegacySyncState(db sql.Executor, timestamp time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateSyncState(db sql.Executor, timestamp time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func Clear(db sql.Executor) error { _ = "STUB: not implemented"; return nil }
