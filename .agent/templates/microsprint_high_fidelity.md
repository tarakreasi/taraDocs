# Microsprint High-Fidelity Template V2.1
**Purpose**: Atomic Sprint for AI Code Generation (e.g., Gemini) with Zero Hallucination.

---

# Sprint X.Y.Z: [Component/Feature Name]
**Parent**: @[Sprint X.Y](./link_to_parent.md)
**Objective**: Create [Name] strictly following the blueprint.
**Status**: PLANNING
**Type**: IMPLEMENTATION
**Executor**: AI Agent (e.g., Gemini - Atomic Mode)

## 🌐 EXECUTION CONTEXT
**WORKING DIRECTORY (CWD)**: `Project Root` (e.g., `/home/username/Project/myApp`)
*CRITICAL*: All file paths in this document are STRICTLY RELATIVE to the Project Root. Do NOT assume execution from a sub-directory like `.agent` or `docs`.

## 📁 FILES TO CREATE
| Path | Description |
|------|-------------|
| `path/to/Component.vue` | The Implementation |
| `path/to/verify.py` | Verification Script |

## 📜 CONTRACT & BLUEPRINT (PREAMBLE)
*(This section defines strict rules for the AI Code Generator. DO NOT REMOVE.)*

### 1. Architecture
**File**: `path/to/Component.vue`
**Type**: Vue 3 / Go / Laravel / Python
**Dependencies**:
```typescript
// Define interfaces and imports explicitly
import { useStore } from '@/stores/myStore';

interface DataShape {
  id: number;
  name: string;
}
```

### 2. Schema Reference (For Backend/DB Tasks)
*(Include this if the task involves Database or API payload)*
```php
// Expected DB Schema (From Migration - Reference ONLY)
$table->string('category'); // NOT category_id
$table->decimal('amount', 15, 2);
```

**API Payload Example** (If applicable):
```json
{
  "title": "Lunch",
  "amount": 50000,
  "category": "Food"
}
```

### 3. API Response Format (For Frontend Consuming API)
*(Include this if the task consumes Backend API)*
**Endpoint**: `GET /api/v1/transactions`
**Response Envelope**:
```json
{
  "data": [
    { "id": 1, "title": "...", "amount": 100 }
  ]
}
```
**Unwrap Logic**: `const items = response.data.data` (Outer `.data` is HTTP, Inner `.data` is Laravel envelope).

### 4. Visual Specification (Tailwind)
*(Required for Frontend Tasks)*
- **Container**: `class names here`
- **Typography**: `class names here`
- **Colors**: Use exact hex codes or theme tokens (e.g. `text-primary-500`)

### 5. Implementation Constraints
- **CRITICAL**: This component MUST be implemented as a SINGLE FILE.
- DO NOT create child components unless explicitly defined in Dependencies.
- Write all HTML inline within `<template>` block (Frontend).
- Use exact variable names, types, and logic defined in Contract.

### 6. Logic Requirements
- Explicit rule 1
- Explicit rule 2 (e.g. "Use Intl.NumberFormat, not external libs")

---

## 📋 TASKS

### Task 1: Generate Implementation
**File**: `path/to/Component.vue`

**Instruction**: 
Generate the code strictly following the **Architecture**, **Schema**, and **Visual Specification** above.
- No hallucinations.
- Use explicit imports defined in Contract.
- If API response format is defined, implement unwrapping logic.

**STRICT RULES**:
- 🚫 **FORBIDDEN**: [List common hallucinations/drifts to avoid]
- ✅ **MANDATORY**: [List specific fields/logic to use from Contract]
- 🔧 **TECH STACK**: [Specific stack rules (e.g. computed, wrap in 'data')]

### Task 2: Verification
**File**: `path/to/verify.py`

**Content**:
```python
import os
import sys
import subprocess

def main():
    print("🔍 Verifying Sprint...")
    # Anchor to the current working directory (Project Root)
    PROJECT_ROOT = os.getcwd()
    path = os.path.join(PROJECT_ROOT, "path/to/Component.vue")
    
    if not os.path.exists(path):
        print(f"❌ File missing: {path}")
        sys.exit(1)
        
    print(f"✅ File created: {path}")
    
    # Content Checks
    with open(path, 'r') as f:
        content = f.read()
        checks = [
            ("expectedString1", "Description1"),
            ("expectedString2", "Description2")
        ]
        for text, desc in checks:
            if text in content:
                print(f"✅ {desc}")
            else:
                print(f"❌ Missing: {desc}")
                sys.exit(1)
    
    # Syntax Check (Optional but Recommended)
    # For PHP:
    # result = subprocess.run(['php', '-l', path], capture_output=True)
    # if result.returncode != 0:
    #     print(f"❌ Syntax Error: {result.stderr.decode()}")
    #     sys.exit(1)
    
    print("✨ Verification Passed")
    sys.exit(0)

if __name__ == "__main__":
    main()
```

## ✅ COMPLETION CRITERIA
1. File created successfully.
2. Verification script passes.
3. No syntax errors (if syntax check enabled).
