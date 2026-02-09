# Data Model: NeoDB Editor Suite

**Feature**: 001-editor-suite
**Date**: 2026-02-09

## Protobuf Entities

The data model reflects the granular gRPC service structure. Each service has its own `.proto` definition and package.

### `neodb.connection.v1.ConnectionService`
*Physical connection management.*

| RPC | Input | Output | Description |
|---|---|---|---|
| `CreateProfile` | `ConnectionProfile` | `CreateProfileResponse` | Save a new connection config. |
| `ListProfiles` | `ListProfilesRequest` | `ListProfilesResponse` | List saved configs. |
| `TestConnection`| `TestConnectionRequest` | `TestConnectionResponse` | Verify connectivity without saving. |

**Core Messages**:
- `ConnectionProfile`: Defines connection parameters (id, name, driver, host, port, etc.).
- `CreateProfileRequest`, `CreateProfileResponse`, `ListProfilesRequest`, `ListProfilesResponse`, `TestConnectionRequest`, `TestConnectionResponse`.

### `neodb.session.v1.SessionService`
*Contextual state management.*

| RPC | Input | Output | Description |
|---|---|---|---|
| `CreateSession` | `CreateSessionRequest` | `CreateSessionResponse` | Start a new session from a profile. |
| `CloseSession` | `CloseSessionRequest` | `CloseSessionResponse` | Close an active session. |
| `GetHistory` | `GetHistoryRequest` | `GetHistoryResponse` | Get executed queries history. |

**Core Messages**:
- `CreateSessionRequest`, `CreateSessionResponse`, `CloseSessionRequest`, `CloseSessionResponse`, `GetHistoryRequest`, `GetHistoryResponse`.
- `HistoryEntry`: Represents an entry in the query history.

### `neodb.discovery.v1.DiscoveryService`
*Metadata introspection.*

| RPC | Input | Output | Description |
|---|---|---|---|
| `GetTables` | `GetTablesRequest` | `GetTablesResponse` | List tables in current schema. |
| `GetColumns` | `GetColumnsRequest` | `GetColumnsResponse` | List columns for a table. |

**Core Messages**:
- `GetTablesRequest`, `GetTablesResponse`, `GetColumnsRequest`, `GetColumnsResponse`.
- `TableInfo`: Metadata about a database table.
- `ColumnInfo`: Metadata about a specific column. (Note: A common `types.proto` might be considered later for shared types like this).

### `neodb.query.v1.QueryService`
*Execution engine.*

| RPC | Input | Output | Description |
|---|---|---|---|
| `ExecuteStream`| `ExecuteRequest` | `stream neodb.result.v1.RowBatch` | Stream rows for SELECT statements. |
| `ExecuteExec` | `ExecuteRequest` | `neodb.result.v1.ExecResult` | Execute non-query (UPDATE/INSERT/DELETE) statements. |

**Core Messages**:
- `ExecuteRequest`: Contains `session_id` and `sql`.

### `neodb.result.v1.ResultService`
*Data processing and export.*

| RPC | Input | Output | Description |
|---|---|---|---|
| `Export` | `ExportRequest` | `stream ExportChunk` | Export result set to a file format. |

**Core Messages**:
- `RowBatch`: Contains a batch of rows and optional column info.
- `Row`: A single row of values.
- `Value`: A union type for different data types.
- `ExecResult`: Summary of an execution (rows affected, error).
- `ExportRequest`, `ExportChunk`.
- `ColumnInfo`: Re-defined for now within `result.proto` to avoid circular dependencies with `discovery.proto` if both import a common `types.proto`.

