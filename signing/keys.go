package signing

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
)

// PrivateKey is an alias to ed25519.PrivateKey.
type PrivateKey = ed25519.PrivateKey

// PrivateKeySize size of the private key in bytes.
const PrivateKeySize = ed25519.PrivateKeySize

// PublicKey is the type describing a public key.
type PublicKey struct {
	ed25519.PublicKey
}

func Public(priv PrivateKey) ed25519.PublicKey {
	_ = "STUB: not implemented"
	return *new(ed25519.PublicKey)
}

// NewPublicKey constructs a new public key instance from a byte array.
func NewPublicKey(pub []byte) *PublicKey { _ = "STUB: not implemented"; return nil }

// Bytes returns the public key as byte array.
func (p *PublicKey) Bytes() []byte {
	_ = "STUB: not implemented"
	// Prevent segfault if unset
	return nil
}

// String returns the public key as a hex representation string.
func (p *PublicKey) String() string { _ = "STUB: not implemented"; return "" }

const shortStringSize = 5

// ShortString returns a representative sub string.
func (p *PublicKey) ShortString() string { _ = "STUB: not implemented"; return "" }

// Equals returns true if the public keys are equal.
func (p *PublicKey) Equals(o *PublicKey) bool { _ = "STUB: not implemented"; return false }
