package mapstructureutil

import (
	"github.com/go-viper/mapstructure/v2"
)

// BigRatDecodeFunc mapstructure decode func for big.Rat.
func BigRatDecodeFunc() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}
