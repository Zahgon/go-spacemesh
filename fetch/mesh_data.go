package fetch

import (
	"context"
	"errors"
	"io"

	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/system"
)

var errBadRequest = errors.New("invalid request")

// GetAtxs gets the data for given atx IDs and validates them. Returns an error if at least one ATX cannot be fetched.
func (f *Fetch) GetAtxs(ctx context.Context, ids []types.ATXID, opts ...system.GetAtxOpt) error {
	_ = "STUB: not implemented"
	return nil
}

type dataReceiver func(context.Context, types.Hash32, p2p.Peer, []byte) error

type getHashesOpt func(*getHashesOpts)

func withLimiter(l limiter) getHashesOpt { _ = "STUB: not implemented"; return *new(getHashesOpt) }

func withHashCallback(callback func(types.Hash32, error)) getHashesOpt {
	_ = "STUB: not implemented"
	return *new(getHashesOpt)
}

func (f *Fetch) getHashes(
	ctx context.Context,
	hashes []types.Hash32,
	hint datastore.Hint,
	receiver dataReceiver,
	opts ...getHashesOpt,
) error {
	_ = "STUB: not implemented"
	return nil
}

// data is available locally

// GetActiveSet downloads activeset.
func (f *Fetch) GetActiveSet(ctx context.Context, set types.Hash32) error {
	_ = "STUB: not implemented"
	return nil
}

// LegacyMalfeasanceProofs gets legacy malfeasance proofs (v1) for the specified NodeIDs and validates them.
func (f *Fetch) LegacyMalfeasanceProofs(ctx context.Context, ids []types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// MalfeasanceProofs gets malfeasance proofs (v2) for the specified NodeIDs and validates them.
func (f *Fetch) MalfeasanceProofs(ctx context.Context, ids []types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBallots gets data for the specified BallotIDs and validates them.
func (f *Fetch) GetBallots(ctx context.Context, ids []types.BallotID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetProposals gets the data for given proposal IDs from peers.
func (f *Fetch) GetProposals(ctx context.Context, ids []types.ProposalID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBlocks gets the data for given block IDs from peers.
func (f *Fetch) GetBlocks(ctx context.Context, ids []types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetProposalTxs fetches the txs provided as IDs and validates them, returns an error if one TX failed to be fetched.
func (f *Fetch) GetProposalTxs(ctx context.Context, ids []types.TransactionID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockTxs fetches the txs provided as IDs and saves them, they will be validated
// before block is applied.
func (f *Fetch) GetBlockTxs(ctx context.Context, ids []types.TransactionID) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Fetch) getTxs(ctx context.Context, ids []types.TransactionID, receiver dataReceiver) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPoetProof gets poet proof from remote peer.
func (f *Fetch) GetPoetProof(ctx context.Context, id types.Hash32) error {
	_ = "STUB: not implemented"
	return nil
}

// data is available locally

// PoET proofs are concurrently stored in DB in two places:
// fetcher and nipost builder. Hence, it might happen that
// a proof had been inserted into the DB while the fetcher
// was fetching.

// LegacyMaliciousIDs gets the malicious IDs from the specified peer. Proofs for those IDs can be fetched via the
// legacy malfeasance proofs protocol (see also LegacyMalfeasanceProofs).
func (f *Fetch) LegacyMaliciousIDs(ctx context.Context, peer p2p.Peer) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MaliciousIDs gets the malicious IDs from the specified peer. Proofs for those IDs can be fetched via the malfeasance
// proof protocol (see also MalfeasanceProofs).
func (f *Fetch) MaliciousIDs(ctx context.Context, peer p2p.Peer) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLayerData get layer data from peers.
func (f *Fetch) GetLayerData(ctx context.Context, peer p2p.Peer, lid types.LayerID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetch) GetLayerOpinions(ctx context.Context, peer p2p.Peer, lid types.LayerID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetch) peerEpochInfoStreamed(ctx context.Context, peer p2p.Peer, epochBytes []byte) (*EpochData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PeerEpochInfo get the epoch info published in the given epoch from the specified peer.
func (f *Fetch) PeerEpochInfo(ctx context.Context, peer p2p.Peer, epoch types.EpochID) (*EpochData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetch) peerMeshHashesStreamed(ctx context.Context, peer p2p.Peer, reqBytes []byte) (*MeshHashes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetch) PeerMeshHashes(ctx context.Context, peer p2p.Peer, req *MeshHashRequest) (*MeshHashes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetch) GetCert(
	ctx context.Context,
	lid types.LayerID,
	bid types.BlockID,
	peers []p2p.Peer,
) (*types.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for generic data fetches by hash (ID for atx/block/proposal/ballot/tx), the check on whether the returned
// data matching the hash was done on the data handlers' path. for block certificate, there is no ID associated
// with it, hence the check here.
// however, certificate doesn't go through that path. it's requested by a separate protocol because a block
// certificate doesn't have an ID.

var ErrIgnore = errors.New("fetch: ignore")

type BatchError struct {
	Errors map[types.Hash32]error
	first  types.Hash32
}

func (b *BatchError) Empty() bool { _ = "STUB: not implemented"; return false }

func (b *BatchError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (b *BatchError) Add(id types.Hash32, err error) { _ = "STUB: not implemented"; return }

func (b *BatchError) Error() string { _ = "STUB: not implemented"; return "" }

func (b *BatchError) Ignore() bool { _ = "STUB: not implemented"; return false }

func (b *BatchError) IsIgnored(hash types.Hash32) bool { _ = "STUB: not implemented"; return false }

func readIDSlice[V any, H scale.DecodablePtr[V]](r io.Reader, slice *[]V, limit uint32) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
