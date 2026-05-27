package config

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type Genesis time.Time

func (g Genesis) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (g Genesis) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *Genesis) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler and is needed for viper.
func (g *Genesis) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// Equal is needed for go-cmp.
func (g Genesis) Equal(other Genesis) bool { _ = "STUB: not implemented"; return false }

// String implements fmt.Stringer.
func (g Genesis) String() string { _ = "STUB: not implemented"; return "" }

// Set implements pflag.Value.Set.
func (g *Genesis) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Type implements pflag.Value.Type.
func (Genesis) Type() string {
	_ = "STUB: not implemented"

	// GenesisConfig contains immutable parameters for the protocol.
	return ""
}

type GenesisConfig struct {
	GenesisTime Genesis           `mapstructure:"genesis-time"`
	ExtraData   string            `mapstructure:"genesis-extra-data"`
	Accounts    map[string]uint64 `mapstructure:"accounts"`
}

// GenesisID computes genesis id from GenesisTime and ExtraData.
func (g *GenesisConfig) GenesisID() types.Hash20 {
	_ = "STUB: not implemented"
	return *new(types.Hash20)
}

func (g *GenesisConfig) GoldenATX() types.Hash32 {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

// Validate GenesisConfig.
func (g *GenesisConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// Diff returns difference between two configs.
func (g *GenesisConfig) Diff(other *GenesisConfig) string { _ = "STUB: not implemented"; return "" }

// LoadFromFile loads config from file.
func (g *GenesisConfig) LoadFromFile(filename string) error { _ = "STUB: not implemented"; return nil }

// WriteToFile writes config content to file.
func (g *GenesisConfig) WriteToFile(filename string) error { _ = "STUB: not implemented"; return nil }

// ToAccounts creates list of types.Account instance from config.
func (g *GenesisConfig) ToAccounts() []types.Account { _ = "STUB: not implemented"; return nil }

// Account1Private is the private key for test account.
const Account1Private = "0x2dcddb8e0ddd2269f536da5768e890790f2b84366e0fb8396bdcd15c0d7c30b" +
	"90002abedccd3ffcbf46f35f11b314d17c05a2905f918d0d72f2f6989640fbb43"

// Account2Private is the private key for second test account.
const Account2Private = "0x0bb3f2936d42f463e597f5fb2c48bbd8475ce74ba91f1eaae97df4084d306b4" +
	"9feaf3d38b6ef430933ebedeb073af7bec018e8d2e379fa47df6a9fa07a6a8344"

// DefaultGenesisConfig is the default configuration for the node.
func DefaultGenesisConfig() GenesisConfig {
	_ = "STUB: not implemented"
	// NOTE(dshulyak) keys in default config are used in some tests
	return *new(GenesisConfig)
}

func generateGenesisAccounts() map[string]uint64 { _ = "STUB: not implemented"; return nil }

// we default to 10^8 SMH per account which is 10^17 smidge
// each genesis account starts off with 10^17 smidge
