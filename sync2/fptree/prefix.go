package fptree

import (
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

const (
	// prefixBytes is the number of bytes in a prefix.
	prefixBytes = rangesync.FingerprintSize
	// maxPrefixLen is the maximum length of a prefix in bits.
	maxPrefixLen = prefixBytes * 8
)

// prefix is a prefix of a key, represented as a bit string.
type prefix struct {
	// the bytes of the prefix, starting from the highest byte.
	b [prefixBytes]byte
	// length of the prefix in bits.
	l uint16
}

// emptyPrefix is the empty prefix (length 0).
var emptyPrefix = prefix{}

// prefixFromKeyBytes returns a prefix made from a key by using the maximum possible
// number of its bytes.
func prefixFromKeyBytes(k rangesync.KeyBytes) (p prefix) {
	_ = "STUB: not implemented"
	return *new(prefix)
}

// len returns the length of the prefix.
func (p prefix) len() int {
	_ = "STUB: not implemented"

	// left returns the prefix with one more 0 bit.
	return 0
}

func (p prefix) left() prefix { _ = "STUB: not implemented"; return *new(prefix) }

// right returns the prefix with one more 1 bit.
func (p prefix) right() prefix { _ = "STUB: not implemented"; return *new(prefix) }

// String implements fmt.Stringer.
func (p prefix) String() string { _ = "STUB: not implemented"; return "" }

// highBit returns the highest bit of the prefix as bool (false=0, true=1).
// If the prefix is empty, it returns false.
func (p prefix) highBit() bool { _ = "STUB: not implemented"; return false }

// minID sets the key to the smallest key with the prefix.
func (p prefix) minID(k rangesync.KeyBytes) { _ = "STUB: not implemented"; return }

// idAfter sets the key to the key immediately after the largest key with the prefix.
// idAfter returns true if the resulting id is zero, meaning wraparound.
func (p prefix) idAfter(k rangesync.KeyBytes) bool { _ = "STUB: not implemented"; return false }

// Copy prefix bits to the key, set all the bits after the prefix to 1, then
// increment the key.

// shift removes the highest bit from the prefix.
func (p prefix) shift() prefix { _ = "STUB: not implemented"; return *new(prefix) }

// match returns true if the prefix matches the key, that is,
// all the prefix bits are equal to the corresponding bits of the key.
func (p prefix) match(b rangesync.KeyBytes) bool { _ = "STUB: not implemented"; return false }

// preFirst0 returns the longest prefix of the key that consists entirely of binary 1s.
func preFirst0(k rangesync.KeyBytes) prefix { _ = "STUB: not implemented"; return *new(prefix) }

// preFirst1 returns the longest prefix of the key that consists entirely of binary 0s.
func preFirst1(k rangesync.KeyBytes) prefix { _ = "STUB: not implemented"; return *new(prefix) }

// commonPrefix returns common prefix between two keys.
func commonPrefix(a, b rangesync.KeyBytes) prefix { _ = "STUB: not implemented"; return *new(prefix) }

// Clear unused bits in the last used prefix byte
