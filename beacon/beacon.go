package beacon

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/spacemeshos/fixed"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/beacon/metrics"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/types/result"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/system"
)

const (
	numEpochsToKeep = 3
)

var (
	errBeaconNotCalculated = errors.New("beacon is not calculated for this epoch")
	errZeroEpochWeight     = errors.New("zero epoch weight provided")
	errDifferentBeacon     = errors.New("different beacons detected")
	errGenesis             = errors.New("genesis")
	errNodeNotSynced       = errors.New("nodes not synced")
	errProtocolRunning     = errors.New("last beacon protocol still running")
	errNoProposals         = errors.New("no proposals")
)

type (
	proposals    = struct{ valid, potentiallyValid proposalSet }
	allVotes     = struct{ support, against proposalSet }
	beaconWeight = struct {
		ballots        map[types.BallotID]struct{}
		totalWeight    fixed.Fixed
		numEligibility int
	}
)

// Opt for configuring beacon protocol.
type Opt func(*ProtocolDriver)

// WithLogger defines logger for beacon.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithConfig defines protocol parameters.
func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func withWeakCoin(wc coin) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// New returns a new ProtocolDriver.
func New(
	publisher pubsub.Publisher,
	edVerifier *signing.EdVerifier,
	vrfVerifier vrfVerifier,
	cdb *datastore.CachedDB,
	clock layerClock,
	opts ...Opt,
) *ProtocolDriver {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

type participant struct {
	signer *signing.EdSigner
	nonce  types.VRFPostIndex
}

func (s participant) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// ProtocolDriver is the driver for the beacon protocol.
type ProtocolDriver struct {
	inProtocol uint64
	logger     *zap.Logger
	eg         errgroup.Group
	startOnce  sync.Once
	closed     chan struct{}

	config    Config
	sync      system.SyncStateProvider
	publisher pubsub.Publisher

	signers  map[types.NodeID]*signing.EdSigner
	weakCoin coin

	edVerifier   *signing.EdVerifier
	vrfVerifier  vrfVerifier
	nonceFetcher nonceFetcher
	theta        *big.Float

	clock    layerClock
	msgTimes *messageTimes
	cdb      *datastore.CachedDB

	mu sync.RWMutex

	// these fields are separate from state because we don't want to pre-maturely create a state
	// for proposals in epoch currentEpoch+1, because we may not have all ATXs for currentEpoch yet.
	roundInProgress                        types.RoundID
	earliestProposalTime, earliestVoteTime time.Time
	// states for the current epoch and the next epoch. we accept early proposals for the next epoch.
	// state for an epoch is created on demand:
	// - when the first proposal of the epoch comes in
	// - when the protocol starts running for the epoch
	// whichever happens first.
	// we start accepting early proposals for the next epoch pd.config.GracePeriodDuration before
	// the next epoch begins.
	states map[types.EpochID]*state

	// beacons store calculated beacons as the result of the beacon protocol.
	// the map key is the target epoch when beacon is used. if a beacon is calculated in epoch N, it will be used
	// in epoch N+1.
	beacons map[types.EpochID]types.Beacon
	// ballotsBeacons store beacons collected from ballots.
	// the map key is the epoch when the ballot is published. the beacon value is calculated in the
	// previous epoch and used in the current epoch.
	ballotsBeacons map[types.EpochID]map[types.Beacon]*beaconWeight

	resultsMtx sync.Mutex
	results    chan result.Beacon

	// metrics
	metricsCollector *metrics.BeaconMetricsCollector
}

// SetSyncState updates sync state provider. Must be executed only once.
func (pd *ProtocolDriver) SetSyncState(sync system.SyncStateProvider) {
	_ = "STUB: not implemented"
	return
}

// Start starts listening for layers and outputs.
func (pd *ProtocolDriver) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) UpdateBeacon(epoch types.EpochID, beacon types.Beacon) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes ProtocolDriver.
func (pd *ProtocolDriver) Close() { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) onResult(epoch types.EpochID, beacon types.Beacon) {
	_ = "STUB: not implemented"
	return
}

// Results notifies waiter that beacon for a target epoch has completed.
func (pd *ProtocolDriver) Results() <-chan result.Beacon { _ = "STUB: not implemented"; return nil }

// isClosed returns true if the beacon protocol is shutting down.
func (pd *ProtocolDriver) isClosed() bool { _ = "STUB: not implemented"; return false }

func (pd *ProtocolDriver) OnAtx(atx *types.ActivationTx) { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) minerAtxHdr(
	epoch types.EpochID,
	nodeID types.NodeID,
) (*types.ActivationTx, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (pd *ProtocolDriver) MinerAllowance(epoch types.EpochID, nodeID types.NodeID) uint32 {
	_ = "STUB: not implemented"
	return 0
}

// ReportBeaconFromBallot reports the beacon value in a ballot along with the smesher's weight unit.
func (pd *ProtocolDriver) ReportBeaconFromBallot(
	epoch types.EpochID,
	ballot *types.Ballot,
	beacon types.Beacon,
	weightPer fixed.Fixed,
) {
	_ = "STUB: not implemented"
	return
}

// already has beacon. i.e. we had participated in the beacon protocol during the last epoch

func (pd *ProtocolDriver) recordBeacon(
	epochID types.EpochID,
	ballot *types.Ballot,
	beacon types.Beacon,
	weightPer fixed.Fixed,
) {
	_ = "STUB: not implemented"
	return
}

// using ballot weight here because we're sampling miners for the beacon value recorded
// in the ballots. we can't just take the entire ATX weight because otherwise, a "whale"
// ATX would dominate the voting.

// checks if we have recorded this ballot before

func (pd *ProtocolDriver) findMajorityBeacon(epoch types.EpochID) types.Beacon {
	_ = "STUB: not implemented"
	return *new(types.Beacon)
}

// GetBeacon returns the beacon for the specified epoch or an error if it doesn't exist.
func (pd *ProtocolDriver) GetBeacon(targetEpoch types.EpochID) (types.Beacon, error) {
	_ = "STUB: not implemented"
	return *new(types.Beacon), nil
}

func (pd *ProtocolDriver) getBeacon(epoch types.EpochID) types.Beacon {
	_ = "STUB: not implemented"
	return *new(types.Beacon)
}

func (pd *ProtocolDriver) setBeacon(targetEpoch types.EpochID, beacon types.Beacon) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) getPersistedBeacon(epoch types.EpochID) (types.Beacon, error) {
	_ = "STUB: not implemented"
	return *new(types.Beacon), nil
}

func (pd *ProtocolDriver) setBeginProtocol(ctx context.Context) { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) setEndProtocol(ctx context.Context) { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) isInProtocol() bool { _ = "STUB: not implemented"; return false }

func (pd *ProtocolDriver) initEpochStateIfNotPresent(logger *zap.Logger, target types.EpochID) (*state, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// w1 is the weight units at δ before the end of the previous epoch, used to calculate `thresholdStrict`
// w2 is the weight units at the end of the previous epoch, used to calculate `threshold`

func (pd *ProtocolDriver) setProposalTimeForNextEpoch() { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) setupEpoch(logger *zap.Logger, epoch types.EpochID) (*state, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pd *ProtocolDriver) cleanupEpoch(epoch types.EpochID) { _ = "STUB: not implemented"; return }

// listens to new layers.
func (pd *ProtocolDriver) listenEpochs(ctx context.Context) { _ = "STUB: not implemented"; return }

func (pd *ProtocolDriver) setRoundInProgress(round types.RoundID) {
	_ = "STUB: not implemented"
	return
}

func (pd *ProtocolDriver) onNewEpoch(ctx context.Context, epoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) runProtocol(ctx context.Context, epoch types.EpochID, st *state) {
	_ = "STUB: not implemented"
	return
}

// K rounds passed
// After K rounds had passed, tally up votes for proposals using simple tortoise vote counting

func calcBeacon(logger *zap.Logger, set proposalSet) types.Beacon {
	_ = "STUB: not implemented"
	return *new(types.Beacon)
}

// Beacon should appear to have the same entropy as the initial proposals, hence cropping it
// to the same size as the proposal

func (pd *ProtocolDriver) runProposalPhase(ctx context.Context, epoch types.EpochID, st *state) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) sendProposal(
	ctx context.Context,
	epoch types.EpochID,
	s participant,
	checker eligibilityChecker,
) {
	_ = "STUB: not implemented"
	return
}

// runConsensusPhase runs K voting rounds and returns result from last weak coin round.
func (pd *ProtocolDriver) runConsensusPhase(ctx context.Context, epoch types.EpochID, st *state) (allVotes, error) {
	_ = "STUB: not implemented"
	return *new(allVotes), nil
}

// For K rounds: In each round that lasts δ, wait for votes to come in.
// For next rounds,
// wait for δ time, and construct a message that points to all messages from previous round received by δ.
// rounds 1 to K

// First round

// shared lock is fine as sorting doesn't modify the state

// Subsequent rounds

// note that votes after this call will _not_ be counted towards our votes
// for this round, as the late votes can be cast after the weak coin is revealed. we
// count them towards our votes in the next round.

func (pd *ProtocolDriver) markProposalPhaseFinished(st *state, finishedAt time.Time) {
	_ = "STUB: not implemented"
	return
}

func (pd *ProtocolDriver) calcVotesBeforeWeakCoin(logger *zap.Logger, st *state) (allVotes, proposalList) {
	_ = "STUB: not implemented"
	return *new(allVotes), *new(proposalList)
}

func (pd *ProtocolDriver) sendFirstRoundVote(
	ctx context.Context,
	msg FirstVotingMessageBody,
	signer *signing.EdSigner,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) getFirstRoundVote(epoch types.EpochID, nodeID types.NodeID) (proposalList, error) {
	_ = "STUB: not implemented"
	return *new(proposalList), nil
}

func (pd *ProtocolDriver) sendFollowingVote(
	ctx context.Context,
	epoch types.EpochID,
	round types.RoundID,
	ownCurrentRoundVotes allVotes,
	signer *signing.EdSigner,
) error {
	_ = "STUB: not implemented"
	return nil
}

type proposalChecker struct {
	threshold       *big.Int
	thresholdStrict *big.Int
}

func createProposalChecker(logger *zap.Logger, conf Config, numEarlyATXs, numATXs int) eligibilityChecker {
	_ = "STUB: not implemented"
	return *new(eligibilityChecker)
}

func (pc *proposalChecker) PassStrictThreshold(proposal types.VrfSignature) bool {
	_ = "STUB: not implemented"
	return false
}

func (pc *proposalChecker) PassThreshold(proposal types.VrfSignature) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO(nkryuchkov): Consider replacing github.com/ALTree/bigfloat.
func atxThresholdFraction(kappa int, q *big.Rat, numATXs int) *big.Float {
	_ = "STUB: not implemented"
	return nil
}

// threshold(k, q, W) = 1 - (2 ^ (- (k/((1-q)*W))))
// Floating point: 1 - math.Pow(2.0, -(float64(tb.config.Kappa)/((1.0-tb.config.Q)*float64(numATXs))))
// Fixed point:

// TODO(nkryuchkov): Consider having a generic function for probabilities.
func atxThreshold(kappa int, q *big.Rat, numATXs int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func buildSignedProposal(
	ctx context.Context,
	logger *zap.Logger,
	signer vrfSigner,
	epoch types.EpochID,
	nonce types.VRFPostIndex,
) types.VrfSignature {
	_ = "STUB: not implemented"
	return *new(types.VrfSignature)
}

func buildProposal(epoch types.EpochID, nonce types.VRFPostIndex) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) sendToGossip(ctx context.Context, protocol string, serialized []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (pd *ProtocolDriver) gatherMetricsData() ([]*metrics.BeaconStats, *metrics.BeaconStats) {
	_ = "STUB: not implemented"
	return nil, nil
}

// whether the beacon for the next epoch is calculated

// messageTimes provides methods to determine the intended send time of
// messages spceific to the beacon protocol.
type messageTimes struct {
	clock layerClock
	conf  Config
}

// proposalSendTime returns the time at which a proposal is sent for an epoch.
func (mt *messageTimes) proposalSendTime(epoch types.EpochID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// firstVoteSendTime returns the time at which the first vote is sent for an epoch.
func (mt *messageTimes) firstVoteSendTime(epoch types.EpochID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// followupVoteSendTime returns the time at which the followup votes are sent for an epoch and round.
func (mt *messageTimes) followupVoteSendTime(epoch types.EpochID, round types.RoundID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// WeakCoinProposalSendTime returns the time at which the weak coin proposals are sent for an epoch and round.
func (mt *messageTimes) WeakCoinProposalSendTime(epoch types.EpochID, round types.RoundID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
