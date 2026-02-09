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

## API Messages (Protobuf mappings)

The internal domain model will map closely to these Protobuf definitions.

### Query Result
Not persisted, transient.

- **RowBatch**: Streamed chunk of rows.
- **Value**: Polymorphic value (string, int, bool, null, bytes).

## Persistence
- **Profiles**: Stored in a local file (e.g., `~/.neodb/config.yaml` or SQLite) for CLI/Desktop. encrypted if possible.