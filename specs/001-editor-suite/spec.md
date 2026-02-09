# Feature Specification: NeoDB Editor Suite

**Feature Branch**: `001-editor-suite`  
**Created**: 2026-02-09  
**Status**: Draft  
**Input**: User description: "Build neodb - database editor applications (cli,web,desktop) that can help developers work with databases"

## User Scenarios & Testing

### User Story 1 - CLI Connection & Introspection (Priority: P1)

As a developer, I want to use the CLI to connect to a database and inspect its structure so that I can quickly verify connectivity and schema without a GUI.

**Why this priority**: Establishes the core connectivity logic and provides immediate utility for scripting/terminal users. Foundation for all other apps.

**Independent Test**: Can be tested by running the CLI binary against a running database Docker container and verifying output matches expected schema.

**Acceptance Scenarios**:

1. **Given** a valid connection string (e.g., PostgreSQL), **When** I run `neodb connect <url>`, **Then** the system establishes a connection and stores the profile.
2. **Given** an active connection, **When** I run `neodb list tables`, **Then** I see a list of all tables in the default schema.
3. **Given** an invalid password, **When** I attempt to connect, **Then** I receive a clear "Authentication Failed" error.

---

### User Story 2 - Ad-hoc Query Execution (Priority: P1)

As a developer, I want to run arbitrary SQL queries via the CLI and see formatted results so that I can retrieve data or perform quick updates.

**Why this priority**: Core functionality of a database tool.

**Independent Test**: Execute a `SELECT` and `UPDATE` command via CLI and verify stdout/stderr and database state changes.

**Acceptance Scenarios**:

1. **Given** an active connection, **When** I run `neodb query "SELECT 1 as id"`, **Then** I see a formatted text table with column `id` and value `1`.
2. **Given** a large result set, **When** I run a query, **Then** the output is paginated or streamed without consuming excessive memory.
3. **Given** a syntax error in SQL, **When** I run the query, **Then** the database error message is displayed clearly.

---

### User Story 3 - Web-based Visual Explorer (Priority: P2)

As a developer, I want to launch a local web interface to visually browse tables and data so that I can explore the database more comfortably than in a terminal.

**Why this priority**: Provides the "Editor" experience for users who prefer GUIs.

**Independent Test**: Launch the web server and use a headless browser (e.g., Playwright) to navigate the UI.

**Acceptance Scenarios**:

1. **Given** the neodb server is running, **When** I visit the local URL, **Then** I see the dashboard with saved connections.
2. **Given** a selected table, **When** I click "Browse Data", **Then** I see a data grid with sorting and filtering options.
3. **Given** a data grid, **When** I edit a cell value and save, **Then** the change is persisted to the database.

---

### User Story 4 - Cross-Platform Desktop App (Priority: P3)

As a developer, I want a standalone desktop application so that I can integrate the tool into my OS workflow (dock, taskbar, offline access).

**Why this priority**: Enhances accessibility and OS integration but depends on the Web/Core logic.

**Independent Test**: Install and launch the desktop bundle; verify it connects and renders the same UI as the web version.

**Acceptance Scenarios**:

1. **Given** the desktop app is installed, **When** I launch it, **Then** it opens in a native window, not a browser tab.
2. **Given** the app is open, **When** I use native OS shortcuts (e.g., Cmd+Q), **Then** the app responds appropriately.

### Edge Cases

- **Network Interruption**: System handles connection loss gracefully during query execution (reconnects or reports error).
- **Large Result Sets**: Queries returning >10MB of data or >1M rows are paginated or streamed; memory limits are enforced.
- **Concurrent Modification**: System warns if the schema has changed since it was last fetched.
- **Unsupported Types**: System displays a placeholder/hex dump for binary or driver-specific types that cannot be rendered as text.
- **Permission Denied**: System clearly communicates when an action (e.g., DROP TABLE) fails due to database permissions.

## Requirements

### Functional Requirements

- **FR-001**: System MUST support connecting to PostgreSQL, MySQL, SQLite, and generic ODBC/JDBC sources.
- **FR-002**: System MUST allow managing multiple named connection profiles.
- **FR-003**: System MUST provide schema inspection (Listing tables, views, columns, types).
- **FR-004**: System MUST allow executing arbitrary SQL statements (SELECT, INSERT, UPDATE, DELETE).
- **FR-005**: CLI output MUST support both human-readable (ASCII table) and machine-parsable (JSON, CSV) formats.
- **FR-006**: Web and Desktop UIs MUST provide a data grid for viewing query results.
- **FR-007**: System MUST support full visual schema editing (CREATE/ALTER TABLE, add/remove columns, constraint management).

### Key Entities

- **ConnectionProfile**: Stores name, engine type, host, port, user, database name (credentials stored securely).
- **QueryResult**: Represents columns, rows, and execution metadata (duration, rows affected).
- **SchemaInfo**: Represents the structure of the connected database (tables, columns, constraints).

## Success Criteria

### Measurable Outcomes

- **SC-001**: CLI startup to query execution time is under 1 second for local databases.
- **SC-002**: Web/Desktop UI renders a table with 1000 rows in under 2 seconds.
- **SC-003**: Application supports simultaneous connections to at least 5 different databases.
- **SC-004**: 90% of common SQL syntax errors are reported with helpful context (line number/position) from the driver.