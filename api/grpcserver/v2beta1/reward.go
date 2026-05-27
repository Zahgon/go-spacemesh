package v2beta1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2beta1 "github.com/spacemeshos/api/release/go/spacemesh/v2beta1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

func NewRewardStreamService(db sql.Executor) *RewardStreamService {
	_ = "STUB: not implemented"
	return nil
}

type RewardStreamService struct {
	db sql.Executor
}

func (s *RewardStreamService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *RewardStreamService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RewardStreamService) Stream(
	request *spacemeshv2beta1.RewardStreamRequest,
	stream spacemeshv2beta1.RewardStreamService_StreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RewardStreamService) fetchFromDB(
	ctx context.Context,
	ops builder.Operations,
) (<-chan *types.Reward, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buffered to avoid blocking, routine should exit immediately after sending an error

// exit if the context is canceled

func (s *RewardStreamService) String() string { _ = "STUB: not implemented"; return "" }

func NewRewardService(db sql.Executor) *RewardService { _ = "STUB: not implemented"; return nil }

type RewardService struct {
	db sql.Executor
}

func (s *RewardService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *RewardService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *RewardService) String() string { _ = "STUB: not implemented"; return "" }

func (s *RewardService) List(
	_ context.Context,
	request *spacemeshv2beta1.RewardRequest,
) (*spacemeshv2beta1.RewardList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toRewardRequest(filter *spacemeshv2beta1.RewardStreamRequest) *spacemeshv2beta1.RewardRequest {
	_ = "STUB: not implemented"
	return nil
}

func toRewardOperations(filter *spacemeshv2beta1.RewardRequest) (builder.Operations, error) {
	_ = "STUB: not implemented"
	return *new(builder.Operations), nil
}

func toReward(reward *types.Reward) *spacemeshv2beta1.Reward { _ = "STUB: not implemented"; return nil }

type rewardsMatcher struct {
	*spacemeshv2beta1.RewardStreamRequest
	ctx context.Context
}

func (m *rewardsMatcher) match(t *types.Reward) bool { _ = "STUB: not implemented"; return false }
