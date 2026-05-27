package registry

import (
	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

// New creates Registry instance.
func New() *Registry { _ = "STUB: not implemented"; return nil }

// Registry stores mapping from address to template handler.
type Registry struct {
	templates map[core.Address]core.Handler
}

// Get template handler for the address if it exists.
func (r *Registry) Get(address core.Address) core.Handler {
	_ = "STUB: not implemented"
	return *new(core.Handler)
}

// Register handler for the address. Panics if address is already taken.
func (r *Registry) Register(address core.Address, handler core.Handler) {
	_ = "STUB: not implemented"
	return
}
