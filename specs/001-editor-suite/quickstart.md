# Quickstart Guide

## Prerequisites

- **Go**: v1.22 or higher
- **Node.js**: v20 or higher (for Web UI)
- **Docker**: Optional, for running test databases
- **Buf**: For working with Protobuf definitions (`brew install bufbuild/buf/buf`)

## Building the Project

### 1. Build the CLI/Server (Backend)

```bash
# From project root
go mod download
go build -o neodb ./cmd/neodb
```

### 2. Build the Web UI

```bash
cd app/neodb-spa
npm install
npm run build
cd ../..
```

## Running the Application

### 1. Start the Server

The `neodb` binary runs the Connect-Go server which serves both the API and the static Web assets.

```bash
./neodb serve --port 8080
```

Open [http://localhost:8080](http://localhost:8080) to use the Visual Explorer.

### 2. Use the CLI

You can use the same binary as a CLI client.

```bash
# Connect to a database
./neodb connect --driver postgres --url "postgres://user:pass@localhost:5432/postgres"

# List tables
./neodb ls

# Execute a query
./neodb query "SELECT * FROM users LIMIT 5"
```

## Development Workflow

### Regenerate Protobufs

If you edit `.proto` files in `pkg/service/*/proto`:

```bash
buf generate
```

### Run Tests

```bash
go test ./...
```