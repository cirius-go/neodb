# Data Model: Interactive SQL TUI

**Feature**: Interactive SQL TUI
**Status**: Draft

## Entities

### Project
*A logical grouping of database connections (e.g., "Personal Blog", "Work Main App").*

| Field | Type | Description |
|---|---|---|
| `ID` | UUID | Unique Identifier |
| `Name` | String | User-friendly name |
| `CreatedAt` | Timestamp | Creation time |

### ConnectionProfile
*Represents a saved database connection configuration.*

| Field | Type | Description |
|---|---|---|
| `ID` | UUID | Unique identifier |
| `ProjectID` | UUID | Foreign Key to Project |
| `Alias` | String | User-friendly name (e.g., "Prod Primary") |
| `Environment` | Enum | `local`, `dev`, `staging`, `prod` |
| `Engine` | Enum | `postgres`, `mysql`, `sqlite` |
| `Host` | String | Hostname or IP |
| `Port` | Integer | Connection port |
| `User` | String | Database username |
| `Database` | String | Database name or file path |
| `SSLMode` | Enum | `disable`, `require`, `verify-full` |
| `IsActive` | Boolean | Soft delete flag |

*Note: Password retrieved via KeyringPort.*

### Session
*Represents a period of interaction with a specific connection.*

| Field | Type | Description |
|---|---|---|
| `ID` | UUID | Unique Identifier |
| `ConnectionID` | UUID | FK to ConnectionProfile |
| `StartedAt` | Timestamp | When session began |
| `EndedAt` | Timestamp | When session closed (nullable) |

### QueryHistory
*A log of executed queries.*

| Field | Type | Description |
|---|---|---|
| `ID` | UUID | Unique Identifier |
| `SessionID` | UUID | FK to Session |
| `QueryText` | Text | The SQL executed |
| `ExecutedAt` | Timestamp | Execution time |
| `DurationMs` | Integer | Execution duration in ms |
| `Status` | Enum | `Success`, `Error` |
| `ErrorMessage` | Text | Error details (if any) |
| `RowsAffected` | Integer | Count of rows changed/returned |

### CachedResult
*A transient storage of result data for quick retrieval/paging in the editor.*

| Field | Type | Description |
|---|---|---|
| `ID` | UUID | Unique Identifier |
| `HistoryID` | UUID | FK to QueryHistory |
| `ColumnsJSON` | JSON | Header definitions |
| `DataJSON` | JSON | The result rows (possibly compressed/truncated) |
| `ExpiresAt` | Timestamp | TTL for cache cleanup |

## State Transitions (UI)

### AppState
- **ProjectSelection**: Choose active project.
- **ConnectionList**: List connections in project.
- **Connecting**: Establishing session.
- **Editor**: Main SQL input (Auto-complete uses History).
- **Results**: Viewing CachedResult.