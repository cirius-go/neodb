package context

import (
	"context"
	"fmt"

	"cirius-go/neodb/cmd/neodb"
	"cirius-go/neodb/cmd/neodb/config"
	"cirius-go/neodb/pkg/common"
	"cirius-go/neodb/pkg/common/logger"
)

// Context represents the application context.
type Context struct {
	reqCtx context.Context
	config config.Config
	logger common.Logger
}

// RequestContext implements neodb.Context.
func (c *Context) RequestContext() context.Context {
	return c.reqCtx
}

// Logger implements neodb.Context.
func (c *Context) Logger() common.Logger {
	return c.logger
}

var _ neodb.Context = (*Context)(nil)

// NewFromContext creates a new application context from the given configuration.
func NewFromContext(ctx context.Context) (*Context, error) {
	appCtx := &Context{
		reqCtx: ctx,
	}
	for _, contextKey := range neodb.ContextKeyValues() {
		var ok bool
		switch contextKey {
		case neodb.ContextKeyLogger:
			appCtx.logger, ok = ctx.Value(contextKey).(common.Logger)
			if !ok {
				appCtx.logger = logger.Get()
			}
		case neodb.ContextKeyConfig:
			appCtx.config, ok = ctx.Value(contextKey).(config.Config)
			if !ok {
				return nil, fmt.Errorf("failed to retrieve config from context")
			}
		default:
			continue
		}
	}
	return appCtx, nil
}
