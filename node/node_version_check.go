package node

import (
	"github.com/spacemeshos/go-spacemesh/config"
)

func (app *App) verifyVersionUpgrades() error { _ = "STUB: not implemented"; return nil }

// v1.5 requires going through v1.4 first as it removed in-code migrations 1 - 3.
func verifyLocalDbMigrations(cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

// if local DB doesnt exist, it's a fresh db and doesn't require in-code migrations
