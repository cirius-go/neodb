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

## Research Area 2: API & Transport (Connect-Go)

**Objective**: Verify if Connect-Go can serve CLI (gRPC) and Web (Browser) clients simultaneously without Envoy.

### Findings
- **Connect-Go**: Designed specifically to serve gRPC, gRPC-Web, and the Connect protocol over a single HTTP/1.1 or HTTP/2 port.
- **Browser Support**: The `connect-web` (or `@connectrpc/connect-web`) client library allows the Angular app to talk directly to the Go server using the Connect protocol (or gRPC-Web) without a proxy.
- **CLI Support**: The Go client can use standard gRPC or Connect protocol.

### Decision
- Use **Connect-Go** for all service definitions.
- Expose a single HTTP server.
- Web UI uses `@connectrpc/connect-web`.
- CLI uses Connect-Go client.

---

## Research Area 3: Desktop Architecture

**Objective**: Choose a strategy for the "Desktop App" requirement (FR-004).

### Options
1.  **Wails**: Wraps Go backend + Web frontend in a native WebView. Produces a single binary. seamless integration.
2.  **Electron**: Requires Node.js. Heavy. We'd have to spawn the Go binary as a sidecar. Complex packaging.
3.  **Lorca / WebView**: Lighter wrappers, but Wails offers better tooling/bindings.
4.  **Local Server + Browser**: Not a "native app".

### Decision
- **Wails**: It aligns perfectly with our stack (Go + Angular). It allows us to reuse the Angular code 100% for the web version and the desktop version. The Go "Service" layer can be bound directly to the frontend in Wails, or we can just keep using localhost HTTP/Connect for consistency between Web and Desktop modes.
- **Implementation Note**: To keep architecture unified, the Desktop app will likely start the Connect server internally and the UI will connect to `localhost`. Or use Wails bindings for "offline" feel. *Refinement*: Using Wails bindings for everything differs from the Web (HTTP) approach. **Hybrid Approach**: The Core Logic is exposed via Connect-Go. The Desktop App (Wails) starts this server in-process. The UI calls the local server. This ensures 100% code reuse between Web and Desktop UI.

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