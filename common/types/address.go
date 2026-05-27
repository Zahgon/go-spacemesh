package types

import (
	"errors"
	"fmt"

	"github.com/spacemeshos/go-scale"
)

const (
	// AddressLength is the expected length of the address.
	AddressLength = 24
	// AddressReservedSpace define how much bytes from top is reserved in address for future.
	AddressReservedSpace = 4
)

var (
	// ErrWrongAddressLength is returned when the length of the address is not correct.
	ErrWrongAddressLength = errors.New("wrong address length")
	// ErrUnsupportedNetwork is returned when a network is not supported.
	ErrUnsupportedNetwork = errors.New("unsupported network")
	// ErrDecodeBech32 is returned when an error occurs during decoding bech32.
	ErrDecodeBech32 = errors.New("error decoding bech32")
	// ErrMissingReservedSpace is returned if top bytes of address is not 0.
	ErrMissingReservedSpace = errors.New("missing reserved space")
)

// Config is the configuration of the address package.
var networkHrp = "sm"

func SetNetworkHRP(update string) { _ = "STUB: not implemented"; return }

func NetworkHRP() string {
	_ = "STUB: not implemented"

	// Address represents the address of a spacemesh account with AddressLength length.
	return ""
}

type Address [AddressLength]byte

// StringToAddress returns a new Address from a given string like `sm1abc...`.
func StringToAddress(src string) (Address, error) {
	_ = "STUB: not implemented"
	return *new(Address), nil
}

// for encoding bech32 uses slice of 5-bit unsigned integers. convert it back it 8-bit uints.

// AddressLength+1 cause ConvertBits append empty byte to the end of the slice.

// check that first 4 bytes are 0.

// Bytes gets the string representation of the underlying address.
func (a Address) Bytes() []byte {
	_ = "STUB: not implemented"

	// IsEmpty checks if address is empty.
	return nil
}

func (a Address) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// String implements fmt.Stringer.
func (a Address) String() string { _ = "STUB: not implemented"; return "" }

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (a Address) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

// EncodeScale implements scale codec interface.
func (a *Address) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (a *Address) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GenerateAddress generates an address from a public key.
func GenerateAddress(publicKey []byte) Address { _ = "STUB: not implemented"; return *new(Address) }

// GetHRPNetwork returns the Human-Readable-Part of bech32 addresses for a networkID.
func (a Address) GetHRPNetwork() string { _ = "STUB: not implemented"; return "" }
