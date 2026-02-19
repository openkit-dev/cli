---
description: Full specification + planning + task breakdown (unifies /specify, /plan, /tasks)
subtask: false
---

# /specify - Specification + Planning + Tasks

$ARGUMENTS

## Overview

Complete workflow from specification to task breakdown. This command unifies (old commands):
- `/specify` (old) - Feature specification
- `/clarify` (old) - Resolve ambiguities
- `/plan` (old) - Create implementation plan
- `/tasks` (old) - Task breakdown

**IMPORTANT:** `/discover` MUST be run before this command.

## If $ARGUMENTS is empty

Use the question tool:

```javascript
question({
  questions: [{
      question: "Describe the feature you want to build.",
      header: "Feature Description"
    }]
})
```

---

## Phase 1: Specification

### Outputs (required)

Create or update:

- `memory/requirements/<feature>/PROBLEM_STATEMENT.md`
- `memory/requirements/<feature>/USER_STORIES.md`
- `memory/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `memory/requirements/<feature>/RISKS.md`

Ensure `memory/requirements/<feature>/` exists before writing.

### Templates

- `.opencode/templates/SDD-ProblemStatement.md`
- `.opencode/templates/SDD-UserStories.md`
- `.opencode/templates/SDD-AcceptanceCriteria.md`
- `.opencode/templates/SDD-Risks.md`

### Workflow

1. Confirm feature name and scope.
2. Create `memory/requirements/<feature>/` if missing.
3. Fill templates for Problem Statement, User Stories, Acceptance Criteria, and Risks.
4. Write prioritized user stories with independent acceptance scenarios.
5. Capture functional requirements and edge cases.
6. Record measurable success criteria (tech-agnostic).
7. Note risks and assumptions.
8. Add Obsidian-compatible links between generated artifacts and include `## Related`.

### Rules

- Do not include tech stack, APIs, or file structure.
- If ambiguity is critical, mark it as `NEEDS CLARIFICATION`.
- Use the question tool for any multi-option decision.

---

## Phase 2: Planning

### Outputs

Create or update:

- `memory/requirements/<feature>/PLAN.md`
- `memory/sprint/Sprint-XX/SPRINT_GOAL.md`
- `memory/sprint/Sprint-XX/BACKLOG.md`
- `memory/sprint/Sprint-XX/RISK_REGISTER.md`

### Workflow

1. Ask user which sprint to use (latest or new):
```javascript
question({
  questions: [{
      question: "Which sprint?",
      header: "Sprint",
      options: [
        { label: "Latest Sprint", description: "Continue most recent work" },
        { label: "New Sprint", description: "Create Sprint-XX" }
      ]
    }]
})
```

2. If no sprint exists, create `memory/sprint/Sprint-01/`.
3. Create PLAN.md with:
   - Technical approach
   - File structure
   - Dependencies
   - Key decisions
4. Create SPRINT_GOAL.md with sprint objective
5. Create BACKLOG.md with user stories
6. Create RISK_REGISTER.md with identified risks

---

## Phase 3: Task Breakdown

### Outputs

Create:
- `memory/sprint/Sprint-XX/TASKS.md`

### Workflow

1. Break down implementation into 5-10 max tasks
2. Each task must have:
   - ID (e.g., `task-01`)
   - Clear description
   - Agent responsible
   - Priority (P0, P1, P2, P3)
   - Dependencies
   - INPUT → OUTPUT → VERIFY criteria

### Task Structure

```markdown
## Tasks

### P0 - Foundation (Critical)
- [ ] task-01: [Description] → Agent: [agent] → Verify: [how to check]

### P1 - Core Backend
- [ ] task-02: [Description] → Agent: [agent] → Verify: [how to check]

### P2 - UI/UX
- [ ] task-03: [Description] → Agent: [agent] → Verify: [how to check]

### P3 - Polish
- [ ] task-04: [Description] → Agent: [agent] → Verify: [how to check]
```

---

## STOP Point

After all phases complete:

```javascript
question({
  questions: [{
      header: "Specification Complete",
      question: "All artifacts created. Proceed to implementation (/create)?",
      options: [
        { label: "Yes, proceed to /create", description: "Start implementing the feature" },
        { label: "Review specification first", description: "Check generated docs" }
      ]
    }]
})
```
