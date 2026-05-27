package rangesync

import (
	"context"
	"io"
	"sync/atomic"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/p2p"
)

type PairwiseSetSyncer struct {
	logger *zap.Logger
	r      Requester
	name   string
	cfg    RangeSetReconcilerConfig
	sent   atomic.Int64
	recv   atomic.Int64
	tracer Tracer
	clock  clockwork.Clock
}

func NewPairwiseSetSyncerInternal(
	logger *zap.Logger,
	r Requester,
	name string,
	cfg RangeSetReconcilerConfig,
	tracer Tracer,
	clock clockwork.Clock,
) *PairwiseSetSyncer {
	_ = "STUB: not implemented"
	return nil
}

func NewPairwiseSetSyncer(
	logger *zap.Logger,
	r Requester,
	name string,
	cfg RangeSetReconcilerConfig,
) *PairwiseSetSyncer {
	_ = "STUB: not implemented"
	return nil
}

func (pss *PairwiseSetSyncer) updateCounts(c *wireConduit) { _ = "STUB: not implemented"; return }

func (pss *PairwiseSetSyncer) createReconciler(os OrderedSet) *RangeSetReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (pss *PairwiseSetSyncer) Probe(
	ctx context.Context,
	peer p2p.Peer,
	os OrderedSet,
	x, y KeyBytes,
) (pr ProbeResult, err error) {
	_ = "STUB: not implemented"
	return *new(ProbeResult), nil
}

// If the conduit is not closed by this point, stop it
// interrupting any ongoing send operations

// Wait for the messages to be sent before closing the conduit

func (pss *PairwiseSetSyncer) requestCallback(
	ctx context.Context,
	stream io.ReadWriter,
	rsr *RangeSetReconciler,
	x, y KeyBytes,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pss *PairwiseSetSyncer) Sync(
	ctx context.Context,
	peer p2p.Peer,
	os OrderedSet,
	x, y KeyBytes,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pss *PairwiseSetSyncer) Serve(ctx context.Context, stream io.ReadWriter, os OrderedSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (pss *PairwiseSetSyncer) Register(d *Dispatcher, os OrderedSet) {
	_ = "STUB: not implemented"
	return
}

func (pss *PairwiseSetSyncer) Sent() int { _ = "STUB: not implemented"; return 0 }

func (pss *PairwiseSetSyncer) Received() int { _ = "STUB: not implemented"; return 0 }
