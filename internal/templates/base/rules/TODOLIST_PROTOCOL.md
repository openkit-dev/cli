---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# TODOLIST PROTOCOL — MANDATORY WORKFLOW

> Mandatory workflow for creating, approving, and managing todo lists across all agents.

---

## Purpose

This rule enforces **correct todolist workflow** with **mandatory user approval** before execution:

1. **Create** todolist with `todowrite`
2. **Request approval** with `question` tool
3. **Execute** only after user approval
4. **Update status** without emojis

**This protocol is MANDATORY for all complex tasks (3+ steps).**

---

## CRITICAL RULES

### 1. **NO Emojis in Status**

** NEVER use emojis in todo items:**
```javascript
//  WRONG - Has emoji
{
  id: "task-1",
  content: "Complete task ",
  status: "completed",
  priority: "high"
}

//  CORRECT - No emoji
{
  id: "task-1",
  content: "Complete task",
  status: "completed",
  priority: "high"
}
```

**Why?** Emojis pollute the task content and make it harder to read/parse.

---

### 2. **Mandatory Approval with `question` Tool**

**After creating todolist, ALWAYS ask for approval:**

```javascript
// Step 1: Create todolist
todowrite({
  todos: [
    { id: "task-1", content: "Design schema", status: "pending", priority: "high" },
    { id: "task-2", content: "Create API", status: "pending", priority: "high" },
    { id: "task-3", content: "Build UI", status: "pending", priority: "medium" }
  ]
})

// Step 2: MANDATORY - Request approval
question({
  questions: [{
    question: "Todo list created with 3 tasks. Proceed with execution?",
    header: "Approve Execution",
    options: [
      { label: "Yes, proceed", description: "Start executing tasks in order" },
      { label: "Modify tasks", description: "I want to adjust the task list" },
      { label: "Cancel", description: "Don't execute" }
    ]
  }]
})

// Step 3: Only execute if approved
// If user says "Yes, proceed", then start execution
```

---

### 3. **Status Values (No Emojis)**

**Valid status values:**
- `"pending"` - Task not started
- `"in_progress"` - Currently working on
- `"completed"` - Task finished
- `"cancelled"` - Task no longer needed

** INVALID (with emojis):**
- `"pending ⏳"`
- `"in_progress "`
- `"completed "`
- `"cancelled "`

**Status is a field, emojis don't belong there.**

---

### 4. **Only ONE Task in_progress**

At any time, only ONE task can have `status: "in_progress"`:

```javascript
//  CORRECT
todowrite({
  todos: [
    { id: "task-1", status: "in_progress", ... },  // Working on this
    { id: "task-2", status: "pending", ... },       // Wait
    { id: "task-3", status: "pending", ... }        // Wait
  ]
})

//  WRONG
todowrite({
  todos: [
    { id: "task-1", status: "in_progress", ... },  // Can't do both
    { id: "task-2", status: "in_progress", ... }   // at same time!
  ]
})
```

---

## Complete Workflow

### Phase 1: Create TodoList

```javascript
// 1. Check for existing todos (if todoread available)
todoread()

// 2. Create todolist with clear, emoji-free content
todowrite({
  todos: [
    {
      id: "task-1",
      content: "Design database schema",  // NO emoji
      status: "pending",
      priority: "high"
    },
    {
      id: "task-2",
      content: "Create API endpoints",    // NO emoji
      status: "pending",
      priority: "high"
    },
    {
      id: "task-3",
      content: "Build frontend UI",       // NO emoji
      status: "pending",
      priority: "medium"
    }
  ]
})
```

---

### Phase 2: Request Approval (MANDATORY)

```javascript
// ALWAYS use question tool after creating todolist
question({
  questions: [{
    question: "Task list created with 3 tasks:\n\n1. Design database schema (high priority)\n2. Create API endpoints (high priority)\n3. Build frontend UI (medium priority)\n\nProceed with execution?",
    header: "Approve Task Execution",
    options: [
      { 
        label: "Yes, proceed", 
        description: "Start executing tasks sequentially" 
      },
      { 
        label: "Modify list", 
        description: "Adjust tasks before execution" 
      },
      { 
        label: "Cancel", 
        description: "Don't execute" 
      }
    ]
  }]
})
```

**Do NOT proceed without approval!**

---

### Phase 3: Execute (After Approval)

```javascript
// Only if user approved with "Yes, proceed"

// Start task 1
todowrite({
  todos: [
    { id: "task-1", content: "Design database schema", status: "in_progress", priority: "high" },
    { id: "task-2", content: "Create API endpoints", status: "pending", priority: "high" },
    { id: "task-3", content: "Build frontend UI", status: "pending", priority: "medium" }
  ]
})

// ... work on task 1 ...

// Complete task 1, start task 2
todowrite({
  todos: [
    { id: "task-1", content: "Design database schema", status: "completed", priority: "high" },
    { id: "task-2", content: "Create API endpoints", status: "in_progress", priority: "high" },
    { id: "task-3", content: "Build frontend UI", status: "pending", priority: "medium" }
  ]
})

// ... work on task 2 ...

// Complete task 2, start task 3
todowrite({
  todos: [
    { id: "task-1", content: "Design database schema", status: "completed", priority: "high" },
    { id: "task-2", content: "Create API endpoints", status: "completed", priority: "high" },
    { id: "task-3", content: "Build frontend UI", status: "in_progress", priority: "medium" }
  ]
})

// ... work on task 3 ...

// Complete all tasks
todowrite({
  todos: [
    { id: "task-1", content: "Design database schema", status: "completed", priority: "high" },
    { id: "task-2", content: "Create API endpoints", status: "completed", priority: "high" },
    { id: "task-3", content: "Build frontend UI", status: "completed", priority: "medium" }
  ]
})
```

---

## Common Mistakes

### Mistake 1: Using Emojis

```javascript
//  WRONG
{ content: "Task complete ", status: "completed" }

//  CORRECT
{ content: "Task complete", status: "completed" }
```

### Mistake 2: Not Requesting Approval

```javascript
//  WRONG - No approval
todowrite({ todos: [...] })
// Start executing immediately ← BAD!

//  CORRECT - Request approval
todowrite({ todos: [...] })
question({ questions: [...] })  // ← MANDATORY
// Wait for user response
// Only execute if approved
```

### Mistake 3: Multiple in_progress

```javascript
//  WRONG
todowrite({
  todos: [
    { id: "task-1", status: "in_progress", ... },
    { id: "task-2", status: "in_progress", ... }  // ← BAD!
  ]
})

//  CORRECT
todowrite({
  todos: [
    { id: "task-1", status: "in_progress", ... },
    { id: "task-2", status: "pending", ... }  // ← Wait
  ]
})
```

### Mistake 4: Not Updating Status

```javascript
//  WRONG - Created but never updated
todowrite({ todos: [...] })
question({ questions: [...] })
// ... do work but forget to update status

//  CORRECT - Update as you go
todowrite({ todos: [...] })
question({ questions: [...] })
todowrite({ todos: [{ status: "in_progress" }] })  // ← Mark started
// ... do work ...
todowrite({ todos: [{ status: "completed" }] })    // ← Mark completed
```

---

## Validation Checklist

Before using todolist, verify:

- [ ] Task has **3+ steps** (otherwise don't use todolist)
- [ ] Task content has **NO emojis** (clean text only)
- [ ] Used **`todoread`** first (if available)
- [ ] Created todolist with **`todowrite`**
- [ ] **MANDATORY:** Used **`question`** tool for approval
- [ ] Waited for user approval before executing
- [ ] Only **ONE task** has `status: "in_progress"`
- [ ] Updated status as tasks complete (NO emojis)

---

## Examples

### Example 1: Simple Feature (3 tasks)

```javascript
// Step 1: Create todolist
todowrite({
  todos: [
    { id: "feat-1", content: "Create database model", status: "pending", priority: "high" },
    { id: "feat-2", content: "Add API endpoint", status: "pending", priority: "high" },
    { id: "feat-3", content: "Update frontend", status: "pending", priority: "medium" }
  ]
})

// Step 2: Request approval (MANDATORY)
question({
  questions: [{
    question: "Created 3-task plan for new feature. Proceed with implementation?",
    header: "Approve Execution",
    options: [
      { label: "Yes, proceed", description: "Start implementation" },
      { label: "Modify", description: "Adjust plan first" }
    ]
  }]
})

// Step 3: Execute (only if approved)
// User says "Yes, proceed"
todowrite({
  todos: [
    { id: "feat-1", content: "Create database model", status: "in_progress", priority: "high" },
    { id: "feat-2", content: "Add API endpoint", status: "pending", priority: "high" },
    { id: "feat-3", content: "Update frontend", status: "pending", priority: "medium" }
  ]
})

// ... work on model ...

todowrite({
  todos: [
    { id: "feat-1", content: "Create database model", status: "completed", priority: "high" },
    { id: "feat-2", content: "Add API endpoint", status: "in_progress", priority: "high" },
    { id: "feat-3", content: "Update frontend", status: "pending", priority: "medium" }
  ]
})

// ... continue ...
```

---

### Example 2: Complex Orchestration (5+ tasks)

```javascript
// Step 1: Create comprehensive todolist
todowrite({
  todos: [
    { id: "phase1-plan", content: "Phase 1: Create planning artifacts", status: "pending", priority: "high" },
    { id: "phase2-db", content: "Phase 2: Database schema", status: "pending", priority: "high" },
    { id: "phase2-api", content: "Phase 2: API implementation", status: "pending", priority: "high" },
    { id: "phase2-ui", content: "Phase 2: UI components", status: "pending", priority: "high" },
    { id: "phase3-test", content: "Phase 3: Tests and verification", status: "pending", priority: "medium" }
  ]
})

// Step 2: Request approval (MANDATORY)
question({
  questions: [{
    question: "Multi-phase plan created with 5 major tasks:\n\nPhase 1: Planning (1 task)\nPhase 2: Implementation (3 tasks)\nPhase 3: Testing (1 task)\n\nThis will coordinate multiple specialist agents. Proceed?",
    header: "Approve Orchestration",
    options: [
      { label: "Yes, proceed", description: "Start multi-agent orchestration" },
      { label: "Review plan", description: "Show detailed breakdown" },
      { label: "Modify", description: "Adjust phases" }
    ]
  }]
})

// Step 3: Execute phases (only if approved)
// ... execution continues ...
```

---

## Integration with Other Tools

### With `question` Tool

**ALWAYS pair todolist creation with question tool:**

```javascript
// Pattern 1: Simple approval
todowrite({ todos: [...] })
question({
  questions: [{
    question: "Proceed with X tasks?",
    header: "Approve",
    options: [
      { label: "Yes", description: "Execute" },
      { label: "No", description: "Cancel" }
    ]
  }]
})

// Pattern 2: With options
todowrite({ todos: [...] })
question({
  questions: [{
    question: "How to proceed?",
    header: "Execution Options",
    options: [
      { label: "Execute all", description: "Run all tasks" },
      { label: "Execute Phase 1 only", description: "Stop after planning" },
      { label: "Modify", description: "Change tasks" }
    ]
  }]
})
```

---

### With `task` Tool (Subagents)

When invoking subagents:

```javascript
// Update todolist before invoking subagent
todowrite({
  todos: [
    { id: "backend-work", status: "in_progress", ... }
  ]
})

// Invoke subagent
task(subagent_type="backend-specialist", prompt="...")

// Update after subagent completes
todowrite({
  todos: [
    { id: "backend-work", status: "completed", ... },
    { id: "frontend-work", status: "in_progress", ... }
  ]
})
```

---

## Standard TodoList ID Schema

**All commands MUST use this standardized ID schema:**

```
{workflow}-{phase}-{step}
```

### ID Prefixes by Command

| Command | ID Prefix | Example |
|---------|-----------|---------|
| `/discover` | `disc-` | `disc-analyze`, `disc-context`, `disc-document` |
| `/specify` | `spec-` | `spec-problem`, `spec-stories`, `spec-criteria`, `spec-plan`, `spec-tasks` |
| `/create` | `create-` | `create-p0`, `create-p1`, `create-p2`, `create-p3` |
| `/verify` | `verify-` | `verify-lint`, `verify-security`, `verify-tests` |
| `/orchestrate` | `orch-` | `orch-phase1`, `orch-phase2`, `orch-phase3` |
| `/debug` | `debug-` | `debug-symptom`, `debug-gather`, `debug-hypothesis`, `debug-resolve` |
| `/deploy` | `deploy-` | `deploy-prep`, `deploy-exec`, `deploy-verify` |

### Standard Phase IDs for OpenKit Workflow

```javascript
// /discover command
{ id: "disc-01-analyze", content: "Analyze project structure", ... }
{ id: "disc-02-context", content: "Generate context documentation", ... }

// /specify command
{ id: "spec-01-problem", content: "Create PROBLEM_STATEMENT.md", ... }
{ id: "spec-02-stories", content: "Create USER_STORIES.md", ... }
{ id: "spec-03-criteria", content: "Create ACCEPTANCE_CRITERIA.md", ... }
{ id: "spec-04-risks", content: "Create RISKS.md", ... }
{ id: "spec-05-plan", content: "Create PLAN.md", ... }
{ id: "spec-06-tasks", content: "Create TASKS.md", ... }
{ id: "sdd-plan-04-sprint", content: "Create sprint artifacts", ... }

// /specify command (includes planning + tasks)
{ id: "spec-01-problem", content: "Create PROBLEM_STATEMENT.md", ... }
{ id: "spec-02-stories", content: "Create USER_STORIES.md", ... }
{ id: "spec-03-criteria", content: "Create ACCEPTANCE_CRITERIA.md", ... }
{ id: "spec-04-risks", content: "Create RISKS.md", ... }
{ id: "spec-05-plan", content: "Create PLAN.md and sprint artifacts", ... }
{ id: "spec-06-tasks", content: "Create TASKS.md with INPUT->OUTPUT->VERIFY", ... }

// /create command
{ id: "create-01-precheck", content: "Pre-implementation checklist", ... }
{ id: "create-02-p0-foundation", content: "P0: Foundation (DB + Security)", ... }
{ id: "create-03-p1-backend", content: "P1: Core Backend", ... }
{ id: "create-04-p2-frontend", content: "P2: UI/UX", ... }
{ id: "create-05-p3-polish", content: "P3: Polish (Tests)", ... }
{ id: "create-06-update", content: "Update sprint TASKS.md", ... }
```

### Standard Phase IDs for Orchestration

```javascript
// /orchestrate command
{ id: "orch-01-analysis", content: "Analyze mission complexity", ... }
{ id: "orch-02-discover", content: "Run /discover", ... }
{ id: "orch-03-specify", content: "Run /specify", ... }
{ id: "orch-04-p0-foundation", content: "P0: Foundation (DB + Security)", ... }
{ id: "orch-05-p1-backend", content: "P1: Core Backend", ... }
{ id: "orch-06-p2-frontend", content: "P2: UI/UX", ... }
{ id: "orch-07-p3-polish", content: "P3: Polish (Tests + Perf)", ... }
{ id: "orch-08-verification", content: "Run /verify", ... }
```

### Why Standardized IDs?

1. **Continuity:** When user runs `/discover` then `/specify` then `/create`, the IDs connect logically
2. **Tracking:** Easy to filter by prefix to see what phase/command we're in
3. **Debugging:** Clear provenance when something fails
4. **Consistency:** All agents speak the same language

---

## Summary

**The TodoList Protocol is:**

1. **Create** todolist with `todowrite` (NO emojis)
2. **Use standard ID schema** (`{workflow}-{phase}-{step}`)
3. **Request approval** with `question` tool (MANDATORY)
4. **Execute** only after user approval
5. **Update status** as you progress (NO emojis)
6. **Only ONE** task `in_progress` at a time

**This is a MANDATORY workflow for all complex tasks (3+ steps).**

---

## References

- **Tool Usage**: `.opencode/rules/TOOL_USAGE.md`
- **TodoList Examples**: `.opencode/rules/TODOLIST_EXAMPLES.md`
- **Master Rules**: `.opencode/rules/MASTER.md`
- **Official Docs**: https://opencode.ai/docs/tools/#todowrite

---

**Last Updated:** 2026-02-04  
**Priority:** P0 (Critical)  
**Enforcement:** MANDATORY for all agents, commands, and skills
