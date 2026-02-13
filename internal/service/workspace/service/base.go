package service

import (
	"context"

	"cirius-go/neodb/internal/common"
	"cirius-go/neodb/internal/service/workspace/domain"
)

// Context represents the service context.
type Context interface {
	ReqCtx() context.Context
	Logger() common.Logger
}

// DefaultBatchSize defines the default batch size for bulk operations.
const DefaultBatchSize = 1000

// TransactionalUnitOfWork represents a transactional unit of work.
type TransactionalUnitOfWork interface {
	// Workspaces returns the workspace repository.
	Workspaces() domain.Workspaces
	// WorkspaceUsers returns the workspace user repository.
	WorkspaceUsers() domain.WorkspaceUsers
	// WorkspaceConnections returns the workspace connection repository.
	WorkspaceConnections() domain.WorkspaceConnections
}

// UnitOfWork represents a unit of work.
type UnitOfWork interface {
	// Embeded the transactional unit of work.
	TransactionalUnitOfWork
	// Transaction executes a function within a transaction.
	Transaction(ctx context.Context, tFn func(ctx context.Context, tuow TransactionalUnitOfWork) error) error
}

// WorkspaceUserService represents the workspace user service.
type WorkspaceUserService interface {
	// AddOwnerInTx adds an owner to the workspace within a transaction.
	AddOwnerInTx(ctx context.Context, txuow TransactionalUnitOfWork, workspaceID, ownerID string) error
}
