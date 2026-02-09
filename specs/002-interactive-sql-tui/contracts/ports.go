package port

import (
	"context"
	"neodb/pkg/domain"
)

// DatabasePort defines the interface for interacting with TARGET database engines.
type DatabasePort interface {
	Connect(ctx context.Context, profile domain.ConnectionProfile, password string) error
	Disconnect() error
	Ping(ctx context.Context) error
	ExecuteQuery(ctx context.Context, query string) (*domain.QueryResult, error)
	ExecuteStatement(ctx context.Context, query string) (*domain.QueryResult, error)
	GetTables(ctx context.Context) ([]string, error)
	GetTableSchema(ctx context.Context, tableName string) ([]domain.ColumnDefinition, error)
}

// RepositoryPort defines access to the INTERNAL SQLite database.
// It manages Projects, Connections, Sessions, and History.
type RepositoryPort interface {
	// Project Management
	CreateProject(ctx context.Context, project domain.Project) error
	ListProjects(ctx context.Context) ([]domain.Project, error)
	GetProject(ctx context.Context, id string) (*domain.Project, error)

	// Connection Management
	SaveConnection(ctx context.Context, conn domain.ConnectionProfile) error
	ListConnections(ctx context.Context, projectID string) ([]domain.ConnectionProfile, error)
	GetConnection(ctx context.Context, id string) (*domain.ConnectionProfile, error)
	DeleteConnection(ctx context.Context, id string) error

	// Session & History
	CreateSession(ctx context.Context, session domain.Session) error
	EndSession(ctx context.Context, sessionID string) error
	LogQuery(ctx context.Context, history domain.QueryHistory) error
	GetHistory(ctx context.Context, sessionID string) ([]domain.QueryHistory, error)
	SearchHistory(ctx context.Context, query string) ([]domain.QueryHistory, error)

	// Result Caching
	SaveCachedResult(ctx context.Context, result domain.CachedResult) error
	GetCachedResult(ctx context.Context, historyID string) (*domain.CachedResult, error)
}

// KeyringPort handles secure password storage.
type KeyringPort interface {
	SetPassword(service, user, password string) error
	GetPassword(service, user string) (string, error)
	DeletePassword(service, user string) error
}