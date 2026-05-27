package tests

import (
	"context"
	"testing"
	"time"

	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	pb2 "github.com/spacemeshos/api/release/go/spacemesh/v2beta1"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/systest/chaos"
	"github.com/spacemeshos/go-spacemesh/systest/cluster"
	"github.com/spacemeshos/go-spacemesh/systest/testcontext"
)

const (
	attempts = 3
)

var retryBackoff = 10 * time.Second

func sendTransactions(
	ctx context.Context,
	logger *zap.Logger,
	cl *cluster.Cluster,
	first, stop uint32,
	receiver types.Address,
	batch, amount int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// give some time for a previous layer to be applied
// TODO(dshulyak) introduce api that simply subscribes to internal clock
// and outputs events when the tick for the layer is available

// retry on failure 3 times

// wait before retrying

func submitTransaction(ctx context.Context, tx []byte, node *cluster.NodeClient) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stateHashStream(
	ctx context.Context,
	node *cluster.NodeClient,
	logger *zap.Logger,
	collector func(*pb.GlobalStateStreamResponse) (bool, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func watchLayers(
	ctx context.Context,
	eg *errgroup.Group,
	node *cluster.NodeClient,
	logger *zap.Logger,
	collector func(*pb.LayerStreamResponse) (bool, error),
) {
	_ = "STUB: not implemented"
	return
}

func layersStream(
	ctx context.Context,
	node *cluster.NodeClient,
	logger *zap.Logger,
	collector func(*pb.LayerStreamResponse) (bool, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func malfeasanceStream(
	ctx context.Context,
	node *cluster.NodeClient,
	logger *zap.Logger,
	collector func(*pb2.MalfeasanceProof) (bool, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func waitGenesis(ctx *testcontext.Context, node *cluster.NodeClient) error {
	_ = "STUB: not implemented"
	return nil
}

func waitLayer(ctx *testcontext.Context, node *cluster.NodeClient, lid uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func waitTransaction(ctx context.Context, eg *errgroup.Group, client *cluster.NodeClient, id []byte) {
	_ = "STUB: not implemented"
	return
}

func watchTransactionResults(
	ctx context.Context,
	client *cluster.NodeClient,
	log *zap.Logger,
	collector func(*pb.TransactionResult) (bool, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func watchProposals(
	ctx context.Context,
	eg *errgroup.Group,
	client *cluster.NodeClient,
	log *zap.Logger,
	collector func(*pb.Proposal) (bool, error),
) {
	_ = "STUB: not implemented"
	return
}

func prettyHex(buf []byte) string { _ = "STUB: not implemented"; return "" }

func scheduleChaos(
	ctx context.Context,
	eg *errgroup.Group,
	client *cluster.NodeClient,
	logger *zap.Logger,
	from, to uint32,
	action func(context.Context) (chaos.Teardown, error),
) {
	_ = "STUB: not implemented"
	return
}

func currentLayer(ctx context.Context, tb testing.TB, client *cluster.NodeClient) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func waitAll(tctx *testcontext.Context, cl *cluster.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func nextFirstLayer(current, size uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func getNonce(ctx context.Context, client *cluster.NodeClient, address types.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func currentBalance(ctx context.Context, client *cluster.NodeClient, address types.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func submitSpawn(ctx context.Context, cluster *cluster.Cluster, account int, client *cluster.NodeClient) error {
	_ = "STUB: not implemented"
	return nil
}

func submitSpend(
	ctx context.Context,
	cluster *cluster.Cluster,
	account int,
	receiver types.Address,
	amount, nonce uint64,
	client *cluster.NodeClient,
) error {
	_ = "STUB: not implemented"
	return nil
}

func syncedNodes(ctx context.Context, cl *cluster.Cluster) []*cluster.NodeClient {
	_ = "STUB: not implemented"
	return nil
}

func isSynced(ctx context.Context, node *cluster.NodeClient) bool {
	_ = "STUB: not implemented"
	return false
}

func getLayer(ctx context.Context, node *cluster.NodeClient, lid uint32) (*pb.Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVerifiedLayer(ctx context.Context, node *cluster.NodeClient) (*pb.Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type txClient struct {
	account cluster.Account
	node    *cluster.NodeClient
}

func (c *txClient) nonce(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *txClient) submit(ctx context.Context, tx []byte) (*txRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type txRequest struct {
	node *cluster.NodeClient
	txid []byte

	rst *pb.TransactionResult
}

func (r *txRequest) wait(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *txRequest) result(ctx context.Context) (*pb.TransactionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// eof without result - transaction wasn't applied yet
