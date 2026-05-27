package types

import "go.uber.org/zap/zapcore"

//go:generate scalegen

// TransactionStatus ...
type TransactionStatus uint8

const (
	// TransactionSuccess is a status for successfully applied transaction.
	TransactionSuccess TransactionStatus = iota
	// TransactionFailure is a status for failed but consumed transaction.
	TransactionFailure
	// TODO(dshulyak) what about TransactionSkipped? We shouldn't store such state
	// but it might be useful to stream it.
)

// String implements human readable representation of the status.
func (t TransactionStatus) String() string { _ = "STUB: not implemented"; return "" }

// TransactionResult is created after consuming transaction.
type TransactionResult struct {
	Status  TransactionStatus
	Message string `scale:"max=1024"` // TODO(mafa): human readable error message, convert to error code
	Gas     uint64
	Fee     uint64
	Block   BlockID
	Layer   LayerID
	// Addresses contains all updated addresses.
	Addresses []Address `scale:"max=10"` // we expect 1-3 addresses to be updated in a transaction
}

// MarshalLogObject implements encoding for the tx result.
func (h *TransactionResult) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// TransactionWithResult is a transaction with attached result.
type TransactionWithResult struct {
	Transaction
	TransactionResult
}
