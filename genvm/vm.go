package vm

import (
	"github.com/spacemeshos/go-scale"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/registry"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

// Opt is for changing VM during initialization.
type Opt func(*VM)

// WithLogger sets logger for VM.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Config defines the configuration options for vm.
type Config struct {
	GasLimit  uint64       `mapstructure:"-"` // not configurable, overwritten by BaseConfig.BlockGasLimit
	GenesisID types.Hash20 `mapstructure:"-"` // not configurable, overwritten by GenesisConfig.GenesisID()
}

// DefaultConfig returns the default RewardConfig.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// WithConfig updates config on the vm.
func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// New returns VM instance.
func New(db sql.StateDatabase, opts ...Opt) *VM { _ = "STUB: not implemented"; return nil }

// VM handles modifications to the account state.
type VM struct {
	logger   *zap.Logger
	db       sql.StateDatabase
	cfg      Config
	registry *registry.Registry
}

// Validation initializes validation request.
func (v *VM) Validation(raw types.RawTx) system.ValidationRequest {
	_ = "STUB: not implemented"
	return *new(system.ValidationRequest)
}

// GetLayerStateRoot returns the state root at a given layer.
func (v *VM) GetLayerStateRoot(lid types.LayerID) (types.Hash32, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// GetLayerApplied returns layer of the applied transaction.
func (v *VM) GetLayerApplied(tid types.TransactionID) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// GetStateRoot gets the current state root hash.
func (v *VM) GetStateRoot() (types.Hash32, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// TODO: reconsider this.
// instead of skipping vm on empty layers, maybe pass empty layer to vm
// and let it persist empty (or previous if we will use cumulative) hash.

// GetAllAccounts returns a dump of all accounts in global state.
func (v *VM) GetAllAccounts() ([]*types.Account, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *VM) revert(lid types.LayerID) error { _ = "STUB: not implemented"; return nil }

// Revert all changes that we made after the layer.
func (v *VM) Revert(lid types.LayerID) error { _ = "STUB: not implemented"; return nil }

// AccountExists returns true if the address exists, spawned or not.
func (v *VM) AccountExists(address core.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetNonce returns expected next nonce for the address.
func (v *VM) GetNonce(address core.Address) (core.Nonce, error) {
	_ = "STUB: not implemented"
	return *new(core.Nonce), nil
}

// GetBalance returns balance for an address.
func (v *VM) GetBalance(address types.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ApplyGenesis saves list of accounts for genesis.
func (v *VM) ApplyGenesis(genesis []types.Account) error { _ = "STUB: not implemented"; return nil }

// Apply transactions.
func (v *VM) Apply(
	layer types.LayerID,
	txs []types.Transaction,
	blockRewards []types.CoinbaseReward,
) ([]types.Transaction, []types.TransactionWithResult, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (v *VM) execute(
	layer types.LayerID,
	ss *core.StagedCache,
	txs []types.Transaction,
) ([]types.TransactionWithResult, []types.Transaction, uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// NOTE this part is executed only for transactions that weren't verified
// when saved into database by txs module

// Request used to implement 2-step validation flow.
// After Parse is executed - conservative cache may do validation and skip Verify
// if transaction can't be executed.
type Request struct {
	vm    *VM
	cache *core.StagedCache

	lid     types.LayerID
	raw     types.RawTx
	decoder *scale.Decoder

	// both ctx and args are set after successful Parse
	ctx  *core.Context
	args scale.Encodable
}

// Parse header from the raw transaction.
func (r *Request) Parse() (*core.Header, error) { _ = "STUB: not implemented"; return nil, nil }

// Verify transaction. Will panic if called without Parse completing successfully.
func (r *Request) Verify() bool { _ = "STUB: not implemented"; return false }

func parse(
	logger *zap.Logger,
	lid types.LayerID,
	reg *registry.Registry,
	loader core.AccountLoader,
	cfg Config,
	raw []byte,
	decoder *scale.Decoder,
) (*core.Header, *core.Context, scale.Encodable, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(scale.Encodable), nil
}

// the transaction is either spawn or self-spawn

// spawn is not possible if principal is not spawned
// so this must be self-spawn, but we can't tell before decoding arguments

// this is any other call transaction

// this is a self spawn. if it fails validation - discard it immediately

func verify(ctx *core.Context, raw []byte, dec *scale.Decoder) bool {
	_ = "STUB: not implemented"
	return false
}
