package v2alpha1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2alpha1 "github.com/spacemeshos/api/release/go/spacemesh/v2alpha1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

type accountConState interface {
	GetProjection(types.Address) (uint64, uint64)
}

func NewAccountService(db sql.Executor, conState accountConState) *AccountService {
	_ = "STUB: not implemented"
	return nil
}

type AccountService struct {
	db       sql.Executor
	conState accountConState
}

func (s *AccountService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *AccountService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *AccountService) String() string { _ = "STUB: not implemented"; return "" }

func (s *AccountService) List(
	_ context.Context,
	request *spacemeshv2alpha1.AccountRequest,
) (*spacemeshv2alpha1.AccountList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAccountOperations(filter *spacemeshv2alpha1.AccountRequest) (builder.Operations, error) {
	_ = "STUB: not implemented"
	return *new(builder.Operations), nil
}
