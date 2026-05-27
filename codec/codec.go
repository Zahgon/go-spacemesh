package codec

import (
	"bytes"
	"errors"
	"io"
	"sync"

	"github.com/spacemeshos/go-scale"
)

var ErrShortRead = errors.New("decode from buffer: not all bytes were consumed")

// Encodable is an interface that must be implemented by a struct to be encoded.
type Encodable = scale.Encodable

// Decodable is an interface that must be implemented bya struct to be decoded.
type Decodable = scale.Decodable

// EncodeTo encodes value to a writer stream.
func EncodeTo(w io.Writer, value Encodable) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// DecodeFrom decodes a value using data from a reader stream.
func DecodeFrom(r io.Reader, value Decodable) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO(dshulyak) this is a temporary solution to improve encoder allocations.
// If this will stay it must be changed to one of the:
// - use buffer with allocations that can be adjusted using stats
// - use multiple buffers that increase in size (e.g. 16, 32, 64, 128 bytes).
var encoderPool = sync.Pool{
	New: func() any {
		b := new(bytes.Buffer)
		b.Grow(64)
		return b
	},
}

func getEncoderBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func putEncoderBuffer(b *bytes.Buffer) { _ = "STUB: not implemented"; return }

func MustEncodeTo(w io.Writer, value Encodable) { _ = "STUB: not implemented"; return }

func MustEncode(value Encodable) []byte { _ = "STUB: not implemented"; return nil }

// Encode value to a byte buffer.
func Encode(value Encodable) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MustDecode(buf []byte, value Decodable) { _ = "STUB: not implemented"; return }

// Decode value from a byte buffer.
func Decode(buf []byte, value Decodable) error { _ = "STUB: not implemented"; return nil }

// EncodeSlice encodes slice with a length prefix.
func EncodeSlice[V any, H scale.EncodablePtr[V]](value []V) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustEncodeSlice encodes slice with a length prefix or panics on error.
func MustEncodeSlice[V any, H scale.EncodablePtr[V]](value []V) []byte {
	_ = "STUB: not implemented"
	return nil
}

// DecodeSliceFromReader accepts a reader and decodes slice with a length prefix.
func DecodeSliceFromReader[V any, H scale.DecodablePtr[V]](r io.Reader) ([]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustDecodeSliceFromReader decodes slice with a length prefix or panics on error.
func MustDecodeSliceFromReader[V any, H scale.DecodablePtr[V]](r io.Reader) []V {
	_ = "STUB: not implemented"
	return nil
}

// DecodeSlice decodes slice from a buffer.
func DecodeSlice[V any, H scale.DecodablePtr[V]](buf []byte) ([]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodeCompact16 encodes uint16 to a buffer.
// ReadSlice decodes slice from am io.Reader.
func ReadSlice[V any, H scale.DecodablePtr[V]](r io.Reader) ([]V, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// EncodeCompact16 encodes uint16 to an io.Writer.
func EncodeCompact16(w io.Writer, value uint16) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeCompact16 decodes uint16 from an io.Reader.
func DecodeCompact16(r io.Reader) (uint16, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

// EncodeStringSlice encodes []string to an io.Writer.
func EncodeStringSlice(w io.Writer, value []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeStringSlice decodes []string from an io.Reader.
func DecodeStringSlice(r io.Reader) ([]string, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// EncodeByteSlice encodes []string to an io.Writer.
func EncodeByteSlice(w io.Writer, value []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeByteSlice decodes []string from an io.Reader.
func DecodeByteSlice(r io.Reader) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// EncodeLen encodes a length value to an io.Writer.
func EncodeLen(w io.Writer, value uint32) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// DecodeLen decodes a length value from an io.Reader.
func DecodeLen(r io.Reader) (uint32, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

// DecodeStringWithLimit decodes a string from an io.Reader, limiting the maximum length.
func DecodeStringWithLimit(r io.Reader, limit uint32) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}
