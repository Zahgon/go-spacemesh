package syncer

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/fetch"
	"github.com/spacemeshos/go-spacemesh/p2p"
)

var errNoPeers = errors.New("no peers")

// DataFetch contains the logic of fetching mesh data.
type DataFetch struct {
	fetcher

	logger *zap.Logger
}

// NewDataFetch creates a new DataFetch instance.
func NewDataFetch(
	fetch fetcher,
	lg *zap.Logger,
) *DataFetch {
	_ = "STUB: not implemented"
	return nil
}

// PollLayerData polls all peers for data in the specified layer.
func (d *DataFetch) PollLayerData(ctx context.Context, lid types.LayerID, peers ...p2p.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// registerLayerHashes registers hashes with the peer that provides these hashes.
func registerLayerHashes(fetcher fetcher, peer p2p.Peer, data *fetch.LayerData) {
	_ = "STUB: not implemented"
	return
}

func (d *DataFetch) PollLayerOpinions(
	ctx context.Context,
	lid types.LayerID,
	needCert bool,
	peers []p2p.Peer,
) ([]*fetch.LayerOpinion, []*types.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// note that we want to fetch block certificate for types.EmptyBlockID as well,
// but we don't need to register hash for the actual block fetching
