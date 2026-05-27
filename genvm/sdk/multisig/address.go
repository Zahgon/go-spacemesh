package multisig

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

// Address compute account address based on the size of the public keys.
func Address(template core.Address, required uint8, pubs ...[]byte) types.Address {
	_ = "STUB: not implemented"
	return *new(types.Address)
}
