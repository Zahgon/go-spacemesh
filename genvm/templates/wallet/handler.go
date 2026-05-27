package wallet

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/registry"
)

func init() {
	TemplateAddress[len(TemplateAddress)-1] = 1
}

// Register Wallet template.
func Register(registry *registry.Registry) { _ = "STUB: not implemented"; return }

var (
	_ core.Handler = (*handler)(nil)
	// TemplateAddress is an address of the Wallet template.
	TemplateAddress core.Address
)

type handler struct{}

// Parse header and arguments.
func (*handler) Parse(method uint8, decoder *scale.Decoder) (output core.ParseOutput, err error) {
	_ = "STUB: not implemented"
	return *new(core.ParseOutput), nil
}

// New instatiates single sig wallet with spawn arguments.
func (*handler) New(args any) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// Load single sig wallet from stored state.
func (*handler) Load(state []byte) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// Exec spawn or spend based on the method selector.
func (*handler) Exec(host core.Host, method uint8, args scale.Encodable) error {
	_ = "STUB: not implemented"
	return nil
}

// Args ...
func (h *handler) Args(method uint8) scale.Type { _ = "STUB: not implemented"; return *new(scale.Type) }
