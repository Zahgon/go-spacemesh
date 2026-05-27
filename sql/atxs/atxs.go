package atxs

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

const (
	CacheKindEpochATXs sql.QueryCacheKind = "epoch-atxs"
	CacheKindATXBlob   sql.QueryCacheKind = "atx-blob"
)

// Query to retrieve ATXs.
// Can't use inner join for the ATX blob here b/c this will break
// filters that refer to the id column.
const fieldsQuery = `select
atxs.id, atxs.nonce, atxs.base_tick_height, atxs.tick_count, atxs.pubkey, atxs.effective_num_units,
atxs.received, atxs.epoch, atxs.sequence, atxs.coinbase, atxs.validity, atxs.commitment_atx, atxs.weight,
atxs.marriage_atx`

const fullQuery = fieldsQuery + ` from atxs`

type decoderCallback func(*types.ActivationTx) bool

func decoder(fn decoderCallback) sql.Decoder { _ = "STUB: not implemented"; return *new(sql.Decoder) }

// Note: received is assigned `0` for checkpointed ATXs.
// We treat `0` as 'zero time'.
// We could use `NULL` instead, but the column has "NOT NULL" constraint.
// In future, consider changing the schema to allow `NULL` for received.

func load(db sql.Executor, query string, enc sql.Encoder) (*types.ActivationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get gets an ATX by a given ATX ID.
func Get(db sql.Executor, id types.ATXID) (*types.ActivationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetByEpochAndNodeID gets any ATX by the specified NodeID published in the given epoch.
func GetByEpochAndNodeID(
	db sql.Executor,
	epoch types.EpochID,
	nodeID types.NodeID,
) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// Has checks if an ATX exists by a given ATX ID.
func Has(db sql.Executor, id types.ATXID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func CommitmentATX(db sql.Executor, nodeID types.NodeID) (id types.ATXID, err error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// IdentityExists checks if an identity has ever published an ATX.
func IdentityExists(db sql.Executor, nodeID types.NodeID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Coinbase retrieves the last coinbase address used by the given node ID.
func Coinbase(db sql.Executor, id types.NodeID) (types.Address, error) {
	_ = "STUB: not implemented"
	return *new(types.Address), nil
}

// GetFirstIDByNodeID gets the initial ATX ID for a given node ID.
func GetFirstIDByNodeID(db sql.Executor, nodeID types.NodeID) (id types.ATXID, err error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// GetLastIDByNodeID gets the last ATX ID for a given node ID.
func GetLastIDByNodeID(db sql.Executor, nodeID types.NodeID) (id types.ATXID, err error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// PrevIDByNodeID returns the previous ATX ID for a given node ID and public epoch, ignoring current if present in the
// DB (to avoid returning the same ATX ID).
// It returns the newest ATX ID containing a PoST of the given node ID that was published in or before the given epoch.
func PrevIDByNodeID(
	db sql.Executor,
	current types.ATXID,
	nodeID types.NodeID,
	pubEpoch types.EpochID,
) (id types.ATXID, err error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// GetIDByEpochAndNodeID gets an ATX ID for a given epoch and node ID.
func GetIDByEpochAndNodeID(db sql.Executor, epoch types.EpochID, nodeID types.NodeID) (id types.ATXID, err error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// GetIDsByEpoch gets ATX IDs for a given epoch.
func GetIDsByEpoch(ctx context.Context, db sql.Executor, epoch types.EpochID) (ids []types.ATXID, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VRFNonce gets the VRF nonce of a smesher for a given epoch.
func VRFNonce(db sql.Executor, id types.NodeID, epoch types.EpochID) (nonce types.VRFPostIndex, err error) {
	_ = "STUB: not implemented"
	return *new(types.VRFPostIndex), nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to ATXs with specified
// ids. For non-existent ATXs, the corresponding items are set to -1.
func GetBlobSizes(db sql.Executor, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads ATX as an encoded blob, ready to be sent over the wire.
func LoadBlob(ctx context.Context, db sql.Executor, id []byte, blob *sql.Blob) (types.AtxVersion, error) {
	_ = "STUB: not implemented"
	return *new(types.AtxVersion), nil
}

// We don't use the provided blob in this case to avoid
// caching references to the underlying slice (subsequent calls would modify it).

func getBlob(ctx context.Context, db sql.Executor, id []byte, blob *sql.Blob) (types.AtxVersion, error) {
	_ = "STUB: not implemented"
	return *new(types.AtxVersion), nil
}

// The migration adding the version column does not set it to 1 for existing ATXs.
// Thus, both values 0 and 1 mean V1.

// Previous gets all previous ATXs for a given ATX ID.
func Previous(db sql.Executor, id types.ATXID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Index is returned in descending order, so the first one defines the length of the slice.

// NonceByID retrieves VRFNonce corresponding to the specified ATX ID.
func NonceByID(db sql.Executor, id types.ATXID) (nonce types.VRFPostIndex, err error) {
	_ = "STUB: not implemented"
	return *new(types.VRFPostIndex), nil
}

func Add(db sql.Executor, atx *types.ActivationTx, blob types.AtxBlob) error {
	_ = "STUB: not implemented"
	return nil
}

func AddBlob(db sql.Executor, id types.ATXID, blob []byte, version types.AtxVersion) error {
	_ = "STUB: not implemented"
	return nil
}

type Filter func(types.ATXID) bool

func FilterAll(types.ATXID) bool {
	_ = "STUB: not implemented"

	// GetIDWithMaxHeight returns the ID of the atx from the last 2 epoch with the highest (or tied for the highest)
	// tick height. It is possible that some poet servers are faster than others and the network ends up having its
	// highest ticked atx still in previous epoch and the atxs building on top of it have not been published yet.
	// Selecting from the last two epochs to strike a balance between being fair to honest miners while not giving
	// unfair advantage for malicious actors who retroactively publish a high tick atx many epochs back.
	return false
}

func GetIDWithMaxHeight(db sql.Executor, pref types.NodeID, filter Filter) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// Results are ordered by height, so we can stop once we see a lower height.

// We can stop on the first ATX if `pref` is empty.

// prefer atxs from `pref`

type CheckpointAtx struct {
	ID             types.ATXID
	Epoch          types.EpochID
	CommitmentATX  types.ATXID
	MarriageATX    *types.ATXID
	VRFNonce       types.VRFPostIndex
	BaseTickHeight uint64
	TickCount      uint64
	SmesherID      types.NodeID
	Sequence       uint64
	Coinbase       types.Address
	// total effective units
	NumUnits uint32
	// actual units of each included smesher
	Units map[types.NodeID]uint32
}

// LatestN returns the latest N ATXs per smesher.
func LatestN(db sql.Executor, n int) ([]CheckpointAtx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddCheckpointed(db sql.Executor, catx *CheckpointAtx) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: should a checkpointed ATX reference its real previous ATX?

// All gets all atx IDs.
func All(db sql.Executor) ([]types.ATXID, error) { _ = "STUB: not implemented"; return nil, nil }

// LatestEpoch with atxs.
func LatestEpoch(db sql.Executor) (types.EpochID, error) {
	_ = "STUB: not implemented"
	return *new(types.EpochID), nil
}

// IterateAtxsData iterate over data used for consensus.
func IterateAtxsData(
	db sql.Executor,
	from, to types.EpochID,
	fn func(
		id types.ATXID,
		node types.NodeID,
		epoch types.EpochID,
		coinbase types.Address,
		weight uint64,
		base uint64,
		height uint64,
		nonce types.VRFPostIndex,
	) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// filtering in CODE is no longer effective on some machines in epoch 29

func SetValidity(db sql.Executor, id types.ATXID, validity types.Validity) error {
	_ = "STUB: not implemented"
	return nil
}

func IterateAtxsOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(*types.ActivationTx) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func CountAtxsByOps(db sql.Executor, operations builder.Operations) (count uint32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// IterateForGrading selects every atx from publish epoch and joins identities to load malfeasance proofs if they exist.
func IterateForGrading(
	db sql.Executor,
	epoch types.EpochID,
	fn func(id types.ATXID, atxTime, proofTime int64, weight uint64) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// no malfeasance

// legacy malfeasance

// new malfeasance

// both legacy and new malfeasance, take oldest one

func IterateAtxsWithMalfeasance(
	db sql.Executor,
	publish types.EpochID,
	fn func(atx *types.ActivationTx, malicious bool) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func IterateAtxIdsWithMalfeasance(
	db sql.Executor,
	publish types.EpochID,
	fn func(id types.ATXID, malicious bool) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PrevATXCollisions returns all ATXs with the same prevATX as the given ATX ID from the same node ID.
// It is used to detect double-publishing and double poet registrations.
// The ATXs returned are ordered by received time so that the first one is the one that was seen first by the node.
func PrevATXCollisions(db sql.Executor, prev types.ATXID, id types.NodeID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Units(db sql.Executor, atxID types.ATXID, nodeID types.NodeID) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FindDoublePublish finds 2 distinct ATXIDs that the given identity contributed PoST to in the given epoch.
//
// It is guaranteed to return 2 distinct ATXs when the error is nil.
// It works by finding an ATX in the given epoch that has a PoST contribution from the given identity.
// - `epoch` is looked up in the `atxs` table by matching atxid.
func FindDoublePublish(db sql.Executor, nodeID types.NodeID, epoch types.EpochID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AllUnits(db sql.Executor, id types.ATXID) (map[types.NodeID]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetPost(
	db sql.Executor,
	atxID, prev types.ATXID,
	prevIndex int,
	id types.NodeID,
	units uint32,
	publish types.EpochID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AtxWithPrevious returns the ATX ID that has the given ATX ID as its previous ATX.
func AtxWithPrevious(db sql.Executor, prev types.ATXID, id types.NodeID) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

// Find 2 distinct merged ATXs (having the same marriage ATX) in the same epoch.
func MergeConflict(db sql.Executor, marriage types.ATXID, publish types.EpochID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
