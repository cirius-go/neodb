# Research: Interactive SQL TUI

**Date**: 2026-02-09
**Status**: Complete

## Decisions

### 1. TUI Framework
- **Decision**: Use **Charmbracelet Stack** (`bubbletea`, `bubbles`, `lipgloss`).
- **Rationale**: 
  - `bubbletea` (The Elm Architecture for Go) is the de-facto standard for modern, interactive Go TUIs.
  - `bubbles` provides ready-made components: `table` (for results), `viewport` (for large text), `textinput`/`textarea` (for SQL editor).
  - `lipgloss` allows for separation of styling and logic, crucial for a maintainable "IDE-like" UI.
  - Fits the "visually appealing" and "developer productivity" goals.

### 2. SQLite Driver (External & Internal)
- **Decision**: Use **`modernc.org/sqlite`** (Pure Go).
- **Rationale**:
  - CGO-free: simplifies cross-compilation and distribution.
  - Used for both connecting to user's SQLite databases AND for the **Application's Internal Storage**.
  - Performance is sufficient for single-user usage.

### 3. Keyring Library
- **Decision**: Use **`github.com/zalando/go-keyring`**.
- **Rationale**:
  - Simple API: `Set(service, user, password)`, `Get(service, user)`.
  - Supports: macOS Keychain, Linux (Secret Service/DBus), Windows Credential Manager.
  - Active maintenance.

### 4. Database Abstraction
- **Decision**: Use Go standard `database/sql` interfaces wrapped in the Domain `DatabasePort`.
- **Rationale**: 
  - The Domain should not depend on specific drivers.
  - `database/sql` provides the necessary abstractions.

### 5. Internal Storage Strategy
- **Decision**: Use a dedicated **Local SQLite Database** (`~/.config/neodb/neodb.db`) managed by GORM or raw SQL, instead of flat YAML config files.
- **Rationale**:
  - **Relational Data**: Projects have Connections; Sessions have Histories.
  - **Query capabilities**: Need to search through thousands of history items efficiently.
  - **State management**: "Sessions" and "Cached Results" require more structured storage than a config file.
  - **ACID**: Prevents config corruption during concurrent writes (though rare in single-user CLI, it's safer).
- **Alternatives Considered**:
  - `YAML/JSON`: Good for simple config, bad for history/session logging and relational data.
  - `BoltDB`/`Badger`: KV stores. Good for Go, but SQL is more natural for a SQL tool, and we already include the SQLite driver.