package common

import (
	"context"
	"io/fs"

	texttemplate "text/template"
)

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

// Aliases for i18n types.
type (
	// LocalizeConfig holds configuration for localizing messages.
	LocalizeConfig struct {
		MessageID    string
		TemplateData any
		PluralCount  any
		Funcs        texttemplate.FuncMap
	}
	// LocalizeUnmarshalFunc defines the function signature for unmarshaling localization files.
	LocalizeUnmarshalFunc = func(data []byte, v any) error
	// Localizer defines the contract for international lozalizer mechanisms.
	Localizer interface {
		// Localize localizes a message based on the provided language tag and configuration.
		Localize(langTag string, conf LocalizeConfig) string
		// RegisterUnmarshalFunc registers a custom unmarshal function for message files.
		RegisterUnmarshalFunc(name string, fn LocalizeUnmarshalFunc)
		// LoadMessageFileFS loads message files from the given filesystem at the specified path.
		LoadMessageFileFS(fsys fs.FS, path ...string) error
	}
)

// Validator defines the contract for data validation mechanisms.
type Validator interface {
	// Struct validates the given struct based on defined tags.
	Struct(s any) error
}

type (
	// Error defines the contract for error types.
	Error[ErrType, DetailType any] interface {
		Error() string
		Unwrap() error
		Is(target error) bool
		WithDetail(detail DetailType) ErrType
		WithInternal(err error) ErrType
	}

	// I18nError represents an internationalization error.
	I18nError[ErrType, DetailType, LocalizedErrType any] interface {
		Error[ErrType, DetailType]
		Localize(langTag string) LocalizedErrType
	}
)
