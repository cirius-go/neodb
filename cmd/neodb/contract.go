package neodb

import (
	"context"

	"cirius-go/neodb/pkg/common"
)

// Context represents the application context.
type Context interface {
	// Logger returns the logger instance.
	Logger() common.Logger
	// RequestContext returns the request context.
	RequestContext() context.Context
}

// ContextKey represents a key type for context values.
// ENUM(config,logger)
//
//go:generate go-enum --marshal --values
type ContextKey string
