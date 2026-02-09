# Implementation Plan: NeoDB Editor Suite

**Branch**: `001-editor-suite` | **Date**: 2026-02-09 | **Spec**: [specs/001-editor-suite/spec.md](specs/001-editor-suite/spec.md)
**Input**: Feature specification from `specs/001-editor-suite/spec.md`

## Summary

Build a multi-platform database editor suite (CLI, Web, Desktop) sharing a common Go core. The frontend applications are managed by an Nx workspace rooted at the project's top-level, with Angular projects (`neodb-spa`, `neodb-desktop`) residing directly within `internal/app`.
- **CLI**: Interactive TUI using `charmbracelet/bubbletea`.
- **Web**: Single Page Application (SPA) built with Angular (latest), Vite, and Tailwind CSS, managed by Nx within `internal/app/neodb-spa`.
- **Desktop**: Native-like application using Wails (Go backend + Angular frontend with Vite and Tailwind CSS), managed by Nx within `internal/app/neodb-desktop`.
- **Core**: Shared business logic divided into micro-services (`connection`, `session`, `discovery`, `query`, `result`) following Hexagonal Architecture.
- **Common**: `internal/common` module defines shared contracts (logging, envloader, UoW factory, i18n).
- **Communication**: gRPC for Frontend <-> Backend communication.
- **Backend Composition**: `internal/app/neodb-server` wires all services together.

## Technical Context

**Language/Version**: Go 1.25+, TypeScript 5.x, Angular (Latest Stable).
**Primary Dependencies**: 
- **Monorepo Management**: Nx (with `nx.json` at project root).
- **Communication**: `connect-rpc` (Go) / `@connectrpc/connect-web` (TS).
- **CLI**: `charmbracelet/bubbletea`.
- **Desktop**: `wails.io/v2`.
- **Frontend Build**: Vite (for Angular).
- **Frontend Styling**: Tailwind CSS.
- **Backend/Core**: `pgx`, `go-sql-driver/mysql`, `mattn/go-sqlite3`.
**Storage**: User config in `~/.config/neodb`.
**Testing**: `smartystreets/goconvey` (Backend), `playwright` (Frontend).
**Target Platform**: Linux, macOS, Windows.
**Project Type**: Polyglot Monorepo (Go + Nx/Angular) with Granular gRPC Services.

## Constitution Check

*GATE: Passed.*

- **Code Quality**: Strict typing via granular Protobuf packages. Nx enforces code quality across frontend.
- **Testing Standards**: Service-level unit tests (GoConvey) and E2E tests (Playwright). Nx facilitates unified testing for frontend apps/libs.
- **UX Consistency**: Unified API ensures consistent behavior across CLI/Web/Desktop. Tailwind CSS facilitates consistent styling.
- **Performance**: Streaming results via gRPC; efficient Go-based processing. Vite for fast frontend build/dev.

## Project Structure

### Documentation (this feature)

```text
specs/001-editor-suite/
├── plan.md              # This file
├── research.md          # Technology decisions
├── data-model.md        # Entities and Schema
├── quickstart.md        # Setup guide
├── contracts/           # Placeholder for generated protos
└── tasks.md             # Implementation tasks
```

### Source Code

```text
# Repository Root
.
├── nx.json                     # Nx configuration file at project root
├── tsconfig.base.json          # Nx base TS config at project root
├── internal/
│   ├── app/                    # Application Entry Points
│   │   ├── neodb/              # CLI Application (Go)
│   │   │   └── cmd/                
│   │   ├── neodb-server/       # Backend Server (Go)
│   │   │   └── cmd/                
│   │   ├── neodb-spa/          # Angular Web SPA (Nx project managed from root nx.json)
│   │   │   └── src/            # Angular app source
│   │   │   └── project.json    # Nx project configuration
│   │   ├── neodb-desktop/      # Angular Desktop (for Wails, Nx project managed from root nx.json)
│   │   │   └── src/            # Angular app source
│   │   │   └── project.json    # Nx project configuration
│   │   └── ... (other Go apps)
│   ├── common/                 # Shared Go Libraries (Contracts)
│   │   ├── envloader/          
│   │   ├── logging/            
│   │   ├── i18n/               
│   │   └── uow/                # UnitOfWork Factory
│   ├── service/                # Core Domain Logic (Hexagonal Architecture)
│   │   ├── connection/             
│   │   │   └── proto/v1/           # Protobuf Contracts
│   │   ├── discovery/              
│   │   │   └── proto/v1/           
│   │   ├── query/                  
│   │   │   └── proto/v1/           
│   │   ├── result/                 
│   │   │   └── proto/v1/           
│   │   └── session/                
│   │       └── proto/v1/           
│   ├── proto/                      # Generated Go Code from Protobufs
└── ... (other top-level config files like go.mod, devenv.nix)

```

**Structure Decision**: 
- The project is a polyglot monorepo with `nx.json` at the root.
- `internal/common` houses shared contracts like logging, envloader, UoW, and i18n.
- Nx-managed Angular frontend applications (`neodb-spa`, `neodb-desktop`) are created directly within `internal/app/<app_name>` paths.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| Hexagonal Architecture | Testability & Decoupling | Layered architecture couples domain to db/frameworks too tightly. |
| 5 Distinct Proto Packages | Separation of Concerns | Single "EditorService" bloats and couples distinct domains. |
| Nx Monorepo (custom paths) | User Requirement | Directly embedding Angular apps allows Nx to manage them while respecting project's `internal/app` structure. |
