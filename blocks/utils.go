package blocks

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/txs"
)

var (
	errNodeHasBadMeshHash   = errors.New("node has different mesh hash from majority")
	errProposalTxMissing    = errors.New("proposal tx not found")
	errProposalTxHdrMissing = errors.New("proposal tx missing header")
	errDuplicateATX         = errors.New("multiple proposals with same ATX")
)

type meshState struct {
	hash  types.Hash32
	count int
}

type proposalMetadata struct {
	ctx        context.Context
	lid        types.LayerID
	proposals  []*types.Proposal
	tids       []types.TransactionID
	tickHeight uint64
	rewards    []types.AnyReward
	optFilter  bool
}

func getProposalMetadata(
	ctx context.Context,
	logger *zap.Logger,
	db sql.StateDatabase,
	atxs *atxsdata.Data,
	cfg Config,
	lid types.LayerID,
	proposals []*types.Proposal,
) (*proposalMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBlockTXs(
	logger *zap.Logger,
	mtxs []*types.MeshTransaction,
	blockSeed []byte,
	gasLimit uint64,
) ([]types.TransactionID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize a Mersenne Twister with the block seed and use it as a source of randomness for
// a Fisher-Yates shuffle of the sorted transaction IDs.

func prune(
	logger *zap.Logger,
	tids []types.TransactionID,
	byTid map[types.TransactionID]*txs.NanoTX,
	gasLimit uint64,
) []types.TransactionID {
	_ = "STUB: not implemented"
	return nil
}

func toUint64Slice(b []byte) []uint64 { _ = "STUB: not implemented"; return nil }

func rewardInfoAndHeight(
	cfg Config,
	db sql.StateDatabase,
	atxs *atxsdata.Data,
	props []*types.Proposal,
) (uint64, []types.AnyReward, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// keys in order so that everyone generates same block
