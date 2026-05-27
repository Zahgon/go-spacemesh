package cmd

import (
	"github.com/spf13/pflag"

	"github.com/spacemeshos/go-spacemesh/config"
)

func AddFlags(flagSet *pflag.FlagSet, cfg *config.Config) (configPath *string) {
	_ = "STUB: not implemented"
	// A workaround to keep the original config intact to avoid
	// overwriting it with the default values.
	return nil
}

/** ======================== Checkpoint Flags ========================== **/

/** ======================== BaseConfig Flags ========================== **/

/** ======================== P2P Flags ========================== **/

/** ======================== TIME Flags ========================== **/

/** ======================== API Flags ========================== **/

/**======================== Hare Eligibility Oracle Flags ========================== **/

/**======================== Beacon Flags ========================== **/

/**======================== Tortoise Flags ========================== **/

// TODO(moshababo): add usage desc

/**======================== Smeshing Flags ========================== **/

/**======================== PoST Proving Flags ========================== **/

/**======================== PoST Verifying Flags ========================== **/

/**======================== Consensus Flags ========================== **/

/**======================== PoET Flags ========================== **/

/**======================== bootstrap data updater Flags ========================== **/

/**======================== testing related flags ========================== **/

/**========================  Deprecated flags ========================== **/
// none at the moment
