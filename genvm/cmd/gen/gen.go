package gen

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"

	"github.com/spacemeshos/go-spacemesh/genvm/templates/multisig"
	"github.com/spacemeshos/go-spacemesh/genvm/templates/vault"
)

type Input struct {
	Keys     []string `json:"keys"`
	Required int      `json:"required"`
	Total    float64  `json:"total"`
}

type Output struct {
	Debug struct {
		VestingAddress string
		VestingArgs    multisig.SpawnArguments
		VaultAddress   string
		VaultArgs      vault.SpawnArguments
	} `json:"-"`

	Address string `json:"address"`
	Balance uint64 `json:"balance"`
}

func Generate(input Input) Output { _ = "STUB: not implemented"; return *new(Output) }

func decodeHexKey(data []byte) [ed25519.PublicKeySize]byte { _ = "STUB: not implemented"; return nil }

func must(err error) { _ = "STUB: not implemented"; return }
