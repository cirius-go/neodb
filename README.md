# neodb: A Multi-purpose Database Tool

`neodb` is a versatile database tool designed to provide a robust Go backend and an interactive SQL TUI. It aims to streamline database interactions and management.

## 📦 Project Overview

`neodb` is structured around a powerful Go backend, designed for efficient data management and serving as the core logic for database operations. It features:

- **Go Backend**: The primary server-side component, written in Go (version 1.25.5).
- **Interactive SQL TUI**: An upcoming interactive Terminal User Interface (TUI) for direct SQL interaction.

## 🚀 Quick Start

To get started with `neodb`, follow these steps:

```bash
# Clone the repository
git clone https://github.com/personal/neodb.git # Assuming this is the repo URL
cd neodb

# Run the Go application
go run cmd/neodb/main.go

# Build the Go application
go build cmd/neodb/main.go

# Run tests
go test ./...
```

## 📁 Project Structure

```
.
├── .envrc
├── .gitignore
├── devenv.lock
├── devenv.nix
├── devenv.yaml
├── GEMINI.md
├── go.mod
├── go.sum
├── README.md
├── Taskfile.yaml
├── .devenv/
├── .direnv/
├── .gemini/
│   └── commands/
├── .git/
├── .idea/
├── .specify/
│   ├── memory/
│   └── scripts/
│       └── bash/
├── .vscode/
│   ├── extensions.json
│   └── launch.json
├── cmd/
│   └── neodb/
├── internal/
│   ├── common/
│   │   ├── envloader/
│   │   ├── errors/
│   │   ├── i18n/
│   │   └── logger/
│   ├── domain/
│   ├── infra/
│   │   └── config/
│   └── service/
├── nix/
│   └── util.nix
└── specs/
    └── 002-interactive-sql-tui/
        ├── checklists/
        └── contracts/
```