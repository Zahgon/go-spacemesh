package multisig

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk"
	"github.com/spacemeshos/go-spacemesh/genvm/templates/multisig"
)

func encode(fields ...scale.Encodable) []byte { _ = "STUB: not implemented"; return nil }

func NewAggregator(unsigned []byte) *Aggregator { _ = "STUB: not implemented"; return nil }

// Aggregator is a signature accumulator.
type Aggregator struct {
	unsigned []byte
	parts    map[uint8]multisig.Part
}

// Add signature parts to the accumulator.
func (tx *Aggregator) Add(parts ...multisig.Part) { _ = "STUB: not implemented"; return }

// Part returns signature part from ref public key.
func (tx *Aggregator) Part(ref uint8) *multisig.Part { _ = "STUB: not implemented"; return nil }

// Raw returns full raw transaction including payload and signature.
func (tx *Aggregator) Raw() []byte { _ = "STUB: not implemented"; return nil }

// buf.Write is not expected to fail

// SelfSpawn returns accumulator for self-spawn transaction.
func SelfSpawn(
	ref uint8,
	pk ed25519.PrivateKey,
	template types.Address,
	required uint8,
	pubs []ed25519.PublicKey,
	nonce core.Nonce,
	opts ...sdk.Opt,
) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}

// Spawn returns accumulator for spawn transaction.
func Spawn(
	ref uint8,
	pk ed25519.PrivateKey,
	principal, template types.Address,
	args scale.Encodable,
	nonce core.Nonce,
	opts ...sdk.Opt,
) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}

// Spend creates spend transaction.
func Spend(
	ref uint8,
	pk ed25519.PrivateKey,
	principal, to types.Address,
	amount uint64,
	nonce types.Nonce,
	opts ...sdk.Opt,
) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}
