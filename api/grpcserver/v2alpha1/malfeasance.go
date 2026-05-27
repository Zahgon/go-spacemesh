package v2alpha1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2alpha1 "github.com/spacemeshos/api/release/go/spacemesh/v2alpha1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/events"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

func NewMalfeasanceService(db sql.StateDatabase, malHandler, legacyHandler malfeasanceInfo) *MalfeasanceService {
	_ = "STUB: not implemented"
	return nil
}

type MalfeasanceService struct {
	db         sql.StateDatabase
	info       malfeasanceInfo
	infoLegacy malfeasanceInfo
}

func (s *MalfeasanceService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *MalfeasanceService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MalfeasanceService) String() string { _ = "STUB: not implemented"; return "" }

func (s *MalfeasanceService) List(
	ctx context.Context,
	request *spacemeshv2alpha1.MalfeasanceRequest,
) (*spacemeshv2alpha1.MalfeasanceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only legacy proofs

// only new proofs

// both legacy and new proofs

type defaultEventProvider struct{}

func (defaultEventProvider) SubscribeMatched(
	request *spacemeshv2alpha1.MalfeasanceStreamRequest,
) (subscription, error) {
	_ = "STUB: not implemented"
	return *new(subscription), nil
}

type malStreamOpts func(*MalfeasanceStreamService)

func withEventProvider(provider eventProvider) malStreamOpts {
	_ = "STUB: not implemented"
	return *new(malStreamOpts)
}

func NewMalfeasanceStreamService(
	db sql.Executor,
	malfeasanceHandler,
	legacyHandler malfeasanceInfo,
	opts ...malStreamOpts,
) *MalfeasanceStreamService {
	_ = "STUB: not implemented"
	return nil
}

type MalfeasanceStreamService struct {
	db         sql.Executor
	info       malfeasanceInfo
	infoLegacy malfeasanceInfo
	events     eventProvider
}

func (s *MalfeasanceStreamService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *MalfeasanceStreamService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MalfeasanceStreamService) String() string { _ = "STUB: not implemented"; return "" }

func (s *MalfeasanceStreamService) Stream(
	request *spacemeshv2alpha1.MalfeasanceStreamRequest,
	stream spacemeshv2alpha1.MalfeasanceStreamService_StreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// process pending events first

// try again with the new handler

// try again with the new handler

func fetchMetaData(
	ctx context.Context,
	info malfeasanceInfo,
	id types.NodeID,
) *spacemeshv2alpha1.MalfeasanceProof {
	_ = "STUB: not implemented"
	return nil
}

func fetchFromDB(
	ctx context.Context,
	db sql.Executor,
	info malfeasanceInfo,
	request *spacemeshv2alpha1.MalfeasanceRequest,
) ([]*spacemeshv2alpha1.MalfeasanceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchLegacyFromDB(
	ctx context.Context,
	db sql.Executor,
	info malfeasanceInfo,
	request *spacemeshv2alpha1.MalfeasanceRequest,
) ([]*spacemeshv2alpha1.MalfeasanceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toMalfeasanceOps(filter *spacemeshv2alpha1.MalfeasanceRequest) (builder.Operations, error) {
	_ = "STUB: not implemented"
	return *new(builder.Operations), nil
}

type malfeasanceMatcher struct {
	*spacemeshv2alpha1.MalfeasanceStreamRequest
}

func (m *malfeasanceMatcher) match(event *events.EventMalfeasance) bool {
	_ = "STUB: not implemented"
	return false
}
