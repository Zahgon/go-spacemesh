package presets

import (
	"github.com/spacemeshos/go-spacemesh/config"
)

func init() {
	register("fastnet", fastnet())
}

func fastnet() config.Config { _ = "STUB: not implemented"; return *new(config.Config) }

// set for systest TestEquivocation

// switch on ATXv2 in epoch 2

// node will select atxs that were received at least 4 seconds before start of the epoch
// for activeset.
// if some atxs weren't received on time it will skew eligibility distribution
// and will make some tests fail.

// ensure that the correct HRP is set when generating the address below

// faster scrypt
// Override proof of work flags to use light mode (less memory intensive)

// RequestTimeout = RequestRetryDelay * 2 * MaxRequestRetries*(MaxRequestRetries+1)/2
