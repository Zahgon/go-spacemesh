package atxsdata

import (
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
)

func Warm(db sql.StateDatabase, keep types.EpochID, logger *zap.Logger, signers ...*signing.EdSigner) (*Data, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Warmup(db sql.Executor, cache *Data, keep types.EpochID, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
