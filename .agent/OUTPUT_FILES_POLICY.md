
> ### ⚠️ IMPORTANT: Output Files Location
> 
> **DO NOT** create output files inside `.agent/` directory!
> 
> - ✅ `.agent/` = **Workflow definitions** (how to do things)
> - ❌ `.agent/` ≠ Output/result files
> 
> **Correct locations:**
> - Sprint output → `docs/dev/sprints/`
> - Spec documents → `docs/specs/`
> - Architecture → `docs/architecture/`
> - API contracts → `docs/specs/API_CONTRACT.md`
> 
> **Wrong:**
> ```bash
> ❌ .agent/workflows/sprint7_1_fix.md
> ❌ .agent/my_spec.md
> ```
> 
> **Correct:**
> ```bash
> ✅ docs/dev/sprints/sprint7_1_fix.md
> ✅ docs/specs/quiz_feature/spec.md
> ```
