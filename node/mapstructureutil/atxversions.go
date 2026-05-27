package mapstructureutil

import (
	"github.com/go-viper/mapstructure/v2"
)

// AtxVersionsDecodeFunc mapstructure decode func for activation.AtxVersions.
func AtxVersionsDecodeFunc() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}
