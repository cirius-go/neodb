# Research & Technical Decisions: NeoDB Editor Suite

**Feature**: 001-editor-suite
**Date**: 2026-02-09

## 1. Service Responsibilities

### Connection Service
- **Role**: Physical access to databases.
- **Responsibilities**: 
  - Manage Driver registry.
  - Parse connection strings (DSN).
  - Maintain connection pools (`*sql.DB` or `pgx.Pool`).
  - Test connectivity.

### Session Service
- **Role**: User context management.
- **Responsibilities**:
  - Track active session state (which Connection ID is in use).
  - Manage Transaction state (Begin/Commit/Rollback).
  - Store Query History.
  - Manage Session Variables (e.g., `SET logic_test_mode = true`).
  - Handle Cursor pagination state (if stateful).

### Discovery Service
- **Role**: Database Introspection.
- **Responsibilities**:
  - Fetch Table/View lists.
  - Fetch Column details (types, constraints).
  - Fetch Indexes/Foreign Keys.
  - Generate DDL for objects.

### Query Service
- **Role**: Execution Engine.
- **Responsibilities**:
  - Validate SQL syntax (basic checks).
  - Inject safeguards (Timeout, Max Rows).
  - Execute query using the Session's active connection.
  - Handle cancelation (Context).

### Result Service
- **Role**: Output Handling.
- **Responsibilities**:
  - Map Database types to gRPC types.
  - Stream rows to the client (batching).
  - Handle dynamic schema (unknown column count/types).
  - Export data to CSV/JSON/Excel.

## 2. API Architecture: Granular gRPC Services with Colocated Protos

### Decision: Separate Proto Packages Colocated with Services
**Context**: Each service will have separated proto package.
**Design**:
- `.proto` files for `ConnectionService` will reside in `service/connection/proto/v1/connection.proto`.
- Similar structure for `session`, `discovery`, `query`, and `result` services.

**Rationale**:
- **Cohesion**: Protocol definitions live directly with the service they define, improving discoverability and maintainability.
- **Versioning**: Clear versioning per service contract.

## 3. Data Streaming

### Decision: Server-Side Streaming for Results
**Context**: `SELECT * FROM large_table`.
**Implementation**: `rpc ExecuteStream(QueryRequest) returns (stream RowBatch)`.
- Sends data in chunks (e.g., 1000 rows per message) to avoid HTTP timeouts and memory spikes.

## 4. Frontend Integration

### Decision: Nx Monorepo for Angular Applications (Custom Paths)
**Context**: No dedicated `frontend` directory; `nx.json` at project root; Angular apps (`neodb-spa`, `neodb-desktop`) directly within `internal/app`.
**Design**:
- Nx will be initialized with `nx.json` at the project's root.
- Angular applications (e.g., `neodb-spa`, `neodb-desktop`) will be created by Nx directly into `internal/app/neodb-spa` and `internal/app/neodb-desktop` respectively, making them Nx projects.
- Nx `libs/` will be used for shared frontend code, potentially colocated with these applications or in a `internal/app/shared-frontend-libs` directory.

**Rationale**:
- **User Requirement**: Directly addresses the user's specific structural preference.
- **Monorepo Benefits**: Still leverages Nx's advantages for managing multiple frontend applications and shared libraries from a single root configuration.

### Decision: Angular Services wrap Granular gRPC Clients
- `ConnectionClient`, `SessionClient`, etc. are injected into Angular components.
- State is managed in Angular Signals (v17+) reflecting the backend Session state.

### Decision: Angular with Vite and Tailwind CSS
**Context**: User requested Angular app created via Vite and using Tailwind CSS.
**Rationale**:
- **Vite**: Faster development server and build times compared to Webpack.
- **Tailwind CSS**: Utility-first CSS framework for rapid and consistent UI development.

## 5. Testing Strategy

### Decision: GoConvey (Backend)
**Context**: Need expressive, BDD-style unit tests for complex service logic.
**Implementation**:
- Use `github.com/smartystreets/goconvey/convey` in `_test.go` files.

### Decision: Playwright (Frontend/E2E)
**Context**: Need reliable cross-browser testing for the Web and Wails frontend.
**Implementation**:
- Playwright tests will be configured within the Nx workspace.
- Run against the local `neodb-web` server.
- Validates critical flows: Connect -> Query -> View Results.

## 6. Shared Infrastructure (`internal/common`)

### Decision: Shared Contracts in `internal/common`
**Context**: "lib of golang is defined in internal/common module".
**Components**:
- `logging`: Interface for structured logging.
- `envloader`: Contract for loading environment variables/config.
- `uow`: Factory interface for creating Unit of Work instances (managing transactions across services).
- `i18n`: Contract for internationalization.

**Rationale**:
- **Standardization**: Ensures all services use consistent patterns for cross-cutting concerns.
- **Decoupling**: Services depend on these interfaces, not concrete implementations (which might live in `infra` or separate packages).
