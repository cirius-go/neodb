package service

import (
	"context"
	"time"

	"cirius-go/neodb/internal/common/errors"
	"cirius-go/neodb/internal/common/uniqueid"
	"cirius-go/neodb/internal/common/validator"
	"cirius-go/neodb/internal/service/workspace/domain"
)

// Workspace represents a workspace service.
type Workspace struct {
	uow       UnitOfWork
	wsUserSvc WorkspaceUserService
}

// NewWorkspace creates a new Workspace service.
func NewWorkspace(uow UnitOfWork, wsUserSvc WorkspaceUserService) *Workspace {
	return &Workspace{
		uow:       uow,
		wsUserSvc: wsUserSvc,
	}
}

// Create the workspace.
func (s *Workspace) Create(ctx Context, ownerID, name string) (*domain.Workspace, error) {
	workspace := &domain.Workspace{
		ID:        uniqueid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		OwnerID:   ownerID,
		Name:      name,
	}
	if err := validator.Struct(workspace); err != nil {
		return nil, err
	}
	if err := s.uow.Transaction(ctx.ReqCtx(), func(ctx context.Context, tuow TransactionalUnitOfWork) error {
		if err := tuow.Workspaces().Insert(ctx, workspace); err != nil {
			if errors.Is(err, domain.ErrUniqueViolation) {
				return domain.ErrWorkspaceNameConflict
			}
			return err
		}
		return s.wsUserSvc.AddOwnerInTx(ctx, tuow, workspace.ID, ownerID)
	}); err != nil {
		return nil, err
	}
	return workspace, nil
}

// UpdateWorkspaceData represents the data to update a workspace.
type UpdateWorkspaceData struct {
	Name    string
	OwnerID string
	UserIDs []string
}

// UpdateByID updates the workspace.
func (s *Workspace) UpdateByID(ctx Context, id string, data UpdateWorkspaceData) error {
	if err := s.uow.Transaction(ctx.ReqCtx(), func(ctx context.Context, tuow TransactionalUnitOfWork) error {
		ws, err := tuow.Workspaces().FindByID(ctx, id)
		if err != nil {
			return err
		}

		ws.Name = data.Name
		ws.OwnerID = data.OwnerID

		wsUsers := make([]*domain.WorkspaceUser, 0, len(data.UserIDs))
		existedOwnerUser := false
		for _, userID := range data.UserIDs {
			if userID == data.OwnerID {
				existedOwnerUser = true
			}
			wsUsers = append(wsUsers, &domain.WorkspaceUser{
				ID:          uniqueid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				WorkspaceID: ws.ID,
				UserID:      userID,
			})
		}
		if !existedOwnerUser {
			return domain.ErrWorkspaceOwnerMustBeMember
		}
		if err := tuow.Workspaces().UpdateByID(ctx, ws); err != nil {
			return err
		}

		return tuow.WorkspaceUsers().ReplaceByWorkspaceID(ctx, ws.ID, wsUsers)
	}); err != nil {
		return err
	}
	return nil
}
