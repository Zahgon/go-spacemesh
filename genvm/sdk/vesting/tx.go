package vesting

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk/multisig"
)

type Aggregator = multisig.Aggregator

var (
	NewAggregator = multisig.NewAggregator

	SelfSpawn = multisig.SelfSpawn
	Spawn     = multisig.Spawn
	Spend     = multisig.Spend
)

// DrainVault creates drain vault transaction.
func DrainVault(
	ref uint8,
	pk ed25519.PrivateKey,
	principal, vault, receiver types.Address,
	amount uint64,
	nonce core.Nonce,
	opts ...sdk.Opt,
) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}
