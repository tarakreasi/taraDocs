---
description: Query the local Librarian for codebase knowledge
---
# 📚 `/librarian` Workflow

Gunakan workflow ini untuk mencari tahu informasi codebase ke Librarian. Librarian menyimpan semantic embeddings dari seluruh file proyek dan dapat menjawab pertanyaan terkait kode, dokumentasi, dan spesifikasi.
Ini akan menghemat waktu Anda dibandingkan harus membaca file satu per satu secara manual.t-term memory (Context):

```bash
uv run python .agent/workflows/automation/scripts/ask_librarian.py "How does the heavy processing work?" --update-context --focus "Perf**2. Automatically Saves Insight**
This writes the insight directly to `.agent/memory/active_context.md`, making it available to the next AI Agent execution automatically.
```

### Basic Query
```
/librarian where is the deploy script?
```
→ `uv run python .agent/workflows/automation/scripts/ask_librarian.py "where is the deploy script?"`

### Implementation Details
```
@librarian how does the EmbeddingService work?
```
→ `uv run python .agent/workflows/automation/scripts/ask_librarian.py "how does the EmbeddingService work?"`

### Specifications
```
/librarian what is the TaraSysDash Windows Upgrade spec?
```
→ `uv run python .agent/workflows/automation/scripts/ask_librarian.py "what is the TaraSysDash Windows Upgrade spec?"`

### Existing Patterns
```
@librarian existing patterns for alert system
```
→ `uv run python .agent/workflows/automation/scripts/ask_librarian.py "existing patterns for alert system"`

---

## When to Use

| Situation | Query Example |
|-----------|---------------|
| 🔍 Finding files | "where is [filename]?" |
| 📖 Understanding code | "how does [function] work?" |
| 📋 Checking specs | "what does [spec] say about [topic]?" |
| 🔄 Reusing patterns | "existing patterns for [feature]" |
| ✅ Verifying existence | "does [thing] exist in codebase?" |

---

## AI Behavior

When you see `@librarian` or `/librarian` in user message:

1.  **Extract the question** after the mention
2.  **Run the command** with the question
3.  **Parse the output** for relevant code snippets
4.  **Use the context** to inform your response

---

## Maintenance

### Re-index after changes
```bash
uv run python .agent/workflows/automation/scripts/index_codebase.py
```

### Check database size
```bash
ls -lh .agent/workflows/automation/memory/librarian.db
```

---

## Technical Details

- **Embedding Model**: `all-MiniLM-Lightweight Dependencies**: HuggingFace transformers + FAISS
- **No Local GPU Required**: Pure Python implementation
- **Incremental Updates**: Only re-embeds what has changed-vec
- **Indexed**: Project code, docs, workflows, specs (~300 files)
- **Excluded**: .venv, node_modules, __pycache__, memory/

---

*The Librarian never forgets. Use it wisely.*
