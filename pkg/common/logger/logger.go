package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"cirius-go/neodb/pkg/common"
)

// Logger is an implementation of common.Logger using slog.
type Logger struct {
	output io.Writer
	level  common.LoggerLevel
	format common.LoggerFormat
	inner  *slog.Logger
}

// Printf implements common.Logger.
func (l *Logger) Printf(format string, args ...any) {
	switch l.level {
	case common.LoggerLevelDebug:
		l.Debugf(format, args...)
	case common.LoggerLevelInfo:
		l.Infof(format, args...)
	case common.LoggerLevelWarn:
		l.Warnf(format, args...)
	case common.LoggerLevelError:
		l.Errorf(format, args...)
	default:
		l.Infof(format, args...)
	}
}

// LogLevel implements common.Logger.
func (l *Logger) LogLevel() common.LoggerLevel {
	return l.level
}

// WithLogLevel implements common.Logger.
func (l *Logger) WithLogLevel(level common.LoggerLevel) common.Logger {
	newL := New(l.output, l.format, level)
	return newL
}

// Debugf implements common.Logger.
func (l *Logger) Debugf(format string, args ...any) {
	l.inner.Debug(fmt.Sprintf(format, args...))
}

// Errorf implements common.Logger.
func (l *Logger) Errorf(format string, args ...any) {
	l.inner.Error(fmt.Sprintf(format, args...))
}

// Infof implements common.Logger.
func (l *Logger) Infof(format string, args ...any) {
	l.inner.Info(fmt.Sprintf(format, args...))
}

// Warnf implements common.Logger.
func (l *Logger) Warnf(format string, args ...any) {
	l.inner.Warn(fmt.Sprintf(format, args...))
}

// Debug implements common.Logger.
func (l *Logger) Debug(msg string, args ...any) {
	l.inner.Debug(msg, args...)
}

// Info implements common.Logger.
func (l *Logger) Info(msg string, args ...any) {
	l.inner.Info(msg, args...)
}

// Warn implements common.Logger.
func (l *Logger) Warn(msg string, args ...any) {
	l.inner.Warn(msg, args...)
}

// Error implements common.Logger.
func (l *Logger) Error(msg string, args ...any) {
	l.inner.Error(msg, args...)
}

// With implements common.Logger.
func (l *Logger) With(args ...any) common.Logger {
	newL := New(l.output, l.format, l.level)

	newL.inner = l.inner.With(args...)
	return newL
}

// Inner returns the inner slog.Logger instance.
func (l *Logger) Inner() *slog.Logger {
	return l.inner
}

var _ common.Logger = (*Logger)(nil)

// New creates a new logger that satisfies the common.Logger
// interface.
func New(output io.Writer, format common.LoggerFormat, level common.LoggerLevel) *Logger {
	l := &Logger{
		output: output,
		format: format,
		level:  level,
		inner:  newSlogLogger(os.Stdout, format, level),
	}
	return l
}

func newSlogLogger(output io.Writer, format common.LoggerFormat, level common.LoggerLevel) *slog.Logger {
	var (
		handler   slog.Handler
		slogLevel = slog.Level(level * 4)
	)
	switch format {
	case common.LoggerFormatJSON:
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: slogLevel,
		})
	case common.LoggerFormatText:
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: slogLevel,
		})
	default:
		panic("unsupported logger format")
	}

	return slog.New(handler)
}
