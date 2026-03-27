---
description: Inject extreme implementation details and verification scripts into a micro-sprint
---
# 🔍 `/sprintdetail` Workflow (AI-Executable)

Gunakan workflow ini untuk mengubah sprint draft menjadi **Instruksi Coding Detail**.

## 📋 Manifest
- **Input**: `sprintX_Y_Z.md` (Draft)
- **Output**: `sprintX_Y_Z.md` (Executable)
- **Next Step**: `Eksekusi`

---

## 🎯 Goal

## 📋 Template Instruksi Detail

Gunakan struktur ini untuk memastikan kode yang akan ditulis sudah lengkap dan benar.

```markdown
# Sprint X.Y.Z: [Title]

**Parent**: @[Sprint X.0](./sprintX_0_parent.md)
**Status**: READY FOR CODING

## 🌐 EXECUTION CONTEXT
**WORKING DIRECTORY (CWD)**: `Project Root` (e.g., `/home/username/Project/myApp`)

## 📁 FILES TO CREATE

| Path | Description |
|------|-------------|
| `path/to/file1.ext` | [Description] |
| `tests/test_file1.py` | Unit Test |

## 🔗 WIRING / INTEGRATION
**Target File**: `path/to/existing/file.go`
**Action**: [Add Route / Inject Dependency]
**Snippet**:
```go
// Add this line
mux.HandleFunc(...)
```

```markdown
# Sprint X.Y.Z: [Title]

**Parent**: @[Sprint X.0](./sprintX_0_parent.md)
**Objective**: [One sentence]
**Executor**: AI Agent (e.g., Gemini)
**Status**: PLANNING

## 🌐 EXECUTION CONTEXT
**WORKING DIRECTORY (CWD)**: `Project Root` (e.g., `/home/username/Project/myApp`)
*CRITICAL*: All file paths in this document are STRICTLY RELATIVE to the Project Root.

## 📁 FILES TO CREATE

| Path | Description |
|------|-------------|
| `path/to/file1.ext` | [Description] |
| `verify_sprint.py` | Verification Script |

## 🔗 WIRING / INTEGRATION
**Target File**: `path/to/existing/file.go`
**Action**: [Add Route / Inject Dependency]
**Snippet**:
```go
// Add this line
mux.HandleFunc(...)
```

## 📋 TASKS

### Task 1: [Action Name]

**File**: `path/to/file1.ext`

**Content**:
```[lang]
[FULL CODE HERE - NO PLACEHOLDERS]
```

### Task 2: [Another Action]

**File**: `path/to/file2.ext`

**Content**:
```[lang]
[FULL CODE HERE]
```

### Task 3: Verification Script

**File**: `verify_sprint.py`

**Content**:
```python
import os
import sys

# Anchor to the Project Root to prevent AI hallucinations
PROJECT_ROOT = os.getcwd()
path = os.path.join(PROJECT_ROOT, "path/to/file")

# ... verification logic ...
```

---

## 📐 Rules of Engagement

### 1. Determinism > Interactivity
- **NEVER** use interactive commands like `npm init`, `composer create-project` (without --no-interaction), or `read input`.
- **PREFER** writing specific config files manually (e.g., writing `package.json` with `cat`) over generating them if generators are flaky.
- **ALWAYS** set `set -e` in shell scripts.

### 2. Code Completeness
- The code inside the `Content` block is what gets written to disk.
- **DO NOT** use `... (rest of code)` or `# TODO`.
- **DO NOT** assume the file exists; always provide the full content or use specific `sed` commands if modifying (but overwriting is safer for automation).

### 3. Verification is Mandatory
- Every sprint MUST produce a `verify_*.py` script.
- This script is the "Test" in TDD.
- It must be self-contained (import only standard libs if possible).

### 4. Schema/Contract Alignment (CRITICAL for AI)
- If a task is for **AI Generation** (not a static Code Block), you MUST provide a `**Context**` or `**Schema**` section.
- Explicitly list **Database Fields** (names, types).
- Explicitly list **API Routes** and **JSON keys**.
- **WHY**: This prevents AI from "hallucinating" extra fields (e.g., adding `user_id` when the database only has `title`).

### 5. Working Directory Precision
- The AI Executor must know exactly where it is.
- ALWAYS specify that execution happens at the Project Root.
- NEVER let the AI assume it's running from `.agent` or `docs/dev/sprints`. Verify scripts MUST use `os.getcwd()` resolving to root.

---

## 🔄 Workflow Steps to Execute /agent-detail

### 1. Context Loading (DNA Check)
Before detailing, you MUST know the rules.
```bash
# 1. Read Automation Rules
cat .agent/rules/CODING_STANDARDS.md

# 2. Read Technical Architecture
cat docs/architecture/ARCHITECTURE.md
```
**Why?**: Your verification scripts must match the project's test runner (e.g., `pytest` vs `unittest`) and language standards.

### 2. Review the Micro-Sprint
Read the high-level plan from `/agent-microsprint` (`docs/dev/sprints/sprintX_Y_Z.md`).

### 3. Technical Checklist (Verify these before writing MD)
- [ ] **Base Classes**: Are `Controller`, `Model`, or `BaseTest` imports included?
- [ ] **Schema**: Are all DB fields/JSON keys explicitly listed?
- [ ] **Standards Compliance**: Does the code follow `CODING_STANDARDS.md`?
- [ ] **CWD Verification**: Does the `verify_*.py` script explicitly anchor to root via `os.getcwd()`?

### 4. Expand Tasks
    - **MANDATORY**: Include a `Local Guardrails` sub-section in the instruction.
    - **STRICT NAMING**: Mention forbidden field names (e.g., "Use title, NOT description").
    - **TECH STACK**: Mention specific implementation rules (e.g., "Use computed for formatting").
4. **Update MD File**: Overwrite the sprint file with this detailed content.
5. **Mark as PLANNING**: Ensure status is PLANNING so the Orchestrator picks it up.
```markdown
### Task 1: [Action]
**Instruction**: 
[General goal]
- 🚫 **FORBIDDEN**: [List common hallucinations]
- ✅ **MANDATORY**: [List specific column names from contract]
```

---

## 🔗 Integration

| Input | Process | Output |
|-------|---------|--------|
| `ARCHITECTURE.md` | **Context** | DNA Compliance |
| `sprintX_Y_Z.md` (Draft) | **`/sprintdetail`** | `sprintX_Y_Z.md` (Detailed) |
| `sprintX_Y_Z.md` (Detailed) | **Manual/AI Coding** | Source Code |

---

*If the AI has to guess, the sprint reflects a failure in planning. Be explicit. Schema context is the key to precision.*
