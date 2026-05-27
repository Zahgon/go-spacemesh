package types

import (
	"fmt"
	"reflect"

	"github.com/spacemeshos/go-scale"
)

const (
	Hash32Length = 32
	Hash20Length = 20
)

var (
	hash20T = reflect.TypeOf(Hash20{})
	hash32T = reflect.TypeOf(Hash32{})

	// EmptyHash32 is the zero hash.
	EmptyHash32 = Hash32{}
)

// Hash32 represents the 32-byte blake3 hash of arbitrary data.
type Hash32 [Hash32Length]byte

// Hash20 represents the 20-byte blake3 hash of arbitrary data.
type Hash20 [Hash20Length]byte

// Bytes gets the byte representation of the underlying hash.
func (h Hash20) Bytes() []byte {
	_ = "STUB: not implemented"

	// String implements the stringer interface and is used also by the logger when
	// doing full logging into a file.
	return nil
}

func (h Hash20) String() string { _ = "STUB: not implemented"; return "" }

// ShortString returns a the first 5 hex-encoded bytes of the hash, for logging purposes.
func (h Hash20) ShortString() string { _ = "STUB: not implemented"; return "" }

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (h Hash20) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

// UnmarshalText parses a hash in hex syntax.
func (h *Hash20) UnmarshalText(input []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON parses a hash in hex syntax.
func (h *Hash20) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText returns the hex representation of h.
func (h Hash20) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetBytes sets the hash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *Hash20) SetBytes(b []byte) { _ = "STUB: not implemented"; return }

// ToHash32 returns a Hash32 whose first 20 bytes are the bytes of this Hash20, it is right-padded with zeros.
func (h Hash20) ToHash32() (h32 Hash32) { _ = "STUB: not implemented"; return *new(Hash32) }

// CalcProposalsHash32 returns the 32-byte blake3 sum of the IDs, sorted in lexicographic order. The pre-image is
// prefixed with additionalBytes.
func CalcProposalsHash32(view []ProposalID, additionalBytes []byte) Hash32 {
	_ = "STUB: not implemented"
	return *new(Hash32)
}

// CalcProposalHash32Presorted returns the 32-byte blake3 sum of the IDs, in the order given. The pre-image is
// prefixed with additionalBytes.
func CalcProposalHash32Presorted(sortedView []ProposalID, additionalBytes []byte) Hash32 {
	_ = "STUB: not implemented"
	return *new(Hash32)
}

// this never returns an error: https://golang.org/pkg/hash/#Hash

// CalcBlockHash32Presorted returns the 32-byte blake3 sum of the IDs, in the order given. The pre-image is
// prefixed with additionalBytes.
func CalcBlockHash32Presorted(sortedView []BlockID, additionalBytes []byte) Hash32 {
	_ = "STUB: not implemented"
	return *new(Hash32)
}

// this never returns an error: https://golang.org/pkg/hash/#Hash

// CalcHash20 returns the 20-byte blake3 sum of the given data.
func CalcHash20(data []byte) Hash20 {
	_ = "STUB: not implemented"
	return *

	// CalcHash32 returns the 32-byte blake3 sum of the given data.
	new(Hash20)
}

func CalcHash32(data []byte) Hash32 {
	_ = "STUB: not implemented"
	return *

	// BytesToHash sets b to hash.
	// If b is larger than len(h), b will be cropped from the left.
	new(Hash32)
}

func BytesToHash(b []byte) Hash32 { _ = "STUB: not implemented"; return *new(Hash32) }

// HexToHash32 sets byte representation of s to hash.
// If b is larger than len(h), b will be cropped from the left.
func HexToHash32(s string) Hash32 { _ = "STUB: not implemented"; return *new(Hash32) }

// Bytes gets the byte representation of the underlying hash.
func (h Hash32) Bytes() []byte {
	_ = "STUB: not implemented"

	// String implements the stringer interface and is used also by the logger when
	// doing full logging into a file.
	return nil
}

func (h Hash32) String() string { _ = "STUB: not implemented"; return "" }

// ShortString returns the first 5 hex-encoded bytes of the hash, for logging purposes.
func (h Hash32) ShortString() string { _ = "STUB: not implemented"; return "" }

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (h Hash32) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

// UnmarshalText parses a hash in hex syntax.
func (h *Hash32) UnmarshalText(input []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON parses a hash in hex syntax.
func (h *Hash32) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText returns the hex representation of h.
func (h Hash32) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetBytes sets the hash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *Hash32) SetBytes(b []byte) { _ = "STUB: not implemented"; return }

// ToHash20 returns a Hash20, whose the 20-byte prefix of this Hash32.
func (h Hash32) ToHash20() (h20 Hash20) { _ = "STUB: not implemented"; return *new(Hash20) }

// EncodeScale implements scale codec interface.
func (h *Hash32) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (h *Hash32) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
