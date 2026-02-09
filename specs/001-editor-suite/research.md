# Research & Decisions - NeoDB Editor Suite

**Status**: FINAL  
**Date**: 2026-02-09

## Research Area 1: Database Drivers

**Objective**: Select the best Go drivers for PostgreSQL, MySQL, and SQLite, prioritizing introspection capabilities and stability.

### Findings
- **PostgreSQL**: `jackc/pgx` is the de-facto standard for modern Go. It offers superior performance and type handling compared to `lib/pq` (which is in maintenance mode). `pgx` also provides a stdlib `database/sql` interface if needed, but its native interface is richer.
- **MySQL**: `go-sql-driver/mysql` is the widely accepted standard. Stable and feature-complete.
- **SQLite**:
  - `mattn/go-sqlite3`: Standard, fast, but requires CGO. Makes cross-compilation (e.g., for Desktop release) harder.
  - `modernc.org/sqlite`: Pure Go translation of SQLite. Slower than CGO version but sufficient for an "Editor" tool (not a high-throughput server). Simplifies build process immensely.
  - `ncruces/go-sqlite3`: Newer pure Go alternative using Wasm, very promising performance.

### Decision
- **PostgreSQL**: Use `jackc/pgx/v5`.
- **MySQL**: Use `go-sql-driver/mysql`.
- **SQLite**: Use `modernc.org/sqlite` initially for easy cross-compilation. If performance becomes an issue for local analysis, switch to `mattn/go-sqlite3` with build tags.

---

## Research Area 2: API & Transport (REST)

**Objective**: Determine the best API strategy for CLI, Web, and Desktop clients.

### Findings
- **Original Plan (Connect-Go)**: Good for type safety, but introduces complexity with Proto generation and frontend dependencies.
- **REST/JSON**: Standard `net/http` in Go is robust and sufficient. Angular has native `HttpClient`. Wails integration is straightforward with REST (or internal Go method bindings).
- **Simplicity**: Removing Protos simplifies the build chain (no `buf`, no `protoc`) and reduces friction for "quick edits" in a local tool context.

### Decision
- **Protocol**: Standard REST API with JSON.
- **Router**: Use Go 1.22+ standard library `http.ServeMux` (or `chi` if middleware needs get complex).
- **Documentation**: OpenAPI 3.0 spec for contract definition.
- **Clients**:
  - Web: Angular `HttpClient`.
  - CLI: Standard Go `net/http` Client.
  - Desktop: Wails can share the same HTTP handlers or bind directly. We will stick to HTTP for consistency across all 3 apps initially.

---

## Research Area 3: Desktop Architecture

**Objective**: Choose a strategy for the "Desktop App" requirement (FR-004).

### Options
1.  **Wails**: Wraps Go backend + Web frontend in a native WebView. Produces a single binary. seamless integration.
2.  **Electron**: Requires Node.js. Heavy. We'd have to spawn the Go binary as a sidecar. Complex packaging.
3.  **Lorca / WebView**: Lighter wrappers, but Wails offers better tooling/bindings.
4.  **Local Server + Browser**: Not a "native app".

### Decision
- **Wails**: It aligns perfectly with our stack (Go + Angular). It allows us to reuse the Angular code 100% for the web version and the desktop version.
- **Architecture**: The Go application will start an HTTP server. The Wails frontend (Angular) will consume this API via `localhost` (or internal bridge). This ensures the Web UI and Desktop UI run the exact same code.

---

## Research Area 4: SQL Parsing & Introspection

**Objective**: How to handle schema validation and autocomplete.

### Findings
- **Full Parsing**: `vitess` or `pg_query_go` are heavy.
- **Introspection**: Querying `information_schema` is standard and reliable across DBs.
- **Autocomplete**: Requires a parser to know context (e.g., "SELECT * FROM [cursor]").

### Decision
- **Phase 1**: Do not implement a full SQL parser. Rely on **Introspection** (querying system catalogs) to get table/column info.
- **Validation**: Send query to DB, parse returned error message.
- **Future**: Evaluate `vitess` parser if complex client-side validation is requested.
