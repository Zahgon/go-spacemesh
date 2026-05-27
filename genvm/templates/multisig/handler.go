package multisig

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/registry"
)

func init() {
	TemplateAddress[len(TemplateAddress)-1] = 2
}

// Register template.
func Register(registry *registry.Registry) { _ = "STUB: not implemented"; return }

var (
	_               core.Handler = (*handler)(nil)
	TemplateAddress core.Address
)

// NewHandler instantiates multisig handler with a particular configuration.
func NewHandler() core.Handler { _ = "STUB: not implemented"; return *new(core.Handler) }

type handler struct{}

// Parse header and arguments.
func (h *handler) Parse(method uint8, decoder *scale.Decoder) (output core.ParseOutput, err error) {
	_ = "STUB: not implemented"
	return *new(core.ParseOutput), nil
}

// New instantiates k-multisig instance.
func (h *handler) New(args any) (core.Template, error) {
	_ = "STUB: not implemented"
	return *new(core.Template), nil
}

// Load k-multisig instance from stored state.
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

// SpendTemplate interface for the template that support Spend method.
type SpendTemplate interface {
	Spend(core.Host, *SpendArguments) error
}
