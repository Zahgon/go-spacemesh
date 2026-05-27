package activation

import (
	"context"
	"errors"
	"time"

	"github.com/spacemeshos/post/config"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/activation/wire"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var ErrPostIndexOutOfRange = errors.New("post index out of range")

type validatorOptions struct {
	postIdx        *int
	postSubsetSeed []byte
	prioritized    bool
}

// PostSubset configures the validator to validate only a subset of the POST indices.
// The `seed` is used to randomize the selection of indices.
func PostSubset(seed []byte) validatorOption {
	_ = "STUB: not implemented"
	return *new(validatorOption)
}

// PostIndex configures the validator to validate only the POST index at the given `idx`.
func PostIndex(idx int) validatorOption { _ = "STUB: not implemented"; return *new(validatorOption) }

func PrioritizeCall() validatorOption { _ = "STUB: not implemented"; return *new(validatorOption) }

// Validator contains the dependencies required to validate NIPosts.
type Validator struct {
	db           sql.Executor
	poetDb       poetDbAPI
	cfg          PostConfig
	scrypt       config.ScryptParams
	postVerifier PostVerifier
}

// NewValidator returns a new NIPost validator.
func NewValidator(
	db sql.Executor,
	poetDb poetDbAPI,
	cfg PostConfig,
	scrypt config.ScryptParams,
	postVerifier PostVerifier,
) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// NIPost validates a NIPost, given a node id and expected challenge. It returns an error if the NIPost is invalid.
//
// Some of the Post metadata fields validation values is ought to eventually be derived from
// consensus instead of local configuration. If so, their validation should be removed to contextual validation,
// while still syntactically-validate them here according to locally configured min/max values.
func (v *Validator) NIPost(
	ctx context.Context,
	nodeId types.NodeID,
	commitmentAtxId types.ATXID,
	nipost *types.NIPost,
	poetChallenge types.Hash32,
	numUnits uint32,
	opts ...validatorOption,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (v *Validator) PoetMembership(
	_ context.Context,
	membership *types.MultiMerkleProof,
	postChallenge types.Hash32,
	poetChallenges [][]byte,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validateMerkleProof(leaf []byte, proof *types.MerkleProof, expectedRoot []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMultiMerkleProof(leaves [][]byte, proof *types.MultiMerkleProof, expectedRoot []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) IsVerifyingFullPost() bool { _ = "STUB: not implemented"; return false }

// Post validates a Proof of Space-Time (PoST). It returns nil if validation passed or an error indicating why
// validation failed.
func (v *Validator) Post(
	ctx context.Context,
	nodeId types.NodeID,
	commitmentAtxId types.ATXID,
	post *types.Post,
	metadata *types.PostMetadata,
	numUnits uint32,
	opts ...validatorOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) PostV2(
	ctx context.Context,
	nodeId types.NodeID,
	commitmentAtxId types.ATXID,
	post *types.Post,
	challenge []byte,
	numUnits uint32,
	opts ...validatorOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Validator) NumUnits(cfg *PostConfig, numUnits uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Validator) LabelsPerUnit(cfg *PostConfig, labelsPerUnit uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) VRFNonce(
	nodeId types.NodeID,
	commitmentAtxId types.ATXID,
	vrfNonce, labelsPerUnit uint64,
	numUnits uint32,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) VRFNonceV2(nodeId types.NodeID, commitment types.ATXID, vrfNonce uint64, numUnits uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) InitialNIPostChallengeV1(
	challenge *wire.NIPostChallengeV1,
	atxs atxProvider,
	goldenATXID types.ATXID,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Validator) NIPostChallengeV1(
	challenge *wire.NIPostChallengeV1,
	prevATX *types.ActivationTx,
	nodeID types.NodeID,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) PositioningAtx(
	id types.ATXID,
	atxs atxProvider,
	goldenATXID types.ATXID,
	pubepoch types.EpochID,
) error {
	_ = "STUB: not implemented"
	return nil
}

type verifyChainOpts struct {
	assumedValidTime time.Time
	trustedNodeID    types.NodeID
	logger           *zap.Logger
}

type verifyChainOptsNs struct{}

var VerifyChainOpts verifyChainOptsNs

type VerifyChainOption func(*verifyChainOpts)

// AssumeValidBefore configures the validator to assume that ATXs received before the given time are valid.
func (verifyChainOptsNs) AssumeValidBefore(val time.Time) VerifyChainOption {
	_ = "STUB: not implemented"
	return *new(VerifyChainOption)
}

// WithTrustedID configures the validator to assume that ATXs created by the given node ID are valid.
func (verifyChainOptsNs) WithTrustedID(val types.NodeID) VerifyChainOption {
	_ = "STUB: not implemented"
	return *new(VerifyChainOption)
}

func (verifyChainOptsNs) WithLogger(log *zap.Logger) VerifyChainOption {
	_ = "STUB: not implemented"
	return *new(VerifyChainOption)
}

type InvalidChainError struct {
	ID  types.ATXID
	src error
}

func (e *InvalidChainError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *InvalidChainError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *InvalidChainError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (v *Validator) VerifyChain(ctx context.Context, id, goldenATXID types.ATXID, opts ...VerifyChainOption) error {
	_ = "STUB: not implemented"
	return nil
}

type atxDeps struct {
	niposts     []types.NIPost
	positioning types.ATXID
	previous    []types.ATXID
	commitment  types.ATXID
}

func (v *Validator) getAtxDeps(ctx context.Context, id types.ATXID) (*atxDeps, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) verifyChainWithOpts(
	ctx context.Context,
	id, goldenATXID types.ATXID,
	opts verifyChainOpts,
) error {
	_ = "STUB: not implemented"
	return nil
}

// validate POST fully

func (v *Validator) verifyChainDeps(
	ctx context.Context,
	deps *atxDeps,
	goldenATXID types.ATXID,
	opts verifyChainOpts,
) error {
	_ = "STUB: not implemented"
	return nil
}

// verify commitment only if arrived at the first ATX in the chain
// to avoid verifying the same commitment ATX multiple times.
