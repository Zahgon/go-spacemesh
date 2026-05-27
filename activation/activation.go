// Package activation is responsible for creating activation transactions and running the mining flow, coordinating
// Post building, sending proofs to PoET and building NIPost structs.
package activation

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/spacemeshos/go-scale"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var (
	errNotFound    = errors.New("not found")
	errNilVrfNonce = errors.New("nil VRF nonce")
)

// PoetConfig is the configuration to interact with the poet server.
type PoetConfig struct {
	// Offset from the epoch start when the poet round starts
	PhaseShift time.Duration `mapstructure:"phase-shift"`
	// CycleGap gives the duration between the end of a PoET round and the start of the next
	CycleGap time.Duration `mapstructure:"cycle-gap"`
	// GracePeriod defines the time before the start of the next PoET round until the node
	// waits before building its NiPoST challenge. Shorter durations allow the node to
	// possibly pick a better positioning ATX, but come with the risk that the node might
	// not be able to validate that ATX and has to fall back to using its own previous ATX.
	GracePeriod       time.Duration `mapstructure:"grace-period"`
	RequestTimeout    time.Duration `mapstructure:"poet-request-timeout"`
	RequestRetryDelay time.Duration `mapstructure:"retry-delay"`
	// Period to find positioning ATX. Must be less, than GracePeriod
	PositioningATXSelectionTimeout time.Duration `mapstructure:"positioning-atx-selection-timeout"`
	InfoCacheTTL                   time.Duration `mapstructure:"info-cache-ttl"`
	PowParamsCacheTTL              time.Duration `mapstructure:"pow-params-cache-ttl"`
	MaxRequestRetries              int           `mapstructure:"retry-max"`
	PoetProofsCache                int           `mapstructure:"poet-proofs-cache"`
}

func DefaultPoetConfig() PoetConfig { _ = "STUB: not implemented"; return *new(PoetConfig) }

const (
	defaultPoetRetryInterval = 5 * time.Second
)

// Config defines configuration for Builder.
type Config struct {
	GoldenATXID      types.ATXID
	RegossipInterval time.Duration
}

// Builder struct is the struct that orchestrates the creation of activation transactions
// it is responsible for initializing post, receiving poet proof and orchestrating nipost after which it will
// calculate total weight and providing relevant view as proof.
type Builder struct {
	accountLock       sync.RWMutex
	coinbaseAccount   types.Address
	conf              Config
	db                sql.Executor
	atxsdata          *atxsdata.Data
	localDB           sql.LocalDatabase
	publisher         pubsub.Publisher
	nipostBuilder     nipostBuilder
	validator         nipostValidator
	layerClock        layerClock
	syncer            syncer
	logger            *zap.Logger
	poets             []PoetService
	poetCfg           PoetConfig
	poetRetryInterval time.Duration
	// delay before PoST in ATX is considered valid (counting from the time it was received)
	postValidityDelay time.Duration
	// ATX versions
	versions []atxVersion

	posAtxFinder positioningAtxFinder

	// post states of each known identity
	postStates PostStates

	// smeshingMutex protects methods like `StartSmeshing` and `StopSmeshing` from concurrent execution
	// since they (can) modify the fields below.
	smeshingMutex sync.Mutex
	signers       map[types.NodeID]*signing.EdSigner
	eg            errgroup.Group
	stop          context.CancelFunc
}

type positioningAtxFinder struct {
	finding sync.Mutex
	found   *struct {
		id         types.ATXID
		forPublish types.EpochID
	}
}

type BuilderOption func(*Builder)

func WithPostValidityDelay(delay time.Duration) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

// WithPoetRetryInterval modifies time that builder will have to wait before retrying ATX build process
// if it failed due to issues with PoET server.
func WithPoetRetryInterval(interval time.Duration) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

// WithPoetConfig sets the poet config.
func WithPoetConfig(c PoetConfig) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

func WithPoets(poets ...PoetService) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

func WithValidator(v nipostValidator) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

func WithPostStates(ps PostStates) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

func BuilderAtxVersions(v AtxVersions) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

// NewBuilder returns an atx builder that will start a routine that will attempt to create an atx upon each new layer.
func NewBuilder(
	conf Config,
	db sql.Executor,
	atxsdata *atxsdata.Data,
	localDB sql.LocalDatabase,
	publisher pubsub.Publisher,
	nipostBuilder nipostBuilder,
	layerClock layerClock,
	syncer syncer,
	log *zap.Logger,
	opts ...BuilderOption,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

// Smeshing returns true if atx builder is smeshing.
func (b *Builder) Smeshing() bool { _ = "STUB: not implemented"; return false }

// PostStates returns the current state of the post service for each registered smesher.
func (b *Builder) PostStates() map[types.IdentityDescriptor]types.PostState {
	_ = "STUB: not implemented"
	return nil
}

// StartSmeshing is the main entry point of the atx builder. It runs the main
// loop of the builder in a new go-routine and shouldn't be called more than
// once without calling StopSmeshing in between. If the post data is incomplete
// or missing, data creation session will be preceded. Changing of the post
// options (e.g., number of labels), after initial setup, is supported. If data
// creation fails for any reason then the go-routine will panic.
func (b *Builder) StartSmeshing(coinbase types.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) startID(ctx context.Context, sig *signing.EdSigner) {
	_ = "STUB: not implemented"
	return
}

// StopSmeshing stops the atx builder.
func (b *Builder) StopSmeshing(deleteFiles bool) error { _ = "STUB: not implemented"; return nil }

// SmesherIDs returns the ID of the smesher that created this activation.
func (b *Builder) SmesherIDs() []types.NodeID { _ = "STUB: not implemented"; return nil }

func (b *Builder) BuildInitialPost(ctx context.Context, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	// Generate the initial POST if we don't have an ATX...
	return nil
}

// ...and if we haven't stored an initial post yet.

// Create the initial post and save it.

func (b *Builder) buildPost(ctx context.Context, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) run(ctx context.Context, sig *signing.EdSigner) {
	_ = "STUB: not implemented"
	return
}

// give node some time to sync in case selecting the positioning ATX caused the challenge to expire

// delete the existing db post
// call build initial post again

// other failures are related to in-process software. we may as well panic here

func (b *Builder) BuildNIPostChallenge(ctx context.Context, nodeID types.NodeID) (*types.NIPostChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to get existing challenge

// Start building new challenge:
// 1. get previous ATX

// no previous ATX
// other errors

// 2. check if we didn't miss beginning of PoET round

// 3. wait if needed till getting closer to PoET round start

// 4. build new challenge

// regular ATX challenge

func (b *Builder) getExistingChallenge(
	logger *zap.Logger,
	currentEpochId types.EpochID,
	nodeID types.NodeID,
) (*types.NIPostChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset the state to idle because we won't be building POST until we get a new PoET proof
// (typically more than epoch time from now).

// challenge is fresh

func (b *Builder) buildInitialNIPostChallenge(
	ctx context.Context,
	logger *zap.Logger,
	nodeID types.NodeID,
	publishEpochId types.EpochID,
) (*types.NIPostChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if initial post is not found, declare it invalid so it is regenerated

func (b *Builder) GetPrevAtx(nodeID types.NodeID) (*types.ActivationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCoinbase sets the address rewardAddress to be the coinbase account written into the activation transaction
// the rewards for blocks made by this miner will go to this address.
func (b *Builder) SetCoinbase(rewardAddress types.Address) { _ = "STUB: not implemented"; return }

// Coinbase returns the current coinbase address.
func (b *Builder) Coinbase() types.Address { _ = "STUB: not implemented"; return *new(types.Address) }

// PublishActivationTx attempts to publish an atx, it returns an error if an atx cannot be created.
func (b *Builder) PublishActivationTx(ctx context.Context, sig *signing.EdSigner) error {
	_ = "STUB: not implemented"
	return nil
}

// try again

func (b *Builder) poetRoundStart(epoch types.EpochID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

type builtAtx interface {
	ID() types.ATXID

	scale.Encodable
	zapcore.ObjectMarshaler
}

func (b *Builder) createAtx(
	ctx context.Context,
	sig *signing.EdSigner,
	challenge *types.NIPostChallenge,
) (builtAtx, error) {
	_ = "STUB: not implemented"
	return *new(builtAtx), nil
}

// initial NIPoST challenge is not discarded; don't return ErrATXChallengeExpired

// `version` is already checked in the beginning of the function
// and it cannot have a different value.

func (b *Builder) broadcast(ctx context.Context, atx scale.Encodable) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// searchPositioningAtx returns atx id with the highest tick height.
// Publish epoch is used for caching the positioning atx.
func (b *Builder) searchPositioningAtx(
	ctx context.Context,
	nodeID types.NodeID,
	publish types.EpochID,
) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// positioning ATX publish epoch must be lower than the publish epoch of built ATX

// getPositioningAtx returns the positioning ATX.
// The provided previous ATX is picked if it has a greater or equal
// tick count as the ATX selected in `searchPositioningAtx`.
func (b *Builder) getPositioningAtx(
	ctx context.Context,
	nodeID types.NodeID,
	publish types.EpochID,
	previous *types.ActivationTx,
) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

func (b *Builder) Regossip(ctx context.Context, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// checkpoint

func (b *Builder) version(publish types.EpochID) types.AtxVersion {
	_ = "STUB: not implemented"
	return *new(types.AtxVersion)
}

func findFullyValidHighTickAtx(
	ctx context.Context,
	atxdata *atxsdata.Data,
	publish types.EpochID,
	goldenATXID types.ATXID,
	validator nipostValidator,
	logger *zap.Logger,
	opts ...VerifyChainOption,
) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *

	// iterate trough epochs, to get first valid, not malicious ATX with the biggest height
	new(types.ATXID), nil
}

// verify ATX-candidate by getting their dependencies (previous Atx, positioning ATX etc.)
// and verifying PoST for every dependency
