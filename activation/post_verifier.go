package activation

import (
	"context"

	"github.com/spacemeshos/post/config"
	"github.com/spacemeshos/post/shared"
	"github.com/spacemeshos/post/verifying"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/events"
)

//go:generate mockgen -typed -source=post_verifier.go -destination=mocks/subscription.go -package=mocks

type verifyPostJob struct {
	ctx      context.Context // context of Verify() call
	proof    *shared.Proof
	metadata *shared.ProofMetadata
	opts     []verifying.OptionFunc
	result   chan error
}

type postStatesGetter interface {
	Get() map[types.NodeID]types.PostState
}

type subscription[T any] interface {
	Close()
	Out() <-chan T
	Full() <-chan struct{}
}
type autoscaler struct {
	sub        subscription[events.UserEvent]
	bufferSize int
	logger     *zap.Logger
	postStates postStatesGetter
}

func newAutoscaler(logger *zap.Logger, postStates postStatesGetter, bufferSize int) *autoscaler {
	_ = "STUB: not implemented"
	return nil
}

func (a *autoscaler) subscribe() error { _ = "STUB: not implemented"; return nil }

func (a autoscaler) run(stop chan struct{}, s scaler, min, target int) {
	_ = "STUB: not implemented"
	return
}

type offloadingPostVerifier struct {
	eg          errgroup.Group
	log         *zap.Logger
	verifier    PostVerifier
	workers     []*postVerifierWorker
	prioritized chan *verifyPostJob
	jobs        chan *verifyPostJob
	stop        chan struct{} // signal to stop all goroutines

	prioritizedIds map[types.NodeID]struct{}
}

type postVerifierWorker struct {
	verifier    PostVerifier
	log         *zap.Logger
	prioritized <-chan *verifyPostJob
	jobs        <-chan *verifyPostJob
	stop        chan struct{} // signal to stop this worker
	stopped     chan struct{} // signal that this worker has stopped
	shutdown    chan struct{} // signal that the verifier is closing
}

type postVerifier struct {
	*verifying.ProofVerifier
	logger *zap.Logger
	cfg    config.Config
}

func (v *postVerifier) Verify(
	_ context.Context,
	p *shared.Proof,
	m *shared.ProofMetadata,
	opts ...postVerifierOptionFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

type postVerifierOpts struct {
	opts           PostProofVerifyingOpts
	prioritizedIds []types.NodeID
	autoscaling    *struct {
		postStates postStatesGetter
	}
}

type PostVerifierOpt func(v *postVerifierOpts)

func WithVerifyingOpts(opts PostProofVerifyingOpts) PostVerifierOpt {
	_ = "STUB: not implemented"
	return *new(PostVerifierOpt)
}

func WithPrioritizedID(id types.NodeID) PostVerifierOpt {
	_ = "STUB: not implemented"
	return *new(PostVerifierOpt)
}

func WithAutoscaling(postStates postStatesGetter) PostVerifierOpt {
	_ = "STUB: not implemented"
	return *new(PostVerifierOpt)
}

// NewPostVerifier creates a new post verifier.
func NewPostVerifier(cfg PostConfig, logger *zap.Logger, opts ...PostVerifierOpt) (PostVerifier, error) {
	_ = "STUB: not implemented"
	return *new(PostVerifier), nil
}

// newOffloadingPostVerifier creates a new post proof verifier with the given number of workers.
// The verifier will distribute incoming proofs between the workers.
// It will block if all workers are busy.
//
// SAFETY: The `verifier` must be safe to use concurrently.
//
// The verifier must be closed after use with Close().
func newOffloadingPostVerifier(
	verifier PostVerifier,
	numWorkers int,
	logger *zap.Logger,
	prioritizedIds ...types.NodeID,
) *offloadingPostVerifier {
	_ = "STUB: not implemented"
	return nil
}

// Turn on automatic scaling of the number of workers.
// The number of workers will be scaled between `min` and `target` (inclusive).
func (v *offloadingPostVerifier) autoscale(min, target int, postStates postStatesGetter) {
	_ = "STUB: not implemented"
	return
}

// Scale the number of workers to the given number.
//
// SAFETY: Must not be called concurrently.
// This is satisfied by the fact that the only caller is the autoscaler,
// which executes scale() serially.
func (v *offloadingPostVerifier) scale(target int) { _ = "STUB: not implemented"; return }

// scale up

// scale down

// Verify creates a Job from given parameters, adds to jobs queue (prioritized or not)
// and waits for result of Job execution.
func (v *offloadingPostVerifier) Verify(
	ctx context.Context,
	p *shared.Proof,
	m *shared.ProofMetadata,
	opts ...postVerifierOptionFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *offloadingPostVerifier) Close() error { _ = "STUB: not implemented"; return nil }

func (w *postVerifierWorker) start() { _ = "STUB: not implemented"; return }

// First try to process a prioritized job.

type noopPostVerifier struct{}

func (v *noopPostVerifier) Verify(
	_ context.Context,
	_ *shared.Proof,
	_ *shared.ProofMetadata,
	_ ...postVerifierOptionFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *noopPostVerifier) Close() error { _ = "STUB: not implemented"; return nil }
