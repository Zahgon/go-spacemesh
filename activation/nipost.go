package activation

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/localsql/nipost"
)

const (
	// Jitter values to avoid all nodes querying the poet at the same time.
	// Note: the jitter values are represented as a percentage of cycle gap.
	//  mainnet cycle-gap: 12h
	//  systest cycle-gap: 30s

	// Minimum jitter value before querying for the proof.
	// Gives the poet service time to generate proof after a round ends (~8s on mainnet).
	//  mainnet -> 8.64s
	//  systest -> 0.36s
	minPoetGetProofJitter = 0.02

	// The maximum jitter value before querying for the proof.
	//  mainnet -> 17.28s
	//  systest -> 0.72s
	maxPoetGetProofJitter = 0.04
)

var ErrInvalidInitialPost = errors.New("invalid initial post")

// NIPostBuilder holds the required state and dependencies to create Non-Interactive Proofs of Space-Time (NIPost).
type NIPostBuilder struct {
	localDB sql.LocalDatabase

	poetProvers map[string]PoetService
	postService postService
	logger      *zap.Logger
	poetCfg     PoetConfig
	layerClock  layerClock
	postStates  PostStates
	validator   nipostValidator
}

type NIPostBuilderOption func(*NIPostBuilder)

func WithPoetServices(clients ...PoetService) NIPostBuilderOption {
	_ = "STUB: not implemented"
	return *new(NIPostBuilderOption)
}

func NipostbuilderWithPostStates(ps PostStates) NIPostBuilderOption {
	_ = "STUB: not implemented"
	return *new(NIPostBuilderOption)
}

// NewNIPostBuilder returns a NIPostBuilder.
func NewNIPostBuilder(
	db sql.LocalDatabase,
	postService postService,
	lg *zap.Logger,
	poetCfg PoetConfig,
	layerClock layerClock,
	validator nipostValidator,
	opts ...NIPostBuilderOption,
) (*NIPostBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nb *NIPostBuilder) ResetState(nodeId types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (nb *NIPostBuilder) Proof(
	ctx context.Context,
	nodeID types.NodeID,
	challenge []byte,
	postChallenge *types.NIPostChallenge,
) (*types.Post, *types.PostInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Wait a few seconds and try connecting again

// every 20 seconds inform user about lost connection (for remote post service)
// TODO(mafa): emit event warning user about lost connection

// we check whether an initial post is included in the challenge
// if so, we verify it to still be valid before creating the post
// e.g. the PoST size might have changed

// We don't set the state to idle here, because we will retry up the stack.

// err == nil

// BuildNIPost uses the given challenge to build a NIPost.
// The process can take considerable time, because it includes waiting for the poet service to
// publish a proof - a process that takes about an epoch.
func (nb *NIPostBuilder) BuildNIPost(
	ctx context.Context,
	signer *signing.EdSigner,
	challenge types.Hash32,
	postChallenge *types.NIPostChallenge,
) (*nipost.NIPostState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: to avoid missing next PoET round, we need to publish the ATX before the next PoET round starts.
//   We can still publish an ATX late (i.e. within publish epoch) and receive rewards, but we will miss one
//   epoch because we didn't submit the challenge to PoET in time for next round.
//                                 PoST
//         ┌─────────────────────┐  ┌┐┌─────────────────────┐
//         │     POET ROUND      │  │││   NEXT POET ROUND   │
// ┌────▲──┴──────────────────┬──▲──┴┴┴─────────────────▲┬──┴─────────────► time
// │    │      EPOCH          │  │   PUBLISH EPOCH      ││  TARGET EPOCH
// └────┼─────────────────────┴──┼──────────────────────┼┴────────────────
//      │                        │                      │
//  WE ARE HERE            PROOF BECOMES         ATX PUBLICATION
//                           AVAILABLE               DEADLINE

// we want to publish before the publish epoch ends or we won't receive rewards

// we want to fetch the PoET proof latest 1 CycleGap before the publish epoch ends
// so that a node that is setup correctly (i.e. can generate a PoST proof within the cycle gap)
// has enough time left to generate a post proof and publish

// Phase 0: Submit challenge to PoET services.
// Deadline: start of PoET round: we will not accept registrations after that

// Phase 1: query PoET services for proofs

// Deadline: the end of the publish epoch minus the cycle gap. A node that is setup correctly (i.e. can
// generate a PoST proof within the cycle gap) has enough time left to generate a post proof and publish.

// Phase 2: Post execution.

// Deadline: the end of the publish epoch. If we do not publish within
// the publish epoch we won't receive any rewards in the target epoch.

// withConditionalTimeout returns a context.WithTimeout if the timeout is greater than 0, otherwise it returns
// the original context.
func withConditionalTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// Submit the challenge (register) to a single PoET.
func (nb *NIPostBuilder) submitPoetChallenge(
	ctx context.Context,
	nodeID types.NodeID,
	deadline time.Time,
	client PoetService,
	prefix, challenge []byte,
	signature types.EdSignature,
) (nipost.PoETRegistration, error) {
	_ = "STUB: not implemented"
	return *new(nipost.PoETRegistration), nil
}

// submitPoetChallenges submit the challenge to registered PoETs
// if some registrations are missing and PoET round didn't start.
func (nb *NIPostBuilder) submitPoetChallenges(
	ctx context.Context,
	signer *signing.EdSigner,
	poetProofDeadline time.Time,
	curPoetRoundStartDeadline time.Time,
	challenge []byte,
) ([]nipost.PoETRegistration, error) {
	_ = "STUB: not implemented"
	// check if some registrations missing or were removed
	return nil, nil
}

// no existing registration at all, drop current registration challenge

// no existing registration for given poets set

// send registrations to missing addresses

// membersContainChallenge verifies that the challenge is included in proof's members.
func membersContainChallenge(members []types.Hash32, challenge types.Hash32) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (nb *NIPostBuilder) getBestProof(
	ctx context.Context,
	nodeID types.NodeID,
	challenge types.Hash32,
	registrations []nipost.PoETRegistration,
) (types.PoetProofRef, *types.MerkleProof, error) {
	_ = "STUB: not implemented"
	return *new(types.PoetProofRef), nil, nil
}

func constructMerkleProof(challenge types.Hash32, members []types.Hash32) (*types.MerkleProof, error) {
	_ = "STUB: not implemented"
	// We are interested only in proofs that we are members of
	return nil, nil
}

func randomDurationInRange(min, max time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Calculate the time to wait before querying for the proof
// We add a jitter to avoid all nodes querying for the proof at the same time.
func proofDeadline(roundEnd time.Time, cycleGap time.Duration) (waitTime time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
