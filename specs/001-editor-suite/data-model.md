# Data Model

## Entities

### ConnectionProfile
Represents a saved database connection configuration.

| Field | Type | Description |
|-------|------|-------------|
| `id` | UUID | Unique identifier |
| `name` | String | User-friendly display name |
| `driver` | Enum | `postgres`, `mysql`, `sqlite` |
| `host` | String | Hostname or IP |
| `port` | Int | Port number |
| `user` | String | Username |
| `password` | String | Encrypted password or reference to secure storage |
| `database` | String | Target database name |
| `ssl_mode` | Enum | `disable`, `require`, `verify-full` |
| `created_at` | Timestamp | |
| `updated_at` | Timestamp | |

### SchemaInfo
Hierarchical representation of the database structure.

- **Database** has many **Schemas** (or just "Default" for MySQL/SQLite)
- **Schema** has many **Tables** and **Views**
- **Table** has many **Columns**

### Column
| Field | Type | Description |
|-------|------|-------------|
| `name` | String | Column name |
| `data_type` | String | Native DB type (e.g., `VARCHAR`, `INT`) |
| `is_nullable` | Boolean | |
| `is_primary_key`| Boolean | |
| `default_value` | String | |

## API Resources (JSON)

The internal domain model will map to these JSON resources in the REST API.

### Query Result
Not persisted, transient response.

- **QueryResult**:
  - `columns`: Array of column metadata (name, type).
  - `rows`: Array of row arrays. Each value is a primitive or null.
  - `rows_affected`: Integer.
  - `duration_ms`: Integer.

## Persistence
- **Profiles**: Stored in a local file (e.g., `~/.neodb/config.yaml` or SQLite) for CLI/Desktop. encrypted if possible.
