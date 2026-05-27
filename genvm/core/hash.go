package core

import (
	"github.com/spacemeshos/go-scale"
)

func SigningBody(genesis, tx []byte) []byte { _ = "STUB: not implemented"; return nil }

// ComputePrincipal address as the last 20 bytes from blake3(scale(template || args)).
func ComputePrincipal(template Address, args scale.Encodable) Address {
	_ = "STUB: not implemented"
	return *new(Address)
}
