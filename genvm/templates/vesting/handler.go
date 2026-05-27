package vesting

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/registry"
)

var TemplateAddress core.Address

func init() {
	TemplateAddress[len(TemplateAddress)-1] = 3
}

// MethodDrainVault is used to relay a call to drain a vault.
const MethodDrainVault = 17

// Register vesting templates.
func Register(reg *registry.Registry) { _ = "STUB: not implemented"; return }

type handler struct {
	multisig core.Handler
}

// Parse header and arguments.
func (h *handler) Parse(method uint8, decoder *scale.Decoder) (output core.ParseOutput, err error) {
	_ = "STUB: not implemented"
	return *new(core.ParseOutput), nil
}

// New instantiates vesting state, note that the state is the same as multisig.
// The difference is that vesting supports one more transaction type.
func (h *handler) New(args any) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// Load instnatiates vesting state from stored state. See comment on New.
func (h *handler) Load(state []byte) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// Exec spawn or spend based on the method selector.
func (h *handler) Exec(host core.Host, method uint8, args scale.Encodable) error {
	_ = "STUB: not implemented"
	return nil
}

// Args ...
func (h *handler) Args(method uint8) scale.Type { _ = "STUB: not implemented"; return *new(scale.Type) }
