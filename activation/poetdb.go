package activation

import (
	"context"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/spacemeshos/poet/shared"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// PoetDbOptions are options for PoetDb.
type PoetDbOptions struct {
	cacheSize int
}

type PoetDbOption func(*PoetDbOptions)

// WithCacheSize sets the cache size for PoetDb.
func WithCacheSize(size int) PoetDbOption { _ = "STUB: not implemented"; return *new(PoetDbOption) }

// PoetDb is a database for PoET proofs.
type PoetDb struct {
	sqlDB               sql.StateDatabase
	poetProofsDbRequest singleflight.Group
	poetProofsLru       *lru.Cache[types.PoetProofRef, *types.PoetProofMessage]
	logger              *zap.Logger
}

// NewPoetDb returns a new PoET handler.
func NewPoetDb(db sql.StateDatabase, log *zap.Logger, opts ...PoetDbOption) (*PoetDb, error) {
	_ = "STUB: not implemented"
	return nil,

		// in last epochs there are 45 proofs per epoch, with each of them nearly 140KB
		// 200 is set not to keep multiple epochs, but to account for unexpected growth
		// select round_id, count(*), max(length(poet)) from poets group by round_id;
		// 25|42|146200
		// 26|45|145936
		// 27|45|145738
		// 28|45|145903
		nil
}

// HasProof returns true if the database contains a proof with the given reference, or false otherwise.
func (db *PoetDb) HasProof(proofRef types.PoetProofRef) bool {
	_ = "STUB: not implemented"
	return false
}

// ValidateAndStore validates and stores a new PoET proof.
func (db *PoetDb) ValidateAndStore(ctx context.Context, proofMessage *types.PoetProofMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateAndStoreMsg validates and stores a new PoET proof.
func (db *PoetDb) ValidateAndStoreMsg(ctx context.Context, expHash types.Hash32, _ p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate validates a new PoET proof.
func (db *PoetDb) Validate(
	root []byte,
	proof types.PoetProof,
	poetID []byte,
	roundID string,
	signature types.EdSignature,
) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(noamnelke): validate signature (or extract public key and use for salting merkle hashes)

// StoreProof saves the poet proof in local db.
func (db *PoetDb) StoreProof(ctx context.Context, ref types.PoetProofRef, proofMessage *types.PoetProofMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *PoetDb) GetProofRef(poetID []byte, roundID string) (types.PoetProofRef, error) {
	_ = "STUB: not implemented"
	return *new(types.PoetProofRef), nil
}

// GetProofMessage returns the originally received PoET proof message.
func (db *PoetDb) GetProofMessage(proofRef types.PoetProofRef) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Proof returns full proof.
func (db *PoetDb) Proof(proofRef types.PoetProofRef) (*types.PoetProof, *types.Hash32, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (db *PoetDb) ProofForRound(poetID []byte, roundID string) (*types.PoetProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calcRoot(leaves []types.Hash32) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func validatePoet(membershipRoot []byte, merkleProof shared.MerkleProof, leafCount uint64) error {
	_ = "STUB: not implemented"
	return nil
}
