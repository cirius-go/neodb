# Quickstart: NeoDB Editor Suite

**Branch**: `001-editor-suite`

## Prerequisites

- **Go**: 1.25+
- **Node.js**: 20+ (for Angular)
- **Wails**: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- **Buf**: `go install github.com/bufbuild/buf/cmd/buf@latest`
- **Playwright**: `npx playwright install`

## Setup

1. **Install Dependencies**:
   ```bash
   go mod download
   cd frontend && npm install
   ```

2. **Generate Protobufs**:
   ```bash
   # From the project root, buf will find all .proto files in the service/*/proto/v1 directories
   buf generate
   ```

3. **Initialize Frontend (if not already done)**:
   ```bash
   # In 'frontend' directory, create Angular app with Vite and add Tailwind CSS
   # Assuming 'frontend' is the root of your Angular project.
   # For a new project:
   cd frontend
   npm create vite@latest . -- --template angular
   npm install -D tailwindcss postcss autoprefixer
   npx tailwindcss init -p
   # Configure tailwind.config.js and postcss.config.js as per Tailwind docs.
   # Add Tailwind directives to src/styles.css (or similar).
   ```

4. **Run Services (Development)**:
   The system is composed of multiple internal services exposed via gRPC.
   ```bash
   # Starts the gRPC server that wires all services
   go run internal/app/neodb-server/cmd/main.go
   ```

5. **Run Frontends**:
   - **CLI**: `go run internal/app/neodb/cmd/main.go`
   - **Web**: `go run internal/app/neodb-spa/cmd/main.go` (This will start the Go server that serves the Angular SPA)
   - **Desktop**: `wails dev -dir internal/app/neodb-desktop` (Assuming `wails.json` is within `internal/app/neodb-desktop`)

## Testing

### Backend (GoConvey)
Run unit tests for core services:
```bash
go test ./internal/service/... 
# OR for the UI with the GoConvey runner
goconvey
```

### Frontend (Playwright)
Run E2E tests:
```bash
cd frontend
npx playwright test
```

## Development Workflow

1. **Modify Contract**: Edit `.proto` files in `service/<name>/proto/v1`.
2. **Regenerate**: `buf generate`.
3. **Implement**: Update `internal/service/{domain}` logic.
4. **Wire**: Update `internal/app/neodb-server` for the server, or `internal/app/neodb` for CLI, etc.
5. **Consume**: Update Angular services or CLI client.
