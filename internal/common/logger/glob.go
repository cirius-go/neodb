package logger

import (
	"os"
	"sync"

	"cirius-go/neodb/internal/common"
)

var (
	mu   sync.Mutex
	glob common.Logger = New(os.Stdout, common.LoggerFormatText, common.LoggerLevelInfo)
)

// Set sets the global logger instance.
func Set(l common.Logger) {
	if l == nil {
		glob.Errorf("set logger level to '%s' instead of set logger by nil", common.LoggerLevelInfo.String())
		return
	}

	mu.Lock()
	defer mu.Unlock()
	glob = l
}

// Get returns the global logger instance.
func Get() common.Logger {
	return glob
}

// Debugf logs a debug message using the global logger.
func Debugf(format string, args ...any) {
	glob.Debugf(format, args...)
}

// Errorf logs an error message using the global logger.
func Errorf(format string, args ...any) {
	glob.Errorf(format, args...)
}

// Infof logs an informational message using the global logger.
func Infof(format string, args ...any) {
	glob.Infof(format, args...)
}

// Warnf logs a warning message using the global logger.
func Warnf(format string, args ...any) {
	glob.Warnf(format, args...)
}

// Debug logs a debug message using the global logger.
func Debug(msg string, args ...any) {
	glob.Debug(msg, args...)
}

// Info logs an informational message using the global logger.
func Info(msg string, args ...any) {
	glob.Info(msg, args...)
}

// Warn logs a warning message using the global logger.
func Warn(msg string, args ...any) {
	glob.Warn(msg, args...)
}

// Error logs an error message using the global logger.
func Error(msg string, args ...any) {
	glob.Error(msg, args...)
}

// SetLogLevel sets the log level of the global logger.
func SetLogLevel(level common.LoggerLevel) {
	if !level.IsValid() {
		return
	}
	Set(glob.WithLogLevel(level))
}

// LogLevel gets the log level of the global logger.
func LogLevel() common.LoggerLevel {
	return glob.LogLevel()
}

// Printf prints a formatted message using the global logger.
func Printf(format string, args ...any) {
	glob.Printf(format, args...)
}
