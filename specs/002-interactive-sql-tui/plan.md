# Implementation Plan: Interactive SQL TUI

**Branch**: `002-interactive-sql-tui` | **Date**: 2026-02-09 | **Spec**: [specs/002-interactive-sql-tui/spec.md](spec.md)
**Input**: Feature specification from `specs/002-interactive-sql-tui/spec.md` + User Request for Internal SQLite.

## Summary

Implement a terminal-based SQL client (TUI) in Go using a Hexagonal Architecture. The application will allow developers to manage database connections (PostgreSQL, MySQL, SQLite), execute queries via a text editor interface, view results in formatted tables, and export data. It focuses on productivity with vim-like navigation and fast startup. **Internal state (Projects, Connections, History) is stored in a local SQLite database.**

## Technical Context

**Language/Version**: Go 1.25+
**Primary Dependencies**: 
- TUI Framework: `charmbracelet/bubbletea`
- TUI Components: `charmbracelet/bubbles`
- Database Drivers: `lib/pq` (Postgres), `go-sql-driver/mysql`, `modernc.org/sqlite` (Pure Go)
- Keyring: `zalando/go-keyring`
- Styling: `charmbracelet/lipgloss`
**Storage**: 
- **Internal DB**: SQLite file at `~/.config/neodb/neodb.db` (Managed via `RepositoryPort`).
- **Snippets**: SQL files in `~/.config/neodb/snippets/` (Optional/Hybrid).
**Testing**: Standard `testing` package, `testify/assert`.
**Target Platform**: Linux, macOS, Windows.
**Project Type**: CLI Application.
**Performance Goals**: Startup < 500ms.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **Code Quality**: Go 1.25+ enforced. Hexagonal architecture.
- **Testing Standards**: Unit tests for Domain and Services. Integration tests for SQLite Repository.
- **User Experience**: TUI consistency via Bubble Tea.
- **Performance**: Internal SQLite ensures fast history search and structured data access.
- **Security**: Password storage via keyring (FR-010).

## Project Structure

### Documentation (this feature)

```text
specs/002-interactive-sql-tui/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
└── tasks.md             # Phase 2 output
```

### Source Code (repository root)

```text
cmd/
└── neodb/                    # Main entry point (wire up adapters)
internal/
├── domain/                   # Core business entities (Project, Connection, Session, History)
├── port/                     # Interfaces (DatabasePort, RepositoryPort, KeyringPort)
├── service/                  # Use cases (ProjectService, SessionService, QueryService)
├── infra/                    # Implementations
│   ├── ui/                   # TUI implementation (Bubble Tea)
│   ├── database/             # Target SQL drivers (Postgres, MySQL, SQLite)
│   ├── keyring/              # System keyring adapter
│   └── filesystem/           # Snippet storage
└── common/                   # Shared utilities (logging, envloader)
```

**Structure Decision**: Hexagonal Architecture. `RepositoryPort` isolates the internal SQLite storage logic. `DatabasePort` isolates external target interactions.
