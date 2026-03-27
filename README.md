# taraDocs

taraDocs is a sleek, offline-first Markdown and HTML documentation viewer for the Tarakreasi ecosystem. It serves as a specialized frontend interface for browsing, creating, and editing local documentation without needing a back-end database.

## Key Features
- **Zero Database Dependency**: Pure file-system based storage.
- **Dual Support**: Read, render, and edit both `.md` (Markdown) and `.html` (Hypertext Markup) files smoothly.
- **Built for Speed**: Powered by **Go (Fiber)** on the backend and **Vue 3 (Inertia.js)** on the frontend.
- **Rich Editor**: Integrated Tiptap editor for seamless WYSIWYG documentation authoring.

## Tech Stack
- **Backend**: Go 1.23+, Fiber v2
- **Frontend**: Vue 3, Vite, Tailwind CSS, Inertia.js
- **Editor**: Tiptap

## Quick Start
### Prerequisites
- Go 1.23+
- Node.js 18+

### Setup
```bash
git clone https://github.com/tarakreasi/taraDocs.git
cd taraDocs

# Install Backend Dependencies
go mod download

# Install Frontend Dependencies
npm install

# Setup Environment
cp .env.example .env

# Run Application
npm start
```
By default, the application runs on `http://localhost:3000`.

## Configuration
You can configure the active source directory serving your documentation by editing `.env`:
```env
DOCS_PATH=/path/to/your/docs
PORT=3000
```
If `DOCS_PATH` is left blank, it defaults to the `docs/` folder in the project root.

## Docker Deployment
Easily deploy taraDocs using the provided multistage Dockerfile.
```bash
# Build Image
docker build -t taradocs .

# Run Container (mount your docs to /app/docs)
docker run -p 3000:3000 -v $(pwd)/docs:/app/docs taradocs
```

## Connect with Me
- **LinkedIn:** [linkedin.com/in/twantoro](https://www.linkedin.com/in/twantoro)
- **GitHub:** [github.com/tarakreasi](https://github.com/tarakreasi)
- **Email:** ajarsinau@gmail.com
