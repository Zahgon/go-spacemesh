package core

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// Context serves 2 purposes:
// - maintains changes to the system state, that will be applied only after successful execution
// - accumulates set of reusable objects and data.
type Context struct {
	Registry HandlerRegistry
	Loader   AccountLoader

	// LayerID of the block.
	LayerID   LayerID
	GenesisID types.Hash20

	PrincipalHandler  Handler
	PrincipalTemplate Template
	PrincipalAccount  Account

	ParseOutput ParseOutput
	Gas         struct {
		BaseGas  uint64
		FixedGas uint64
	}
	Header Header
	Args   scale.Encodable

	// consumed is in gas units and will be used
	consumed uint64
	// fee is in coins units
	fee uint64
	// an amount transferred to other accounts
	transferred uint64

	touched []Address
	changed map[Address]*Account
}

// Principal returns address of the account that signed transaction.
func (c *Context) Principal() Address { _ = "STUB: not implemented"; return *new(Address) }

// Layer returns block layer id.
func (c *Context) Layer() LayerID {
	_ = "STUB: not implemented"

	// GetGenesisID returns genesis id.
	return *new(LayerID)
}

func (c *Context) GetGenesisID() Hash20 {
	_ = "STUB: not implemented"

	// Balance returns the account balance.
	return *new(Hash20)
}

func (c *Context) Balance() uint64 { _ = "STUB: not implemented"; return 0 }

// Template of the principal account.
func (c *Context) Template() Template { _ = "STUB: not implemented"; return *new(Template) }

// Handler of the principal account.
func (c *Context) Handler() Handler {
	_ = "STUB: not implemented"
	return *

	// Spawn account.
	new(Handler)
}

func (c *Context) Spawn(args scale.Encodable) error { _ = "STUB: not implemented"; return nil }

// Transfer amount to the address after validation passes.
func (c *Context) Transfer(to Address, amount uint64) error { _ = "STUB: not implemented"; return nil }

func (c *Context) transfer(from *Account, to Address, amount, max uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// noop. only gas is consumed

// Relay call to the remote account.
func (c *Context) Relay(remoteTemplate, address Address, call func(Host) error) error {
	_ = "STUB: not implemented"
	return nil
}

// ideally such changes would be serialized once for the whole block execution
// but it requires more changes in the cache, so can be done as an optimization
// if it proves meaningful (most likely wont)

// Consume gas from the account after validation passes.
func (c *Context) Consume(gas uint64) (err error) { _ = "STUB: not implemented"; return nil }

// Apply is executed if transaction was consumed.
func (c *Context) Apply(updater AccountUpdater) error { _ = "STUB: not implemented"; return nil }

// Consumed gas.
func (c *Context) Consumed() uint64 {
	_ = "STUB: not implemented"

	// Fee computed from consumed gas.
	return 0
}

func (c *Context) Fee() uint64 {
	_ = "STUB: not implemented"

	// Updated list of addresses.
	return 0
}

func (c *Context) Updated() []types.Address { _ = "STUB: not implemented"; return nil }

func (c *Context) load(address types.Address) (*Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Context) change(account *Account) { _ = "STUB: not implemented"; return }

// RemoteContext ...
type RemoteContext struct {
	*Context
	remote   *Account
	handler  Handler
	template Template
}

// Balance returns the remote account balance.
func (r *RemoteContext) Balance() uint64 { _ = "STUB: not implemented"; return 0 }

// Template ...
func (r *RemoteContext) Template() Template {
	_ = "STUB: not implemented"

	// Handler ...
	return *new(Template)
}

func (r *RemoteContext) Handler() Handler {
	_ = "STUB: not implemented"

	// Transfer ...
	return *new(Handler)
}

func (r *RemoteContext) Transfer(to Address, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}
