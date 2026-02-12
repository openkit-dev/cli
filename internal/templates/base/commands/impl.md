---
description: Execute implementation from plan files with multi-agent coordination
subtask: true
---

**IMPLEMENTATION MODE ACTIVATED**

**User input:** $ARGUMENTS

**If it starts with "from":**
- Extract the path (e.g., "from docs/sprint/Sprint-17/TASKS.md")
- Read that plan and execute its tasks

**If $ARGUMENTS is empty or missing "from":**
Use the question tool to ask:
```javascript
question({
  questions: [{
      question: "Which sprint should I execute?",
      header: "Sprint Selection",
      options: [
        { label: "Latest Sprint", description: "Continue most recent work" },
        { label: "Custom Path", description: "Provide full path" }
      ]
    }]
})
```

**Execution Protocol:**

### Phase 0: TodoList Setup (MANDATORY)

Before any implementation work:

1. **Read existing todolist:**
   ```javascript
   todoread()
   ```

2. **Check if planning todolist exists:**
   - If exists (from `/plan`): Continue with implementation todolist
   - If not exists: Create implementation todolist from scratch

3. **Create implementation todolist (using standard ID schema):**
   ```javascript
   todowrite({
     todos: [
       {
         id: "sdd-impl-01-precheck",
         content: "Pre-implementation checklist validation",
         status: "in_progress",
         priority: "high"
       },
       {
         id: "sdd-impl-02-analysis",
         content: "Read context and identify scope",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-impl-03-p0-foundation",
         content: "P0: Foundation (DB + Security)",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-impl-04-p1-backend",
         content: "P1: Core Backend",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-impl-05-p2-frontend",
         content: "P2: UI/UX",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-impl-06-p3-polish",
         content: "P3: Polish (Tests + Perf)",
         status: "pending",
         priority: "medium"
       },
       {
         id: "sdd-impl-07-update",
         content: "Update sprint TASKS.md",
         status: "pending",
         priority: "medium"
       }
     ]
   })
   ```

4. **Update todolist** as you complete each phase

### Pre-Implementation Checklist
- [ ] `docs/` exists
- [ ] Spec artifacts exist (`docs/requirements/<feature>/`)
- [ ] Plan exists (`docs/requirements/<feature>/PLAN.md`)
- [ ] Selected sprint defined
- [ ] `TASKS.md` exists and is valid
- [ ] User approval obtained
- [ ] Execution order determined

**SDD Gate (MANDATORY):**
If any required spec or plan artifact is missing, STOP and direct the user to run `/specify`, `/clarify`, and `/plan` first.

**After checklist:**
- Update todolist: Mark "sdd-impl-01-precheck" as `completed`
- Mark "sdd-impl-02-analysis" as `in_progress`

### Execution Order (P0 → P1 → P2 → P3)

**P0 - Foundation (Required if applicable):**
-  Invoke `database-architect` (if DB is needed)
-  Invoke `security-auditor` (always for auth/security)

** STOP (use question tool):**
```javascript
question({
  questions: [{
    header: "P0 Complete",
    question: "P0 (Foundation) phase complete. Proceed to P1 (Core Backend)?",
    options: [
      { label: "Yes, proceed to P1", description: "Continue with backend implementation" },
      { label: "Review P0 first", description: "Check foundation work" }
    ]
  }]
})
```
- Update todolist: Mark "sdd-impl-03-p0-foundation" as `completed`
- Mark "sdd-impl-04-p1-backend" as `in_progress`

**P1 - Core Backend:**
-  Invoke `backend-specialist`

** STOP (use question tool):**
```javascript
question({
  questions: [{
    header: "P1 Complete",
    question: "P1 (Core Backend) phase complete. Proceed to P2 (UI/UX)?",
    options: [
      { label: "Yes, proceed to P2", description: "Continue with UI/UX implementation" },
      { label: "Review P1 first", description: "Check backend work" }
    ]
  }]
})
```
- Update todolist: Mark "sdd-impl-04-p1-backend" as `completed`
- Mark "sdd-impl-05-p2-frontend" as `in_progress`

**P2 - UI/UX:**
-  Invoke `frontend-specialist` (WEB projects)
-  Invoke `mobile-developer` (MOBILE projects)

** STOP (use question tool):**
```javascript
question({
  questions: [{
    header: "P2 Complete",
    question: "P2 (UI/UX) phase complete. Proceed to P3 (Polish)?",
    options: [
      { label: "Yes, proceed to P3", description: "Continue with tests and polish" },
      { label: "Review P2 first", description: "Check UI/UX work" }
    ]
  }]
})
```
- Update todolist: Mark "sdd-impl-05-p2-frontend" as `completed`
- Mark "sdd-impl-06-p3-polish" as `in_progress`

**P3 - Polish:**
-  Invoke `test-engineer`
-  Invoke `performance-optimizer` (if needed)

**After P3:**
- Update todolist: Mark "sdd-impl-06-p3-polish" as `completed`
- Mark "sdd-impl-07-update" as `in_progress`

### Progress Updates

Mark each task in `docs/sprint/Sprint-XX/TASKS.md` as completed:
```markdown
- [x] Task 1: [Name]  COMPLETE
- [ ] Task 2: [Name]  IN PROGRESS
```

**After updating TASKS.md:**
- Update todolist: Mark "sdd-impl-07-update" as `completed`

### FINAL STOP POINT

After implementation:
- Final todolist update: Verify all tasks marked as `completed`

Use the question tool:
```javascript
question({
  questions: [{
    header: "Implementation Complete",
    question: "All implementation phases complete. Run final verification (Phase 3) with all validation scripts?",
    options: [
      { label: "Yes, run verification", description: "Execute security scan, lint, tests, and performance checks" },
      { label: "Later", description: "Wait for manual /test" },
      { label: "Skip verification", description: "Mark complete without verification (not recommended)" }
    ]
  }]
})
```

---

## Workflow Reference (Migrated)

# /impl - Implementation

$ARGUMENTS

---

## Overview

This command implements features, fixes bugs, or enhances existing code. It combines "Codegen" (New Features) and "Enhance" (Iterative Updates).

---

## Workflow

### 1. Analysis & Preparation
- ** Reasoning Loop**:
  - *Context*: Confirm `docs/sprint` and relevant `docs/requirements`.
  - *Strategy*: Decide between new file vs modification and pick the pattern (MVC, Component, Hook).
  - *Dependencies*: Identify other files impacted by the change.
- **Read Context**: Check `docs/sprint/` tasks and `docs/requirements/`.
- **Identify Scope**: Decide whether this is a new feature (Codegen) or an update (Enhance).
- **Load State**: Understand current codebase state.

**After analysis:**
- Update todolist: Mark "sdd-impl-02-analysis" as `completed`
- Mark "sdd-impl-03-p0-foundation" as `in_progress`

### 2. Execution (The Loop)

**For New Features (Codegen Mode):**
1. **Before React**: Invoke `nextjs-react-expert`.
2. **Implement**: Write FE/BE code (complete, no TODOs).
3. **Docs**: Update API docs or README if needed.
4. **Status**: Update `docs/sprint/<Sprint-N>/TASKS.md` (MUST follow `rules/MASTER.md` documentation rule).

**For Updates (Enhance Mode):**
1. **Plan Changes**: Determine affected files.
2. **Present Plan** (if complex): "I will modify X files. proceed?"
3. **Apply**: Edit files using `replace_file_content` or `multi_replace`.
4. **Test**: Verify locally if possible.

---

## Output Requirements

- **Production Ready**: No `TODO` comments, strict typing, handled errors.
- **Security**: Validations, Sanitization, Auth checks (follow stack-appropriate patterns).
- **Best Practices**:
    - Use the stack defined in `docs/requirements/<feature>/TECH_STACK.md` or existing project stack
    - Follow patterns appropriate for chosen technologies
    - Maintain consistency with existing codebase

**Note:** Best practices are technology-specific. For example:
- With TanStack Query: Use query keys, staleTime, cache management
- With Redux: Use selectors, action creators, proper state structure
- With FastAPI: Use Pydantic schemas, dependency injection, routers
- With Express: Use middleware, proper error handling, async/await

Always follow the patterns that match the chosen stack.

---

## Usage Examples

```bash
/impl create user profile page
/impl add dark mode to settings
/impl fix validation error in checkout
/impl refactor auth service
```
