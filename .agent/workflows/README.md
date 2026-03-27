# 📚 Workflow Guide - TaraLes Project

Panduan memilih workflow yang tepat untuk development task.

---

## 🎯 Quick Decision Tree

```
Mau develop apa?
│
├─ Feature baru yang kompleks?
│  └─ Pakai: /sprint-frontend → /sprint-backend
│
├─ Feature kecil/bug fix?
│  └─ Pakai: /sprint1 → /sprintmicro
│
├─ Debug masalah?
│  └─ Pakai: /debug
│
├─ Refactor/optimize code?
│  └─ Pakai: /optimizer
│
└─ Project baru dari 0?
   └─ Pakai: /architect → /idea-to-code
```

---

## 🚀 Primary Workflows (Development)

### 1. `/sprint-frontend` + `/sprint-backend` (Frontend-First) ⭐ NEW

**Kapan pakai:**
- Feature baru yang **kompleks** (multi-page, banyak form, API banyak)
- Team/QC **perfeksionis** yang mau review UI dulu sebelum backend
- Mau **test UX early** sebelum invest waktu ke backend
- Mau **API contract** yang jelas dari awal (no revision)

**Flow:**
```
1. /sprint-frontend
   → Buat UI dengan mock data (hardcoded JSON)
   → Define API contract (request/response format)
   → Define database schema (belum implement)
   → Request user approval
   
2. User review & approve
   
3. /sprint-backend
   → Implement database sesuai schema
   → Implement handlers sesuai API contract
   → Ganti mock data → real API di frontend
   → Integration test
```

**Contoh use case:**
- Dashboard analytics baru (banyak charts, filters)
- Fitur quiz dengan multiple steps (list → detail → submit → result)
- Form wizard yang kompleks

**Output:**
- `web/static/[feature].html` (dengan mock data)
- `docs/specs/API_CONTRACT.md` (approved)
- `internal/server/handler_[feature].go`
- `internal/database/[feature].go`

---

### 2. `/sprint1` → `/sprintmicro` (Traditional)

**Kapan pakai:**
- Feature **kecil/sedang** yang straightforward
- Sudah tahu persis struktur data & API yang dibutuhkan
- Bug fix yang perlu refactor besar
- Maintenance task (update dependencies, dll)

**Flow:**
```
1. /sprint1
   → Read spec
   → Breakdown jadi micro-sprints (1.1, 1.2, 1.3)
   → Define test strategy
   
2. /sprintmicro
   → Generate individual sprint files
   → Each micro-sprint = 1-3 files max
   
3. /sprintdetail (optional)
   → Add verification scripts
```

**Contoh use case:**
- Add "Delete Quiz" button (simple CRUD)
- Fix grading calculation bug
- Add export to CSV feature
- Update user profile form

**Output:**
- `docs/dev/sprints/sprintX_0_parent.md`
- `docs/dev/sprints/sprintX_1_microtask1.md`
- `docs/dev/sprints/sprintX_2_microtask2.md`

---

## 🛠️ Utility Workflows

### `/debug` - Root Cause Analysis
**Kapan pakai:**
- Ada bug yang **sulit ditemukan**
- Error yang **inconsistent** (kadang muncul, kadang nggak)
- Performance issue

**Output:** Diagnosis + fix recommendations

---

### `/optimizer` - Code Quality Improvement
**Kapan pakai:**
- Code review menemukan **technical debt**
- Perlu **refactor** untuk readability
- Performance optimization

**Output:** Refactored code with better structure

---

### `/verifier` - Implementation Verification
**Kapan pakai:**
- Setelah selesai development, mau **QA check**
- Pastikan **semua requirement** terpenuhi
- Generate test report

**Output:** Verification report + checklist

---

## 📋 Planning Workflows

### `/spec-initializer` → `/spec-writer`
**Kapan pakai:**
- Ide masih **mentah** (belum jelas requirementnya)
- Perlu gather requirements dari stakeholder

**Flow:**
```
1. /spec-initializer → Save raw idea
2. /spec-shaper → Gather detailed requirements via Q&A
3. /spec-writer → Create formal spec document
4. /spec-verifier → Validate spec completeness
```

**Output:** `docs/specs/YYYY-MM-DD-feature/spec.md`

---

### `/architect` - Define Tech Stack
**Kapan pakai:**
- **Project baru** dari 0
- Mau define architecture, database, folder structure

**Output:** `docs/architecture/ARCHITECTURE.md`

---

## 🔀 Combined Workflows (End-to-End)

### `/idea-to-code` - Full Pipeline
**Kapan pakai:**
- Ada **ide baru**, langsung jadi working code
- Master orchestrator untuk semua workflow

**Flow:**
```
Idea → Spec → Planning → Implementation → Verification → Done
```

---

## 📊 Comparison Table

| Workflow | Complexity | Time | Best For |
|----------|-----------|------|----------|
| `/sprint-frontend` + `/sprint-backend` | High | 2-4 hours | Complex features, Frontend-first |
| `/sprint1` + `/sprintmicro` | Medium | 1-2 hours | Standard features, Bug fixes |
| `/debug` | Low | 30-60 min | Bug hunting |
| `/optimizer` | Medium | 1-2 hours | Refactoring |
| `/verifier` | Low | 15-30 min | QA check |
| `/idea-to-code` | Very High | 4-8 hours | New project/major feature |

---

## 💡 Recommended Flow for TaraLes

Untuk project **TaraLes** yang sudah mature, gunakan:

### Scenario 1: Feature Baru (e.g., "Attendance System")
```bash
1. /spec-initializer       # Save idea
2. /spec-writer            # Create spec
3. /sprint-frontend        # UI mockup + API contract
   # → User approve
4. /sprint-backend         # Implement backend
5. /verifier               # QA check
6. Commit & PR
```

### Scenario 2: Bug Fix (e.g., "Score calculation wrong")
```bash
1. /debug                  # Find root cause
2. Fix code manually       # Small change
3. Test                    # Verify fix
4. Commit
```

### Scenario 3: Refactor (e.g., "Improve quiz engine")
```bash
1. /optimizer              # Get refactor suggestions
2. Apply changes           # Refactor code
3. /verifier               # Ensure nothing broken
4. Commit
```

---

## ⚠️ Common Pitfalls

### ❌ **Jangan** pakai `/sprint-frontend` untuk bug fix
- Terlalu overkill
- Pakai `/debug` atau `/sprint1` saja

### ❌ **Jangan** skip API contract di `/sprint-frontend`
- Nanti backend akan banyak revision
- Contract harus **approved** sebelum lanjut backend

### ❌ **Jangan** coding backend di fase frontend
- Sesuai workflow, frontend STOP dulu untuk approval
- Backend di fase terpisah

---

## 📝 Best Practices

### 1. Selalu buat spec dulu untuk feature baru
```bash
# Good
/spec-writer → /sprint-frontend → /sprint-backend

# Bad
Langsung coding tanpa spec → banyak revision
```

### 2. Gunakan `/verifier` sebelum commit besar
```bash
# Before commit
/verifier

# Review checklist
# Fix issues
# Then commit
```

### 3. Document API contract dengan detail
```bash
# In API_CONTRACT.md
- Exact endpoint path: GET /api/quiz/{id}
- Exact request body: { "quiz_id": int, "answers": map }
- Exact response: { "score": int }
- Error cases: 404 if quiz not found, 400 if invalid input
```

---

## 🎓 Learning Path

**Untuk yang baru belajar coding (seperti Pak):**

1. Mulai dari bug fix kecil → pakai `/debug`
2. Lalu feature sederhana → pakai `/sprint1`
3. Kalau sudah comfortable, coba `/sprint-frontend` + `/sprint-backend`
4. Terakhir, orchestrate full pipeline dengan `/idea-to-code`

**Tips:**
- Jangan overwhelmed, pakai workflow sesuai kebutuhan
- Kalau bingung, default ke `/sprint1`
- Save workflow output di `docs/dev/` untuk reference

---

## 📞 Quick Reference

| I want to... | Use this workflow |
|--------------|-------------------|
| Build complex feature (UI-first) | `/sprint-frontend` → `/sprint-backend` |
| Add simple feature | `/sprint1` → `/sprintmicro` |
| Fix a bug | `/debug` |
| Improve code quality | `/optimizer` |
| Verify implementation | `/verifier` |
| Plan from idea | `/spec-initializer` → `/spec-writer` |
| Full automation | `/idea-to-code` |

---

**Last Updated:** 2026-01-30  
**Maintainer:** AI Assistant  
**Project:** TaraLes Learning Management System
