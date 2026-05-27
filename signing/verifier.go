package signing

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
)

type edVerifierOption struct {
	prefix []byte
}

// VerifierOptionFunc to modify verifier.
type VerifierOptionFunc func(*edVerifierOption)

// WithVerifierPrefix sets the prefix used by PubKeyVerifier. This usually is the Network ID.
func WithVerifierPrefix(prefix []byte) VerifierOptionFunc {
	_ = "STUB: not implemented"
	return *new(VerifierOptionFunc)
}

// EdVerifier extracts public keys from signatures.
type EdVerifier struct {
	prefix []byte
}

func NewEdVerifier(opts ...VerifierOptionFunc) *EdVerifier { _ = "STUB: not implemented"; return nil }

// Verify verifies that a signature matches public key and message.
func (es *EdVerifier) Verify(d Domain, nodeID types.NodeID, m []byte, sig types.EdSignature) bool {
	_ = "STUB: not implemented"
	return false
}
