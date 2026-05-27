package vault

import (
	"errors"

	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

var (
	// ErrNotOwner is raised if Spend is not executed by a principal that matches owner.
	ErrNotOwner = errors.New("vault: not an owner")
	// ErrAmountNotAvailable is raised if Spend overlows available amount (see method with the same name).
	ErrAmountNotAvailable = errors.New("vault: amount not available")
	// ErrMisconfigured is raised if a Vault account is misconfigured.
	ErrMisconfigured = errors.New("vault: account is misconfigured")
)

const (
	VAULT_STATE_SIZE = core.ACCOUNT_HEADER_SIZE + 56
	DRAINED_SIZE     = 8
)

//go:generate scalegen

type Vault struct {
	Owner               core.Address
	TotalAmount         uint64
	InitialUnlockAmount uint64
	VestingStart        core.LayerID
	VestingEnd          core.LayerID
}

func (v *Vault) isOwner(address core.Address) bool { _ = "STUB: not implemented"; return false }

func (v *Vault) Vested(lid core.LayerID) uint64 { _ = "STUB: not implemented"; return 0 }

// Note: VestingStart may equal VestingEnd but division by zero is not possible here since in this case
// one of the first two conditionals above would have been triggered and the method would already have
// returned.

// Spend transaction.
func (v *Vault) Spend(host core.Host, to core.Address, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// sanity checks

// cannot vest more than initial endowment

// account must contain at least unvested portion of initial endowment

// current account balance minus unvested portion of initial endowment equals unspent, vested coins
// plus coins received. in simpler pseudocode:
//   unvested_portion = v.TotalAmount - vested
//   spendable_portion = host.balance() - unvested_portion
//   if amount > spendable_portion { ... }

// MaxSpend is noop for this template type, principal of this account type can't submit transactions.
func (v *Vault) MaxSpend(uint8, any) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *Vault) BaseGas(uint8) uint64 { _ = "STUB: not implemented"; return 0 }

func (v *Vault) LoadGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (v *Vault) ExecGas(uint8) uint64 {
	_ = "STUB: not implemented"

	// Verify always returns false.
	return 0
}

func (v *Vault) Verify(core.Host, []byte, *scale.Decoder) bool {
	_ = "STUB: not implemented"
	return false
}
