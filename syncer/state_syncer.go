package syncer

import (
	"context"
	"errors"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
)

var errMeshHashDiverged = errors.New("mesh hash diverged with peer")

type peerOpinion struct {
	prevAggHash types.Hash32
	peer        p2p.Peer
}

func (s *Syncer) stateSynced() bool { _ = "STUB: not implemented"; return false }

func (s *Syncer) processLayers(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// used to make sure we only resync from the same peer once during each run.

// layers should be processed in order. once we skip one layer, there is no point
// continuing with later layers. return on error

// certificate is effective only within hare distance, outside it we don't vote according to other rules.

// there is no point in tortoise counting after every single layer, in fact it is wasteful.
// we periodically invoke counting to evict executed layers.

// we try once as we cannot assume that all blocks that reported as missing are valid
// we need to continue download layers, as they may be deemed as invalid after counting more votes

func (s *Syncer) processLayerOpinions(ctx context.Context, lid types.LayerID, resyncPeers map[p2p.Peer]struct{}) {
	_ = "STUB: not implemented"
	return
}

func (s *Syncer) needCert(ctx context.Context, lid types.LayerID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Syncer) layerOpinions(
	ctx context.Context,
	lid types.LayerID,
) ([]*peerOpinion, []*types.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *Syncer) checkMeshAgreement(
	ctx context.Context,
	lid types.LayerID,
	opinions []*peerOpinion,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) adopt(ctx context.Context, lid types.LayerID, certs []*types.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) certCutoffLayer() types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

func (s *Syncer) adoptCert(ctx context.Context, lid types.LayerID, cert *types.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// it is safer to ask for block after certificate was downloaded, as we know that we ask for block signed by
// committee so we should not reorder this

// in GetBlocks call block will be also passed to tortoise.OnBlock

// see https://github.com/spacemeshos/go-spacemesh/issues/2507 for implementation rationale.
func (s *Syncer) ensureMeshAgreement(
	ctx context.Context,
	diffLayer types.LayerID,
	opinions []*peerOpinion,
	resyncPeers map[p2p.Peer]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// getting the atx IDs targeting this epoch

// node and peer has different state. check if peer has valid ATXs to back up its opinions

// if the node cannot download the ATXs claimed by this peer, it does not trust this peer's mesh

// find the divergent layer and adopt the peer's mesh from there

// ideally syncer should import opinions from peer and let tortoise decide whether to rerun
// verifying tortoise or enter full tortoise mode.
// however, for genesis running full tortoise with 10_000 layers as a sliding window is completely
// viable. so here we don't sync opinions, just the data.
// see https://github.com/spacemeshos/go-spacemesh/issues/2507

// clear the agreement cache after syncing new data
