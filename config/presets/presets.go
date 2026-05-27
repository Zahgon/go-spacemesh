package presets

import (
	"github.com/spacemeshos/go-spacemesh/config"
)

var presets = map[string]config.Config{}

func register(name string, preset config.Config) { _ = "STUB: not implemented"; return }

// Options returns list of registered options.
func Options() []string { _ = "STUB: not implemented"; return nil }

// Get return one of the available preset.
func Get(name string) (config.Config, error) {
	_ = "STUB: not implemented"
	return *new(config.Config), nil
}
