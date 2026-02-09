# Quickstart: Interactive SQL TUI

## Prerequisites

- Go 1.25+ installed
- A running database (PostgreSQL, MySQL) or a SQLite file.

## Installation

```bash
# Clone the repository
git clone https://github.com/personal/neodb.git
cd neodb

# Build the binary
go build -o neodb cmd/neodb/main.go

# Verify installation
./neodb --version
```

## Running the TUI

1. Start the application:
   ```bash
   ./neodb
   ```

2. **Add a Connection**:
   - Navigate to "New Connection" (or press `n`).
   - Enter details:
     - **Engine**: postgres
     - **Host**: localhost
     - **Port**: 5432
     - **User**: myuser
     - **Database**: mydb
   - You will be prompted for a password (saved securely).

3. **Connect**:
   - Select the profile from the list and press `Enter`.

4. **Execute Query**:
   - Type `SELECT * FROM my_table;` in the editor.
   - Press `Ctrl+Enter` to execute.
   - Use Arrow Keys/PageUp/PageDown to scroll results.

5. **Export**:
   - Press `Ctrl+E` to export results to CSV.

## Troubleshooting

- **"Keyring not found"**: Ensure your OS has a credential manager active (Gnome Keyring, KWallet, macOS Keychain).
- **"Connection failed"**: Check if the database server is running and reachable.
