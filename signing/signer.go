package signing

import (
	"io"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type Domain byte

const (
	ATX Domain = 0

	PROPOSAL = 1
	BALLOT   = 2
	HARE     = 3
	POET     = 4
	MARRIAGE = 5

	BEACON_FIRST_MSG    = 10
	BEACON_FOLLOWUP_MSG = 11
)

// String returns the string representation of a domain.
func (d Domain) String() string { _ = "STUB: not implemented"; return "" }

type edSignerOption struct {
	priv   PrivateKey
	file   string
	prefix []byte
}

// EdSignerOptionFunc modifies EdSigner.
type EdSignerOptionFunc func(*edSignerOption) error

// WithPrefix sets the prefix used by EdSigner. This usually is the Network ID.
func WithPrefix(prefix []byte) EdSignerOptionFunc {
	_ = "STUB: not implemented"
	return *new(EdSignerOptionFunc)
}

// ToFile writes the private key to a file after creation.
func ToFile(path string) EdSignerOptionFunc {
	_ = "STUB: not implemented"
	return *new(EdSignerOptionFunc)
}

// FromFile loads the private key from a file.
func FromFile(path string) EdSignerOptionFunc {
	_ = "STUB: not implemented"
	return *new(EdSignerOptionFunc)
}

// read hex data from file

// WithPrivateKey sets the private key used by EdSigner.
func WithPrivateKey(priv PrivateKey) EdSignerOptionFunc {
	_ = "STUB: not implemented"
	return *new(EdSignerOptionFunc)
}

// WithKeyFromRand sets the private key used by EdSigner using predictable randomness source.
func WithKeyFromRand(rand io.Reader) EdSignerOptionFunc {
	_ = "STUB: not implemented"
	return *new(EdSignerOptionFunc)
}

// EdSigner represents an ED25519 signer.
type EdSigner struct {
	priv PrivateKey
	file string

	prefix []byte
}

// NewEdSigner returns an auto-generated ed signer.
func NewEdSigner(opts ...EdSignerOptionFunc) (*EdSigner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// continue

// err == nil

// Sign signs the provided message.
func (es *EdSigner) Sign(d Domain, m []byte) types.EdSignature {
	_ = "STUB: not implemented"
	return *new(types.EdSignature)
}

// NodeID returns the node ID of the signer.
func (es *EdSigner) NodeID() types.NodeID { _ = "STUB: not implemented"; return *new(types.NodeID) }

// PublicKey returns the public key of the signer.
func (es *EdSigner) PublicKey() *PublicKey { _ = "STUB: not implemented"; return nil }

// PrivateKey returns private key.
func (es *EdSigner) PrivateKey() PrivateKey {
	_ = "STUB: not implemented"

	// Name returns the name of the signer. This is the filename of the identity file.
	return *new(PrivateKey)
}

func (es *EdSigner) Name() string { _ = "STUB: not implemented"; return "" }

// VRFSigner wraps same ed25519 key to provide ecvrf.
func (es *EdSigner) VRFSigner() *VRFSigner { _ = "STUB: not implemented"; return nil }

func (es *EdSigner) Prefix() []byte {
	_ = "STUB: not implemented"

	// Matches implements the gomock.Matcher interface for testing.
	return nil
}

func (es *EdSigner) Matches(x any) bool { _ = "STUB: not implemented"; return false }

func (es *EdSigner) String() string { _ = "STUB: not implemented"; return "" }
