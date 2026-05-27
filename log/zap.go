package log

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is an exported type that embeds our logger.
type Log struct {
	logger *zap.Logger
	name   string
}

// Exported from Log basic logging options.

// Info prints formatted info level log message.
func (l Log) Info(format string, args ...any) { _ = "STUB: not implemented"; return }

// Debug prints formatted debug level log message.
func (l Log) Debug(format string, args ...any) { _ = "STUB: not implemented"; return }

// Error prints formatted error level log message.
func (l Log) Error(format string, args ...any) { _ = "STUB: not implemented"; return }

// Warning prints formatted warning level log message.
func (l Log) Warning(format string, args ...any) { _ = "STUB: not implemented"; return }

// Panic prints the log message and then panics.
func (l Log) Panic(format string, args ...any) { _ = "STUB: not implemented"; return }

// Fatal prints formatted fatal level log message.
func (l Log) Fatal(format string, args ...any) { _ = "STUB: not implemented"; return }

// Wrap and export field logic

// Field is a log field holding a name and value.
type Field zap.Field

// Field satisfies loggable field interface.
func (f Field) Field() Field {
	_ = "STUB: not implemented"

	// FieldNamed returns a field with the provided name instead of the default.
	return *new(Field)
}

func FieldNamed(name string, field LoggableField) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func (f Field) AddTo(enc zapcore.ObjectEncoder) { _ = "STUB: not implemented"; return }

// String returns a string Field.
func String(name, val string) Field { _ = "STUB: not implemented"; return *new(Field) }

// Strings returns a strings Field.
func Strings(name string, val []string) Field { _ = "STUB: not implemented"; return *new(Field) }

// Stringer returns an fmt.Sringer Field.
func Stringer(name string, val fmt.Stringer) Field { _ = "STUB: not implemented"; return *new(Field) }

type ShortString interface {
	ShortString() string
}

type shortStringAdapter struct {
	val ShortString
}

func (a shortStringAdapter) String() string { _ = "STUB: not implemented"; return "" }

func ShortStringer(name string, val ShortString) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func ZShortStringer(name string, val ShortString) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// Uint16 returns an uint32 Field.
func Uint16(name string, val uint16) Field { _ = "STUB: not implemented"; return *new(Field) }

// Uint32 returns an uint32 Field.
func Uint32(name string, val uint32) Field { _ = "STUB: not implemented"; return *new(Field) }

// Time returns a field for time.Time struct value.
func Time(name string, val time.Time) Field { _ = "STUB: not implemented"; return *new(Field) }

// Duration returns a duration field.
func Duration(name string, val time.Duration) Field { _ = "STUB: not implemented"; return *new(Field) }

// Err returns an error field.
func Err(err error) Field { _ = "STUB: not implemented"; return *new(Field) }

// Inline for inline logging.
func Inline(object zapcore.ObjectMarshaler) Field { _ = "STUB: not implemented"; return *new(Field) }

// Array for logging array efficiently.
func Array(name string, array zapcore.ArrayMarshaler) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func ZContext(ctx context.Context) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }

func NiceZapError(err error) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }

func Any(key string, value any) Field { _ = "STUB: not implemented"; return *new(Field) }

type marshalledContext struct {
	context.Context
}

func (c *marshalledContext) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// LoggableField as an interface to enable every type to be used as a log field.
type LoggableField interface {
	Field() Field
}

func unpack(fields []LoggableField) []zap.Field { _ = "STUB: not implemented"; return nil }

// FieldLogger is a logger that only logs messages with fields. It does not support formatting.
type FieldLogger struct {
	l    *zap.Logger
	name string
}

// With returns a logger object that logs fields.
func (l Log) With() FieldLogger { _ = "STUB: not implemented"; return *new(FieldLogger) }

// SetLevel returns a logger with level as the log level derived from l.
func (l Log) SetLevel(level *zap.AtomicLevel) Log {
	_ = "STUB: not implemented"
	// Warn if the new level is lower than the parent level
	return *new(Log)
}

// Check if following level is supported by the logger.
func (l Log) Check(level zapcore.Level) bool { _ = "STUB: not implemented"; return false }

// WithName appends a name to a current name.
func (l Log) WithName(prefix string) Log { _ = "STUB: not implemented"; return *new(Log) }

// Named overwrites name.
func (l Log) Named(name string) Log { _ = "STUB: not implemented"; return *new(Log) }

// WithFields returns a logger with fields permanently appended to it.
func (l Log) WithFields(fields ...LoggableField) Log { _ = "STUB: not implemented"; return *new(Log) }

// WithContext creates a Log from an existing log and a context object.
func (l Log) WithContext(ctx context.Context) Log { _ = "STUB: not implemented"; return *new(Log) }

// Zap returns internal zap logger.
func (l Log) Zap() *zap.Logger {
	_ = "STUB: not implemented"

	// note: we construct the fieldset on the fly, below, rather than simply adding `name' as a field since it may change
	// if a child logger is created from a parent. once a field has been added to a logger it cannot be changed or removed.
	// see WithName, above.
	return nil
}

// Info prints message with fields.
func (fl FieldLogger) Info(msg string, fields ...LoggableField) { _ = "STUB: not implemented"; return }

// Debug prints message with fields.
func (fl FieldLogger) Debug(msg string, fields ...LoggableField) { _ = "STUB: not implemented"; return }

// Error prints message with fields.
func (fl FieldLogger) Error(msg string, fields ...LoggableField) { _ = "STUB: not implemented"; return }

// Warning prints message with fields.
func (fl FieldLogger) Warning(msg string, fields ...LoggableField) {
	_ = "STUB: not implemented"
	return
}

// Panic prints message with fields.
func (fl FieldLogger) Panic(msg string, fields ...LoggableField) { _ = "STUB: not implemented"; return }

// DebugField is only added if debug level is enabled.
func DebugField(logger *zap.Logger, field zap.Field) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}
