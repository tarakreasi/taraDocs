# 📚 Semi-Agentic Toolkit - Complete Documentation Index

> **"Collaborative Intelligence"** — A system where Specialized Agents (The Brain) drive Deterministic Automation (The Body).

---

## 🗺️ Quick Navigation

| Goal | Document | Time |
|------|----------|------|
| **Start a Project** | [IDEA_TO_CODE_GUIDE.md](../docs/dev/IDEA_TO_CODE_GUIDE.md) | 5 min |
| **Understand the Flow** | [Workflows README](workflows/README.md) | 10 min |
| **Fix Bugs** | [Incident Reporting](workflows/agent-debug.md) | 5 min |
| **Clean Code** | [Optimization](workflows/agent-optimizer.md) | 5 min |
| **Manual Setup** | [INSTALL.md](INSTALL.md) | 10 min |

---

## 📦 The Agentic Ecosystem

This toolkit is organized by **Roles**. Each agent has a specific job in the software lifecycle.

### 🧠 The Strategists (Definition Phase)
| Role | Command | Purpose |
|------|---------|---------|
| **The Researcher** | `/agent-research` | **Context Loading**. Ingests docs, PDFs, and competitors. |
| **The Planner** | `/product-planner` | **Scope Definition**. Translates "ideas" into "missions". |
| **The Architect** | `/agent-architect` | **Tech Strategy**. Defines the stack and coding standards. |
| **The Designer** | `/agent-design` | **Visuals**. Mocks UI, defines flows, generates components. |
| **The Reviver** | `/agent-code` | **Archaeology**. Audits and revives legacy/stalled projects. |

### 🏗️ The Builders (Design & Plan Phase)
| Role | Command | Purpose |
|------|---------|---------|
| **The Writer** | `/agent-spec` | **Blueprint**. Writes detailed specifications. |
| **The Manager** | `/agent-init` | **Roadmap**. Creates the Master Sprint Plan. |
| **The Scrum Master**| `/agent-scrum-master`| **Guardian**. Breaks down tasks & enforces contracts. |

### 🤖 The Executors (Build Phase)
| Role | Command | Purpose |
|------|---------|---------|
| **The Detailer** | `/agent-detail` | **Translation**. Converts tasks into machine-executable plans. |
| **The Builder** | `Supervisor` | **Action**. The automated loop that writes code. |
| **The QA** | `/agent-verifier` | **Gatekeeper**. Runs linting, tests, and arch checks. |

### 🚑 The Support Team (Maintenance Phase)
| Role | Command | Purpose |
|------|---------|---------|
| **The Detective** | `/agent-debug` | **RCA**. Reproduce -> Analyze -> Fix. |
| **The Optimizer** | `/agent-optimizer` | **Refactor**. Pays down technical debt. |

---

## 📂 System Structure

```bash
.agent/
├── 📖 Documentation
│   ├── README.md              # You are here
│   ├── PM.md                  # The deterministic philosophy
│   └── docs/dev/IDEA_TO_CODE_GUIDE.md # The User Manual
│
├── ⚙️ Workflows (The Brains)
│   └── workflows/
│       ├── idea-to-code.md    # � The Master Pipeline
│       ├── agent-*.md         # Individual Agent Instructions
│       └── automation/        # Python Scripts (The Body)
│
└── 🧠 Memory & Standards
    ├── rules/
    │   └── CODING_STANDARDS.md # The "Law" for the project
    └── system/
        └── standards/         # Agent OS Global Patterns
```

---

## � Getting Started

### 1. The "Greenfield" Path (New Project)
Start here if you have a fresh idea.
```bash
# 1. Define the Product
/product-planner

# 2. Design the System
/agent-architect

# 3. Follow the Guide
cat docs/dev/IDEA_TO_CODE_GUIDE.md
```

### 2. The "Brownfield" Path (Legacy Project)
Start here if you are rescuing an old project.
```bash
# 1. Audit the Code
/agent-code

# 2. Initialize Sprint from Recovery Plan
/agent-init
```

---

## ⚡ Quick Reference

### Core Commands
```bash
# Check Agent Status or execute a sprint script
python .agent/workflows/automation/scripts/verify_sprint.py

# Query Context (Librarian)
uv run python .agent/workflows/automation/scripts/ask_librarian.py "How does auth work?"
```

### Support Commands
```bash
# "It's broken!"
/agent-debug

# "It's messy!"
/agent-optimizer
```

---

## 📞 Troubleshooting

- **Check Sprints**: Run `cat docs/dev/sprints/sprintX_Y_Z.md` for pending checks `[ ]`.
- **AI Hallucinating?**: Check `.agent/memory/active_context.md` to refresh context.
- **Wrong Stack?**: Re-run `/agent-architect` to update `CODING_STANDARDS.md`.

---

**Last Updated**: 2026-02-20
**Version**: 5.1 (Cloud Agentic Evolution)
**Status**: ✅ Operational

