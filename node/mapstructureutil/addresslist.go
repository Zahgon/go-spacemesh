package mapstructureutil

import (
	"github.com/go-viper/mapstructure/v2"
)

// AddressListDecodeFunc mapstructure decode func for p2p.AddressList.
// AddressList can be represented either by a string or by a slice of strings.
func AddressListDecodeFunc() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}
