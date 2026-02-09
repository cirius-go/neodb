package common

import "context"

// LoggerLevel represents the level of logging.
// ENUM(Debug=-1,Info,Warn,Error)
//
//go:generate go-enum --marshal --values
type LoggerLevel int

// LoggerFormat represents the format of the logger output.
// ENUM(Text,JSON)
//
//go:generate go-enum --marshal --values
type LoggerFormat int

type (
	// Logger defines the contract for a logging mechanism.
	Logger interface {
		// Info logs an informational message.
		Info(msg string, args ...any)
		// Infof logs a formatted informational message.
		Infof(format string, args ...any)
		// Warn logs a warning message.
		Warn(msg string, args ...any)
		// Warnf logs a formatted warning message.
		Warnf(format string, args ...any)
		// Debug logs a debug message.
		Debug(msg string, args ...any)
		// Debugf logs a formatted debug message.
		Debugf(format string, args ...any)
		// Error logs an error message.
		Error(msg string, args ...any)
		// Errorf logs a formatted error message.
		Errorf(format string, args ...any)
		// With adds contextual information to the logger.
		With(args ...any) Logger
		// WithLogLevel with a specific log level.
		WithLogLevel(level LoggerLevel) Logger
		// LogLevel gets the current log level.
		LogLevel() LoggerLevel
		// Printf logs a formatted message (generic).
		Printf(format string, args ...any)
	}
)

// EnvLoader defines a function type for loading environment variables.
type EnvLoader func(ctx context.Context) error
