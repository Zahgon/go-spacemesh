package activation

import (
	"errors"
)

var (
	// errATXChallengeExpired is returned when atx missed its publication window and needs to be regenerated.
	errATXChallengeExpired = errors.New("builder: atx expired")
	// errPoetProofNotReceived is returned when no poet proof was received.
	errPoetProofNotReceived = errors.New("builder: didn't receive any poet proof")
)

// PoetSvcUnstableError means there was a problem communicating
// with a Poet service. It wraps the source error.
type PoetSvcUnstableError struct {
	// additional contextual information
	msg string
	// the source (if any) that caused the error
	source error
}

func (e PoetSvcUnstableError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *PoetSvcUnstableError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type PoetRegistrationMismatchError struct {
	registrations   []string
	configuredPoets []string
}

func (e PoetRegistrationMismatchError) Error() string { _ = "STUB: not implemented"; return "" }
