package v2alpha1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2alpha1 "github.com/spacemeshos/api/release/go/spacemesh/v2alpha1"
	"github.com/spacemeshos/go-scale"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
	"github.com/spacemeshos/go-spacemesh/system"
)

// transactionConState is an API to validate transaction.
type transactionConState interface {
	Validation(raw types.RawTx) system.ValidationRequest
	HasEvicted(tid types.TransactionID) (bool, error)
}

// transactionValidator is the API to validate and cache transactions.
type transactionValidator interface {
	VerifyAndCacheTx(context.Context, []byte) error
}

func NewTransactionStreamService() *TransactionStreamService { _ = "STUB: not implemented"; return nil }

type TransactionStreamService struct{}

func (s *TransactionStreamService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *TransactionStreamService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TransactionStreamService) Stream(
	request *spacemeshv2alpha1.TransactionStreamRequest,
	stream spacemeshv2alpha1.TransactionStreamService_StreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TransactionStreamService) String() string { _ = "STUB: not implemented"; return "" }

func NewTransactionService(db sql.Executor, conState transactionConState,
	syncer syncer, validator transactionValidator,
	publisher pubsub.Publisher,
) *TransactionService {
	_ = "STUB: not implemented"
	return nil
}

type TransactionService struct {
	db        sql.Executor
	conState  transactionConState
	syncer    syncer
	validator transactionValidator
	publisher pubsub.Publisher // P2P Swarm
}

func (s *TransactionService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *TransactionService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *TransactionService) String() string { _ = "STUB: not implemented"; return "" }

func (s *TransactionService) List(
	ctx context.Context,
	request *spacemeshv2alpha1.TransactionRequest,
) (*spacemeshv2alpha1.TransactionList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TransactionService) ParseTransaction(
	ctx context.Context,
	request *spacemeshv2alpha1.ParseTransactionRequest,
) (*spacemeshv2alpha1.ParseTransactionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TransactionService) SubmitTransaction(
	ctx context.Context,
	request *spacemeshv2alpha1.SubmitTransactionRequest,
) (*spacemeshv2alpha1.SubmitTransactionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TransactionService) EstimateGas(
	ctx context.Context,
	request *spacemeshv2alpha1.EstimateGasRequest,
) (*spacemeshv2alpha1.EstimateGasResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toTransactionOperations(filter *spacemeshv2alpha1.TransactionRequest) (builder.Operations, error) {
	_ = "STUB: not implemented"
	return *new(builder.Operations), nil
}

func (s *TransactionService) toTx(
	ctx context.Context,
	tx *types.MeshTransaction,
	result *types.TransactionResult,
	includeResult, includeState bool,
) *spacemeshv2alpha1.TransactionResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *TransactionService) convertTxResult(
	result *types.TransactionResult,
) spacemeshv2alpha1.TransactionResult_Status {
	_ = "STUB: not implemented"
	return *new(spacemeshv2alpha1.TransactionResult_Status)
}

func (s *TransactionService) convertTxState(
	ctx context.Context, tx *types.MeshTransaction,
) *spacemeshv2alpha1.TransactionState {
	_ = "STUB: not implemented"
	return nil
}

func decodeTxArgs(decoder *scale.Decoder) (uint8, *core.Address, scale.Encodable, error) {
	_ = "STUB: not implemented"
	return 0, nil, *new(scale.Encodable), nil
}

func toTxContents(rawTx []byte) (*spacemeshv2alpha1.TransactionContents,
	spacemeshv2alpha1.Transaction_TransactionType, error,
) {
	_ = "STUB: not implemented"
	return nil, *new(spacemeshv2alpha1.Transaction_TransactionType), nil
}
