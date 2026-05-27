package testcontext

type stringToString map[string]string

func (s stringToString) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (s stringToString) Type() string { _ = "STUB: not implemented"; return "" }

func (s stringToString) String() string { _ = "STUB: not implemented"; return "" }
