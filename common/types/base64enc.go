package types

type Base64Enc struct {
	inner []byte
}

func (b *Base64Enc) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (b Base64Enc) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Base64Enc) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (b *Base64Enc) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func NewBase64Enc(b []byte) Base64Enc { _ = "STUB: not implemented"; return *new(Base64Enc) }

func Base64FromString(s string) (Base64Enc, error) {
	_ = "STUB: not implemented"
	return *new(Base64Enc), nil
}

func MustBase64FromString(s string) Base64Enc { _ = "STUB: not implemented"; return *new(Base64Enc) }
