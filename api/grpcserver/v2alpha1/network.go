package v2alpha1

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2alpha1 "github.com/spacemeshos/api/release/go/spacemesh/v2alpha1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

func NewNetworkService(
	genesisTime time.Time,
	genesisID types.Hash20,
	layerDuration time.Duration,
	labelsPerUnit uint64,
) *NetworkService {
	_ = "STUB: not implemented"
	return nil
}

type NetworkService struct {
	genesisTime   time.Time
	genesisID     types.Hash20
	layerDuration time.Duration
	labelsPerUnit uint64
}

func (s *NetworkService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *NetworkService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *NetworkService) String() string { _ = "STUB: not implemented"; return "" }

func (s *NetworkService) Info(context.Context,
	*spacemeshv2alpha1.NetworkInfoRequest,
) (*spacemeshv2alpha1.NetworkInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
