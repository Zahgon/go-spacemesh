package rangesync

import (
	"errors"

	"github.com/spacemeshos/go-scale"
)

const (
	compactHashTypeNil  = 0
	compactHashType32   = 1
	compactHashSizeBits = 6
	maxCompactHashSize  = 32
)

var errInvalidCompactHash = errors.New("invalid compact hash")

// CompactHash encodes hashes in a compact form, skipping trailing zeroes.
// It also supports a nil hash (no value).
// The encoding format is as follows:
// byte 0:     spec byte
// bytes 1..n: data bytes
//
// The format of the spec byte is as follows:
// bits 0..5:  number of non-zero leading bytes
// bits 6..7:  hash type
//
// The following hash types are supported:
// 0:    nil hash
// 1:    32-byte hash
// 2,3:  reserved
//
// NOTE: when adding new hash types, we need to add a mechanism that makes sure that every
// received hash is of the expected type. Alternatively, we need to add some kind of
// context to the scale.Decoder / scale.Encoder, which may contain the size of hashes to
// be used.
type CompactHash struct {
	H KeyBytes
}

// DecodeScale implements scale.Decodable.
func (c *CompactHash) DecodeScale(dec *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EncodeScale implements scale.Encodable.
func (c *CompactHash) EncodeScale(enc *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *CompactHash) ToOrdered() KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }

func chash(h KeyBytes) CompactHash { _ = "STUB: not implemented"; return *new(CompactHash) }

const (
	keyCollectionLimit = 100000000
)

// KeyCollection represents a collection of keys of the same size.
type KeyCollection struct {
	Keys []KeyBytes
}

// DecodeScale implements scale.Decodable.
func (c *KeyCollection) DecodeScale(dec *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EncodeScale implements scale.Encodable.
func (c *KeyCollection) EncodeScale(enc *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
