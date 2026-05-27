package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/transactions"
)

// TransactionService exposes transaction data, and a submit tx endpoint.
type TransactionService struct {
	db        sql.StateDatabase
	publisher pubsub.Publisher // P2P Swarm
	mesh      meshAPI          // Mesh
	conState  conservativeState
	syncer    syncer
	txHandler txValidator
}

// RegisterService registers this service with a grpc server instance.
func (s *TransactionService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *TransactionService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (s *TransactionService) String() string { _ = "STUB: not implemented"; return "" }

// NewTransactionService creates a new grpc service using config data.
func NewTransactionService(
	db sql.StateDatabase,
	publisher pubsub.Publisher,
	msh meshAPI,
	conState conservativeState,
	syncer syncer,
	txHandler txValidator,
) *TransactionService {
	_ = "STUB: not implemented"
	return nil
}

func (s *TransactionService) ParseTransaction(
	ctx context.Context,
	in *pb.ParseTransactionRequest,
) (*pb.ParseTransactionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubmitTransaction allows a new tx to be submitted.
func (s *TransactionService) SubmitTransaction(
	ctx context.Context,
	in *pb.SubmitTransactionRequest,
) (*pb.SubmitTransactionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get transaction and status for a given txid. It's not an error if we cannot find the tx,
// we just return all nils.
func (s *TransactionService) getTransactionAndStatus(
	txID types.TransactionID,
) (*types.Transaction, pb.TransactionState_TransactionState) {
	_ = "STUB: not implemented"
	return nil, *new(pb.TransactionState_TransactionState)
}

// TransactionsState returns current tx data for one or more txs.
func (s *TransactionService) TransactionsState(
	_ context.Context,
	in *pb.TransactionsStateRequest,
) (*pb.TransactionsStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert the incoming txid into a known type

// Look up data for this tx. If it's unknown to us, status will be zero (unspecified).

// If the tx is unknown to us, add an empty placeholder

// STREAMS

// TransactionsStateStream exposes a stream of tx data.
func (s *TransactionService) TransactionsStateStream(
	in *pb.TransactionsStateStreamRequest,
	stream pb.TransactionService_TransactionsStateStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// The tx channel tells us about newly received and newly created transactions
// The layer channel tells us about status updates

// Filter

// If the tx was just invalidated, we already know its state.
// If not, read it from the database.

// Don't match on any other transactions

// Transaction objects do not have an associated status. The status we assign them here is based on the
// status of the layer (really, the block) they're contained in. Here, we receive updates to layer status.
// In order to update tx status, we have to read every transaction in the layer.

// In order to read transactions, we first need to read layer blocks

// Filter for any matching transactions in the reported layer

// create a set for the block transaction IDs

// Since the txid coming in from the API does not have a fixed length (see
// https://github.com/spacemeshos/api/issues/130), we need to convert it from a slice to a fixed
// size array before we can convert it into a TransactionID object. We don't need to worry about
// error handling, since copy intelligently copies only what it can. If the resulting TransactionID
// is invalid, an error will be thrown below.

// if there is an ID corresponding to inputTxID in the block

// StreamResults allows to query historical results and subscribe to live data using the same filter.
func (s *TransactionService) StreamResults(
	in *pb.TransactionResultsRequest,
	stream pb.TransactionService_StreamResultsServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func castResult(rst *types.TransactionWithResult) *pb.TransactionResult {
	_ = "STUB: not implemented"
	return nil
}

type resultsMatcher transactions.ResultsFilter

func (m resultsMatcher) match(rst *types.TransactionWithResult) bool {
	_ = "STUB: not implemented"
	return false
}
