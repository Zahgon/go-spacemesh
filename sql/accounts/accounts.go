package accounts

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

// Has the account in the database.
func Has(db sql.Executor, address types.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Latest latest account data for an address.
func Latest(db sql.Executor, address types.Address) (types.Account, error) {
	_ = "STUB: not implemented"
	return *new(types.Account), nil
}

// TODO(mafa): returning `sql.ErrNotFound` causes a bunch of tests to fail, some even panic
// this needs to be investigated and fixed
//
// if account.Address != address {
// 	return types.Account{}, sql.ErrNotFound
// }
// without this tests are failing not only assertions but are also panicking

// Get account data that was valid at the specified layer.
func Get(db sql.Executor, address types.Address, layer types.LayerID) (types.Account, error) {
	_ = "STUB: not implemented"
	return *new(types.Account), nil
}

// TODO(mafa): returning `sql.ErrNotFound` causes a bunch of tests to fail, some even panic
// this needs to be investigated and fixed
//
// if account.Address != address {
// 	return types.Account{}, sql.ErrNotFound
// }
// without this tests are failing not only assertions but are also panicking

// All returns all latest accounts.
func All(db sql.Executor) ([]*types.Account, error) { _ = "STUB: not implemented"; return nil, nil }

func Snapshot(db sql.Executor, layer types.LayerID) ([]*types.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update account state at a certain layer.
func Update(db sql.Executor, to *types.Account) error { _ = "STUB: not implemented"; return nil }

// Revert state after the layer.
func Revert(db sql.Executor, after types.LayerID) error { _ = "STUB: not implemented"; return nil }

func IterateAccountsOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(account *types.Account) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
