package wallet

import (
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk"
	"github.com/spacemeshos/go-spacemesh/signing"
)

func encode(fields ...scale.Encodable) []byte { _ = "STUB: not implemented"; return nil }

// SelfSpawn creates a self-spawn transaction.
func SelfSpawn(pk signing.PrivateKey, nonce core.Nonce, opts ...sdk.Opt) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Spawn creates a spawn transaction.
func Spawn(
	pk signing.PrivateKey,
	template core.Address,
	args scale.Encodable,
	nonce core.Nonce,
	opts ...sdk.Opt,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

// note that principal is computed from pk

// Spend creates spend transaction.
func Spend(pk signing.PrivateKey, to types.Address, amount uint64, nonce types.Nonce, opts ...sdk.Opt) []byte {
	_ = "STUB: not implemented"
	return nil
}
