package p2p

import (
	ma "github.com/multiformats/go-multiaddr"
)

// AddressList represents a list of addresses.
type AddressList []ma.Multiaddr

// AddressListFromStringSlice parses strings in the slice into
// Multiaddr values and returns them as an AddressList.
func AddressListFromStringSlice(s []string) (AddressList, error) {
	_ = "STUB: not implemented"
	return *new(AddressList), nil
}

// AddressListFromString parses a string containing a Multiaddr
// and returns it as an AddressList.
func AddressListFromString(s string) (AddressList, error) {
	_ = "STUB: not implemented"
	return *new(AddressList), nil
}

// MustParseAddresses parses multiaddr strings into AddressList, and
// panics upon any parse errors.
func MustParseAddresses(s ...string) AddressList {
	_ = "STUB: not implemented"
	return *new(AddressList)
}

// String implements Stringer.
func (al AddressList) String() string { _ = "STUB: not implemented"; return "" }

func writeAsCSV(vals AddressList) (string, error) { _ = "STUB: not implemented"; return "", nil }
