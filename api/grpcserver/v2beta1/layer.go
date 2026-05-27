package v2beta1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2beta1 "github.com/spacemeshos/api/release/go/spacemesh/v2beta1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/events"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
	"github.com/spacemeshos/go-spacemesh/sql/layers"
)

func NewLayerStreamService(db sql.Executor) *LayerStreamService {
	_ = "STUB: not implemented"
	return nil
}

type LayerStreamService struct {
	db sql.Executor
}

func (s *LayerStreamService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *LayerStreamService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *LayerStreamService) Stream(
	request *spacemeshv2beta1.LayerStreamRequest,
	stream spacemeshv2beta1.LayerStreamService_StreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *LayerStreamService) fetchFromDB(
	ctx context.Context,
	ops builder.Operations,
) (<-chan *spacemeshv2beta1.Layer, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buffered to avoid blocking, routine should exit immediately after sending an error

// exit if the stream context is canceled

func toLayerRequest(filter *spacemeshv2beta1.LayerStreamRequest) *spacemeshv2beta1.LayerRequest {
	_ = "STUB: not implemented"
	return nil
}

func (s *LayerStreamService) String() string { _ = "STUB: not implemented"; return "" }

func NewLayerService(db sql.Executor) *LayerService { _ = "STUB: not implemented"; return nil }

type LayerService struct {
	db sql.Executor
}

func (s *LayerService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *LayerService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *LayerService) String() string { _ = "STUB: not implemented"; return "" }

func (s *LayerService) List(
	ctx context.Context,
	request *spacemeshv2beta1.LayerRequest,
) (*spacemeshv2beta1.LayerList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toLayerOperations(filter *spacemeshv2beta1.LayerRequest) (builder.Operations, error) {
	_ = "STUB: not implemented"
	return *new(builder.Operations), nil
}

func toLayer(layer *layers.Layer) *spacemeshv2beta1.Layer { _ = "STUB: not implemented"; return nil }

type layersMatcher struct {
	*spacemeshv2beta1.LayerStreamRequest
}

func (m *layersMatcher) match(l *events.LayerUpdate) bool { _ = "STUB: not implemented"; return false }

func convertEventStatus(eventStatus int) (status spacemeshv2beta1.Layer_LayerStatus) {
	_ = "STUB: not implemented"
	return *new(spacemeshv2beta1.Layer_LayerStatus)
}
