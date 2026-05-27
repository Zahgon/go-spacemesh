package types

import "go.uber.org/zap/zapcore"

//go:generate scalegen

// TxHeader is a transaction header, with some of the fields defined directly in the tx
// and the rest is computed by the template based on immutable state and method arguments.
type TxHeader struct {
	Principal       Address
	TemplateAddress Address
	Method          uint8
	Nonce           Nonce
	LayerLimits     LayerLimits
	MaxGas          uint64
	GasPrice        uint64
	MaxSpend        uint64
}

// Fee is a MaxGas multiplied by a GasPrice.
func (h *TxHeader) Fee() uint64 { _ = "STUB: not implemented"; return 0 }

// Spending is Fee() + MaxSpend.
func (h *TxHeader) Spending() uint64 { _ = "STUB: not implemented"; return 0 }

// MarshalLogObject implements encoding for the tx header.
func (h *TxHeader) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// LayerLimits if defined restricts in what layers transaction may be applied.
type LayerLimits struct {
	Min, Max uint32
}

// Nonce alias to uint64.
type Nonce = uint64
