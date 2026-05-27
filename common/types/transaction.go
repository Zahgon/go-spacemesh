package types

import (
	"time"

	"github.com/spacemeshos/go-scale"
)

//go:generate scalegen -types Transaction,Reward,RawTx

// TransactionID is a 32-byte blake3 sum of the transaction, used as an identifier.
type TransactionID Hash32

const (
	// TransactionIDSize in bytes.
	TransactionIDSize = Hash32Length
)

// Hash32 returns the TransactionID as a Hash32.
func (id TransactionID) Hash32() Hash32 {
	_ = "STUB: not implemented"

	// ShortString returns a the first 10 characters of the ID, for logging purposes.
	return *new(Hash32)
}

func (id TransactionID) ShortString() string { _ = "STUB: not implemented"; return "" }

// String returns a hexadecimal representation of the TransactionID with "0x" prepended, for logging purposes.
// It implements the fmt.Stringer interface.
func (id TransactionID) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the TransactionID as a byte slice.
func (id TransactionID) Bytes() []byte {
	_ = "STUB: not implemented"

	// Compare returns true if other (the given TransactionID) is less than this TransactionID, by lexicographic comparison.
	return nil
}

func (id TransactionID) Compare(other TransactionID) bool { _ = "STUB: not implemented"; return false }

// EncodeScale implements scale codec interface.
func (id *TransactionID) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale codec interface.
func (id *TransactionID) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Transaction is an alias to RawTx.
type Transaction struct {
	RawTx
	*TxHeader
}

// GetRaw returns raw bytes of the transaction with id.
func (t Transaction) GetRaw() RawTx {
	_ = "STUB: not implemented"

	// Verified returns true if header is set.
	return *new(RawTx)
}

func (t Transaction) Verified() bool { _ = "STUB: not implemented"; return false }

// Hash32 returns the TransactionID as a Hash32.
func (t *Transaction) Hash32() Hash32 {
	_ = "STUB: not implemented"
	return *

	// ShortString returns the first 5 characters of the ID, for logging purposes.
	new(Hash32)
}

func (t *Transaction) ShortString() string { _ = "STUB: not implemented"; return "" }

// ToTransactionIDs returns a slice of TransactionID corresponding to the given transactions.
func ToTransactionIDs(txs []*Transaction) []TransactionID { _ = "STUB: not implemented"; return nil }

// TransactionIDsToHashes turns a list of TransactionID into their Hash32 representation.
func TransactionIDsToHashes(ids []TransactionID) []Hash32 { _ = "STUB: not implemented"; return nil }

// TXState describes the state of a transaction.
type TXState uint32

const (
	// PENDING represents the state when a transaction is syntactically valid, but its nonce and
	// the principal's ability to cover gas have not been verified yet.
	PENDING TXState = iota
	// MEMPOOL represents the state when a transaction is in mempool.
	MEMPOOL
	// APPLIED represents the state when a transaction is applied to the state.
	APPLIED
)

// MeshTransaction is stored in the mesh and included in the block.
type MeshTransaction struct {
	Transaction
	LayerID  LayerID
	BlockID  BlockID
	State    TXState
	Received time.Time
}

// Reward is a virtual reward transaction, which the node keeps track of for the gRPC api.
type Reward struct {
	Layer       LayerID
	TotalReward uint64
	LayerReward uint64
	Coinbase    Address
	SmesherID   NodeID
}

// NewRawTx computes id from raw bytes and returns the object.
func NewRawTx(raw []byte) RawTx { _ = "STUB: not implemented"; return *new(RawTx) }

// RawTx stores an identity and a pointer to raw bytes.
type RawTx struct {
	ID  TransactionID
	Raw []byte `scale:"max=4096"` // transactions should always be less than 4kb
}

// AddressNonce is an (address, nonce) named tuple.
type AddressNonce struct {
	Address Address
	Nonce   Nonce
}
