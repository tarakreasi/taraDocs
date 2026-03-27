---
description: Define architecture and tech stack for new projects built from scratch
---
# 🏗️ `/architect` Workflow

Use this workflow **before** `/agent-spec` when starting a **completely new project from scratch**.

## 📋 Manifest
- **Input**: `docs/product/mission.md`
- **Output**: `docs/architecture/ARCHITECTURE.md` + `.agent/rules/CODING_STANDARDS.md`
- **Next Step**: `/agent-design` (Form) OR `/agent-research` (Feasibility)

## When to Use
- Building a new project from zero (no existing codebase)
- Major rewrite of an existing system
- Technology evaluation before commitment
- Early ideation phase where tech stack is undecided

---

## Steps

### 1️⃣ Clarify Project Vision (Check Context)

**Action**: Check if a Product Mission exists.
```bash
cat docs/product/mission.md 2>/dev/null || echo "No mission file found."
```

- **If Found**: Use it as the primary source of truth.
- **If Missing**: Ask these fundamental questions:
1. **What problem does this solve?**
2. **Who is the target user?**
3. **What is the expected scale?**
4. **What is the deployment target?**

### 2️⃣ Tech Stack Decision

Evaluate and document choices for each layer:

```markdown
# Architecture: [Project Name]

## Tech Stack Decision

### Backend
- **Language**: [Go / Python / Node.js / etc.]
- **Framework**: [Gin / FastAPI / Express / etc.]
- **Rationale**: [Why this choice?]

### Database
- **Primary**: [PostgreSQL / SQLite / MongoDB / etc.]
- **Cache**: [Redis / None]
- **Rationale**: [Why this choice?]

### Frontend (if applicable)
- **Framework**: [Vue / React / Svelte / None]
- **Styling**: [TailwindCSS / Vanilla CSS / etc.]
- **Rationale**: [Why this choice?]

### Infrastructure
- **Deployment**: [Docker / Binary / Serverless]
- **CI/CD**: [GitHub Actions / None]
- **Monitoring**: [Prometheus / Custom / None]
```

### 3️⃣ Workflow Strategy Definition

Define how development will proceed. Choose one:

#### Option A: Unified Workflow (Small/Medium Projects)
- **Best for**: Monoliths, simple tools, single developer
- **Strategy**: Single `/agent-spec` covering everything
- **Structure**:
  ```
  docs/specs/
  └── feature-name/
      └── spec.md (Contains both BE and FE)
  ```

#### Option B: Split Workflow (Complex/Team Projects)
- **Best for**: Separate Backend/Frontend, scalable apps
- **Strategy**: Separate specs and sprints for BE and FE
- **Strict Rule**: **API Contract First** (OpenAPI/Swagger)
- **Structure**:
  ```
  docs/specs/
  ├── backend/
  │   └── feature-name/
  └── frontend/
      └── feature-name/
  ```

---

### 4️⃣ Architecture Diagram

Create a component diagram. Emphasis on boundaries:

```
┌──────────────┐       ┌──────────────┐
│  FRONTEND    │       │   BACKEND    │
│  (Vue/React) │◀─────▶│  (Go/Python) │
│  [View Layer]│  API  │ [Logic Layer]│
└──────────────┘ Contract └──────────────┘
                               │
                               ▼
                        ┌──────────────┐
                        │  DATABASE    │
                        │  (Storage)   │
                        └──────────────┘
```

### 5️⃣ Define Automation DNA (Coding Standards)

Define the "Rulebook" for the Automated Executor. 

**CRITICAL**: Do NOT invent standards if they exist in `agent-os`.
1. Check Agent OS Standards:
   ```bash
   ls .agent/workflows/system/profiles/default/standards/
   ```
2. **Inherit** these rules into your project-specific `CODING_STANDARDS.md`.
   - If using Vue, copy/reference `.agent/system/profiles/default/standards/frontend/vue.md`.
   - If using Python, copy/reference `.agent/system/profiles/default/standards/backend/python.md`.

Create `.agent/rules/CODING_STANDARDS.md`:

```markdown
# Coding Standards for [Project Name]

## Core Policy
- Inherited from: [.agent/workflows/system/profiles/default/standards/...]

## Project Specific Overrides
- It contains the specific tech stack, the coding style, and rules for test coverage.
```

**Critical**: This file MUST be created. It acts as the "Brain Implant" for the AI Agent (e.g., Gemini). The executor will always check this file before writing code.

### 6️⃣ Project Structure Template

Define the folder structure based on chosen strategy:

#### For Split Workflow (Recommended):
```
[project-name]/
├── backend/                # Independent Backend
│   ├── src/
│   ├── tests/
│   └── Makefile
├── frontend/               # Independent Frontend
│   ├── src/
│   ├── public/
│   └── package.json
├── docs/
│   ├── architecture/
│   └── specs/
│       ├── backend/
│       └── frontend/
└── docker-compose.yml      # Orchestration
```

---

### 7️⃣ Environment Prerequisites

Document what needs to be installed:

```markdown
## Prerequisites

### Common
- [ ] Git
- [ ] Docker (optional)

### Backend Environment
- [ ] Python 3.12+ / Go 1.21+
- [ ] Database (PostgreSQL/SQLite)
- [ ] Migration tools

### Frontend Environment
- [ ] Node.js 20+
- [ ] pnpm (recommended)
```

### 8️⃣ Save Architecture Document

Save to: `docs/architecture/ARCHITECTURE.md`

---

## Output Checklist

Before proceeding to `/agent-spec`:
- [ ] Tech stack documented with rationale
- [ ] **Workflow Strategy selected (Unified vs Split)**
- [ ] Architecture diagram created
- [ ] Folder structure defined
- [ ] Prerequisites listed

---

## Example Usage

**User Prompt**:
> "/agent-architect for a scalable E-commerce app"

**AI Actions**:
1. Clarify requirements (scale? platforms?)
2. **Select Option B (Split Workflow):**
   - Backend: Go + Postgres
   - Frontend: Nuxt.js
3. Define folder structure with clear separation
4. Save `docs/architecture/ARCHITECTURE.md`
5. Recommend next step: "Create Backend Spec first"

---

*A solid architecture separates concerns and prevents spaghetti code.*
