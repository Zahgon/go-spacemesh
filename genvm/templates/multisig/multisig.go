package multisig

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

//go:generate scalegen

// MultiSig K/N template.
type MultiSig struct {
	Required   uint8
	PublicKeys []core.PublicKey `scale:"max=10"`
}

func (ms *MultiSig) BaseGas(method uint8) uint64 { _ = "STUB: not implemented"; return 0 }

func (ms *MultiSig) LoadGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms *MultiSig) ExecGas(method uint8) uint64 { _ = "STUB: not implemented"; return 0 }

// MaxSpend returns amount specified in the SpendArguments.
func (ms *MultiSig) MaxSpend(method uint8, args any) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Verify that transaction is signed has k valid signatures.
func (ms *MultiSig) Verify(host core.Host, raw []byte, dec *scale.Decoder) bool {
	_ = "STUB: not implemented"
	return false
}

// Spend transfers an amount to the address specified in SpendArguments.
func (ms *MultiSig) Spend(host core.Host, args *SpendArguments) error {
	_ = "STUB: not implemented"
	return nil
}
