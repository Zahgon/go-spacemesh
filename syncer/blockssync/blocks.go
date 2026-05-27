package blockssync

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

//go:generate mockgen -typed -package=blockssync -destination=./mocks.go -source=./blocks.go

type blockFetcher interface {
	GetBlocks(context.Context, []types.BlockID) error
}

// Sync requests last specified blocks in background.
func Sync(ctx context.Context, logger *zap.Logger, requests <-chan []types.BlockID, fetcher blockFetcher) error {
	_ = "STUB: not implemented"
	return nil
}
