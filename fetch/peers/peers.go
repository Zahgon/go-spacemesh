package peers

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"go.uber.org/zap/zapcore"
)

type data struct {
	id                peer.ID
	success, failures int
	failRate          float64
	averageLatency    float64
	protocols         func() []protocol.ID
}

func (d *data) latency(global float64) float64 { _ = "STUB: not implemented"; return 0 }

// to prioritize trying out new peer

func (p *data) less(other *data, global float64) bool { _ = "STUB: not implemented"; return false }

func New() *Peers { _ = "STUB: not implemented"; return nil }

type Peers struct {
	mu    sync.Mutex
	peers map[peer.ID]*data

	// globalLatency is the average latency of all successful responses from peers.
	// It is used as a reference value for new peers.
	// And to adjust average peer latency based on failure rate.
	globalLatency float64
}

func (p *Peers) Contains(id peer.ID) bool { _ = "STUB: not implemented"; return false }

func (p *Peers) Add(id peer.ID, protocols func() []protocol.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Peers) Delete(id peer.ID) { _ = "STUB: not implemented"; return }

// OnLatency updates average peer and global latency.
func (p *Peers) onLatency(id peer.ID, size int, latency time.Duration, failed bool) {
	_ = "STUB: not implemented"
	// We assume that latency is proportional to the size of the message
	// and define it as a duration to transmit 1kiB.
	// To account for the additional overhead of transmitting small messages,
	// we treat them as if they were 1kiB.
	return
}

// 86% of the value is the last 19

// 86% of the value is the last 49

func (p *Peers) OnFailure(id peer.ID, size int, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// OnLatency updates average peer and global latency.
func (p *Peers) OnLatency(id peer.ID, size int, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// SelectBest peer with preferences.
func (p *Peers) SelectBestFrom(peers []peer.ID) peer.ID {
	_ = "STUB: not implemented"
	return *new(peer.ID)
}

// SelectBest selects at most n peers sorted by responsiveness and latency.
//
// SelectBest parametrized by N because sync protocol relies on receiving data from redundant
// connections to guarantee that it will get complete data set.
// If it doesn't get complete data set it will have to fallback into hash resolution, which is
// generally more expensive.
func (p *Peers) SelectBest(n int) []peer.ID { _ = "STUB: not implemented"; return nil }

// SelectBestWithProtocols is similar to SelectBest but filters peers by supported protocols.
// If protocols is empty, it returns the best peers regardless of the protocol.
// If protocols is not empty, it returns the best peers that support at least one of the protocols.
func (p *Peers) SelectBestWithProtocols(n int, protocols []protocol.ID) []peer.ID {
	_ = "STUB: not implemented"
	return nil
}

func (p *Peers) selectBest(n int, protocols []protocol.ID) []peer.ID {
	_ = "STUB: not implemented"
	return nil
}

func (p *Peers) Total() int { _ = "STUB: not implemented"; return 0 }

func (p *Peers) Stats() Stats { _ = "STUB: not implemented"; return *new(Stats) }

type Stats struct {
	Total                int
	GlobalAverageLatency time.Duration
	BestPeers            []PeerStats
}

func (s *Stats) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type PeerStats struct {
	ID       peer.ID
	Success  int
	Failures int
	Latency  time.Duration
}

func (p *PeerStats) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
