# Contributing to taraDocs

Thank you for your interest in contributing to taraDocs!

## 1. Development Setup

### Prerequisites
- Go 1.23+
- Node.js 18+

### Quick Start
```bash
# 1. Install Dependencies
go mod download
npm install

# 2. Setup Environment
cp .env.example .env

# 3. Build Frontend
npm run build

# 4. Run Server (Concurrently runs Go & Vite)
npm start
```

## 2. Directory Structure
- **Handlers**: `internal/handlers` (Core file reader operations)
- **Routes**: `cmd/server/main.go`
- **Frontend**: `resources/` (Vue 3 / Inertia pages)
- **Docs**: `docs/` (Sandbox markdown/html folder)

## 3. Workflow
- Always format code with `go fmt ./...` before committing.
- Ensure `make build-linux` passes locally.
- For frontend logic, run `npm run build` to update public assets.

## 4. Docker Testing
To test the production build locally without development dependencies:
```bash
docker build -t taradocs .
docker run -p 3000:3000 taradocs
```
