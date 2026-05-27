package blocks

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hare3/eligibility"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

var (
	errMultipleCerts      = errors.New("multiple valid certificates")
	errInvalidCert        = errors.New("invalid certificate")
	errInvalidCertMsg     = errors.New("invalid cert msg")
	errUnexpectedMsg      = errors.New("unexpected lid")
	errBeaconNotAvailable = errors.New("beacon not available")
)

// CertConfig is the config for Certifier.
type CertConfig struct {
	CommitteeSize    int    `mapstructure:"committee-size"`
	CertifyThreshold int    `mapstructure:"-"` // not configurable, overwritten by CommitteeSize/2 + 1
	LayerBuffer      uint32 `mapstructure:"-"` // not configurable, overwritten by tortoise.Config.Zdist
	NumLayersToKeep  uint32 `mapstructure:"-"` // not configurable, overwritten by tortoise.Config.Zdist * 2
}

func defaultCertConfig() CertConfig { _ = "STUB: not implemented"; return *new(CertConfig) }

// CertifierOpt for configuring Certifier.
type CertifierOpt func(*Certifier)

// WithCertConfig defines cfg for Certifier.
func WithCertConfig(cfg CertConfig) CertifierOpt {
	_ = "STUB: not implemented"
	return *new(CertifierOpt)
}

// WithCertifierLogger defines logger for Certifier.
func WithCertifierLogger(logger *zap.Logger) CertifierOpt {
	_ = "STUB: not implemented"
	return *new(CertifierOpt)
}

type certInfo struct {
	registered, done bool
	totalEligibility uint16
	signatures       []types.CertifyMessage
}

// Certifier collects enough CertifyMessage for a given hare output and generate certificate.
type Certifier struct {
	logger *zap.Logger
	cfg    CertConfig
	once   sync.Once
	eg     errgroup.Group

	stop    func()
	stopped atomic.Bool

	db         sql.StateDatabase
	oracle     eligibility.Rolacle
	signers    map[types.NodeID]*signing.EdSigner
	edVerifier *signing.EdVerifier
	publisher  pubsub.Publisher
	layerClock layerClock
	beacon     system.BeaconGetter
	tortoise   system.Tortoise

	mu          sync.Mutex
	certifyMsgs map[types.LayerID]map[types.BlockID]*certInfo
	certCount   map[types.EpochID]int
}

// NewCertifier creates new block certifier.
func NewCertifier(
	db sql.StateDatabase,
	o eligibility.Rolacle,

	v *signing.EdVerifier,
	p pubsub.Publisher,
	lc layerClock,
	b system.BeaconGetter,
	tortoise system.Tortoise,
	opts ...CertifierOpt,
) *Certifier {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

// Start starts the background goroutine for periodic pruning.
func (c *Certifier) Start(ctx context.Context) {
	_ = "STUB: not implemented"

	// TODO(mafa): fix this
	return
}

// nolint:fatcontext

// Stop stops the outstanding goroutines.
func (c *Certifier) Stop() { _ = "STUB: not implemented"; return }

// not started

func (c *Certifier) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Certifier) prune() { _ = "STUB: not implemented"; return }

func (c *Certifier) createIfNeeded(lid types.LayerID, bid types.BlockID) *certInfo {
	_ = "STUB: not implemented"
	return nil
}

// RegisterForCert register to generate a certificate for the specified layer/block.
func (c *Certifier) RegisterForCert(ctx context.Context, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// CertifyIfEligible signs the hare output, along with its role proof as a certifier, and gossip the CertifyMessage
// if the node is eligible to be a certifier.
func (c *Certifier) CertifyIfEligible(ctx context.Context, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) certifySingleSigner(
	ctx context.Context,
	s *signing.EdSigner,
	lid types.LayerID,
	bid types.BlockID,
	beacon types.Beacon,
) error {
	_ = "STUB: not implemented"
	return nil
}

// not eligible

func newCertifyMsg(
	s *signing.EdSigner,
	lid types.LayerID,
	bid types.BlockID,
	proof types.VrfSignature,
	eligibility uint16,
) *types.CertifyMessage {
	_ = "STUB: not implemented"
	return nil
}

// NumCached returns the number of layers being cached in memory.
func (c *Certifier) NumCached() int { _ = "STUB: not implemented"; return 0 }

// HandleSyncedCertificate handles Certificate from sync.
func (c *Certifier) HandleSyncedCertificate(ctx context.Context, lid types.LayerID, cert *types.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) validateCert(ctx context.Context, logger *zap.Logger, cert *types.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) certified(lid types.LayerID, bid types.BlockID) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Certifier) expected(lid types.LayerID) bool { _ = "STUB: not implemented"; return false }

// only accept early msgs within a range and with limited size to prevent DOS

func (c *Certifier) HandleCertifyMessage(ctx context.Context, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleCertifyMessage is the gossip receiver for certify message.
func (c *Certifier) handleCertifyMessage(ctx context.Context, _ p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// should still gossip this msg to peers even when this node has created a certificate

func (c *Certifier) validate(ctx context.Context, logger *zap.Logger, msg types.CertifyMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) saveMessage(ctx context.Context, logger *zap.Logger, msg types.CertifyMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) tryGenCert(
	ctx context.Context,
	logger *zap.Logger,
	lid types.LayerID,
	bid types.BlockID,
	info *certInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

// do not try to generate a certificate for this block.
// wait for syncer to download from peers

func (c *Certifier) checkAndSave(
	ctx context.Context,
	logger *zap.Logger,
	lid types.LayerID,
	cert *types.Certificate,
) error {
	_ = "STUB: not implemented"
	return nil
}

// just verified

// the original live hare output should be trumped by a valid certificate

// stop processing certify message for this block

func (c *Certifier) addCertCount(epoch types.EpochID) { _ = "STUB: not implemented"; return }

func (c *Certifier) CertCount() map[types.EpochID]int { _ = "STUB: not implemented"; return nil }

func (c *Certifier) save(
	ctx context.Context,
	lid types.LayerID,
	cert *types.Certificate,
	valid, invalid []types.BlockID,
) error {
	_ = "STUB: not implemented"
	return nil
}
