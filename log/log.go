// Package log provides the both file and console (general) logging capabilities
// to spacemesh modules such as app and identity.
package log

import (
	"context"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// mainLoggerName is a name of the global logger.
const mainLoggerName = "00000.defaultLogger"

// where logs go by default.
var logWriter io.Writer = os.Stdout

// Logger is an interface for our logging API.
type Logger interface {
	Info(format string, args ...any)
	Debug(format string, args ...any)
	Panic(format string, args ...any)
	Error(format string, args ...any)
	Warning(format string, args ...any)
	With() FieldLogger
	WithContext(context.Context) Log
	WithName(string) Log
}

// AppLog is the local app singleton logger.
var (
	mu     sync.RWMutex
	AppLog Log
)

// GetLogger gets logger.
func GetLogger() Log { _ = "STUB: not implemented"; return *new(Log) }

// SetupGlobal overwrites global logger.
func SetupGlobal(logger Log) { _ = "STUB: not implemented"; return }

func init() {
	SetupGlobal(NewWithLevel(mainLoggerName,
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	))
}

// NewNop creates silent logger.
func NewNop() Log { _ = "STUB: not implemented"; return *new(Log) }

// NewWithLevel creates a logger with a fixed level and with a set of (optional) hooks.
func NewWithLevel(module string,
	level zap.AtomicLevel,
	encoder zapcore.Encoder,
	hooks ...func(zapcore.Entry) error,
) Log {
	_ = "STUB: not implemented"
	return *new(Log)
}

// NewFromLog creates a Log from an existing zap-compatible log.
func NewFromLog(l *zap.Logger) Log {
	_ = "STUB: not implemented"
	return *

	// public wrappers abstracting away logging lib impl
	new(Log)
}

// Info prints formatted info level log message.
func Info(msg string, args ...any) { _ = "STUB: not implemented"; return }

// Debug prints formatted debug level log message.
func Debug(msg string, args ...any) { _ = "STUB: not implemented"; return }

// Warning prints formatted warning level log message.
func Warning(msg string, args ...any) { _ = "STUB: not implemented"; return }

// With returns a FieldLogger which you can append fields to.
func With() FieldLogger { _ = "STUB: not implemented"; return *new(FieldLogger) }

// Panic writes the log message and then panics.
func Panic(msg string, args ...any) { _ = "STUB: not implemented"; return }
