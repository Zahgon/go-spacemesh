package wire

import (
	"github.com/spacemeshos/merkle-tree"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

//go:generate scalegen

type ActivationTxV2 struct {
	PublishEpoch   types.EpochID
	PositioningATX types.ATXID
	Coinbase       types.Address

	// only present in initial ATX
	Initial      *InitialAtxPartsV2
	PreviousATXs PrevATXs `scale:"max=256"`
	NIPosts      NIPosts  `scale:"max=4"`

	// The VRF nonce must be valid for the collected space of all included IDs.
	VRFNonce uint64

	// The list of marriages with other IDs.
	// A marriage is permanent and cannot be revoked or repeated.
	// All new IDs that are married to this ID are added to the equivocation set
	// that this ID belongs to.
	// It must contain a self-marriage certificate (needed for malfeasance proofs).
	Marriages MarriageCertificates `scale:"max=256"`

	// The ID of the ATX containing marriage for the included IDs.
	// Only required when the ATX includes married IDs.
	MarriageATX *types.ATXID

	SmesherID types.NodeID
	Signature types.EdSignature

	// cached fields to avoid repeated calculations
	id   types.ATXID
	blob []byte
}

func (atx *ActivationTxV2) Blob() types.AtxBlob {
	_ = "STUB: not implemented"
	return *new(types.AtxBlob)
}

func DecodeAtxV2(blob []byte) (*ActivationTxV2, error) { _ = "STUB: not implemented"; return nil, nil }

func (atx *ActivationTxV2) Sign(signer *signing.EdSigner) { _ = "STUB: not implemented"; return }

func (atx *ActivationTxV2) TotalNumUnits() uint32 { _ = "STUB: not implemented"; return 0 }

func (atx *ActivationTxV2) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (atx *ActivationTxV2) merkleTree(tree *merkle.Tree) { _ = "STUB: not implemented"; return }

func (atx *ActivationTxV2) merkleProof(leafIndex MerkleTreeIndex) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

// ID returns the ATX ID. It is the root of the ATX merkle tree.
func (atx *ActivationTxV2) ID() types.ATXID { _ = "STUB: not implemented"; return *new(types.ATXID) }

func (atx *ActivationTxV2) PublishEpochProof() PublishEpochProof {
	_ = "STUB: not implemented"
	return *new(PublishEpochProof)
}

type PublishEpochProof []types.Hash32

func (p PublishEpochProof) Valid(atxID types.ATXID, publishEpoch types.EpochID) bool {
	_ = "STUB: not implemented"
	return false
}

func (atx *ActivationTxV2) PositioningATXProof() []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

func (atx *ActivationTxV2) CoinbaseProof() []types.Hash32 { _ = "STUB: not implemented"; return nil }

func (atx *ActivationTxV2) InitialPostRootProof() InitialPostRootProof {
	_ = "STUB: not implemented"
	return *new(InitialPostRootProof)
}

type InitialPostRootProof []types.Hash32

func (p InitialPostRootProof) Valid(atxID types.ATXID, initialPostRoot InitialPostRoot) bool {
	_ = "STUB: not implemented"
	return false
}

func (atx *ActivationTxV2) PreviousATXsRootProof() PrevATXsRootProof {
	_ = "STUB: not implemented"
	return *new(PrevATXsRootProof)
}

type PrevATXsRootProof []types.Hash32

func (p PrevATXsRootProof) Valid(atxID types.ATXID, prevATXsRoot PrevATXsRoot) bool {
	_ = "STUB: not implemented"
	return false
}

func (atx *ActivationTxV2) NIPostsRootProof() NIPostsRootProof {
	_ = "STUB: not implemented"
	return *new(NIPostsRootProof)
}

type NIPostsRootProof []types.Hash32

func (p NIPostsRootProof) Valid(atxID types.ATXID, niPostsRoot NIPostsRoot) bool {
	_ = "STUB: not implemented"
	return false
}

func (atx *ActivationTxV2) VRFNonceProof() []types.Hash32 { _ = "STUB: not implemented"; return nil }

func (atx *ActivationTxV2) MarriagesRootProof() MarriageCertificatesRootProof {
	_ = "STUB: not implemented"
	return *new(MarriageCertificatesRootProof)
}

type MarriageCertificatesRootProof []types.Hash32

func (p MarriageCertificatesRootProof) Valid(atxID types.ATXID, marriagesRoot MarriageCertificatesRoot) bool {
	_ = "STUB: not implemented"
	return false
}

func (atx *ActivationTxV2) MarriageATXProof() MarriageATXProof {
	_ = "STUB: not implemented"
	return *new(MarriageATXProof)
}

type MarriageATXProof []types.Hash32

func (p MarriageATXProof) Valid(atxID, marriageATX types.ATXID) bool {
	_ = "STUB: not implemented"
	return false
}

type InitialAtxPartsV2 struct {
	CommitmentATX types.ATXID
	Post          PostV1
}

func (parts *InitialAtxPartsV2) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (parts *InitialAtxPartsV2) merkleTree(tree *merkle.Tree) { _ = "STUB: not implemented"; return }

func (parts *InitialAtxPartsV2) merkleProof(leafIndex InitialPostTreeIndex) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

type InitialPostRoot types.Hash32

func (parts *InitialAtxPartsV2) Root() InitialPostRoot {
	_ = "STUB: not implemented"
	return *new(InitialPostRoot)
}

func (parts *InitialAtxPartsV2) CommitmentATXProof() CommitmentATXProof {
	_ = "STUB: not implemented"
	return *new(CommitmentATXProof)
}

type CommitmentATXProof []types.Hash32

func (p CommitmentATXProof) Valid(initialPostRoot InitialPostRoot, commitmentATX types.ATXID) bool {
	_ = "STUB: not implemented"
	return false
}

func (parts *InitialAtxPartsV2) PostProof() []types.Hash32 { _ = "STUB: not implemented"; return nil }

type PrevATXs []types.ATXID

func (prevATXs PrevATXs) merkleTree(tree *merkle.Tree) { _ = "STUB: not implemented"; return }

type PrevATXsRoot types.Hash32

func (prevATXs PrevATXs) Root() PrevATXsRoot { _ = "STUB: not implemented"; return *new(PrevATXsRoot) }

func (prevATXs PrevATXs) Proof(index int) PrevATXsProof {
	_ = "STUB: not implemented"
	return *new(PrevATXsProof)
}

type PrevATXsProof []types.Hash32

func (p PrevATXsProof) Valid(prevATXsRoot PrevATXsRoot, index int, prevATX types.ATXID) bool {
	_ = "STUB: not implemented"
	return false
}

type NIPosts []NIPostV2

func (nps NIPosts) merkleTree(tree *merkle.Tree, prevATXs []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

// Add empty NiPoSTs up to the max scale limit.
// This must be updated when the max scale limit is changed.

type NIPostsRoot types.Hash32

func (nps NIPosts) Root(prevATXs []types.ATXID) NIPostsRoot {
	_ = "STUB: not implemented"
	return *new(NIPostsRoot)
}

func (nps NIPosts) Proof(index int, prevATXs []types.ATXID) NIPostRootProof {
	_ = "STUB: not implemented"
	return *new(NIPostRootProof)
}

type NIPostRootProof []types.Hash32

func (p NIPostRootProof) Valid(niPostsRoot NIPostsRoot, index int, nipostRoot NIPostRoot) bool {
	_ = "STUB: not implemented"
	return false
}

type NIPostV2 struct {
	// Single membership proof for all IDs in `Posts`.
	Membership MerkleProofV2
	// The root of the PoET proof, that serves as the challenge for PoSTs.
	Challenge types.Hash32
	Posts     SubPostsV2 `scale:"max=256"` // support merging up to 256 IDs
}

func (np *NIPostV2) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// skip membership proof

func (np *NIPostV2) merkleTree(tree *merkle.Tree, prevATXs []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

func (np *NIPostV2) merkleProof(leafIndex NIPostTreeIndex, prevATXs []types.ATXID) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

type NIPostRoot types.Hash32

func (np *NIPostV2) Root(prevATXs []types.ATXID) NIPostRoot {
	_ = "STUB: not implemented"
	return *new(NIPostRoot)
}

func (np *NIPostV2) MembershipProof(prevATXs []types.ATXID) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

func (np *NIPostV2) ChallengeProof(prevATXs []types.ATXID) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

type ChallengeProof []types.Hash32

func (p ChallengeProof) Valid(nipostRoot NIPostRoot, challenge types.Hash32) bool {
	_ = "STUB: not implemented"
	return false
}

func (np *NIPostV2) PostsRootProof(prevATXs []types.ATXID) SubPostsRootProof {
	_ = "STUB: not implemented"
	return *new(SubPostsRootProof)
}

type SubPostsRootProof []types.Hash32

func (p SubPostsRootProof) Valid(nipostRoot NIPostRoot, postsRoot SubPostsRoot) bool {
	_ = "STUB: not implemented"
	return false
}

// MerkleProofV2 proves membership of multiple challenges in a PoET membership merkle tree.
type MerkleProofV2 struct {
	// Nodes on path from leaf to root (not including leaf)
	Nodes []types.Hash32 `scale:"max=32"`
}

func (mp *MerkleProofV2) Root() (result types.Hash32) {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

type SubPostsV2 []SubPostV2

func (sp SubPostsV2) merkleTree(tree *merkle.Tree, prevATXs []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

// if root is nil it will be handled like 0x00...00
// this will still generate a valid ID for the ATX,
// but syntactical validation will catch the invalid subPost and
// consider the ATX invalid

type SubPostsRoot types.Hash32

func (sp SubPostsV2) Root(prevATXs []types.ATXID) SubPostsRoot {
	_ = "STUB: not implemented"
	return *new(SubPostsRoot)
}

func (sp SubPostsV2) Proof(index int, prevATXs []types.ATXID) SubPostRootProof {
	_ = "STUB: not implemented"
	return *new(SubPostRootProof)
}

type SubPostRootProof []types.Hash32

func (p SubPostRootProof) Valid(subPostsRoot SubPostsRoot, index int, subPostRoot SubPostRoot) bool {
	_ = "STUB: not implemented"
	return false
}

type SubPostV2 struct {
	// Index of marriage certificate for this ID in the 'Marriages' slice. Only valid for merged ATXs.
	// Can be used to extract the nodeID and verify if it is married with the smesher of the ATX.
	// Must be 0 for non-merged ATXs.
	MarriageIndex uint32
	PrevATXIndex  uint32 // Index of the previous ATX in the `ActivationTxV2.PreviousATXs` slice
	// Index of the leaf for this ID's challenge in the poet membership tree.
	// IDs might shared the same index if their nipost challenges are equal.
	// This happens when the IDs are continuously merged (they share the previous ATX).
	MembershipLeafIndex uint64
	Post                PostV1
	NumUnits            uint32
}

func (post *SubPostV2) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *SubPostV2) merkleTree(tree *merkle.Tree, prevATX types.ATXID) {
	_ = "STUB: not implemented"
	return
}

func (sp *SubPostV2) merkleProof(leafIndex SubPostTreeIndex, prevATXs []types.ATXID) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

// special case for initial ATX: prevATXs is empty

// not the full set of prevATXs is provided, proof cannot be generated

type SubPostRoot types.Hash32

func (sp *SubPostV2) Root(prevATXs []types.ATXID) SubPostRoot {
	_ = "STUB: not implemented"
	return *new(SubPostRoot)
}

// special case for initial ATX: prevATXs is empty

// prevATXIndex is out of range, don't fail ATXID generation
// will be detected by syntactical validation

func (sp *SubPostV2) MarriageIndexProof(prevATXs []types.ATXID) MarriageIndexProof {
	_ = "STUB: not implemented"
	return *new(MarriageIndexProof)
}

type MarriageIndexProof []types.Hash32

func (p MarriageIndexProof) Valid(subPostRoot SubPostRoot, marriageIndex uint32) bool {
	_ = "STUB: not implemented"
	return false
}

func (sp *SubPostV2) PrevATXIndexProof(prevATXs []types.ATXID) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

func (sp *SubPostV2) PrevATXProof(prevATX types.ATXID) PrevATXProof {
	_ = "STUB: not implemented"
	return *new(PrevATXProof)
}

type PrevATXProof []types.Hash32

func (p PrevATXProof) Valid(subPostRoot SubPostRoot, prevATX types.ATXID) bool {
	_ = "STUB: not implemented"
	return false
}

func (sp *SubPostV2) MembershipLeafIndexProof(prevATXs []types.ATXID) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

func (sp *SubPostV2) PostProof(prevATXs []types.ATXID) PostRootProof {
	_ = "STUB: not implemented"
	return *new(PostRootProof)
}

type PostRootProof []types.Hash32

func (p PostRootProof) Valid(subPostRoot SubPostRoot, postRoot PostRoot) bool {
	_ = "STUB: not implemented"
	return false
}

func (sp *SubPostV2) NumUnitsProof(prevATXs []types.ATXID) NumUnitsProof {
	_ = "STUB: not implemented"
	return *new(NumUnitsProof)
}

type NumUnitsProof []types.Hash32

func (p NumUnitsProof) Valid(subPostRoot SubPostRoot, numUnits uint32) bool {
	_ = "STUB: not implemented"
	return false
}

type MarriageCertificates []MarriageCertificate

func (mcs MarriageCertificates) merkleTree(tree *merkle.Tree) { _ = "STUB: not implemented"; return }

type MarriageCertificatesRoot types.Hash32

func (mcs MarriageCertificates) Root() MarriageCertificatesRoot {
	_ = "STUB: not implemented"
	return *new(MarriageCertificatesRoot)
}

func (mcs MarriageCertificates) Proof(index int) MarriageCertificateProof {
	_ = "STUB: not implemented"
	return *new(MarriageCertificateProof)
}

type MarriageCertificateProof []types.Hash32

func (p MarriageCertificateProof) Valid(marriageRoot MarriageCertificatesRoot, index int, mc MarriageCertificate) bool {
	_ = "STUB: not implemented"
	return false
}

// MarriageCertificate proves the will of ID to be married with the ID that includes this certificate.
// A marriage allows for publishing a merged ATX, which can contain PoST for all married IDs.
// Any ID from the marriage can publish a merged ATX on behalf of all married IDs.
type MarriageCertificate struct {
	// An ATX of the NodeID that marries. It proves that the NodeID exists.
	// Note: the reference ATX does not need to be from the previous epoch.
	// It only needs to prove the existence of the Identity.
	//
	// In the case of a self signed certificate that is included in the Marriage ATX by the Smesher signing the ATX,
	// this can be `types.EmptyATXID`.
	ReferenceAtx types.ATXID
	// Signature over the other ID that this ID marries with
	// If Alice marries Bob, then Alice signs Bob's ID
	// and Bob includes this certificate in his ATX.
	Signature types.EdSignature
}

func (mc *MarriageCertificate) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (mc *MarriageCertificate) Root() (result types.Hash32) {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

func atxTreeHash(buf, lChild, rChild []byte) []byte { _ = "STUB: not implemented"; return nil }

func createRoot(addLeaves func(tree *merkle.Tree)) types.Hash32 {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

func createProof(leafIndex uint64, addLeaves func(tree *merkle.Tree)) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

func validateProof(root, leaf types.Hash32, proof []types.Hash32, leafIndex uint64) bool {
	_ = "STUB: not implemented"
	return false
}
