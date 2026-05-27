package types

const (
	// BeaconSize in bytes.
	BeaconSize = 4
)

// Beacon defines the beacon value. A beacon is generated once per epoch and is used to
// - verify smesher's VRF signature for proposal/ballot eligibility
// - determine good ballots in verifying tortoise.
type Beacon [BeaconSize]byte

// EmptyBeacon is a canonical empty Beacon.
var EmptyBeacon = Beacon{}

// String implements the stringer interface and is used also by the logger when
// doing full logging into a file.
func (b Beacon) String() string { _ = "STUB: not implemented"; return "" }

// Bytes gets the byte representation of the underlying hash.
func (b Beacon) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Beacon) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Beacon) UnmarshalText(buf []byte) error { _ = "STUB: not implemented"; return nil }

// BytesToBeacon sets the first BeaconSize bytes of b to the Beacon's data.
func BytesToBeacon(b []byte) Beacon { _ = "STUB: not implemented"; return *new(Beacon) }

// HexToBeacon sets byte representation of s to a Beacon.
func HexToBeacon(s string) Beacon { _ = "STUB: not implemented"; return *new(Beacon) }
