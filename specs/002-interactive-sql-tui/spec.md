# Feature Specification: Interactive SQL TUI

**Feature Branch**: `002-interactive-sql-tui`
**Created**: 2026-02-09
**Status**: Draft
**Input**: User description: "NeoDB is a cross-database CLI application..."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Manage Database Connections (Priority: P1)

As a developer, I want to save, list, and connect to multiple database configurations (PostgreSQL, MySQL, SQLite) so that I can quickly switch between local, staging, and production environments without re-entering credentials.

**Why this priority**: Core prerequisite. Users cannot run queries without connecting.

**Independent Test**: Can be fully tested by creating a profile, listing it, and successfully establishing a connection to a live database.

**Acceptance Scenarios**:

1. **Given** no existing connections, **When** I add a new PostgreSQL connection with valid credentials, **Then** it is saved to the configuration and listed.
2. **Given** a list of saved connections, **When** I select one, **Then** the application establishes a session with that database.
3. **Given** a saved connection, **When** I choose to edit it, **Then** I can update the host, user, or password.

---

### User Story 2 - Interactive Query Execution (Priority: P1)

As a developer, I want to write SQL in a terminal editor and execute it, viewing the results in a scrollable, formatted table, so that I can inspect data and verify logic.

**Why this priority**: The primary function of the tool is to query databases.

**Independent Test**: Can be tested by connecting to a DB, typing `SELECT 1`, and verifying the result is displayed in a table.

**Acceptance Scenarios**:

1. **Given** an active connection, **When** I type a valid SQL query and press execute, **Then** the results are displayed in a formatted table with headers.
2. **Given** an active connection, **When** I run a query with a syntax error, **Then** a clear error message is displayed without crashing the app.
3. **Given** a large result set, **When** I scroll down, **Then** I can see subsequent rows (pagination/scrolling).

---

### User Story 3 - Schema Browsing (Priority: P2)

As a developer, I want to browse the database structure (tables, columns, indexes) in a side panel or menu so that I can write queries without context switching to external docs.

**Why this priority**: enhancing productivity and "IDE-like" feel.

**Independent Test**: Connect to a DB with known tables and verify the list matches.

**Acceptance Scenarios**:

1. **Given** an active connection, **When** I open the schema view, **Then** a list of all tables in the current database is shown.
2. **Given** a selected table, **When** I inspect it, **Then** I see its columns, data types, and primary keys.

---

### User Story 4 - Result Export (Priority: P3)

As a developer, I want to export the current query results to CSV or JSON so that I can share them or use them in other tools.

**Why this priority**: Useful utility but not blocking core usage.

**Independent Test**: Run a query, export to CSV, verify file content matches results.

**Acceptance Scenarios**:

1. **Given** a query result displayed, **When** I choose "Export to CSV", **Then** a file is created containing the result data in CSV format.
2. **Given** a query result, **When** I choose "Export to JSON", **Then** a JSON file is created.

### Edge Cases

- **Connection Loss**: If the database connection drops while the app is open, the system MUST display a "Connection Lost" status and offer a "Reconnect" action.
- **Corrupt Config**: If the configuration file is malformed (e.g., manual edit error), the system MUST display a descriptive error on startup and offer to back it up and start fresh.
- **Permission Denied**: If exporting to a file/directory without write permissions, the system MUST display a "Permission Denied" error without crashing.
- **Query Timeout**: If a query takes longer than the configured timeout, the system MUST allow the user to cancel the operation.
- **Large Blobs**: If a result contains large text/binary data, the system MUST truncate display in the grid but allow full inspection on selection.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST support connections to PostgreSQL, MySQL, and SQLite databases.
- **FR-002**: System MUST persist connection profiles (alias, type, host, port, user, dbname) between sessions.
- **FR-003**: System MUST provide a text editor interface within the terminal for writing SQL.
- **FR-004**: System MUST display query results in a grid/table format that aligns columns.
- **FR-005**: System MUST allow users to navigate through large result sets (scrolling or paging).
- **FR-006**: System MUST display error messages returned by the database engine clearly.
- **FR-007**: System MUST provide a mechanism to list all tables in the connected database.
- **FR-008**: System MUST allow exporting the current result set to a local CSV or JSON file.
- **FR-009**: System MUST allow storing and recalling saved query snippets as individual local files in a user-defined directory (defaulting to `~/.config/neodb/snippets/`).
- **FR-010**: System MUST handle database passwords securely using the operating system's native keyring/credential store.

### Key Entities

- **ConnectionProfile**: Stores connection details (Alias, Engine, Host, Port, User, Password/Auth, DatabaseName).
- **QuerySnippet**: A named block of SQL code saved by the user.
- **QueryResult**: The data returned from a query, consisting of Headers (Column definitions) and Rows (Values).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can establish a connection to a local database in under 15 seconds (including configuration time for a new profile).
- **SC-002**: Application starts and is ready for input in under 500ms.
- **SC-003**: Query results for 1000 rows render in the TUI in under 1 second.
- **SC-004**: 100% of standard SQL syntax errors from the DB are displayed to the user.

## Assumptions

- The environment is a standard Linux/macOS/Windows terminal supporting ANSI colors.
- Users have network access to the target databases.
- "Debugging" refers to error inspection and potentially `EXPLAIN` execution, not stepping through stored procedures.
- "Scripting" refers to executing SQL scripts, not an embedded scripting language.