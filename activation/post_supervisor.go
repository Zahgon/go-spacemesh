package activation

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
	"testing"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

// DefaultPostServiceConfig returns the default config for post service.
func DefaultPostServiceConfig() PostSupervisorConfig {
	_ = "STUB: not implemented"
	return *new(PostSupervisorConfig)
}

// DefaultTestPostServiceConfig returns the default config for post service in tests.
func DefaultTestPostServiceConfig(tb testing.TB) PostSupervisorConfig {
	_ = "STUB: not implemented"
	return *new(PostSupervisorConfig)
}

// PostSupervisorConfig holds the configuration for the post service.
// This is not intended to be configurable by the user.
type PostSupervisorConfig struct {
	PostServiceCmd string // path to the post service binary, configurable for testing
	NodeAddress    string // address of the node, configurable for testing, set automatically during startup
	MaxRetries     int    // only for testing, not used in production

	CACert string // only for testing, not used in production
	Cert   string // only for testing, not used in production
	Key    string // only for testing, not used in production
}

// PostSupervisor manages a local post service.
type PostSupervisor struct {
	logger *zap.Logger

	postCfg     PostConfig
	provingOpts PostProvingOpts

	postSetupProvider postSetupProvider
	atxBuilder        atxBuilder

	pid atomic.Int64 // pid of the running post service, only for tests.

	mtx  sync.Mutex         // protects fields below
	eg   errgroup.Group     // eg manages post service goroutines.
	stop context.CancelFunc // stops the running command.
}

// NewPostSupervisor returns a new post service.
func NewPostSupervisor(
	logger *zap.Logger,
	postCfg PostConfig,
	provingOpts PostProvingOpts,
	postSetupProvider postSetupProvider,
	atxBuilder atxBuilder,
) *PostSupervisor {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PostSupervisor) Config() PostConfig {
	_ = "STUB: not implemented"

	// Providers returns a list of available compute providers for Post setup.
	return *new(PostConfig)
}

func (*PostSupervisor) Providers() ([]PostSetupProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Benchmark runs a short benchmarking session for a given provider to evaluate its performance.
func (*PostSupervisor) Benchmark(p PostSetupProvider) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ps *PostSupervisor) Status() *PostSetupStatus { _ = "STUB: not implemented"; return nil }

func (ps *PostSupervisor) Start(cmdCfg PostSupervisorConfig, opts PostSetupOpts, sig *signing.EdSigner) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(mafa): verify that opts don't delete existing files

// reset errgroup to allow restarts.

// If it returns any error other than context.Canceled
// (which is how we signal it to stop) then we shutdown.

// Stop stops the post service.
func (ps *PostSupervisor) Stop(deleteFiles bool) error { _ = "STUB: not implemented"; return nil }

// captureCmdOutput returns a function that reads from the given pipe and logs the output.
// It returns when the pipe is closed.
func (ps *PostSupervisor) captureCmdOutput(pipe io.ReadCloser, smesherId types.NodeID) func() error {
	_ = "STUB: not implemented"
	return nil
}

// remove line delimiters at end of input

func (ps *PostSupervisor) runCmd(
	ctx context.Context,
	cmdCfg PostSupervisorConfig,
	postCfg PostConfig,
	postOpts PostSetupOpts,
	provingOpts PostProvingOpts,
	smesherId types.NodeID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// safe: cmdCfg is not configurable by the user.

// safe: is checked to be a valid directory in Start().

// safe: cmdCfg is not configurable by the user.

// safe: cmdCfg is not configurable by the user.

// safe: cmdCfg is not configurable by the user.

// safe: cmdCfg is not configurable by the user.

// arguments are shell escaped by exec.Command
