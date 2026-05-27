package flags

import (
	"github.com/spacemeshos/go-spacemesh/p2p"
)

// AddressListValue wraps an AddressList to make it possible to use it
// like a slice value in pflag.
type AddressListValue struct {
	value   *p2p.AddressList
	changed bool
}

// NewAddressListValue creates an AddressListValue that wraps the
// specified AddressList.
func NewAddressListValue(val p2p.AddressList, p *p2p.AddressList) *AddressListValue {
	_ = "STUB: not implemented"
	return nil
}

// Set implements pflag.Value.
func (alv *AddressListValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

// String implements pflag.Value.
func (alv *AddressListValue) String() string { _ = "STUB: not implemented"; return "" }

// Type implements pflag.Value.
func (alv *AddressListValue) Type() string { _ = "STUB: not implemented"; return "" }

// Append implements pflag.SliceValue.
func (alv *AddressListValue) Append(val string) error { _ = "STUB: not implemented"; return nil }

// GetSlice implements pflag.SliceValue.
func (alv *AddressListValue) GetSlice() []string { _ = "STUB: not implemented"; return nil }

// Replace implements pflag.SliceValue.
func (alv *AddressListValue) Replace(val []string) error { _ = "STUB: not implemented"; return nil }

func readAsCSV(val string) (p2p.AddressList, error) {
	_ = "STUB: not implemented"
	return *new(p2p.AddressList), nil
}
