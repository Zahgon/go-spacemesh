package wallet

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

// New returns Wallet instance with SpawnArguments.
func New(args *SpawnArguments) *Wallet { _ = "STUB: not implemented"; return nil }

//go:generate scalegen

// Wallet is a single-key wallet.
type Wallet struct {
	PublicKey core.PublicKey
}

// MaxSpend returns amount specified in the SpendArguments for Spend method.
func (s *Wallet) MaxSpend(method uint8, args any) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Verify that transaction is signed by the owner of the PublicKey using ed25519.
func (s *Wallet) Verify(host core.Host, raw []byte, dec *scale.Decoder) bool {
	_ = "STUB: not implemented"
	return false
}

// Spend transfers an amount to the address specified in SpendArguments.
func (s *Wallet) Spend(host core.Host, args *SpendArguments) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Wallet) BaseGas(method uint8) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Wallet) LoadGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Wallet) ExecGas(method uint8) uint64 { _ = "STUB: not implemented"; return 0 }
