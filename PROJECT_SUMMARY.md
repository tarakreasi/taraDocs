# Project Summary: taraDocs

**taraDocs** is an offline-first documentation viewer for the Tarakreasi ecosystem. It serves as a specialized frontend interface for browsing Markdown documentation.

## Architecture
- **Unified Application**: Frontend (Vue 3 + Vite) and Backend API (Go Fiber) are monolithic and located at `~/Project/taraDocs`.
- **Content Source**: Markdown files mapped via `DOCS_PATH` (default: `~/Project/taraDocs/docs`).

## Key Features
- **Decoupled Architecture**: Frontend consumes backend API via `X-Inertia` headers.
- **Offline Capable**: Designed to run locally without internet dependency.
- **Markdown Rendering**: Client-side parsing using `markdown-it`.

## Usage
1.  **Start Application**:
    ```bash
    cd ~/Project/taraDocs
    ./run_backend.sh
    ```
    *(Backend Go fiber and Frontend Vite will run concurrently via `npm start` inside `backend/`)*
