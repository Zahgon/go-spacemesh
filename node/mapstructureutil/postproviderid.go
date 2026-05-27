package mapstructureutil

import (
	"github.com/go-viper/mapstructure/v2"
)

// PostProviderIDDecodeFunc mapstructure decode func for activation.PostProviderID.
func PostProviderIDDecodeFunc() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}
