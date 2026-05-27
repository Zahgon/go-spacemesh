package fetch

import (
	"context"
	"io"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type handler struct {
	logger *zap.Logger
	db     sql.StateDatabase
	bs     *datastore.BlobStore
}

func newHandler(
	db sql.StateDatabase,
	bs *datastore.BlobStore,
	lg *zap.Logger,
) *handler {
	_ = "STUB: not implemented"
	return nil
}

// handleLegacyMaliciousIDsReq returns the IDs of all known malicious nodes.
func (h *handler) handleLegacyMaliciousIDsReq(ctx context.Context, _ p2p.Peer, _ []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleMaliciousIDsReq returns the IDs of all known malicious nodes.
func (h *handler) handleMaliciousIDsReq(ctx context.Context, _ p2p.Peer, _ []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) handleLegacyMaliciousIDsReqStream(ctx context.Context, _ p2p.Peer, _ []byte, s io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *handler) handleMaliciousIDsReqStream(ctx context.Context, _ p2p.Peer, _ []byte, s io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// handleEpochInfoReq returns the ATXs published in the specified epoch.
func (h *handler) handleEpochInfoReq(ctx context.Context, _ p2p.Peer, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleEpochInfoReq streams the ATXs published in the specified epoch.
func (h *handler) handleEpochInfoReqStream(ctx context.Context, _ p2p.Peer, msg []byte, s io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	retrieveCallback func(total int, id []byte) error
	retrieveFunc     func(retrieveCallback) error
)

func (h *handler) streamIDs(ctx context.Context, s io.ReadWriter, retrieve retrieveFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// If any IDs were sent:
// Response.Data already sent
// Response.Error has length 0

// If no ATX IDs were sent:
// Response.Data is just a single zero byte (length 0),
// but the length of Response.Data is 1 so we must send it
// Response.Error has length 0

// handleLayerDataReq returns all data in a layer, described in LayerData.
func (h *handler) handleLayerDataReq(ctx context.Context, _ p2p.Peer, req []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) handleLayerOpinionsReq2(ctx context.Context, _ p2p.Peer, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) handleCertReq(ctx context.Context, lid types.LayerID, bid types.BlockID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) handleHashReq(ctx context.Context, _ p2p.Peer, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) doHandleHashReq(ctx context.Context, data []byte, hint datastore.Hint) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this will iterate all requests and populate appropriate Responses, if there are any missing items they will not
// be included in the response at all

// add response to batch

func (h *handler) handleHashReqStream(ctx context.Context, _ p2p.Peer, msg []byte, s io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *handler) doHandleHashReqStream(
	ctx context.Context,
	msg []byte,
	s io.ReadWriter,
	hint datastore.Hint,
) error {
	_ = "STUB: not implemented"
	return nil
}

// At this point, nothing has been written yet, so we can report
// and error

func (h *handler) handleMeshHashReq(ctx context.Context, _ p2p.Peer, reqData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) handleMeshHashReqStream(ctx context.Context, _ p2p.Peer, reqData []byte, s io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}
