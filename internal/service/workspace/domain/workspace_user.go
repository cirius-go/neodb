package domain

import (
	"context"
	"time"
)

// WorkspaceUser represents a user associated with a workspace.
type WorkspaceUser struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	WorkspaceID string
	UserID      string
}

// WorkspaceUsers represents a collection of workspace users.
type WorkspaceUsers interface {
	// Insert a new workspace user with the given workspace user.
	Insert(ctx context.Context, wu WorkspaceUser) error
	// DeleteAllMembersByWorkspaceID deletes all workspace users associated with the given
	// workspace ID.
	DeleteAllMembersByWorkspaceID(ctx context.Context, workspaceID string) error
	// BulkInsert inserts multiple workspace users in a single operation.
	BulkInsert(ctx context.Context, wus []WorkspaceUser, batchSize int) error
	// FindMemberByWorkspaceAndUserID finds a workspace user by workspace ID and user ID.
	FindMemberByWorkspaceAndUserID(ctx context.Context, workspaceID, userID string) (WorkspaceUser, error)
	// ReplaceByWorkspaceID replaces all workspace users associated with the given workspace ID.
	ReplaceByWorkspaceID(ctx context.Context, workspaceID string, workspaceUsers []*WorkspaceUser) error
}
