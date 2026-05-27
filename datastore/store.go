package datastore

import (
	"context"
	"errors"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/proposals/store"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/activesets"
	"github.com/spacemeshos/go-spacemesh/sql/atxs"
	"github.com/spacemeshos/go-spacemesh/sql/ballots"
	"github.com/spacemeshos/go-spacemesh/sql/blocks"
	"github.com/spacemeshos/go-spacemesh/sql/identities"
	"github.com/spacemeshos/go-spacemesh/sql/poets"
	"github.com/spacemeshos/go-spacemesh/sql/transactions"
)

var ErrNotFound = errors.New("not found")

type VrfNonceKey struct {
	ID    types.NodeID
	Epoch types.EpochID
}

// CachedDB is simply a database injected with cache.
type CachedDB struct {
	sql.Database

	// cache is optional in tests. It MUST be set for the 'App'
	// for properly checking malfeasance.
	atxsdata *atxsdata.Data

	atxCache      *lru.Cache[types.ATXID, *types.ActivationTx]
	vrfNonceCache *lru.Cache[VrfNonceKey, types.VRFPostIndex]

	// used to coordinate db update and cache
	mu               sync.Mutex
	malfeasanceCache *lru.Cache[types.NodeID, []byte]
}

type Config struct {
	// ATXSize must be larger than the sum of all ATXs in last 2 epochs to be effective
	ATXSize         int `mapstructure:"atx-size"`
	MalfeasanceSize int `mapstructure:"malfeasance-size"`
}

func DefaultConfig() Config {
	_ = "STUB: not implemented"

	// NOTE(dshulyak) there are several places where this cache is used, but none of them require to hold
	// all atxs in memory. those places should eventually be refactored to load necessary data from db.
	return *new(Config)
}

type cacheOpts struct {
	cfg      Config
	atxsdata *atxsdata.Data
}

type Opt func(*cacheOpts)

func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithConsensusCache(c *atxsdata.Data) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// NewCachedDB create an instance of a CachedDB.
func NewCachedDB(db sql.StateDatabase, lg *zap.Logger, opts ...Opt) *CachedDB {
	_ = "STUB: not implemented"
	return nil
}

// MalfeasanceProof returns the malfeasance proof for the given node ID. This function is thread safe and will return
// an error if the proof is not found in the ATX DB.
// Deprecated: use functions in the `sql/identities` and `sql/malfeasance` packages.
func (db *CachedDB) MalfeasanceProof(id types.NodeID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CacheMalfeasanceProof caches the malfeasance proof for the given node ID. This function is thread safe.
// Deprecated: caching is done by the sql database automatically.
func (db *CachedDB) CacheMalfeasanceProof(id types.NodeID, proof []byte) {
	_ = "STUB: not implemented"
	return
}

// VRFNonce returns the VRF nonce of for the given node in the given epoch. This function is thread safe and will
// return an error if the nonce is not found in the ATX DB.
func (db *CachedDB) VRFNonce(id types.NodeID, epoch types.EpochID) (types.VRFPostIndex, error) {
	_ = "STUB: not implemented"
	return *new(types.VRFPostIndex), nil
}

// GetAtx returns the ATX by the given ID. This function is thread safe and will return an error if the ID
// is not found in the ATX DB.
func (db *CachedDB) GetAtx(id types.ATXID) (*types.ActivationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Previous retrieves the list of previous ATXs for the given ATX ID.
// Deprecated: replaced by atxs.Previous.
func (db *CachedDB) Previous(id types.ATXID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// IterateMalfeasanceProofs iterates over all malfeasance proofs in the database and calls the provided callback on
	// each.
	// Deprecated: replaced by identities.IterateOps and malfeasance.IterateOps.
}

func (db *CachedDB) IterateMalfeasanceProofs(
	iter func(types.NodeID, []byte) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// MaxHeightAtx returns the ATX ID with the maximum height.
// Deprecated: replaced by atxs.GetIDWithMaxHeight.
func (db *CachedDB) MaxHeightAtx() (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// Hint marks which DB should be queried for a certain provided hash.
type Hint string

// DB hints per DB.
const (
	NoHint            Hint = ""
	BallotDB          Hint = "ballotDB"
	BlockDB           Hint = "blocksDB"
	ProposalDB        Hint = "proposalDB"
	ATXDB             Hint = "ATXDB"
	TXDB              Hint = "TXDB"
	POETDB            Hint = "POETDB"
	LegacyMalfeasance Hint = "malfeasance"
	Malfeasance       Hint = "malfeasance2"
	ActiveSet         Hint = "activeset"
)

// NewBlobStore returns a BlobStore.
func NewBlobStore(db sql.StateDatabase, proposals *store.Store) *BlobStore {
	_ = "STUB: not implemented"
	return nil
}

// SetMalfeasanceProvider sets the malfeasance provider dependency.
//
// TODO(mafa): this is a hack because of a cyclic dependency between the packages
//
//	malfeasance2 -> fetcher -> datastore -> malfeasance2
func (bs *BlobStore) SetMalfeasanceProvider(p MalfeasanceProvider) {
	_ = "STUB: not implemented"
	return

	//go:generate mockgen -typed -package=datastore -destination=./mocks.go -source=./store.go
}

type MalfeasanceProvider interface {
	ProofByID(ctx context.Context, nodeID types.NodeID) ([]byte, error)
}

// BlobStore gets data as a blob to serve direct fetch requests.
type BlobStore struct {
	DB          sql.StateDatabase
	proposals   *store.Store
	malfeasance MalfeasanceProvider
}

type (
	loadBlobFunc func(ctx context.Context, db sql.Executor, key []byte, blob *sql.Blob) error
	blobSizeFunc func(db sql.Executor, ids [][]byte) (sizes []int, err error)
)

var loadBlobDispatch = map[Hint]loadBlobFunc{
	ATXDB: func(ctx context.Context, db sql.Executor, key []byte, blob *sql.Blob) error {
		_, err := atxs.LoadBlob(ctx, db, key, blob)
		return err
	},
	BallotDB:          ballots.LoadBlob,
	BlockDB:           blocks.LoadBlob,
	TXDB:              transactions.LoadBlob,
	POETDB:            poets.LoadBlob,
	LegacyMalfeasance: identities.LoadMalfeasanceBlob,
	ActiveSet:         activesets.LoadBlob,
}

var blobSizeDispatch = map[Hint]blobSizeFunc{
	ATXDB:             atxs.GetBlobSizes,
	BallotDB:          ballots.GetBlobSizes,
	BlockDB:           blocks.GetBlobSizes,
	TXDB:              transactions.GetBlobSizes,
	POETDB:            poets.GetBlobSizes,
	LegacyMalfeasance: identities.GetBlobSizes,
	ActiveSet:         activesets.GetBlobSizes,
}

func (bs *BlobStore) loadProposal(key []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BlobStore) proposalSizes(keys [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bs *BlobStore) loadMalfeasance(key []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BlobStore) malfeasanceSizes(keys [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob gets an blob as bytes by an object ID as bytes.
func (bs *BlobStore) LoadBlob(ctx context.Context, hint Hint, key []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to the specified ids. For
// non-existent objects, the corresponding items are set to -1.
func (bs *BlobStore) GetBlobSizes(hint Hint, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bs *BlobStore) Has(hint Hint, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
