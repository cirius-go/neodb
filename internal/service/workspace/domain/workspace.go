package domain

import (
	"context"
	"time"

	"cirius-go/neodb/internal/common/errors"
)

// Workspace represents a workspace in the system.
type Workspace struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	OwnerID   string
	Name      string
}

// Workspaces represents a collection of workspaces.
type Workspaces interface {
	// Insert a new workspace with the given workspace.
	Insert(ctx context.Context, w *Workspace) error
	// UpdateByID updates a workspace by its ID with the given data.
	UpdateByID(ctx context.Context, ws Workspace) error
	// ListByOwnerID lists all workspaces owned by the given owner ID.
	ListByOwnerID(ctx context.Context, params ListingParams[string]) ([]Workspace, int64, error)
	// FindByID finds a workspace by its ID.
	FindByID(ctx context.Context, id string) (Workspace, error)
}

var (
	ErrWorkspaceNotFound          = errors.NewDomain("WorkspaceNotFound")
	ErrWorkspaceNameConflict      = errors.NewDomain("WorkspaceNameConflict")
	ErrWorkspaceOwnerMustBeMember = errors.NewDomain("WorkspaceOwnerMustBeMember")
)
