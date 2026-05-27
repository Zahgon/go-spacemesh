package mapstructureutil

import (
	"github.com/go-viper/mapstructure/v2"
)

// Deprecated is an interface for deprecated config fields.
// A deprecated config field should implement this interface (with a value receiver)
// and return a message explaining the deprecation.
type Deprecated interface {
	DeprecatedMsg() string
}

func DeprecatedHook() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}
