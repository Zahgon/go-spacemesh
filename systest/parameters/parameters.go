package parameters

import (
	"time"
)

// New creates Parameters instance with default values.
func New() *Parameters { _ = "STUB: not implemented"; return nil }

// FromValues instantiates parameters with a cope of values.
func FromValues(values map[string]string) *Parameters { _ = "STUB: not implemented"; return nil }

type Parameters struct {
	values map[string]string
}

func (p *Parameters) Update(values map[string]string) { _ = "STUB: not implemented"; return }

type Parser[T any] func(string) (T, error)

type Parameter[T any] struct {
	name, desc string
	defaults   T
	parser     Parser[T]
}

func (p *Parameter[T]) Get(params *Parameters) T { _ = "STUB: not implemented"; return *new(T) }

// TODO(dshulyak) move this to part to validation that is executed
// when values are updated (Parameters.Update)

func NewParameter[T any](name, desc string, defaults T, parser Parser[T]) Parameter[T] {
	_ = "STUB: not implemented"
	return nil
}

func String(name, desc, defaults string) Parameter[string] { _ = "STUB: not implemented"; return nil }

func Bytes(name, desc string, defaults []byte) Parameter[[]byte] {
	_ = "STUB: not implemented"
	return nil
}

func Duration(name, desc string, defaults time.Duration) Parameter[time.Duration] {
	_ = "STUB: not implemented"
	return nil
}

func Int(name, desc string, defaults int) Parameter[int] { _ = "STUB: not implemented"; return nil }

func Bool(name, desc string) Parameter[bool] { _ = "STUB: not implemented"; return nil }

func toBytes(value string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func toString(value string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func toInt(value string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func toDuration(value string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
