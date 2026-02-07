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

3. **Create implementation todolist:**
   ```javascript
   todowrite({
     todos: [
       {
         id: "impl-01-precheck",
         content: "Pre-implementation checklist validation",
         status: "in_progress",
         priority: "high"
       },
       {
         id: "impl-02-analysis",
         content: "Read context and identify scope",
         status: "pending",
         priority: "high"
       },
       {
         id: "impl-03-p0-foundation",
         content: "P0: Foundation (DB architect + Security auditor)",
         status: "pending",
         priority: "high"
       },
       {
         id: "impl-04-p1-backend",
         content: "P1: Core Backend (backend-specialist)",
         status: "pending",
         priority: "high"
       },
       {
         id: "impl-05-p2-frontend",
         content: "P2: UI/UX (frontend-specialist or mobile-developer)",
         status: "pending",
         priority: "high"
       },
       {
         id: "impl-06-p3-polish",
         content: "P3: Polish (test-engineer + performance-optimizer)",
         status: "pending",
         priority: "medium"
       },
       {
         id: "impl-07-tasks-update",
          content: "Update sprint TASKS.md with completion status",
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
- Update todolist: Mark "impl-01-precheck" as `completed`
- Mark "impl-02-analysis" as `in_progress`

### Execution Order (P0 → P1 → P2 → P3)

**P0 - Foundation (Required if applicable):**
-  Invoke `database-architect` (if DB is needed)
-  Invoke `security-auditor` (always for auth/security)

** STOP:**
- Update todolist: Mark "impl-03-p0-foundation" as `completed`
- Mark "impl-04-p1-backend" as `in_progress`
- Use the question tool to ask "P0 phase complete. Proceed to P1 (Core Backend)?"

**P1 - Core Backend:**
-  Invoke `backend-specialist`

** STOP:**
- Update todolist: Mark "impl-04-p1-backend" as `completed`
- Mark "impl-05-p2-frontend" as `in_progress`
- Use the question tool to ask "P1 phase complete. Proceed to P2 (UI/UX)?"

**P2 - UI/UX:**
-  Invoke `frontend-specialist` (WEB projects)
-  Invoke `mobile-developer` (MOBILE projects)

** STOP:**
- Update todolist: Mark "impl-05-p2-frontend" as `completed`
- Mark "impl-06-p3-polish" as `in_progress`
- Use the question tool to ask "P2 phase complete. Proceed to P3 (Polish)?"

**P3 - Polish:**
-  Invoke `test-engineer`
-  Invoke `performance-optimizer` (if needed)

**After P3:**
- Update todolist: Mark "impl-06-p3-polish" as `completed`
- Mark "impl-07-tasks-update" as `in_progress`

### Progress Updates

Mark each task in `docs/sprint/Sprint-XX/TASKS.md` as completed:
```markdown
- [x] Task 1: [Name]  COMPLETE
- [ ] Task 2: [Name]  IN PROGRESS
```

**After updating TASKS.md:**
- Update todolist: Mark "impl-07-tasks-update" as `completed`

### FINAL STOP POINT

After implementation:
- Final todolist update: Verify all tasks marked as `completed`
> **STOP:** Use the question tool to ask "Implementation complete!  Run final verification (Phase X) with all validation scripts?"

**User options:**
- "yes" → Run Phase X (add verification tasks to todolist)
- "no" → Wait for manual `/test`

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
- Update todolist: Mark "impl-02-analysis" as `completed`
- Mark "impl-03-p0-foundation" as `in_progress`

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
    - Use the stack defined in `docs/requirements/<feature>/TechStack.md` or existing project stack
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
