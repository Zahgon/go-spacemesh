package atxsync

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

func getMissing(db sql.StateDatabase, set []types.ATXID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Download specified set of atxs from peers in the network.
//
// Actual retry interval will be between [retryInterval, 2*retryInterval].
func Download(
	ctx context.Context,
	retryInterval time.Duration,
	logger *zap.Logger,
	db sql.StateDatabase,
	fetcher system.AtxFetcher,
	set []types.ATXID,
) error {
	_ = "STUB: not implemented"
	return nil
}
