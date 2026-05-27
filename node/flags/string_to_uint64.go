package flags

// StringToUint64Value is a flag type for string to uint64 values.
type StringToUint64Value struct {
	value *map[string]uint64
}

// NewStringToUint64Value creates instance.
func NewStringToUint64Value(p *map[string]uint64) *StringToUint64Value {
	_ = "STUB: not implemented"
	return nil
}

// Set expects value as "smth=101,else=102".
func (s *StringToUint64Value) Set(val string) error { _ = "STUB: not implemented"; return nil }

// Type returns stringToUint64 type.
func (s *StringToUint64Value) Type() string { _ = "STUB: not implemented"; return "" }

// String marshals value of the StringToUint64Value instance.
func (s *StringToUint64Value) String() string { _ = "STUB: not implemented"; return "" }
