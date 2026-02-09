---
description: "Task list for NeoDB Editor Suite implementation"
---

# Tasks: NeoDB Editor Suite

**Input**: Design documents from `/specs/001-editor-suite/`
**Prerequisites**: plan.md, spec.md, data-model.md, contracts/, research.md

**Tests**: Included as requested in the plan (unit tests + integration tests with Docker).

**Organization**: Tasks are grouped by user story. Phase 2 (Foundational) blocks all stories. Stories 1-4 can sequentially or partially parallelize.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Parallelizable
- **[Story]**: [US1]..[US4]

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and structure validation

- [ ] T001 Verify project structure (pkg/, cmd/, app/) matches plan
- [ ] T002 [P] Initialize Go module dependencies (pgx, sqlite, connect-go)
- [ ] T003 [P] Configure Buf workspace for `pkg/service` protos
- [ ] T004 [P] Setup Wails project structure (if not already present) or placeholder

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure MUST be complete before user stories

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T005 Generate Go code from Protos (`connection.proto`, `query.proto`, `introspection.proto`)
- [ ] T006 Create generic `Driver` interface in `pkg/core/driver/driver.go` (Connect, Introspect, Query)
- [ ] T007 Define `ConnectionProfile` and `SchemaInfo` core structs in `pkg/core/models/models.go`
- [ ] T008 Implement config loader (save/load profiles) in `pkg/core/config/config.go`
- [ ] T009 Setup Connect-Go server boilerplate (Mux, CORS, Reflection) in `cmd/neodb/server/server.go`
- [ ] T010 [P] Setup `testcontainers-go` helper for integration tests in `pkg/testutil/db.go`

**Checkpoint**: Core interfaces and server scaffold ready.

---

## Phase 3: User Story 1 - CLI Connection & Introspection (Priority: P1) 🎯 MVP

**Goal**: Connect to DB, list tables via CLI.

**Independent Test**: `neodb connect <url>` -> success; `neodb ls` -> lists tables.

### Tests for US1
- [ ] T011 [P] [US1] Integration test for Postgres connection in `tests/integration/postgres_test.go`
- [ ] T012 [P] [US1] Integration test for SQLite connection in `tests/integration/sqlite_test.go`

### Implementation for US1
- [ ] T013 [P] [US1] Implement Postgres Driver (generic interface) in `pkg/drivers/postgres/postgres.go`
- [ ] T014 [P] [US1] Implement SQLite Driver (generic interface) in `pkg/drivers/sqlite/sqlite.go`
- [ ] T015 [US1] Implement `ConnectionService` (GRPC) in `pkg/service/connection/service.go`
- [ ] T016 [US1] Implement `IntrospectionService` (GRPC) in `pkg/service/introspection/service.go`
- [ ] T017 [US1] Implement CLI `root` and `connect` commands in `cmd/neodb/cmd/connect.go`
- [ ] T018 [US1] Implement CLI `ls` command in `cmd/neodb/cmd/ls.go`

**Checkpoint**: Can connect and list tables from CLI.

---

## Phase 4: User Story 2 - Ad-hoc Query Execution (Priority: P1)

**Goal**: Run SQL queries and see results.

**Independent Test**: `neodb query "SELECT 1"` -> output table.

### Tests for US2
- [ ] T019 [P] [US2] Integration test for `ExecuteStream` in `tests/integration/query_test.go`

### Implementation for US2
- [ ] T020 [US2] Update Driver interface to support `Query` and `Exec` if not done
- [ ] T021 [US2] Implement `QueryService` (ExecuteStream) in `pkg/service/query/service.go`
- [ ] T022 [US2] Create CLI `query` command in `cmd/neodb/cmd/query.go`
- [ ] T023 [P] [US2] Implement ASCII table formatter in `pkg/common/format/table.go`
- [ ] T024 [P] [US2] Implement JSON output formatter in `pkg/common/format/json.go`

**Checkpoint**: Can run queries and see formatted output.

---

## Phase 5: User Story 3 - Web-based Visual Explorer (Priority: P2)

**Goal**: Web UI for browsing data.

**Independent Test**: Browser opens, can navigate tables and view data.

### Tests for US3
- [ ] T025 [P] [US3] Unit tests for Angular ConnectionService in `app/neodb-spa/src/app/services/connection.service.spec.ts`

### Implementation for US3
- [ ] T026 [US3] Generate TypeScript/Angular clients from Protos
- [ ] T027 [US3] Implement `ConnectionService` (Angular) in `app/neodb-spa/src/app/services/connection.service.ts`
- [ ] T028 [US3] Implement `QueryService` (Angular) in `app/neodb-spa/src/app/services/query.service.ts`
- [ ] T029 [US3] Create Connection Manager Component in `app/neodb-spa/src/app/features/connection/connection.component.ts`
- [ ] T030 [US3] Create Data Grid/Table Browser Component in `app/neodb-spa/src/app/features/table-browser/table-browser.component.ts`
- [ ] T031 [US3] Configure Go server to serve SPA static files in `cmd/neodb/server/static.go`

**Checkpoint**: Web UI is functional.

---

## Phase 6: User Story 4 - Cross-Platform Desktop App (Priority: P3)

**Goal**: Standalone Wails app.

**Independent Test**: Binary launches native window, same functionality as Web.

### Implementation for US4
- [ ] T032 [US4] Configure Wails to wrap the Go server and Angular dist in `main.go` (Wails entrypoint)
- [ ] T033 [US4] Ensure Wails bindings (if any) or HTTP proxy setup works in `pkg/desktop/app.go`
- [ ] T034 [US4] Build Linux desktop bundle

**Checkpoint**: Desktop app runs.

---

## Phase 7: Polish & Cross-Cutting Concerns

- [ ] T035 [P] Add error handling interceptors (gRPC to user-friendly)
- [ ] T036 Update README.md with usage instructions
- [ ] T037 [P] Polish UI styling (Material Design adjustments)
- [ ] T038 Verify all acceptance criteria from spec.md

---

## Dependencies & Execution Order

1. **Setup & Foundational** (T001-T010) must be done first.
2. **US1** (T011-T018) builds on Foundation.
3. **US2** (T019-T024) can start after US1 is stable (shares Driver interface).
4. **US3** (T025-T031) depends on Backend Services (US1/US2) being available, but UI dev can mock.
5. **US4** (T032-T034) depends on US3 (Web UI) being ready to wrap.

### Parallel Opportunities
- Drivers (Postgres/SQLite) can be built in parallel.
- Web UI (US3) can be built in parallel with CLI (US1/US2) if API contract is stable.
- Tests can be written in parallel with implementation.
