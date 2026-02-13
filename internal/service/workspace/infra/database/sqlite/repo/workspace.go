package repo

import (
	"context"

	"cirius-go/neodb/internal/service/workspace/domain"

	"gorm.io/gorm"
)

// Workspaces represents a collection of workspace.
type Workspaces struct {
	db *gorm.DB
}

// NewWorkspaces creates a new Workspaces repository.
func NewWorkspaces(db *gorm.DB) *Workspaces {
	return &Workspaces{
		db: db,
	}
}

// UpdateByID updates a workspace by its ID with the given data.
func (r *Workspaces) UpdateByID(ctx context.Context, ws domain.Workspace) error {
	if _, err := gorm.G[domain.Workspace](r.db).Where("id = ?", ws.ID).Select("name", "owner_id").Updates(ctx, ws); err != nil {
		return domain.ErrDatabaseFailure.WithCause(err)
	}
	return nil
}
