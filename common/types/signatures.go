package types

import (
	"github.com/spacemeshos/go-scale"
)

const (
	EdSignatureSize  = 64
	VrfSignatureSize = 80
)

type EdSignature [EdSignatureSize]byte

// EmptyEdSignature is a canonical empty EdSignature.
var EmptyEdSignature EdSignature

// EncodeScale implements scale codec interface.
func (s *EdSignature) EncodeScale(encoder *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (s *EdSignature) DecodeScale(decoder *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// String returns a string representation of the Signature, for logging purposes.
// It implements the Stringer interface.
func (s EdSignature) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the byte representation of the Signature.
func (s *EdSignature) Bytes() []byte { _ = "STUB: not implemented"; return nil }

type VrfSignature [VrfSignatureSize]byte

// EmptyVrfSignature is a canonical empty VrfSignature.
var EmptyVrfSignature VrfSignature

// String returns a string representation of the Signature, for logging purposes.
// It implements the Stringer interface.
func (s VrfSignature) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the byte representation of the Signature.
func (s *VrfSignature) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Cmp compares s and x and returns:
//
//	-1 if s <  x
//	 0 if s == x
//	+1 if s >  x
//
// The comparison is done in little endian order.
// Additionally, if x is nil, -1 is returned.
func (s *VrfSignature) Cmp(x *VrfSignature) int { _ = "STUB: not implemented"; return 0 }

// VRF signatures are little endian, so we need to compare in reverse order.

// LSB returns the least significant bit of the signature, so either 0 or 1.
func (s *VrfSignature) LSB() byte { _ = "STUB: not implemented"; return 0 }
