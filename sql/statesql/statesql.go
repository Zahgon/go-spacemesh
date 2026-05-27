package statesql

import (
	"embed"
	"testing"

	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate go run ../schemagen -dbtype state -output schema/schema.sql

//go:embed schema/schema.sql
var schemaScript string

//go:embed schema/migrations/*.sql
var migrations embed.FS

type database struct {
	sql.Database
}

var _ sql.StateDatabase = &database{}

func (db *database) IsStateDatabase() {
	_ = "STUB: not implemented"

	// Schema returns the schema for the state database.
	return
}

func Schema(inCodeMigrations ...sql.Migration) (*sql.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: coded state migrations can be added here
// They can be a part of this localsql package

// Open opens a state database.
func Open(uri string, opts ...sql.Opt) (sql.StateDatabase, error) {
	_ = "STUB: not implemented"
	return *new(sql.StateDatabase), nil
}

// Open opens an in-memory state database.
func InMemory(opts ...sql.Opt) sql.StateDatabase {
	_ = "STUB: not implemented"
	return *new(sql.StateDatabase)
}

// InMemoryTest returns an in-mem database for testing and ensures database is closed during `tb.Cleanup`.
func InMemoryTest(tb testing.TB, opts ...sql.Opt) sql.StateDatabase {
	_ = "STUB: not implemented"
	return *new(sql.StateDatabase)
}
