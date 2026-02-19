---
description: Execute implementation from plan with multi-agent coordination
subtask: true
---

# /create - Implementation

$ARGUMENTS

## Overview

Execute implementation tasks from the specification. This command:
- Reads tasks from `openkit-memory/sprint/Sprint-XX/TASKS.md`
- Executes in priority order (P0 → P1 → P2 → P3)
- Coordinates multiple agents as needed

**IMPORTANT:** `/specify` MUST be complete before running this command.

**For new applications:** Use `/orchestrate` instead.

## If $ARGUMENTS is empty or "from" syntax

```javascript
question({
  questions: [{
      question: "Which sprint to implement?",
      header: "Sprint Selection",
      options: [
        { label: "Latest Sprint", description: "Continue most recent work" },
        { label: "Custom Path", description: "Provide full path" }
      ]
    }]
})
```

**Syntax:** `from openkit-memory/sprint/Sprint-XX/TASKS.md`

---

## Pre-Implementation Checklist

- [ ] `/discover` has been run
- [ ] `/specify` is complete
- [ ] Sprint is selected
- [ ] `TASKS.md` exists and is valid

**SDD Gate:** If spec is missing, STOP and direct to `/specify` first.

---

## TodoList Setup

```javascript
todowrite({
  todos: [
    { id: "impl-01", content: "Precheck validation", status: "in_progress", priority: "high" },
    { id: "impl-02", content: "P0: Foundation", status: "pending", priority: "high" },
    { id: "impl-03", content: "P1: Backend", status: "pending", priority: "high" },
    { id: "impl-04", content: "P2: Frontend", status: "pending", priority: "high" },
    { id: "impl-05", content: "P3: Polish", status: "pending", priority: "medium" },
    { id: "impl-06", content: "Update TASKS.md", status: "pending", priority: "medium" }
  ]
})
```

---

## Execution Order (P0 → P1 → P2 → P3)

### P0 - Foundation

- Invoke `database-architect` (if DB needed)
- Invoke `security-auditor` (always)

**STOP:**
```javascript
question({
  questions: [{
      header: "P0 Complete",
      question: "Foundation complete. Proceed to P1 (Backend)?",
      options: [
        { label: "Yes, proceed to P1", description: "Continue with backend" },
        { label: "Review P0", description: "Check foundation work" }
      ]
    }]
})
```

### P1 - Core Backend

- Invoke `backend-specialist`

**STOP:**
```javascript
question({
  questions: [{
      header: "P1 Complete",
      question: "Backend complete. Proceed to P2 (Frontend)?",
      options: [
        { label: "Yes, proceed to P2", description: "Continue with UI/UX" },
        { label: "Review P1", description: "Check backend work" }
      ]
    }]
})
```

### P2 - UI/UX

- Invoke `frontend-specialist` (WEB projects)
- Invoke `mobile-developer` (MOBILE projects)

**STOP:**
```javascript
question({
  questions: [{
      header: "P2 Complete",
      question: "Frontend complete. Proceed to P3 (Polish)?",
      options: [
        { label: "Yes, proceed to P3", description: "Continue with polish" },
        { label: "Review P2", description: "Check UI/UX work" }
      ]
    }]
})
```

### P3 - Polish

- Invoke `test-engineer`
- Invoke `performance-optimizer` (if needed)

---

## Progress Updates

Mark tasks in `openkit-memory/sprint/Sprint-XX/TASKS.md`:

```markdown
- [x] task-01: [Name]  COMPLETE
- [ ] task-02: [Name]  IN PROGRESS
```

---

## Final STOP

```javascript
question({
  questions: [{
      header: "Implementation Complete",
      question: "All phases complete. Run verification (/verify)?",
      options: [
        { label: "Yes, run /verify", description: "Execute security, lint, tests" },
        { label: "Later", description: "Skip verification for now" }
      ]
    }]
})
```
