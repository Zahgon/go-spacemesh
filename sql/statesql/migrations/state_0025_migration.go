package migrations

import (
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type migration0025 struct {
	edVerifier *signing.EdVerifier
}

var _ sql.Migration = &migration0025{}

func New0025Migration(config config.Config) *migration0025 { _ = "STUB: not implemented"; return nil }

func (*migration0025) Name() string { _ = "STUB: not implemented"; return "" }

func (*migration0025) Order() int { _ = "STUB: not implemented"; return 0 }

func (*migration0025) Rollback() error { _ = "STUB: not implemented"; return nil }

func (m *migration0025) Apply(db sql.Executor, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// we only care about invalid prev ATX proofs

func (m *migration0025) Validate(data wire.ProofData) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}
