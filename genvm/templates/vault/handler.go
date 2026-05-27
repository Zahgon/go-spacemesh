package vault

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/registry"
)

// TemplateAddress is an address of the vault template.
var TemplateAddress core.Address

func init() {
	TemplateAddress[len(TemplateAddress)-1] = 4
}

// Register vault template.
func Register(reg *registry.Registry) { _ = "STUB: not implemented"; return }

type handler struct{}

// Parse is noop on vault template.
func (h *handler) Parse(method uint8, decoder *scale.Decoder) (core.ParseOutput, error) {
	_ = "STUB: not implemented"
	return *new(core.ParseOutput), nil
}

// New instantiates vault state.
func (h *handler) New(args any) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// InitialUnlockAmount is no longer used per SMIP-0002

// Load vault from state.
func (h *handler) Load(state []byte) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// Exec supports only MethodSpend.
func (h *handler) Exec(host core.Host, method uint8, args scale.Encodable) error {
	_ = "STUB: not implemented"
	return nil
}

// Args ...
func (h *handler) Args(method uint8) scale.Type { _ = "STUB: not implemented"; return *new(scale.Type) }
