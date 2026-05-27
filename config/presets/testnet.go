package presets

import (
	"github.com/spacemeshos/go-spacemesh/config"
)

func init() {
	register("testnet", testnet())
}

func testnet() config.Config { _ = "STUB: not implemented"; return *new(config.Config) }

// NOTE(dshulyak) i forgot to set protocol name for testnet when we configured it manually.
// we can't do rolling upgrade if protocol name changes, so lets keep it like that temporarily.

// https://github.com/spacemeshos/go-spacemesh/issues/4559
// 3000 of spends

// RequestRetryDelay * 2 * MaxRequestRetries*(MaxRequestRetries+1)/2

// NOTE(dshulyak) this is intentional. we increased committee size with hare3 upgrade
// but certifier continues to use 200 committee size.
// this will be upgraded in future with scheduled upgrade.
