package compat

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hare4"
)

type weakCoin interface {
	Set(types.LayerID, bool) error
}

func ReportWeakcoin(ctx context.Context, logger *zap.Logger, from <-chan hare4.WeakCoinOutput, to weakCoin) {
	_ = "STUB: not implemented"
	return
}
