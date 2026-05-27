package certifier

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type PoetCert struct {
	Data      []byte
	Signature []byte
}

func AddCertificate(db sql.Executor, nodeID types.NodeID, cert PoetCert, cerifierID []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteCertificate(db sql.Executor, nodeID types.NodeID, certifierID []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func Certificate(db sql.Executor, nodeID types.NodeID, certifierID []byte) (*PoetCert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
