# Coding Standards for taraLes

## Core Philosophy
**"Complexity is the Enemy"**
This project serves non-technical users on local hardware. Code must be verifiable, robust, and easy to read.

## 1. Backend (Go)
-   **Standard Lib First**: Use `net/http` features before reaching for a framework.
-   **Error Handling**: Wrap errors with context. `fmt.Errorf("doing x: %w", err)`.
-   **SQL**:
    -   Write raw SQL in `internal/database/queries.go`.
    -   Use prepared statements always (Security).
-   **Structure**:
    -   `Handler` -> `Service` -> `Repo` pattern is allowed but keep it shallow.
    -   Don't over-abstract.

## 2. Frontend (Vanilla)
-   **No Build Step**: If you change a JS file, you must be able to refresh the browser to see it (in dev mode).
-   **CSS**:
    -   Use `var(--token-name)` for colors.
    -   Avoid inline styles.
-   **JS**:
    -   Use ES6 Modules (`<script type="module">`).
    -   Use `fetch` for API calls.

## 3. Database (SQLite)
-   **Naming**: Snake_case (`user_id`, `created_at`).
-   **WAL Mode**: Always enable Write-Ahead Logging for concurrency.

## 4. Automation Rules
-   **Atomic Commits**: One feature per commit.
-   **Verification**: Every feature must include a verification step.

## 5. Configuration & Quality
-   **12-Factor App**: No hardcoded ports or paths. Use `os.Getenv` with defaults.
    -   Example: `port := getEnv("PORT", "8080")`
    -   Prefer `subprocess.run(["go", "build", ...])` over `grep` source code.

## 6. Test Hygiene (CRITICAL)
-   **Isolation**: Tests must NOT depend on the state left by previous runs.
-   **Cleanup**: Always clean up artifacts (DB files, temp dirs) in a `defer` or `t.Cleanup()` block.
    -   Example: `defer os.Remove(dbPath)` immediately after defining the path.
-   **Idempotency**: A test running twice should pass twice.

## 7. Integration & Wiring
-   **Read Before Write**: Before modifying an existing file (e.g., `server.go`), READ it to understand the current state.
-   **Duplicate Check**: Ensure you are not adding a duplicate route or function.
-   **Context**: When planning a modification, explicitly state *where* the code goes (e.g., "After line 50").
