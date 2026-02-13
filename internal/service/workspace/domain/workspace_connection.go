package domain

import "time"

// WorkspaceConnection represents a connection related to a workspace.
type WorkspaceConnection struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	WorkspaceID  string
	ConnectionID string
}

// WorkspaceConnections represents a collection of workspace connections.
type WorkspaceConnections interface{}
