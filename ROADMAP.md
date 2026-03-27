# Roadmap

This document outlines the planned development direction for taraDocs.

## Released

### v2.0.0 (2026-03-27)
- Standalone Go (Fiber) + Vue 3 (Inertia.js) documentation viewer
- Dual-format support: `.md` and `.html` files
- Tiptap WYSIWYG editor with autosave
- Dynamic DOCS_PATH configuration via environment variable
- Security hardening: path traversal protection, CORS config, rate limiting
- Unit test coverage for all core handlers

## Planned

### v2.1.0 - Search and Navigation
- Full-text search across all documents
- Table of contents auto-generation from headings
- Keyboard shortcut navigation (j/k/g/G)
- Breadcrumb trail for deep folder structures

### v2.2.0 - Themes and Customization
- Multiple built-in themes (light, dark, sepia)
- Custom CSS injection via config
- Font size and line-height user preference
- Print-friendly view

### v2.3.0 - Productivity Features
- Drag-and-drop file reordering in sidebar
- In-browser file rename and move
- Batch file operations (copy, delete, move)
- Markdown frontmatter metadata support (title, date, tags)

### v3.0.0 - Multi-user and Collaboration
- Optional authentication layer (Basic Auth / token)
- Read-only public mode with optional write access
- Change history with simple diff view
- Webhook integration for external notifications on file change

## Long-term Ideas
- Plugin system for custom renderers (diagrams, math, etc.)
- Mobile-optimized reader view
- Export to PDF or static HTML bundle
- Docker Compose stack with Nginx reverse proxy examples
