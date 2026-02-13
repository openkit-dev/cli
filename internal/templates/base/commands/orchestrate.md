---
description: Universal orchestrator for complex multi-agent missions
subtask: false
---

# /orchestrate - Universal Orchestrator

$ARGUMENTS

## Overview

The orchestrator coordinates multiple specialized agents for complex tasks. This command unifies:
- `/engineer` (old) - Original orchestrator
- `/ui-ux` (old) - Design work
- For new applications, use `/orchestrate` directly

## If $ARGUMENTS is empty

Use the question tool:

```javascript
question({
  questions: [{
      question: "Describe the complex task you need executed.",
      header: "Task Description"
    }]
})
```

---

## Mode Detection

### Router Mode (Simple Tasks)

**Triggers when:**
- Keywords: "specify", "clarify", "plan", "verify", "test", "debug", "deploy"
- Single domain tasks

**Action:** Redirect to appropriate command:
- "test..." → `/verify`
- "debug..." → `/debug`
- "deploy..." → `/deploy`
- "discover..." → `/discover`

### Orchestrator Mode (Complex Tasks)

**Triggers when:**
- Multiple domains (backend + frontend + database + security)
- Keywords: "build", "create", "full", "system", "platform"
- Large feature or new project

**Action:** Start multi-phase orchestration

---

## Orchestration Protocol

### Phase 1: Analysis & Planning

1. **Chain of Thought (MANDATORY):**
   - What did the user ask?
   - What is the implicit goal?
   - Which specialists are required?
   - Why is this complex enough?

2. **Create todolist:**
```javascript
todowrite({
  todos: [
    { id: "orch-01", content: "Analyze mission", status: "in_progress", priority: "high" },
    { id: "orch-02", content: "Discovery", status: "pending", priority: "high" },
    { id: "orch-03", content: "Specification", status: "pending", priority: "high" },
    { id: "orch-04", content: "P0: Foundation", status: "pending", priority: "high" },
    { id: "orch-05", content: "P1: Backend", status: "pending", priority: "high" },
    { id: "orch-06", content: "P2: Frontend", status: "pending", priority: "high" },
    { id: "orch-07", content: "P3: Polish", status: "pending", priority: "medium" },
    { id: "orch-08", content: "Verification", status: "pending", priority: "high" }
  ]
})
```

3. **Run `/discover`** if not already done

4. **Run `/specify`** for full specification

---

### Phase 2: Implementation

**Execution Order (P0 → P1 → P2 → P3):**

**P0 - Foundation:**
- `database-architect` (if DB needed)
- `security-auditor` (always)

** STOP:**
```javascript
question({
  questions: [{
      header: "P0 Complete",
      question: "Foundation phase complete. Proceed to P1 (Backend)?",
      options: [
        { label: "Yes, proceed to P1", description: "Continue with backend" },
        { label: "Review P0", description: "Check foundation work" }
      ]
    }]
})
```

**P1 - Core Backend:**
- `backend-specialist`

** STOP:**
```javascript
question({
  questions: [{
      header: "P1 Complete",
      question: "Backend phase complete. Proceed to P2 (Frontend)?",
      options: [
        { label: "Yes, proceed to P2", description: "Continue with UI/UX" },
        { label: "Review P1", description: "Check backend work" }
      ]
    }]
})
```

**P2 - UI/UX:**
- `frontend-specialist` (WEB projects)
- `mobile-developer` (MOBILE projects)

** STOP:**
```javascript
question({
  questions: [{
      header: "P2 Complete",
      question: "Frontend phase complete. Proceed to P3 (Polish)?",
      options: [
        { label: "Yes, proceed to P3", description: "Continue with polish" },
        { label: "Review P2", description: "Check UI/UX work" }
      ]
    }]
})
```

**P3 - Polish:**
- `test-engineer`
- `performance-optimizer` (if needed)

---

### Phase 3: Verification

Run `/verify` automatically

---

## Available Specialist Agents

| Agent | Domain |
|-------|--------|
| `database-architect` | Schema, migrations |
| `security-auditor` | Vulnerabilities, auth |
| `backend-specialist` | API, business logic |
| `frontend-specialist` | React, UI |
| `mobile-developer` | iOS, Android, RN |
| `test-engineer` | QA, E2E |
| `performance-optimizer` | Web Vitals |
| `devops-engineer` | Docker, CI/CD |
| `debugger` | Root cause |

---

## STOP Point

```javascript
question({
  questions: [{
      header: "Orchestration Complete",
      question: "All phases executed. Mark project complete?",
      options: [
        { label: "Yes, complete", description: "Project finished" },
        { label: "Review results", description: "Check all work" }
      ]
    }]
})
```
