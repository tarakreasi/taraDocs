---
description: Phase 2 - Backend Implementation Based on Approved Contract
---
# ⚙️ `/sprint-backend` Workflow

**Phase 2 of 2**: Implement backend sesuai API contract yang sudah di-approve.
**Prerequisites**: Frontend + API Contract sudah selesai dan di-approve user.

## 📋 Manifest
- **Input**: 
  - `docs/specs/API_CONTRACT.md` (approved)
  - Frontend files dari Phase 1
- **Output**: 
  - Backend handlers, database, routes
  - Frontend updated (ganti mock data → real API)
- **Next Step**: Integration testing

---

## Steps

### 0. Verify Prerequisites
Sebelum mulai, pastikan:
- [ ] `docs/specs/API_CONTRACT.md` exists and is approved
- [ ] Frontend dengan mock data sudah jalan
- [ ] User sudah approve API contract

**If not → STOP. Kembali ke `/sprint-frontend`**

---

### 1. Read API Contract
Baca `docs/specs/API_CONTRACT.md` dan catat:
- Semua endpoints (method + path)
- Request/response format untuk setiap endpoint
- Database schema yang dibutuhkan
- Validation rules

---

### 2. Create Database Schema & Migration

Buat file `internal/database/migration_[feature].go`:

```go
func MigrateFeatureName(db *sql.DB) error {
    schema := `
    CREATE TABLE IF NOT EXISTS quizzes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        subject TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS questions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        quiz_id INTEGER NOT NULL,
        text TEXT NOT NULL,
        option_a TEXT NOT NULL,
        option_b TEXT NOT NULL,
        option_c TEXT NOT NULL,
        option_d TEXT NOT NULL,
        correct_option TEXT NOT NULL,
        FOREIGN KEY(quiz_id) REFERENCES quizzes(id)
    );
    `
    _, err := db.Exec(schema)
    return err
}
```

**Call migration** di `internal/database/schema.go`:
```go
func RunMigrations(db *sql.DB) error {
    // ... existing migrations
    if err := MigrateFeatureName(db); err != nil {
        return err
    }
    return nil
}
```

---

### 3. Create Database Functions

Buat `internal/database/[feature].go` dengan functions sesuai contract:

```go
package database

import "database/sql"

// GetQuizzes returns all quizzes (for GET /api/quizzes)
func GetQuizzes(db *sql.DB) ([]Quiz, error) {
    rows, err := db.Query("SELECT id, title, subject, created_at FROM quizzes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var quizzes []Quiz
    for rows.Next() {
        var q Quiz
        if err := rows.Scan(&q.ID, &q.Title, &q.Subject, &q.CreatedAt); err != nil {
            continue
        }
        quizzes = append(quizzes, q)
    }
    return quizzes, nil
}

// GetQuiz returns single quiz with questions (for GET /api/quiz/{id})
func GetQuiz(db *sql.DB, id int) (QuizDetail, error) {
    // ... implementation
}

// SaveQuizSubmission saves answers and calculates score (for POST /api/quiz/submit)
func SaveQuizSubmission(db *sql.DB, quizID int, answers map[int]string, studentName string) (int, error) {
    // ... implementation
}
```

**Rule**: 1 function = 1 endpoint action

---

### 4. Create Handlers

Buat `internal/server/handler_[feature].go` dengan handlers sesuai contract:

```go
package server

import (
    "encoding/json"
    "net/http"
    "github.com/twantoro/taraLes/internal/database"
)

// GET /api/quizzes
func (s *Server) handleGetQuizzes(w http.ResponseWriter, r *http.Request) {
    quizzes, err := database.GetQuizzes(s.db)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "data": quizzes,
    })
}

// GET /api/quiz/{id}
func (s *Server) handleGetQuiz(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    // ... implementation
}

// POST /api/quiz/submit
func (s *Server) handleSubmitQuiz(w http.ResponseWriter, r *http.Request) {
    // ... implementation
}
```

**Rule**: 
- Handler harus **persis match** dengan API contract (method, path, request, response)
- Gunakan struct yang sama dengan mock data di frontend

---

### 5. Register Routes

Update `internal/server/server.go`:

```go
// Quiz Routes
mux.Handle("GET /api/quizzes", AuthMiddleware(http.HandlerFunc(s.handleGetQuizzes)))
mux.Handle("GET /api/quiz/{id}", AuthMiddleware(http.HandlerFunc(s.handleGetQuiz)))
mux.Handle("POST /api/quiz/submit", AuthMiddleware(http.HandlerFunc(s.handleSubmitQuiz)))
```

**Rule**: Order harus sesuai dengan API contract document

---

### 6. Update Frontend (Remove Mock Data)

Ganti hardcoded mock data dengan real API calls:

**Before (Mock)**:
```javascript
const mockQuizzes = [
  { id: 1, title: "Math Quiz", subject: "Math", total_questions: 10 }
];
renderQuizzes(mockQuizzes);
```

**After (Real API)**:
```javascript
async function loadQuizzes() {
    const res = await fetch('/api/quizzes');
    const data = await res.json();
    renderQuizzes(data.data); // Same structure as mock!
}
loadQuizzes();
```

**Rule**: Frontend code harus **minimal changes** karena struktur data sama dengan mock

---

### 7. Testing Checklist

Sebelum mark selesai:

- [ ] Compile berhasil (`go build`)
- [ ] Semua endpoint accessible (test pakai browser/curl)
- [ ] Response format match dengan API contract
- [ ] Database queries tidak error
- [ ] Frontend bisa load data dari backend (bukan mock lagi)
- [ ] Error handling works (test dengan bad input)

---

### 8. Integration Test

Buat simple test script:

```bash
#!/bin/bash
# Test all endpoints

# 1. List quizzes
curl http://localhost:8000/api/quizzes

# 2. Get quiz detail
curl http://localhost:8000/api/quiz/1

# 3. Submit quiz
curl -X POST http://localhost:8000/api/quiz/submit \
  -H "Content-Type: application/json" \
  -d '{"quiz_id":1,"answers":{"101":"b"},"student_name":"Test"}'
```

Run test, pastikan semua return 200 dan data sesuai contract.

---

## 9. Deliverables Checklist

Sebelum finalize:

- [ ] Database migration created
- [ ] Database functions created (1:1 dengan endpoints)
- [ ] Handlers created (match API contract)
- [ ] Routes registered
- [ ] Frontend updated (mock → real API)
- [ ] All endpoints tested manually
- [ ] No breaking changes to API contract
- [ ] Error responses follow contract format

---

## 10. Documentation Update

Update `CHANGELOG.md`:

```markdown
## [Version X.Y] - 2026-01-30

### Added
- **[Feature Name]** backend implementation
  - Endpoints: GET /api/quizzes, GET /api/quiz/{id}, POST /api/quiz/submit
  - Database: tables `quizzes`, `questions`
  - Frontend integrated with real API (mock data removed)

### Changed
- Frontend now uses real backend API (was mock data)
```

---

## ⚠️ Important Rules

### Must Follow API Contract
- **Endpoint paths** harus exact match dengan contract
- **Request/response format** harus exact match
- **HTTP methods** harus exact match
- **Status codes** harus sesuai contract

### No Frontend Changes (Except API Calls)
- UI/UX tidak boleh berubah dari Phase 1
- Hanya ganti `mockData` → `fetch('/api/...')`
- Kalau perlu ubah UI → kembali ke Phase 1 dan revisi contract

### Case Sensitivity
- Database column names: **lowercase with underscore** (created_at)
- JSON fields: **camelCase** or **snake_case** (sesuai contract)
- SQL comparisons: **use COLLATE NOCASE** atau Go `strings.EqualFold()`

---

## Next Steps

Setelah backend selesai:
1. Commit changes
2. Create PR
3. Deploy to staging
4. User acceptance testing

---

## Related Workflows

| Workflow | Purpose |
|----------|---------|
| `/sprint-frontend` | Phase 1 (prerequisite) |
| `/verifier` | Verify implementation quality |
| `/debug` | If bugs found during testing |
