package p2p

import (
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

const keyFilename = "p2p.key"

type identityInfo struct {
	Key []byte
	ID  peer.ID // this is needed only to simplify integration with some testing tools
}

func genIdentity() (crypto.PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey), nil
}

func identityInfoFromDir(dir string) (*identityInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdentityInfoFromDir returns a printable ID from a given identity directory.
func IdentityInfoFromDir(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// EnsureIdentity generates an identity key file in given directory.
func EnsureIdentity(dir string) (crypto.PrivKey, error) {
	_ = "STUB: not implemented"
	// TODO add crc check
	return *new(crypto.PrivKey), nil
}
