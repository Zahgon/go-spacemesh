package ballots

import (
	"bytes"
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

func decodeBallot(id types.BallotID, body *bytes.Reader) (*types.Ballot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add ballot to the database.
func Add(db sql.Executor, ballot *types.Ballot) error { _ = "STUB: not implemented"; return nil }

// Has a ballot in the database.
func Has(db sql.Executor, id types.BallotID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to ballots with specified
// ids. For non-existent ballots, the corresponding items are set to -1.
func GetBlobSizes(db sql.Executor, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads ballot as an encoded blob, ready to be sent over the wire.
func LoadBlob(ctx context.Context, db sql.Executor, id []byte, b *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

// Get ballot with id from database.
func Get(db sql.Executor, id types.BallotID) (rst *types.Ballot, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(mafa): ideally there would be 2 types of ballots - one for persisting that maps to the db
// and one for the in-memory representation. The in-memory representation would have the information
// about the maliciousness of the smesher and is fetched via a service that adds malfeasance information
// either via DB query or from cached data. This is a temporary solution until we have a better way to handle this.

// Layer returns full body ballot for layer.
// NOTE: this function does not mark the ballot as malicious, if the smesher that published the ATX is!
func Layer(db sql.Executor, lid types.LayerID) (rst []*types.Ballot, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IDsInLayer returns ballots ids in the layer.
func IDsInLayer(db sql.Executor, lid types.LayerID) (rst []types.BallotID, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LayerBallotByNodeID returns any ballot by the specified NodeID in a given layer.
func LayerBallotByNodeID(db sql.Executor, lid types.LayerID, nodeID types.NodeID) (*types.Ballot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LatestLayer gets the highest layer with ballots.
func LatestLayer(db sql.Executor) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// FirstInEpoch returns the first ballot referencing the specified ATX in the epoch.
// NOTE: it does not mark the ballot as malicious, if the smesher that published the ATX is!
func FirstInEpoch(db sql.Executor, atx types.ATXID, epoch types.EpochID) (*types.Ballot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastInEpoch returns the last ballot referencing the specified ATX in the epoch.
// NOTE: it does not mark the ballot as malicious, if the smesher that published the ATX is!
func LastInEpoch(db sql.Executor, atx types.ATXID, epoch types.EpochID) (*types.Ballot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AllFirstInEpoch(db sql.Executor, epoch types.EpochID) ([]*types.Ballot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
