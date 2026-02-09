# Implementation Plan - NeoDB Editor Suite

**Feature**: NeoDB Editor Suite  
**Status**: DRAFT  
**Branch**: 001-editor-suite-refactor

## Technical Context

> [!IMPORTANT]
> - **Language**: Go (Backend/CLI), TypeScript/Angular (Frontend)
> - **Frameworks**: Cobra/Viper (CLI), Standard `net/http` (REST API), Angular + Material (Web)
> - **Infrastructure**: Docker for local DBs, Nx for Monorepo, Wails (Desktop)
> - **Key Libraries**: `pgx` (PostgreSQL), `go-sqlite3` (SQLite), `sqlx` or `squirrel` (Query building)

### Constraints & Unknowns

- **Database Drivers**: Validated `pgx` (Postgres), `go-sql-driver/mysql` (MySQL), and `modernc.org/sqlite` (SQLite, pure Go).
  - *Status*: RESOLVED (See `research.md`)
- **API Strategy**: Switched to Standard REST/JSON for simplicity and direct frontend/Wails integration.
  - *Status*: RESOLVED (See `research.md`)
- **Desktop Bundling**: Selected Wails for unified Go+Angular binary.
  - *Status*: RESOLVED (See `research.md`)
- **Cross-Platform SQL Parsing**: Decision to rely on DB introspection and execution errors for Phase 1.
  - *Status*: RESOLVED (See `research.md`)

## Constitution Check

> [!CAUTION]
> **Core Principles Review**:
> 1.  **Code Quality**: Will use idiomatic Go (standard project layout) and Angular (Nx standards).
> 2.  **Testing**: Will implement unit tests for all core logic and integration tests for DB connectors using Docker containers (testcontainers-go?).
> 3.  **UX**: CLI will support JSON output (FR-005). Errors will be typed and actionable.
> 4.  **Performance**: Will benchmark result streaming for large datasets (SC-002).

## Gate: Phase 0 (Research)

- [x] **Research Task 1**: Evaluate Go database drivers for introspection capabilities (Postgres, MySQL, SQLite).
- [x] **Research Task 2**: Prototype Connect-Go vs REST (Decision: REST).
- [x] **Research Task 3**: Investigate Wails vs Electron vs Local Server for the "Desktop" experience.
- [x] **Research Task 4**: Determine strategy for SQL parsing/validation (Client-side vs Server-side vs DB-side).

## Gate: Phase 1 (Design)

- [ ] `data-model.md`: Define `ConnectionProfile`, `Schema`, `Table`, `Column`, `QueryResult` structs/messages.
- [ ] `contracts/openapi.yaml`: Define the REST API endpoints using OpenAPI 3.0.
- [ ] `quickstart.md`: Guide for running the CLI and Web UI locally.
- [ ] Agent Context Update: Run `.specify/scripts/bash/update-agent-context.sh`.

## Gate: Phase 2 (Implementation)

*(To be detailed after Design phase)*