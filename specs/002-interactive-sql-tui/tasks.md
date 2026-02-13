# Tasks: Interactive SQL TUI

**Feature Branch**: `002-interactive-sql-tui`
**Spec**: [specs/002-interactive-sql-tui/spec.md](spec.md)
**Plan**: [specs/002-interactive-sql-tui/plan.md](plan.md)

## Phase 1: Setup
**Goal**: Initialize project structure and tooling.

- [x] T001 Initialize Go module `neodb` in root (if not exists)
- [x] T002 Create directory structure (`cmd`, `internal/domain`, `internal/port`, `internal/service`, `internal/infra`, `internal/common`)
- [ ] T003 Setup `Taskfile.yaml` or `Makefile` for build, lint, and test commands
- [ ] T004 Configure `golangci-lint` in `.golangci.yml` (or verify existing)
- [ ] T005 Create `cmd/neodb/main.go` with basic "Hello World" to verify build

## Phase 2: Foundation (Core Domain & Ports)
**Goal**: Implement the core domain entities, ports, and the internal SQLite persistence layer required for US1.

- [ ] T010 [P] Define `Project` and `ConnectionProfile` entities in `internal/domain/project.go` and `internal/domain/connection.go`
- [ ] T011 [P] Define `Session`, `QueryHistory`, and `CachedResult` entities in `internal/domain/session.go` and `internal/domain/query.go`
- [ ] T012 [P] Define `DatabasePort` interface in `internal/port/database.go`
- [ ] T013 [P] Define `RepositoryPort` interface in `internal/port/repository.go`
- [ ] T014 [P] Define `KeyringPort` interface in `internal/port/keyring.go`
- [ ] T015 Implement `KeyringAdapter` using `zalando/go-keyring` in `internal/infra/keyring/keyring.go`
- [ ] T016 Setup Internal SQLite connection logic (GORM or raw SQL) in `internal/infra/repository/sqlite.go`
- [ ] T017 Implement `RepositoryAdapter` methods for Project Management (Create, List, Get) in `internal/infra/repository/project_repo.go`
- [ ] T018 Implement `RepositoryAdapter` methods for Connection Management (Save, List, Get, Delete) in `internal/infra/repository/connection_repo.go`
- [ ] T019 Implement `RepositoryAdapter` methods for Session & History in `internal/infra/repository/session_repo.go`
- [ ] T020 Create unit tests for `RepositoryAdapter` (integration test with temporary SQLite file) in `internal/infra/repository/repository_test.go`

## Phase 3: Manage Database Connections (US1)
**Goal**: Allow users to save, list, and connect to database profiles.
**Priority**: P1

- [ ] T030 [US1] Create unit tests for `ProjectService` in `internal/service/project_service_test.go`
- [ ] T031 [US1] Implement `ProjectService` (CRUD logic) in `internal/service/project_service.go`
- [ ] T032 [US1] Create unit tests for `ConnectionService` in `internal/service/connection_service_test.go`
- [ ] T033 [US1] Implement `ConnectionService` (CRUD + Keyring integration) in `internal/service/connection_service.go`
- [ ] T034 [US1] Initialize Bubbletea TUI structure (Model, Update, View) in `internal/infra/ui/model.go`
- [ ] T035 [US1] Implement `ProjectListView` component (Select Project) in `internal/infra/ui/project_list.go`
- [ ] T036 [US1] Implement `ConnectionListView` component (Select Connection) in `internal/infra/ui/connection_list.go`
- [ ] T037 [US1] Implement `ConnectionForm` component (Inputs for Host, User, etc.) in `internal/infra/ui/connection_form.go`
- [ ] T038 [US1] Wire up Services to TUI in `cmd/neodb/main.go` and ensure navigation between Project/Connection lists works

## Phase 4: Interactive Query Execution (US2)
**Goal**: Connect to target databases and execute SQL queries.
**Priority**: P1

- [ ] T050 [US2] Implement `DatabaseAdapter` for PostgreSQL (using `lib/pq`) in `internal/infra/database/postgres.go`
- [ ] T051 [US2] Implement `DatabaseAdapter` for MySQL (using `go-sql-driver/mysql`) in `internal/infra/database/mysql.go`
- [ ] T052 [US2] Implement `DatabaseAdapter` for SQLite (using `modernc.org/sqlite`) in `internal/infra/database/sqlite.go`
- [ ] T053 [US2] Implement `DatabaseFactory` to return correct adapter based on engine type in `internal/infra/database/factory.go`
- [ ] T054 [US2] Create unit tests for `SessionService` in `internal/service/session_service_test.go`
- [ ] T055 [US2] Implement `SessionService` (Connect, Disconnect, Log Query) in `internal/service/session_service.go`
- [ ] T056 [US2] Implement `EditorView` using `bubbles/textarea` in `internal/infra/ui/editor.go`
- [ ] T057 [US2] Implement `ResultView` using `bubbles/table` in `internal/infra/ui/result.go`
- [ ] T058 [US2] Implement `QueryService` (Execute query via DatabasePort, return QueryResult) in `internal/service/query_service.go`
- [ ] T059 [US2] Wire up Execution Flow: Editor -> Ctrl+Enter -> QueryService -> ResultView in `internal/infra/ui/main_layout.go`

## Phase 5: Schema Browsing (US3)
**Goal**: Display database tables and columns in a side panel.
**Priority**: P2

- [ ] T070 [US3] Implement `GetTables` and `GetTableSchema` in `internal/infra/database/postgres.go`
- [ ] T071 [US3] Implement `GetTables` and `GetTableSchema` in `internal/infra/database/mysql.go`
- [ ] T072 [US3] Implement `GetTables` and `GetTableSchema` in `internal/infra/database/sqlite.go`
- [ ] T073 [US3] Implement `SchemaView` component (Tree/List of tables) in `internal/infra/ui/schema.go`
- [ ] T074 [US3] Integrate `SchemaView` into `MainLayout` (Toggleable Sidebar) in `internal/infra/ui/main_layout.go`

## Phase 6: Result Export (US4)
**Goal**: Export query results to CSV/JSON.
**Priority**: P3

- [ ] T080 [US4] Implement `ExportService` (Marshaling logic) in `internal/service/export_service.go`
- [ ] T081 [US4] Add Export Keybindings (Ctrl+E) and Format Selection Modal in `internal/infra/ui/result.go`
- [ ] T082 [US4] Integrate ExportService with UI to write files in `internal/infra/ui/main_layout.go`

## Phase 7: Polish & Cross-Cutting
**Goal**: Finalize UI styling and ensure robustness.

- [ ] T090 Apply Lipgloss styling to all TUI components for consistent theme in `internal/infra/ui/style.go`
- [ ] T091 Ensure error handling is graceful (Error Modals/Toast messages) across all views
- [ ] T092 Verify cross-platform builds (Linux, macOS, Windows)

## Dependencies

- **US1** requires **Phase 2** (Foundation) to be complete.
- **US2** requires **US1** (to establish connection) and **Phase 2**.
- **US3** requires **US2**'s active connection.
- **US4** requires **US2**'s query results.

## Implementation Strategy

1.  **MVP**: Complete Phases 1, 2, 3, and 4. This gives a functional SQL runner.
2.  **Enhanced**: Complete Phase 5 (Schema) and 6 (Export).
3.  **Polish**: Refine UI and UX.
