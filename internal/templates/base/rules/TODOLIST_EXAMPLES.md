# TodoList Usage Examples

> Practical examples demonstrating correct usage of `todowrite` and `todoread` tools.

** IMPORTANT:** This file shows examples. For complete protocol, see `.opencode/rules/TODOLIST_PROTOCOL.md`

**Key rules:**
-  NO emojis in task content
-  MUST use `question` tool for approval after creating todolist
-  Only ONE task `in_progress` at a time

---

## When to Use TodoList

**Use todolist tools when:**
-  Task has **3+ steps**
-  Task is **complex** (requires planning)
-  Task has **dependencies** between steps
-  Task involves **multiple agents** (orchestration)
-  Task has **STOP points** requiring approval
-  Task requires **progress tracking** over time

**Don't use todolist when:**
-  Single, immediate operation (e.g., "run npm install")
-  Simple read/analysis task (e.g., "read this file")
-  Straightforward question/answer

---

## Complete Examples

### Example 1: Multi-File Refactoring (5+ steps)

**User request:** "Refactor authentication system to use JWT instead of sessions"

**Correct approach:**

```javascript
// Step 1: Check existing todos (if available)
todoread()

// Step 2: Create todolist
todowrite({
  todos: [
    {
      id: "refactor-1",
      content: "Analyze current session-based auth code",
      status: "pending",
      priority: "high"
    },
    {
      id: "refactor-2",
      content: "Install JWT dependencies",
      status: "pending",
      priority: "high"
    },
    {
      id: "refactor-3",
      content: "Create JWT utility functions",
      status: "pending",
      priority: "high"
    },
    {
      id: "refactor-4",
      content: "Update login endpoint to generate JWT",
      status: "pending",
      priority: "high"
    },
    {
      id: "refactor-5",
      content: "Update auth middleware to verify JWT",
      status: "pending",
      priority: "high"
    },
    {
      id: "refactor-6",
      content: "Remove session dependencies",
      status: "pending",
      priority: "medium"
    },
    {
      id: "refactor-7",
      content: "Update tests",
      status: "pending",
      priority: "medium"
    }
  ]
})

// Step 3: Start first task
todowrite({
  todos: [
    {
      id: "refactor-1",
      content: "Analyze current session-based auth code",
      status: "in_progress", // ← Mark as in_progress
      priority: "high"
    },
    // ... rest pending
  ]
})

// Perform analysis...
glob("**/*auth*.ts")
read("src/auth/session.ts")

// Step 4: Complete first task, start next
todowrite({
  todos: [
    {
      id: "refactor-1",
      content: "Analyze current session-based auth code",
      status: "completed", // ← Mark completed
      priority: "high"
    },
    {
      id: "refactor-2",
      content: "Install JWT dependencies",
      status: "in_progress", // ← Start next
      priority: "high"
    },
    // ... rest pending
  ]
})

// Continue through all tasks...

// Final step: All tasks completed
todowrite({
  todos: [
    {
      id: "refactor-1",
      content: "Analyze current session-based auth code",
      status: "completed",
      priority: "high"
    },
    // ... all marked completed
    {
      id: "refactor-7",
      content: "Update tests",
      status: "completed",
      priority: "medium"
    }
  ]
})
```

---

### Example 2: Orchestrator Managing Multiple Agents

**User request:** "Build a complete e-commerce checkout flow with Stripe"

**Correct approach:**

```javascript
// Step 1: Check existing todos
todoread()

// Step 2: Create orchestration todolist
todowrite({
  todos: [
    {
      id: "phase1-discovery",
      content: "Phase 1: Analyze requirements",
      status: "pending",
      priority: "high"
    },
    {
      id: "phase1-planning",
      content: "Phase 1: Create planning artifacts",
      status: "pending",
      priority: "high"
    },
    {
      id: "phase2-database",
      content: "Phase 2: Database schema (database-architect)",
      status: "pending",
      priority: "high"
    },
    {
      id: "phase2-security",
      content: "Phase 2: Security audit (security-auditor)",
      status: "pending",
      priority: "high"
    },
    {
      id: "phase2-backend",
      content: "Phase 2: Stripe API integration (backend-specialist)",
      status: "pending",
      priority: "high"
    },
    {
      id: "phase2-frontend",
      content: "Phase 2: Checkout UI (frontend-specialist)",
      status: "pending",
      priority: "high"
    },
    {
      id: "phase3-tests",
      content: "Phase 3: E2E tests (test-engineer)",
      status: "pending",
      priority: "medium"
    },
    {
      id: "phase3-verification",
      content: "Phase 3: Final verification",
      status: "pending",
      priority: "high"
    }
  ]
})

// Step 3: Start Phase 1 - Discovery
todowrite({
  todos: [
    {
      id: "phase1-discovery",
      content: "Phase 1: Analyze requirements",
      status: "in_progress",
      priority: "high"
    },
    // ... rest pending
  ]
})

// Perform discovery...
glob("**/*.ts")
grep("checkout", "**/*.ts")

// Step 4: Complete discovery, start planning
todowrite({
  todos: [
    {
      id: "phase1-discovery",
      content: "Phase 1: Analyze requirements",
      status: "completed",
      priority: "high"
    },
    {
      id: "phase1-planning",
      content: "Phase 1: Create planning artifacts",
      status: "in_progress",
      priority: "high"
    },
    // ... rest pending
  ]
})

// Invoke project-planner...
task(subagent_type="project-planner", prompt="Create plan for e-commerce checkout")

// Step 5: Complete planning, start Phase 2 (parallel agents)
todowrite({
  todos: [
    {
      id: "phase1-discovery",
      content: "Phase 1: Analyze requirements",
      status: "completed",
      priority: "high"
    },
    {
      id: "phase1-planning",
      content: "Phase 1: Create planning artifacts",
      status: "completed",
      priority: "high"
    },
    {
      id: "phase2-database",
      content: "Phase 2: Database schema (database-architect)",
      status: "in_progress", // ← Start foundation tasks
      priority: "high"
    },
    {
      id: "phase2-security",
      content: "Phase 2: Security audit (security-auditor)",
      status: "in_progress", // ← Parallel with database
      priority: "high"
    },
    // ... rest pending
  ]
})

// STOP POINT: Ask user for approval
question({
  questions: [{
    question: "Planning complete. Proceed with Phase 2 implementation?",
    header: "Approval Required",
    options: [
      { label: "Yes, proceed", description: "Start implementation" },
      { label: "Review plan first", description: "Show me the plan details" }
    ]
  }]
})

// If approved, invoke parallel agents...
task(subagent_type="database-architect", prompt="Design checkout schema")
task(subagent_type="security-auditor", prompt="Audit Stripe integration")

// Continue updating todolist as each phase completes...
```

---

### Example 3: Creating New Feature (Planning + Implementation)

**User request:** "Add user profile page with avatar upload"

**Correct approach:**

```javascript
// Step 1: Check existing todos
todoread()

// Step 2: Create feature todolist
todowrite({
  todos: [
    {
      id: "profile-1",
      content: "Design database schema for profiles",
      status: "pending",
      priority: "high"
    },
    {
      id: "profile-2",
      content: "Create profile API endpoints",
      status: "pending",
      priority: "high"
    },
    {
      id: "profile-3",
      content: "Set up file upload service (avatars)",
      status: "pending",
      priority: "high"
    },
    {
      id: "profile-4",
      content: "Build profile UI component",
      status: "pending",
      priority: "high"
    },
    {
      id: "profile-5",
      content: "Add avatar upload component",
      status: "pending",
      priority: "high"
    },
    {
      id: "profile-6",
      content: "Write unit tests",
      status: "pending",
      priority: "medium"
    },
    {
      id: "profile-7",
      content: "Write E2E tests",
      status: "pending",
      priority: "medium"
    }
  ]
})

// Step 3: Start with schema design
todowrite({
  todos: [
    {
      id: "profile-1",
      content: "Design database schema for profiles",
      status: "in_progress",
      priority: "high"
    },
    // ... rest pending
  ]
})

// Design schema...
read("backend/models/user.ts")
edit("backend/models/user.ts", oldString, newString)

// Step 4: Complete schema, start API
todowrite({
  todos: [
    {
      id: "profile-1",
      content: "Design database schema for profiles",
      status: "completed",
      priority: "high"
    },
    {
      id: "profile-2",
      content: "Create profile API endpoints",
      status: "in_progress",
      priority: "high"
    },
    // ... rest pending
  ]
})

// Create API...
write("backend/routes/profile.ts", apiContent)

// Continue through all tasks...
```

---

### Example 4: Bug Investigation and Fix (4 steps)

**User request:** "Login page crashes on submit"

**Correct approach:**

```javascript
// Step 1: Check existing todos
todoread()

// Step 2: Create debugging todolist
todowrite({
  todos: [
    {
      id: "debug-1",
      content: "Reproduce bug and capture logs",
      status: "pending",
      priority: "high"
    },
    {
      id: "debug-2",
      content: "Identify root cause",
      status: "pending",
      priority: "high"
    },
    {
      id: "debug-3",
      content: "Implement fix",
      status: "pending",
      priority: "high"
    },
    {
      id: "debug-4",
      content: "Verify fix and add regression test",
      status: "pending",
      priority: "high"
    }
  ]
})

// Step 3: Start reproduction
todowrite({
  todos: [
    {
      id: "debug-1",
      content: "Reproduce bug and capture logs",
      status: "in_progress",
      priority: "high"
    },
    // ... rest pending
  ]
})

// Reproduce...
read("frontend/pages/login.tsx")
grep("onSubmit", "frontend/**/*.tsx")

// Step 4: Complete reproduction, identify cause
todowrite({
  todos: [
    {
      id: "debug-1",
      content: "Reproduce bug and capture logs",
      status: "completed",
      priority: "high"
    },
    {
      id: "debug-2",
      content: "Identify root cause",
      status: "in_progress",
      priority: "high"
    },
    // ... rest pending
  ]
})

// Continue through fix and verification...
```

---

## Anti-Patterns (What NOT to Do)

### Anti-Pattern 1: Using todolist for simple tasks

```javascript
//  WRONG - Too simple for todolist
todowrite({
  todos: [
    {
      id: "simple-1",
      content: "Read file",
      status: "pending",
      priority: "high"
    }
  ]
})

//  CORRECT - Just do it
read("file.ts")
```

### Anti-Pattern 2: Not using todoread before todowrite

```javascript
//  WRONG - Creates duplicate todos
todowrite({
  todos: [...]
})

//  CORRECT - Check existing first
todoread()
// Then decide: create new or update existing
todowrite({
  todos: [...]
})
```

### Anti-Pattern 3: Multiple tasks in_progress

```javascript
//  WRONG - Can't do 2 things at once
todowrite({
  todos: [
    { id: "task1", status: "in_progress", ... },
    { id: "task2", status: "in_progress", ... } // ← BAD!
  ]
})

//  CORRECT - One at a time
todowrite({
  todos: [
    { id: "task1", status: "in_progress", ... },
    { id: "task2", status: "pending", ... } // ← Wait your turn
  ]
})
```

### Anti-Pattern 4: Forgetting to update todolist

```javascript
//  WRONG - Created todolist but never updated
todowrite({ todos: [...] })
// ... do work ...
// ... do more work ...
// ... forgot to mark completed!

//  CORRECT - Update as you go
todowrite({ todos: [...] })
// Step 1: Mark in_progress
todowrite({ todos: [{ status: "in_progress" }] })
// ... do work ...
// Step 2: Mark completed
todowrite({ todos: [{ status: "completed" }] })
```

### Anti-Pattern 5: Inline tracking for complex tasks

```javascript
//  WRONG - Complex task needs todolist
"Working on feature...
Step 1: Design schema
Step 2: Create API
Step 3: Build UI
Step 4: Write tests
Step 5: Deploy"

//  CORRECT - Use todolist
todowrite({
  todos: [
    { id: "step1", content: "Design schema", ... },
    { id: "step2", content: "Create API", ... },
    { id: "step3", content: "Build UI", ... },
    { id: "step4", content: "Write tests", ... },
    { id: "step5", content: "Deploy", ... }
  ]
})
```

---

## Decision Tree: Should I Use TodoList?

```
Is the task complex? (3+ steps)
    ├─ Yes → Does it involve multiple phases?
    │         ├─ Yes → Does it need STOP points?
    │         │         ├─ Yes →  USE TODOLIST (orchestration pattern)
    │         │         └─ No →  USE TODOLIST (tracking pattern)
    │         └─ No →  USE TODOLIST (simple tracking)
    │
    └─ No → Is it a single command/read?
              ├─ Yes →  DON'T USE TODOLIST (inline is fine)
              └─ No → Maybe 2-3 steps?
                        ├─ Complex? →  USE TODOLIST
                        └─ Simple? →  DON'T USE TODOLIST
```

---

## Quick Reference Table

| Task Type | Steps | Use TodoList? | Example |
|-----------|-------|---------------|---------|
| Simple read | 1 |  No | "Read this file" |
| Single command | 1 |  No | "Run npm install" |
| Analysis task | 1-2 |  No | "Analyze auth code" |
| Small fix | 2-3 |  Maybe | "Fix typo in 2 files" |
| Feature addition | 3-5 |  Yes | "Add profile page" |
| Refactoring | 4-7 |  Yes | "Refactor auth to JWT" |
| Multi-agent task | 5+ |  Yes | "Build checkout flow" |
| Full feature | 7+ |  Yes | "E-commerce with Stripe" |

---

## Best Practices

1. **Always check existing todos first** (if `todoread` available)
   ```javascript
   todoread() // Check what's already there
   ```

2. **Create todolist at the START of complex tasks**
   ```javascript
   // As soon as you understand the scope
   todowrite({ todos: [...] })
   ```

3. **Only ONE task `in_progress` at a time**
   ```javascript
   // Complete current before starting next
   ```

4. **Mark completed IMMEDIATELY after finishing**
   ```javascript
   // Don't wait to batch updates
   todowrite({ todos: [{ status: "completed" }] })
   ```

5. **Use descriptive but concise content** (max 50 chars)
   ```javascript
   { content: "Create user profile API endpoint" } //  Good
   { content: "Do stuff with the database" } //  Too vague
   ```

6. **Add checkmarks to completed tasks**
   ```javascript
   { content: "Design schema", status: "completed" }
   ```

7. **Update todolist at STOP points**
   ```javascript
   // Before asking user for approval
   todowrite({ todos: [...] })
   question({ questions: [...] })
   ```

---

## References

- **Tool Usage**: `.opencode/rules/TOOL_USAGE.md`
- **Master Rules**: `.opencode/rules/MASTER.md`
- **Official Docs**: https://opencode.ai/docs/tools/#todowrite

---

**Use this guide as a reference when deciding whether and how to use todolist tools in your workflow.**
