# Changelog

All notable changes to this project will be documented in this file.

## [v2.0.0] - 2026-03-27
### Architecture Redesign (taraDocs Evolution)
- **Migrated**: Transformed the monolithic `taraNote Go` full-stack app into `taraDocs` (a flat-file, offline-first Documentation CMS).
- **Removed**: Eliminated SQLite database, authentication sessions, and `users`/`notes` GORM models for a purely file-based CMS operation.
- **Added**: Dual-extension support. The engine now parses, renders, and allows native editing of `.html` files alongside `.md` files.
- **Updated**: Revamped `DocsReader` and `DocsEditor` to natively load HTML DOM strings without forcing markdown recompilation if the file is an HTML file.
- **Refactored**: Flattened application directory by merging the `taraDocs` frontend and backend workspaces into a single root folder.
- **Changed**: Simplified `Makefile` and `Dockerfile` to completely remove legacy database migration logic. Quick Vite static bundling introduced in Docker multi-stage.
- **Fixed**: Resolved Inertia.js navigation state preservation issues. Sidebar navigation now freezes UI state (`preserve-state` and `preserve-scroll`), ensuring folder trees remain expanded while reading content.
- **Fixed**: Render engine CSS bug for codeblocks nested inside `<pre>` tags during Light Mode view.
