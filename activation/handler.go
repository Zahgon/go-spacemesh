package activation

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/system"
)

var (
	errKnownAtx      = errors.New("known atx")
	errMalformedData = fmt.Errorf("%w: malformed data", pubsub.ErrValidationReject)
	errWrongHash     = fmt.Errorf("%w: incorrect hash", pubsub.ErrValidationReject)
	errMaliciousATX  = errors.New("malicious atx")
)

type atxVersion struct {
	// epoch since this version is valid
	publish types.EpochID
	types.AtxVersion
}

type AtxVersions map[types.EpochID]types.AtxVersion

func (v AtxVersions) asSlice() []atxVersion { _ = "STUB: not implemented"; return nil }

func (v AtxVersions) Validate() error { _ = "STUB: not implemented"; return nil }

// Handler processes the atxs received from all nodes and their validity status.
type Handler struct {
	local    p2p.Peer
	logger   *zap.Logger
	versions []atxVersion

	// inProgress is used to avoid processing the same ATX multiple times in parallel.
	inProgress singleflight.Group

	v1 *HandlerV1
	v2 *HandlerV2
}

// HandlerOption is a functional option for the handler.
type HandlerOption func(*Handler)

func WithAtxVersions(v AtxVersions) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func WithTickSize(tickSize uint64) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

// NewHandler returns a data handler for ATX.
func NewHandler(
	local p2p.Peer,
	cdb *datastore.CachedDB,
	atxsdata *atxsdata.Data,
	edVerifier *signing.EdVerifier,
	c layerClock,
	fetcher system.Fetcher,
	goldenATXID types.ATXID,
	nipostValidator nipostValidator,
	malPublisher atxMalfeasancePublisher,
	legacyMalPublisher legacyMalfeasancePublisher,
	beacon atxReceiver,
	tortoise system.Tortoise,
	lg *zap.Logger,
	opts ...HandlerOption,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

// HandleSyncedAtx handles atxs received by sync.
func (h *Handler) HandleSyncedAtx(ctx context.Context, expHash types.Hash32, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleGossipAtx handles the atx gossip data channel.
func (h *Handler) HandleGossipAtx(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) determineVersion(msg []byte) (*types.AtxVersion, error) {
	_ = "STUB: not implemented"
	// The first field of all ATXs is the publish epoch, which
	// we use to determine the version of the ATX.
	return nil, nil
}

type opaqueAtx interface {
	ID() types.ATXID
}

func (h *Handler) decodeATX(msg []byte) (atx opaqueAtx, err error) {
	_ = "STUB: not implemented"
	return *new(opaqueAtx), nil
}

func (h *Handler) handleAtx(ctx context.Context, expHash types.Hash32, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}
