package rangesync

const (
	// FingerprintSize is the size of a fingerprint in bytes.
	FingerprintSize = 12
)

// Fingerprint represents a fingerprint of a set of keys.
// The fingerprint is obtained by XORing together the keys in the set.
type Fingerprint [FingerprintSize]byte

// String implements log.ShortString.
func (fp Fingerprint) ShortString() string { _ = "STUB: not implemented"; return "" }

// Compare compares two fingerprints.
func (fp Fingerprint) Compare(other Fingerprint) int { _ = "STUB: not implemented"; return 0 }

// String implements fmt.Stringer.
func (fp Fingerprint) String() string { _ = "STUB: not implemented"; return "" }

// Update includes the byte slice in the fingerprint.
func (fp *Fingerprint) Update(h []byte) { _ = "STUB: not implemented"; return }

// BitFromLeft returns the n-th bit from the left in the fingerprint.
func (fp *Fingerprint) BitFromLeft(i int) bool { _ = "STUB: not implemented"; return false }

// CombineFingerprints combines two fingerprints into one.
func CombineFingerprints(a, b Fingerprint) Fingerprint {
	_ = "STUB: not implemented"
	return *

	// RandomFingerprint generates a random fingerprint.
	new(Fingerprint)
}

func RandomFingerprint() Fingerprint { _ = "STUB: not implemented"; return *new(Fingerprint) }

// EmptyFingerprint returns an empty fingerprint.
func EmptyFingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *

	// MustParseHexFingerprint converts a hex string to Fingerprint.
	new(Fingerprint)
}

func MustParseHexFingerprint(s string) Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
