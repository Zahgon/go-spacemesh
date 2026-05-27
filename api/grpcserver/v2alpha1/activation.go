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

func NewActivationStreamService(db sql.Executor) *ActivationStreamService {
	_ = "STUB: not implemented"
	return nil
}

type ActivationStreamService struct {
	db sql.Executor
}

func (s *ActivationStreamService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *ActivationStreamService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivationStreamService) String() string { _ = "STUB: not implemented"; return "" }

func (s *ActivationStreamService) Stream(
	request *spacemeshv2alpha1.ActivationStreamRequest,
	stream spacemeshv2alpha1.ActivationStreamService_StreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivationStreamService) fetchFromDB(
	ctx context.Context,
	ops builder.Operations,
) (<-chan *types.ActivationTx, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buffered to avoid blocking, routine should exit immediately after sending an error

// exit if the context is canceled

func toAtx(atx *types.ActivationTx) *spacemeshv2alpha1.Activation {
	_ = "STUB: not implemented"
	return nil
}

func NewActivationService(db sql.Executor, goldenAtx types.ATXID) *ActivationService {
	_ = "STUB: not implemented"
	return nil
}

type ActivationService struct {
	goldenAtx types.ATXID
	db        sql.Executor
}

func (s *ActivationService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *ActivationService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *ActivationService) String() string { _ = "STUB: not implemented"; return "" }

func (s *ActivationService) List(
	ctx context.Context,
	request *spacemeshv2alpha1.ActivationRequest,
) (*spacemeshv2alpha1.ActivationList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// every full atx is ~1KB. 100 atxs is ~100KB.

func (s *ActivationService) ActivationsCount(
	ctx context.Context,
	request *spacemeshv2alpha1.ActivationsCountRequest,
) (*spacemeshv2alpha1.ActivationsCountResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ActivationService) Highest(
	_ context.Context,
	_ *spacemeshv2alpha1.HighestRequest,
) (*spacemeshv2alpha1.HighestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAtxRequest(filter *spacemeshv2alpha1.ActivationStreamRequest) *spacemeshv2alpha1.ActivationRequest {
	_ = "STUB: not implemented"
	return nil
}

func toAtxOperations(filter *spacemeshv2alpha1.ActivationRequest) (builder.Operations, error) {
	_ = "STUB: not implemented"
	return *new(builder.Operations), nil
}

type atxsMatcher struct {
	*spacemeshv2alpha1.ActivationStreamRequest
	ctx context.Context
}

func (m *atxsMatcher) match(t *events.ActivationTx) bool { _ = "STUB: not implemented"; return false }
