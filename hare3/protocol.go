package hare3

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
)

type grade uint8

const (
	grade0 grade = iota
	grade1
	grade2
	grade3
	grade4
	grade5
)

func isSubset(s, v []types.ProposalID) bool { _ = "STUB: not implemented"; return false }

func toHash(proposals []types.ProposalID) types.Hash32 {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

type messageKey struct {
	IterRound
	Sender types.NodeID
}

type input struct {
	*Message
	atxgrade  grade
	malicious bool
	msgHash   types.Hash32
}

func (i *input) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type output struct {
	coin       *bool              // set based on preround messages right after preround completes in 0 iter
	result     []types.ProposalID // set based on notify messages at the start of next iter
	terminated bool               // protocol participates in one more iteration after outputing result
	message    *Message
}

func (o *output) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newProtocol(threshold uint16, logger *zap.Logger) *protocol {
	_ = "STUB: not implemented"
	return nil
}

type protocol struct {
	mu sync.Mutex
	IterRound
	coinout        bool
	coin           *types.VrfSignature // smallest vrf from preround messages. not a part of paper
	initial        []types.ProposalID  // Si
	result         *types.Hash32       // set after waiting for notify messages. Case 1
	locked         *types.Hash32       // Li
	hardLocked     bool
	validProposals map[types.Hash32][]types.ProposalID // Ti
	gossip         gossip
	logger         *zap.Logger
}

func (p *protocol) OnInitial(proposals []types.ProposalID) { _ = "STUB: not implemented"; return }

func (p *protocol) OnInput(msg *input) (bool, *wire.HareProof) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *protocol) thresholdProposals(ir IterRound, grade grade) (*types.Hash32, []types.ProposalID) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *protocol) commitExists(iter uint8, match types.Hash32, grade grade) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *protocol) execution(out *output) {
	_ = "STUB: not implemented"
	// 4.3 Protocol Execution
	return
}

// receiver expects non-nil result

// condition (d) is realized by ordering proposals by vrf

// condition (a) and (b)
// grade0 proposals are not added to the set

// condition (c)

// condition (e)

// condition (f)

// condition (g)

// condition (h)

func (p *protocol) Next() output { _ = "STUB: not implemented"; return *new(output) }

// skips hardlock unlike softlock in the paper.
// this makes no practical difference from correctness.
// but allows to simplify assignment in validValues

func (p *protocol) Stats() *stats { _ = "STUB: not implemented"; return nil }

// preround messages that are received after the very first iteration
// has no impact on protocol

// stats are collected at the start of current iteration (p.Iter)
// we expect 2 network delays to pass since commit messages were broadcasted

// we are not interested in any other grade for notify message as they have no impact on protocol execution

type gossipInput struct {
	*input
	received      IterRound
	otherReceived *IterRound
}

// Protocol 1. Graded-gossip. Page 10.
type gossip struct {
	threshold uint16
	state     map[messageKey]*gossipInput
}

func (g *gossip) receive(current IterRound, input *input) (bool, *wire.HareProof) {
	_ = "STUB: not implemented"
	// Case 1: will be discarded earlier
	return false, nil
}

// Protocol 3. thresh-gossip. keep one with the maximal grade.

// Case 3

// Case 2. but also we filter duplicates from p2p layer here

// Case 4

type gset struct {
	values   []types.ProposalID
	grade    grade
	smallest types.VrfSignature
}

// Protocol 2. Gradecast. Page 13.
func (g *gossip) gradecast(target IterRound) []gset {
	_ = "STUB: not implemented"
	// unlike paper we use 5-graded gossip for gradecast as well
	return nil
}

// 2 (a)

// 2 (b)

// 3 (a)

// 3 (b)

// hare expects to receive multiple proposals. expected number of leaders is set to 5.
// we need to choose the same one for commit across the cluster.
// we do that by ordering them by vrf value, and picking one that passes other checks (see commit in execution).
// in hare3 paper look for p-Weak leader election property.

func tallyProposals(all map[types.ProposalID]proposalTally, inp *gossipInput) {
	_ = "STUB: not implemented"
	return
}

// Protocol 3. Thresh-gossip. Page 15.
// Output returns union of sorted proposals received
// in the given round with minimal specified grade.
func (g *gossip) thresholdGossip(filter IterRound, grade grade, logger *zap.Logger) []types.ProposalID {
	_ = "STUB: not implemented"
	return nil
}

func tallyRefs(all map[types.Hash32]refTally, inp *gossipInput) { _ = "STUB: not implemented"; return }

// thresholdGossipRef returns all references to proposals in the given round with minimal grade.
func (g *gossip) thresholdGossipRef(filter IterRound, grade grade, logger *zap.Logger) []types.Hash32 {
	_ = "STUB: not implemented"
	return nil
}

func thresholdGossip[T interface {
	comparable
	fmt.Stringer
}](
	tallies map[T]tallyStats[T], threshold uint16, logger *zap.Logger,
) []T {
	_ = "STUB: not implemented"
	return nil
}

// valid > 0 and total >= f
// atleast one non-equivocating vote and crossed committee/2 + 1

func thresholdTallies[T interface {
	comparable
	fmt.Stringer
}](
	state map[messageKey]*gossipInput,
	filter IterRound,
	msgGrade grade,
	tally func(tally map[T]tallyStats[T], inp *gossipInput),
) map[T]tallyStats[T] {
	_ = "STUB: not implemented"
	return nil
}

// pick min atx grade from non equivocating identity.

// tally votes for valid and malicious messages

type preroundStats struct {
	grade   grade
	tallies []proposalTally
}

func (s *preroundStats) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type tallyStats[T fmt.Stringer] struct {
	id    T
	total uint16
	valid uint16
}

func (s *tallyStats[T]) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	proposalTally = tallyStats[types.ProposalID]
	refTally      = tallyStats[types.Hash32]
)

type proposeStats struct {
	grade     grade
	ref       types.Hash32
	proposals []types.ProposalID
}

func (s *proposeStats) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type commitStats struct {
	grade   grade
	tallies []refTally
}

func (s *commitStats) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type notifyStats struct {
	grade   grade
	tallies []refTally
}

func (n *notifyStats) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type stats struct {
	iter      uint8
	threshold uint16
	preround  []preroundStats
	propose   []proposeStats
	commit    []commitStats
	notify    []notifyStats
}

func (s *stats) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
