package sql

import (
	"io/fs"
	"regexp"

	"go.uber.org/zap"
)

// MigrationList denotes a list of migrations.
type MigrationList []Migration

// AddMigration adds a Migration to the MigrationList, overriding the migration with the
// same order number if it already exists. The function returns updated migration list.
// The state of the original migration list is undefined after calling this function.
func (l MigrationList) AddMigration(migration Migration) MigrationList {
	_ = "STUB: not implemented"
	return *new(MigrationList)
}

// Version returns database version for the specified migration list.
func (l MigrationList) Version() int { _ = "STUB: not implemented"; return 0 }

type sqlMigration struct {
	order   int
	name    string
	content string
}

var sqlCommentRx = regexp.MustCompile(`(?m)--.*$`)

func (m *sqlMigration) Apply(db Executor, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: use more advanced approach to split the SQL script
// into commands

func (m *sqlMigration) Name() string { _ = "STUB: not implemented"; return "" }

func (m *sqlMigration) Order() int { _ = "STUB: not implemented"; return 0 }

func (sqlMigration) Rollback() error {
	_ = "STUB: not implemented"
	// handled by the DB itself
	return nil
}

func version(db Executor) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func LoadSQLMigrations(fsys fs.FS) (MigrationList, error) {
	_ = "STUB: not implemented"
	return *new(MigrationList), nil
}
