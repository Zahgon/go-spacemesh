package wire

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate scalegen

// MarryProof is a proof that a NodeID is married to another NodeID.
type MarryProof struct {
	// MarriageCertificatesRoot and its proof that it is contained in the ATX.
	MarriageCertificatesRoot  MarriageCertificatesRoot
	MarriageCertificatesProof MarriageCertificatesRootProof `scale:"max=32"`

	// The signature of the certificate and the proof that the certificate is contained in the MarriageRoot at
	// the given index.
	Certificate      MarriageCertificate
	CertificateProof MarriageCertificateProof `scale:"max=32"`
	CertificateIndex uint32
}

func createMarryProof(db sql.Executor, atx *ActivationTxV2, nodeID types.NodeID) (MarryProof, error) {
	_ = "STUB: not implemented"
	return *new(MarryProof), nil
}

// special case of the self signed certificate of the ATX publisher

// Valid returns an error if the proof is invalid. It checks that `nodeID` signed a certificate to marry `smesherID`
// and it was included in the ATX with the given `atxID`.
func (p MarryProof) Valid(
	malValidator MalfeasanceValidator,
	atxID types.ATXID,
	smesherID types.NodeID,
	nodeID types.NodeID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// MarriageProof is a proof for two identities to be married via a marriage ATX.
type MarriageProof struct {
	// MarriageATX and its proof that it is contained in the ATX.
	MarriageATX      types.ATXID
	MarriageATXProof MarriageATXProof `scale:"max=32"`
	// MarriageATXSmesherID is the ID of the smesher that published the marriage ATX.
	MarriageATXSmesherID types.NodeID

	// NodeIDMarryProof is the proof that NodeID married in MarriageATX.
	NodeIDMarryProof MarryProof
	// SmesherIDMarryProof is the proof that SmesherID married in MarriageATX.
	SmesherIDMarryProof MarryProof
}

func createMarriageProof(db sql.Executor, atx *ActivationTxV2, nodeID types.NodeID) (MarriageProof, error) {
	_ = "STUB: not implemented"
	return *new(MarriageProof), nil
}

// we don't need a marriage proof if the node ID is the same as the smesher ID

func (p MarriageProof) Valid(
	malValidator MalfeasanceValidator,
	atxID types.ATXID,
	nodeID, smesherID types.NodeID,
) error {
	_ = "STUB: not implemented"
	return nil
}
