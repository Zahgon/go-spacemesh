package flags

// Deprecated is an interface for deprecated config fields.
// A deprecated config field should implement this interface (with a value receiver)
// and return a message explaining the deprecation.
type Deprecated interface {
	DeprecatedMsg() string
}
type deprecatedFlag struct {
	inner Deprecated
}

func (f *deprecatedFlag) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (f *deprecatedFlag) Type() string { _ = "STUB: not implemented"; return "" }

func (f *deprecatedFlag) String() string { _ = "STUB: not implemented"; return "" }

// NewDeprecatedFlag returns a flag that returns an error when used.
func NewDeprecatedFlag(v Deprecated) *deprecatedFlag { _ = "STUB: not implemented"; return nil }
