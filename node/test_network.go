package node

import (
	"testing"

	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/log"
)

// NewTestNetwork creates a network of fully connected nodes.
func NewTestNetwork(tb testing.TB, conf config.Config, l log.Log, size int) []*TestApp {
	_ = "STUB: not implemented"
	// We need to set this global state
	return nil
}

// set to generate coinbase

// To save an epoch of startup time, we bootstrap (meaning we manually set
// it) the beacon for epoch 2 so that in epoch 3 hare can start.

// This context is used to call Start on a node and canceling it will
// shutdown the node. (Hence no timeout has been set).

// Copy config, services don't modify their config, so we just need to
// be careful here when we modify any pointer values in the config.

// Note that we must call cleanup after all calls to t.TempDir since calls
// to cleanup are executed in LIFO fashion (similar to defer) and t.TempDir
// internally calls Cleanup to delete the dir. By calling Cleanup here
// we ensure that the apps have been shut-down before attempting to delete
// the temp dirs.

// Wait for nodes to shutdown

// we are in a cleanup function and `tb.Context()` is already canceled
// nolint:usetesting

// Connect all nodes to each other

func NewApp(tb testing.TB, conf *config.Config, l log.Log) *App {
	_ = "STUB: not implemented"
	return nil
}

type TestApp struct {
	*App
	Conn *grpc.ClientConn
}
