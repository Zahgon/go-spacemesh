package peersync

import (
	"time"
)

type timedResponse struct {
	Response
	receiveTimestamp int64
}

type round struct {
	ID                uint64
	Timestamp         int64
	RequiredResponses int
	responses         []timedResponse
}

func (r *round) AddResponse(resp Response, timestamp int64) { _ = "STUB: not implemented"; return }

func (r *round) Ready() bool { _ = "STUB: not implemented"; return false }

func (r *round) Offset() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
