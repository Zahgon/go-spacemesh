package rangesync

// KeyBytes represents an item (key) in a reconciliable set.
type KeyBytes []byte

// String implements fmt.Stringer.
func (k KeyBytes) String() string { _ = "STUB: not implemented"; return "" }

// String implements log.ShortString.
func (k KeyBytes) ShortString() string { _ = "STUB: not implemented"; return "" }

// Clone returns a copy of the key.
func (k KeyBytes) Clone() KeyBytes {
	_ = "STUB: not implemented"
	return *

	// Compare compares two keys.
	new(KeyBytes)
}

func (k KeyBytes) Compare(other KeyBytes) int { _ = "STUB: not implemented"; return 0 }

// BitFromLeft returns the n-th bit from the left in the key.
func (k KeyBytes) BitFromLeft(i int) bool { _ = "STUB: not implemented"; return false }

// Inc returns the key with the same number of bytes as this one, obtained by incrementing
// the key by one. It returns true if the increment has caused an overflow.
func (k KeyBytes) Inc() (overflow bool) { _ = "STUB: not implemented"; return false }

// Zero sets all bytes in the key to zero.
func (k KeyBytes) Zero() { _ = "STUB: not implemented"; return }

// IsZero returns true if all bytes in the key are zero.
func (k KeyBytes) IsZero() bool { _ = "STUB: not implemented"; return false }

// Trim zeroes all the bits in the key starting with the given bit index.
func (k KeyBytes) Trim(bit int) { _ = "STUB: not implemented"; return }

// RandomKeyBytes generates random data in bytes for testing.
func RandomKeyBytes(size int) KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }

// MustParseHexKeyBytes converts a hex string to KeyBytes.
func MustParseHexKeyBytes(s string) KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }
