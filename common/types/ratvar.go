package types

import (
	"math/big"
)

// RatVar is a wrapper for big.Rat to use it with the pflag package.
type RatVar big.Rat

// String returns a string representation of big.Rat.
func (r *RatVar) String() string { _ = "STUB: not implemented"; return "" }

// Set sets the value of big.Rat to a string.
func (r *RatVar) Set(s string) error { _ = "STUB: not implemented"; return nil }

// Type returns *big.Rat type.
func (r *RatVar) Type() string { _ = "STUB: not implemented"; return "" }
