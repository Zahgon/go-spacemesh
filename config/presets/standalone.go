package presets

import (
	"github.com/spacemeshos/go-spacemesh/config"
)

func init() {
	register("standalone", standalone())
}

func standalone() config.Config { _ = "STUB: not implemented"; return *new(config.Config) }

// switch on ATXv2 in epoch 2

// ensure that the correct HRP is set when generating the address below

// RequestTimeout = RequestRetryDelay * 2 * MaxRequestRetries*(MaxRequestRetries+1)/2
