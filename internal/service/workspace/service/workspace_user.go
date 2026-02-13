package service

import (
	"context"
	"time"

	"cirius-go/neodb/internal/common"
	"cirius-go/neodb/internal/common/errors"
	"cirius-go/neodb/internal/service/workspace/domain"
)

// WorkspaceUser represents a workspace user service.
type WorkspaceUser struct {
	uow       UnitOfWork
	logger    common.Logger
	validator common.Validator
}

// NewWorkspaceUser creates a new instance of WorkspaceUser service.
func NewWorkspaceUser(uow UnitOfWork, logger common.Logger, validator common.Validator) *WorkspaceUser {
	return &WorkspaceUser{
		uow:       uow,
		logger:    logger.With("service", "workspace_user"),
		validator: validator,
	}
}

// AddOwnerInTx adds an owner to the workspace within a transaction.
func (s *WorkspaceUser) AddOwnerInTx(ctx context.Context, txuow TransactionalUnitOfWork, workspaceID, ownerID string) error {
	return txuow.WorkspaceUsers().Insert(ctx, domain.WorkspaceUser{
		ID:          common.NewUUID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserID:      ownerID,
		WorkspaceID: workspaceID,
	})
}

// UpdateMemberListInTx updates the member list of a workspace.
func (s *WorkspaceUser) UpdateMemberListInTx(ctx context.Context, txuow TransactionalUnitOfWork, ws domain.Workspace, userIDs []string) error {
	if err := txuow.WorkspaceUsers().DeleteAllMembersByWorkspaceID(ctx, ws.ID); err != nil {
		return errors.ErrInternalServer.WithInternal(err)
	}
	newMembers := make([]domain.WorkspaceUser, 0, len(userIDs))
	for _, userID := range userIDs {
		newMembers = append(newMembers, domain.WorkspaceUser{
			ID:          common.NewUUID(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			WorkspaceID: ws.ID,
			UserID:      userID,
		})
	}
	if err := txuow.WorkspaceUsers().BulkInsert(ctx, newMembers, DefaultBatchSize); err != nil {
		return errors.ErrInternalServer.WithInternal(err)
	}
	return nil
}
