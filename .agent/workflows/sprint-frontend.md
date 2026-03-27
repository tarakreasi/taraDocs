---
description: Phase 1 - Frontend Mockup & API Contract Design (No Backend Code)
---
# 🎨 `/sprint-frontend` Workflow

**Phase 1 of 2**: Design frontend UI dengan mock data + definisikan API contract.
**User MUST approve** sebelum lanjut ke `/sprint-backend`.

## 📋 Manifest
- **Input**: `docs/specs/.../spec.md`
- **Output**: 
  - Frontend HTML/CSS/JS dengan mock data
  - `docs/specs/API_CONTRACT.md`
- **Next Step**: User review → `/sprint-backend`

---

## Steps

### 1. Read Spec & Understand Requirements
- Baca `docs/specs/YYYY-MM-DD-feature/spec.md`
- Pahami user flow dan data yang dibutuhkan
- Catat semua actions yang user bisa lakukan

### 2. Design UI Components (Frontend Only)
Buat file HTML/CSS/JS di `web/static/` dengan:
- **Hardcoded mock data** (JSON array dalam `<script>`)
- UI yang fully functional (bisa diklik, ada animasi, dll)
- **Tidak ada koneksi ke backend** (semua data static)

**Example Mock Data**:
```javascript
// Hardcode di HTML, bukan fetch dari API
const mockQuizzes = [
  { id: 1, title: "Math Quiz", subject: "Math", total_questions: 10 },
  { id: 2, title: "English Quiz", subject: "English", total_questions: 5 }
];
```

**Deliverables**:
- `web/static/quiz_list.html` (dengan mock data)
- `web/static/quiz_detail.html` (dengan mock data)
- `web/static/css/quiz.css`
- `web/static/js/quiz_mock.js` (helper untuk render mock)

### 3. Define API Contract
Buat `docs/specs/API_CONTRACT.md` yang mendefinisikan:

```markdown
# API Contract: [Feature Name]

## Endpoints

### GET /api/quizzes
**Description**: List all quizzes
**Request**: None
**Response**:
\`\`\`json
{
  "data": [
    {
      "id": 1,
      "title": "Math Quiz",
      "subject": "Math",
      "total_questions": 10,
      "created_at": "2026-01-30T12:00:00Z"
    }
  ]
}
\`\`\`

### GET /api/quiz/{id}
**Description**: Get quiz details
**Request**: None
**Response**:
\`\`\`json
{
  "data": {
    "id": 1,
    "title": "Math Quiz",
    "questions": [
      {
        "id": 101,
        "text": "What is 2+2?",
        "options": [
          {"label": "a", "text": "3"},
          {"label": "b", "text": "4"},
          {"label": "c", "text": "5"}
        ],
        "correct_option": "b"
      }
    ]
  }
}
\`\`\`

### POST /api/quiz/submit
**Request**:
\`\`\`json
{
  "quiz_id": 1,
  "answers": {
    "101": "b",
    "102": "a"
  },
  "student_name": "John Doe"
}
\`\`\`
**Response**:
\`\`\`json
{
  "score": 100
}
\`\`\`

## Database Schema (for Backend Phase)
\`\`\`sql
CREATE TABLE quizzes (
  id INTEGER PRIMARY KEY,
  title TEXT NOT NULL,
  subject TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE questions (
  id INTEGER PRIMARY KEY,
  quiz_id INTEGER NOT NULL,
  text TEXT NOT NULL,
  option_a TEXT NOT NULL,
  option_b TEXT NOT NULL,
  option_c TEXT NOT NULL,
  option_d TEXT NOT NULL,
  correct_option TEXT NOT NULL,
  FOREIGN KEY(quiz_id) REFERENCES quizzes(id)
);
\`\`\`
```

### 4. Create User Flow Documentation
Tambahkan di `API_CONTRACT.md`:

```markdown
## User Flows

### Flow 1: Student Takes Quiz
1. User opens `/learn` → sees list of quizzes (mock data)
2. User clicks quiz → opens `/learn/quiz/1` (mock data)
3. User answers questions → clicks Submit
4. Mock: Alert showing "Score: 80%" (hardcoded calculation)
5. Real (Phase 2): POST to `/api/quiz/submit` → get real score

### Flow 2: Teacher Creates Quiz
1. User opens `/dashboard/cms` → sees quiz list (mock)
2. User clicks "Add Quiz" → form appears
3. User fills form → clicks Save
4. Mock: Add to local array, show success message
5. Real (Phase 2): POST to `/api/quiz` → save to database
```

---

## 5. Deliverables Checklist

Before requesting user review:

- [ ] All HTML pages created with mock data
- [ ] UI fully functional (buttons work, forms validate)
- [ ] CSS styled (looks good, responsive)
- [ ] `API_CONTRACT.md` document complete
- [ ] Database schema defined
- [ ] All request/response formats documented
- [ ] User flows documented

---

## 6. Request User Review

Call `notify_user` with:
- PathsToReview: All HTML files + `API_CONTRACT.md`
- Message: 
  > "Phase 1 selesai! Saya sudah buat:
  > 1. Frontend UI dengan mock data (bisa dicoba di browser)
  > 2. API Contract yang mendefinisikan semua endpoint
  > 
  > Silakan review:
  > - Apakah UI sudah sesuai kebutuhan?
  > - Apakah API contract sudah lengkap?
  > 
  > Setelah approve, saya lanjut ke Phase 2 (Backend Implementation)."

**STOP HERE**. Wait for user approval.

---

## ⚠️ PENTING: Jangan Coding Backend!

Pada fase ini **DILARANG**:
- ❌ Buat handler Go (handler_*.go)
- ❌ Buat database migration
- ❌ Register routes di server.go
- ❌ Tulis SQL queries

**HANYA**:
- ✅ HTML/CSS/JS dengan mock data
- ✅ Dokumentasi API contract
- ✅ Definisi schema (teks, bukan SQL file)

---

## Next Workflow

Setelah user approve → `/sprint-backend`
