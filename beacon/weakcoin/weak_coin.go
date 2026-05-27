package weakcoin

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

var (
	errNotGenerated = errors.New("weakcoin not generated")
	errNotSmallest  = errors.New("proposal not smallest")
)

func defaultConfig() config { _ = "STUB: not implemented"; return *new(config) }

// ~1mb given the size of Message is ~100b

type config struct {
	Threshold           types.VrfSignature
	NextRoundBufferSize int
	MaxRound            types.RoundID
}

//go:generate scalegen -types Message,VrfMessage

// Message defines weak coin message format.
type Message struct {
	Epoch        types.EpochID
	Round        types.RoundID
	Unit         uint32
	NodeID       types.NodeID
	VRFSignature types.VrfSignature
}

// VrfMessage is the payload for the signature of `Message`.
type VrfMessage struct {
	Type  types.EligibilityType // always types.EligibilityBeaconWC
	Nonce types.VRFPostIndex
	Epoch types.EpochID
	Round types.RoundID
	Unit  uint32
}

// OptionFunc for optional configuration adjustments.
type OptionFunc func(*WeakCoin)

// WithLog changes logger.
func WithLog(logger *zap.Logger) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithMaxRound changes max round.
func WithMaxRound(round types.RoundID) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithThreshold changes signature threshold.
func WithThreshold(threshold types.VrfSignature) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithNextRoundBufferSize changes size of the buffer for messages from future rounds.
func WithNextRoundBufferSize(size int) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// messageTime interface exists so that we can pass an object from the beacon
// package to the weakCoinPackage (as does allowance), this is indicative of a
// circular dependency, probably the weak coin should be merged with the beacon
// package.
// Issue: https://github.com/spacemeshos/go-spacemesh/issues/4199
type messageTime interface {
	WeakCoinProposalSendTime(epoch types.EpochID, round types.RoundID) time.Time
}

// New creates an instance of weak coin protocol.
func New(
	publisher pubsub.Publisher,
	verifier vrfVerifier,
	nonceFetcher nonceFetcher,
	allowance allowance,
	msgTime messageTime,
	opts ...OptionFunc,
) *WeakCoin {
	_ = "STUB: not implemented"
	return nil
}

// WeakCoin implementation of the protocol.
type WeakCoin struct {
	logger       *zap.Logger
	config       config
	verifier     vrfVerifier
	nonceFetcher nonceFetcher
	publisher    pubsub.Publisher

	mu                         sync.RWMutex
	epochStarted, roundStarted bool
	epoch                      types.EpochID
	round                      types.RoundID
	smallest                   *types.VrfSignature
	allowance                  allowance
	// nextRoundBuffer is used to optimistically buffer messages from the next round.
	nextRoundBuffer []Message
	coins           map[types.RoundID]bool
	msgTime         messageTime
}

// Get the result of the coin flip in this round. It is only valid in between StartEpoch/EndEpoch
// and only after CompleteRound was called.
func (wc *WeakCoin) Get(ctx context.Context, epoch types.EpochID, round types.RoundID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// StartEpoch notifies that epoch is started and we can accept messages for this epoch.
func (wc *WeakCoin) StartEpoch(ctx context.Context, epoch types.EpochID) {
	_ = "STUB: not implemented"
	return
}

// FinishEpoch completes epoch.
func (wc *WeakCoin) FinishEpoch(ctx context.Context, epoch types.EpochID) {
	_ = "STUB: not implemented"
	return
}

type Participant struct {
	Signer vrfSigner
	Nonce  types.VRFPostIndex
}

// StartRound process any buffered messages for this round and broadcast our proposal.
func (wc *WeakCoin) StartRound(ctx context.Context, round types.RoundID, participants []Participant) {
	_ = "STUB: not implemented"
	return
}

func (wc *WeakCoin) updateProposal(ctx context.Context, message Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (wc *WeakCoin) prepareProposal(
	epoch types.EpochID,
	signer vrfSigner,
	nonce types.VRFPostIndex,
	round types.RoundID,
) ([]byte, types.VrfSignature) {
	_ = "STUB: not implemented"
	return nil, *new(types.VrfSignature)
}

func (wc *WeakCoin) publishProposal(
	ctx context.Context,
	epoch types.EpochID,
	signer vrfSigner,
	nonce types.VRFPostIndex,
	round types.RoundID,
) {
	_ = "STUB: not implemented"
	return
}

// FinishRound computes coinflip based on proposals received in this round.
// After it is called new proposals for this round won't be accepted.
func (wc *WeakCoin) FinishRound(ctx context.Context) { _ = "STUB: not implemented"; return }

func (wc *WeakCoin) updateSmallest(ctx context.Context, sig types.VrfSignature) error {
	_ = "STUB: not implemented"
	return nil
}

func (wc *WeakCoin) aboveThreshold(proposal types.VrfSignature) bool {
	_ = "STUB: not implemented"
	return false
}

func (wc *WeakCoin) encodeProposal(
	epoch types.EpochID,
	nonce types.VRFPostIndex,
	round types.RoundID,
	unit uint32,
) []byte {
	_ = "STUB: not implemented"
	return nil
}
