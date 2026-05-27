package signing

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// VRFSigner is a signer for VRF purposes.
type VRFSigner struct {
	privateKey ed25519.PrivateKey
	nodeID     types.NodeID
}

// Sign signs a message for VRF purposes.
func (s VRFSigner) Sign(msg []byte) types.VrfSignature {
	_ = "STUB: not implemented"
	return *new(types.VrfSignature)
}

// NodeID of the signer.
func (s VRFSigner) NodeID() types.NodeID {
	_ = "STUB: not implemented"

	// PublicKey of the signer.
	return *new(types.NodeID)
}

func (s VRFSigner) PublicKey() *PublicKey { _ = "STUB: not implemented"; return nil }

type VRFVerifier func(types.NodeID, []byte, types.VrfSignature) bool

func NewVRFVerifier() VRFVerifier {
	_ = "STUB: not implemented"

	// Verify verifies that a signature matches public key and message.
	return *new(VRFVerifier)
}

func (v VRFVerifier) Verify(nodeID types.NodeID, msg []byte, sig types.VrfSignature) bool {
	_ = "STUB: not implemented"
	return false

	// VRFVerify verifies that a signature matches public key and message.
}

func VRFVerify(nodeID types.NodeID, msg []byte, sig types.VrfSignature) bool {
	_ = "STUB: not implemented"
	return false
}
