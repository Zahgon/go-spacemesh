package internal

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	localDbFile = "local.sql"

	keyDir                  = "identities"
	supervisedIDKeyFileName = "local.key"
)

func MergeDBs(ctx context.Context, dbLog *zap.Logger, from, to string) error {
	_ = "STUB: not implemented"
	// Open the target database
	return nil
}

// target database does not exist, create it

// target database exists, check if there is at least one key in the target key directory
// not named supervisedIDKeyFileName

// Open the source database

// check for name collisions

// skip files that are not identity files

// skip files that are not identity files

// copy files from `from` to `to`

// skip subdirectories and files in them

// skip files that are not identity files

func openDB(dbLog *zap.Logger, path string) (sql.LocalDatabase, error) {
	_ = "STUB: not implemented"
	return *new(sql.LocalDatabase), nil
}

func checkIdentities(dbLog *zap.Logger, path string) error { _ = "STUB: not implemented"; return nil }
