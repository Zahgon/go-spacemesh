package logtest

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/log"
)

const testLogLevel = "TEST_LOG_LEVEL"

// New creates log.Log instance that will use testing.TB.Log internally.
func New(tb testing.TB, override ...zapcore.Level) log.Log {
	_ = "STUB: not implemented"
	return *new(log.Log)
}
