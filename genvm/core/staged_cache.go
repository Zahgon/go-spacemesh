package core

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type DBLoader struct {
	sql.Executor
}

func (db DBLoader) Get(address types.Address) (types.Account, error) {
	_ = "STUB: not implemented"
	return *new(types.Account), nil
}

// NewStagedCache returns instance of the staged cache.
func NewStagedCache(loader AccountLoader) *StagedCache { _ = "STUB: not implemented"; return nil }

// StagedCache is a passthrough cache for accounts state and enforces order for updated accounts.
type StagedCache struct {
	loader AccountLoader
	// list of changed accounts. preserving order
	touched []Address
	cache   map[Address]stagedAccount
}

// Get a copy of the Account state for the address.
func (ss *StagedCache) Get(address Address) (Account, error) {
	_ = "STUB: not implemented"
	return *new(Account), nil
}

// Update cache with a copy of the account state.
func (ss *StagedCache) Update(account Account) error { _ = "STUB: not implemented"; return nil }

// IterateChanged accounts in the order they were updated.
func (ss *StagedCache) IterateChanged(f func(*Account) bool) { _ = "STUB: not implemented"; return }

type stagedAccount struct {
	Account
	Changed bool
}
