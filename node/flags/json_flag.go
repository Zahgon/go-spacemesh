package flags

type JSONFlag struct {
	Value any
}

func (f *JSONFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *JSONFlag) Set(v string) error { _ = "STUB: not implemented"; return nil }

func (f *JSONFlag) Type() string { _ = "STUB: not implemented"; return "" }
