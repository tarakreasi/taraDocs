# Contributing to taraDocs

Thank you for your interest in contributing to taraDocs! This guide will help you get started.

## Ways to Contribute

- Report bugs or request features via [GitHub Issues](https://github.com/tarakreasi/taraDocs/issues)
- Improve documentation
- Fix bugs or implement features from the [Roadmap](ROADMAP.md)
- Review pull requests

## Development Setup

### Prerequisites
- Go 1.23+
- Node.js 20+

### Quick Start
```bash
git clone https://github.com/tarakreasi/taraDocs.git
cd taraDocs

# Install Backend Dependencies
go mod download

# Install Frontend Dependencies
npm install

# Setup Environment
cp .env.example .env

# Run Application (Go backend + Vite frontend)
npm start
```

The app will be available at `http://localhost:3000`.

## Project Structure

```
taraDocs/
├── cmd/server/         # Go entrypoint (main.go)
├── internal/
│   ├── handlers/       # HTTP controllers (docs.go, docs_test.go)
│   ├── routes/         # Route setup
│   ├── utils/          # Inertia, Vite, slug helpers
│   └── config/         # Session config (legacy, kept for reference)
├── resources/
│   ├── js/Pages/Docs/  # Vue 3 page components
│   ├── js/Components/  # Shared components
│   └── css/            # Tailwind stylesheet
├── views/              # Inertia root HTML template
├── docs/               # Default documentation folder (DOCS_PATH)
└── public/             # Compiled static assets
```

## Workflow

1. Fork the repository and create a feature branch:
   ```bash
   git checkout -b feat/your-feature-name
   ```
2. Make your changes.
3. Run tests before committing:
   ```bash
   go test ./...
   go vet ./...
   ```
4. Run a frontend build check:
   ```bash
   npm run build
   ```
5. Commit using [Conventional Commits](https://www.conventionalcommits.org/):
   ```bash
   git commit -m "feat: add table of contents generation"
   git commit -m "fix: prevent path traversal on create-file"
   ```
6. Open a Pull Request against `main`.

## Code Style

- Go: follow standard `gofmt` formatting. Run `go fmt ./...` before committing.
- Vue: use Composition API (`<script setup>`). Keep component files focused.
- No em/en dashes in code comments or documentation.
- No emoji in code files.

## Docker Testing

To verify the production Docker build locally:
```bash
docker build -t taradocs .
docker run -p 3000:3000 -v $(pwd)/docs:/app/docs taradocs
```

## Reporting a Bug

Please include:
- OS and Go version
- Steps to reproduce
- Expected vs actual behavior
- Relevant logs or screenshots
