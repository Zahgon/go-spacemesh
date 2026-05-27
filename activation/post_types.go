package activation

import (
	"github.com/spacemeshos/post/config"
)

type PowDifficulty [32]byte

func (d PowDifficulty) String() string { _ = "STUB: not implemented"; return "" }

// Set implements pflag.Value.Set.
func (d *PowDifficulty) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Type implements pflag.Value.Type.
func (PowDifficulty) Type() string { _ = "STUB: not implemented"; return "" }

func (d *PowDifficulty) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type PostProviderID struct {
	value *uint32
}

// String implements pflag.Value.String.
func (id PostProviderID) String() string { _ = "STUB: not implemented"; return "" }

// Type implements pflag.Value.Type.
func (PostProviderID) Type() string { _ = "STUB: not implemented"; return "" }

// Set implements pflag.Value.Set.
func (id *PostProviderID) Set(value string) error { _ = "STUB: not implemented"; return nil }

// SetInt64 sets the value of the PostProviderID to the given int64.
func (id *PostProviderID) SetUint32(value uint32) {
	_ = "STUB: not implemented"

	// Value returns the value of the PostProviderID as a pointer to uint32.
	return
}

func (id *PostProviderID) Value() *uint32 { _ = "STUB: not implemented"; return nil }

type PostPowFlags config.PowFlags

// String implements pflag.Value.String.
func (f PostPowFlags) String() string { _ = "STUB: not implemented"; return "" }

// Type implements pflag.Value.Type.
func (PostPowFlags) Type() string { _ = "STUB: not implemented"; return "" }

// Set implements pflag.Value.Set.
func (f *PostPowFlags) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (f *PostPowFlags) Value() config.PowFlags {
	_ = "STUB: not implemented"
	return *new(config.PowFlags)
}

type PostRandomXMode string

const (
	PostRandomXModeFast  PostRandomXMode = "fast"
	PostRandomXModeLight PostRandomXMode = "light"
)

// String implements pflag.Value.String.
func (m PostRandomXMode) String() string {
	_ = "STUB: not implemented"

	// Type implements pflag.Value.Type.
	return ""
}

func (PostRandomXMode) Type() string { _ = "STUB: not implemented"; return "" }

// Set implements pflag.Value.Set.
func (m *PostRandomXMode) Set(value string) error { _ = "STUB: not implemented"; return nil }
