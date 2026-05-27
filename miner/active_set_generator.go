package miner

import (
	"context"
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type activesetGenOpt func(*activeSetGenerator)

func withWallClock(clock clockwork.Clock) activesetGenOpt {
	_ = "STUB: not implemented"
	return *new(activesetGenOpt)
}

func newActiveSetGenerator(
	cfg config,
	log *zap.Logger,
	db, localdb sql.Executor,
	atxsdata atxsData,
	clock layerClock,
	opts ...activesetGenOpt,
) *activeSetGenerator {
	_ = "STUB: not implemented"
	return nil
}

type activeSetGenerator struct {
	cfg config
	log *zap.Logger

	db, localdb sql.Executor
	atxsdata    atxsData
	clock       layerClock
	wallclock   clockwork.Clock

	fallback struct {
		sync.Mutex
		data map[types.EpochID][]types.ATXID
	}

	// we use it to avoid running `prepare` method in parallel
	linearized sync.Mutex
}

func (p *activeSetGenerator) updateFallback(target types.EpochID, set []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

// ensure tries to generate active set for the target epoch in configured number of tries,
// each try will be repeated after configured retry interval.
func (a *activeSetGenerator) ensure(ctx context.Context, target types.EpochID) {
	_ = "STUB: not implemented"
	return
}

// we run it here for side effects

// generate generates activeset.
//
// It persists it in the local database, so that when node is restarted
// it doesn't have to redo the work.
//
// The method is expected to be called at any point in current epoch, and at the very end of the previous epoch.
func (p *activeSetGenerator) generate(
	current types.LayerID,
	target types.EpochID,
) (types.Hash32, uint64, []types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), 0, nil, nil
}

// see how it can be further improved https://github.com/spacemeshos/go-spacemesh/issues/5560

type gradedActiveSet struct {
	// Set includes activations with highest grade.
	Set []types.ATXID
	// Weight of the activations in the Set.
	Weight uint64
	// Total number of activations in the database that targets requests epoch.
	Total int
}

// activeSetFromGrades includes activations with the highest grade.
// Such activations were received at least 4 network delays before the epoch start, and no malfeasance proof for
// identity was received before the epoch start.
//
// On mainnet we use 30minutes as a network delay parameter.
func activeSetFromGrades(
	db sql.Executor,
	target types.EpochID,
	epochStart time.Time,
	networkDelay time.Duration,
) (gradedActiveSet, error) {
	_ = "STUB: not implemented"
	return *new(gradedActiveSet), nil
}

func getSetWeight(atxsdata atxsData, target types.EpochID, set []types.ATXID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// atxGrade describes the grade of an ATX as described in
// https://community.spacemesh.io/t/grading-atxs-for-the-active-set/335
//
// let s be the start of the epoch, and δ the network propagation time.
// Grade 0: ATX was received at time t >= s-3δ, or an equivocation proof was received by time s-δ.
// Grade 1: ATX was received at time t < s-3δ before the start of the epoch, and no equivocation proof by time s-δ.
// Grade 2: ATX was received at time t < s-4δ, and no equivocation proof was received for that id until time s.
type atxGrade int

const (
	evil atxGrade = iota
	acceptable
	good
)

func gradeAtx(epochStart time.Time, networkDelay time.Duration, atxNsec, proofNsec int64) atxGrade {
	_ = "STUB: not implemented"
	return *new(atxGrade)
}

func ActiveSetFromEpochFirstBlock(db sql.Executor, epoch types.EpochID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func activeSetFromBlock(db sql.Executor, bid types.BlockID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the active set is the union of all active sets recorded in rewarded miners' ref ballot
