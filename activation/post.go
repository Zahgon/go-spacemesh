package activation

import (
	"context"
	"sync"
	"time"

	"github.com/spacemeshos/post/config"
	"github.com/spacemeshos/post/initialization"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// PostSetupProvider represent a compute provider for Post setup data creation.
type PostSetupProvider initialization.Provider

// PostConfig is the configuration of the Post protocol, used for data creation, proofs generation and validation.
type PostConfig struct {
	MinNumUnits   uint32 `mapstructure:"post-min-numunits"`
	MaxNumUnits   uint32 `mapstructure:"post-max-numunits"`
	LabelsPerUnit uint64 `mapstructure:"post-labels-per-unit"`
	K1            uint   `mapstructure:"post-k1"`
	K2            uint   `mapstructure:"post-k2"`
	// size of the subset of labels to verify in POST proofs
	// lower values will result in faster ATX verification but increase the risk
	// as the node must depend on malfeasance proofs to detect invalid ATXs
	K3            uint          `mapstructure:"post-k3"`
	PowDifficulty PowDifficulty `mapstructure:"post-pow-difficulty"`
}

func (c PostConfig) ToConfig() config.Config { _ = "STUB: not implemented"; return *new(config.Config) }

// PostSetupOpts are the options used to initiate a Post setup data creation session,
// either via the public smesher API, or on node launch (via cmd args).
type PostSetupOpts struct {
	DataDir          string              `mapstructure:"smeshing-opts-datadir"`
	NumUnits         uint32              `mapstructure:"smeshing-opts-numunits"`
	MaxFileSize      uint64              `mapstructure:"smeshing-opts-maxfilesize"`
	ProviderID       PostProviderID      `mapstructure:"smeshing-opts-provider"`
	Throttle         bool                `mapstructure:"smeshing-opts-throttle"`
	Scrypt           config.ScryptParams `mapstructure:"smeshing-opts-scrypt"`
	ComputeBatchSize uint64              `mapstructure:"smeshing-opts-compute-batch-size"`
}

// PostProvingOpts are the options controlling POST proving process.
type PostProvingOpts struct {
	// Number of threads used in POST proving process.
	Threads uint `mapstructure:"smeshing-opts-proving-threads"`

	// Number of nonces tried in parallel in POST proving process.
	Nonces uint `mapstructure:"smeshing-opts-proving-nonces"`

	// RandomXMode is the mode used for RandomX computations.
	RandomXMode PostRandomXMode `mapstructure:"smeshing-opts-proving-randomx-mode"`
}

func DefaultPostProvingOpts() PostProvingOpts {
	_ = "STUB: not implemented"
	return *new(PostProvingOpts)
}

// PostProofVerifyingOpts are the options controlling POST proving process.
type PostProofVerifyingOpts struct {
	// Disable verifying POST proofs. Experimental.
	// Use with caution, only on private nodes with a trusted public peer that
	// validates the proofs.
	Disabled bool `mapstructure:"smeshing-opts-verifying-disable"`

	// Number of workers spawned to verify proofs.
	Workers int `mapstructure:"smeshing-opts-verifying-workers"`
	// The minimum number of verifying workers to keep
	// while POST is being generated in parallel.
	//
	// Caps at the value of `Workers` (then scaling is disabled).
	MinWorkers int `mapstructure:"smeshing-opts-verifying-min-workers"`
	// Flags used for the PoW verification.
	Flags PostPowFlags `mapstructure:"smeshing-opts-verifying-powflags"`
}

func DefaultPostVerifyingOpts() PostProofVerifyingOpts {
	_ = "STUB: not implemented"
	return *new(PostProofVerifyingOpts)
}

func DefaultTestPostVerifyingOpts() PostProofVerifyingOpts {
	_ = "STUB: not implemented"
	return *new(PostProofVerifyingOpts)
}

// PostSetupStatus represents a status snapshot of the Post setup.
type PostSetupStatus struct {
	State            PostSetupState
	NumLabelsWritten uint64
	LastOpts         *PostSetupOpts
}

type PostSetupState int32

const (
	PostSetupStateNotStarted PostSetupState = 1 + iota
	PostSetupStatePrepared
	PostSetupStateInProgress
	PostSetupStateStopped
	PostSetupStateComplete
	PostSetupStateError
)

// DefaultPostConfig defines the default configuration for Post.
func DefaultPostConfig() PostConfig { _ = "STUB: not implemented"; return *new(PostConfig) }

// The default is to verify all K2 indices.

// DefaultPostSetupOpts defines the default options for Post setup.
func DefaultPostSetupOpts() PostSetupOpts { _ = "STUB: not implemented"; return *new(PostSetupOpts) }

func (o PostSetupOpts) ToInitOpts() config.InitOpts {
	_ = "STUB: not implemented"
	return *new(config.InitOpts)
}

// PostSetupManager implements the PostProvider interface.
type PostSetupManager struct {
	commitmentAtxId types.ATXID
	syncer          syncer

	cfg         PostConfig
	logger      *zap.Logger
	db          sql.Executor
	atxsdata    *atxsdata.Data
	goldenATXID types.ATXID
	validator   nipostValidator

	mu       sync.Mutex                  // mu protects setting the values below.
	lastOpts *PostSetupOpts              // the last options used to initiate a Post setup session.
	state    PostSetupState              // state is the current state of the Post setup.
	init     *initialization.Initializer // init is the current initializer instance.

	// delay before PoST in ATX is considered valid (counting from the time it was received)
	// used to decide whether to fully verify a candidate for commitment ATX
	postValidityDelay time.Duration
}

type PostSetupManagerOpt func(*PostSetupManager)

// PostValidityDelay sets the delay before PoST in ATX is considered valid.
func PostValidityDelay(delay time.Duration) PostSetupManagerOpt {
	_ = "STUB: not implemented"
	return *new(PostSetupManagerOpt)
}

// NewPostSetupManager creates a new instance of PostSetupManager.
func NewPostSetupManager(
	cfg PostConfig,
	logger *zap.Logger,
	db sql.Executor,
	atxsdata *atxsdata.Data,
	goldenATXID types.ATXID,
	syncer syncer,
	validator nipostValidator,
	opts ...PostSetupManagerOpt,
) (*PostSetupManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Status returns the setup current status.
func (mgr *PostSetupManager) Status() *PostSetupStatus { _ = "STUB: not implemented"; return nil }

// StartSession starts (or continues) a PoST session. It supports resuming a
// previously started session, and will return an error if a session is already
// in progress. It must be ensured that PrepareInitializer is called once
// before each call to StartSession and that the node is ATX synced.
func (mgr *PostSetupManager) StartSession(ctx context.Context, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	// Ensure only one goroutine can execute initialization at a time.
	return nil
}

// PrepareInitializer prepares the initializer to begin the initialization
// process, it needs to be called before each call to StartSession. Having this
// function be separate from StartSession provides a means to understand if the
// post configuration is valid before kicking off a very long running task
// (StartSession can take days to complete). After the first call to this
// method subsequent calls to this method will return an error until
// StartSession has completed execution.
func (mgr *PostSetupManager) PrepareInitializer(ctx context.Context, opts PostSetupOpts, id types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *PostSetupManager) commitmentAtx(ctx context.Context, dataDir string, id types.NodeID) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// if this node has already published an ATX, get its initial ATX and from it the commitment ATX

// if this node has not published an ATX select the best ATX with `findCommitmentAtx`

// findCommitmentAtx determines the best commitment ATX to use for the node.
// It will use the ATX with the highest height seen by the node and defaults to the goldenATX,
// when no ATXs have yet been published.
func (mgr *PostSetupManager) findCommitmentAtx(ctx context.Context) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// Reset deletes the data file(s).
func (mgr *PostSetupManager) Reset() error { _ = "STUB: not implemented"; return nil }

// Reset internal state.
