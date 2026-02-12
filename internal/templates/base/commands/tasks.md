---
description: Generate executable tasks from the approved plan (PRIMARY command for TASKS.md creation).
subtask: true
---

# /tasks - Task Breakdown

$ARGUMENTS

## Purpose

Convert the plan into a detailed task list that can be executed by `/impl`.

**IMPORTANT:** This is the PRIMARY and ONLY command that creates `docs/sprint/Sprint-XX/TASKS.md`.
- `/plan` creates high-level artifacts (SPRINT_GOAL.md, BACKLOG.md, RISK_REGISTER.md)
- `/tasks` creates detailed executable tasks with INPUT->OUTPUT->VERIFY criteria

## Inputs

Read from (MANDATORY):

- `docs/requirements/<feature>/PLAN.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/DATA_CONTRACTS.md` (required if data is involved)


If data is involved and `DATA_CONTRACTS.md` is missing, STOP and direct the user to update `/plan`.
If APIs/events are involved and `DATA_CONTRACTS.md` is missing, STOP and direct the user to create it using `.opencode/templates/SDD-Contracts.md`.

## Output

- `docs/sprint/Sprint-XX/TASKS.md`

## Template

- `.opencode/templates/SDD-Tasks.md`

## Required Paths

- Always write tasks to `docs/sprint/Sprint-XX/TASKS.md`
- Include feature reference in the header and link to `docs/requirements/<feature>/PLAN.md`

## Phase 0: TodoList Setup (MANDATORY)

Before any task breakdown work:

1. **Read existing todolist:**
   ```javascript
   todoread()
   ```

2. **Create task breakdown todolist (using standard ID schema):**
   ```javascript
   todowrite({
     todos: [
       {
         id: "sdd-tasks-01-read",
         content: "Read plan and requirements artifacts",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-tasks-02-sprint",
         content: "Determine target sprint",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-tasks-03-breakdown",
         content: "Create detailed task breakdown by story",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-tasks-04-verify",
         content: "Add INPUT->OUTPUT->VERIFY for each task",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-tasks-05-write",
         content: "Write TASKS.md to sprint folder",
         status: "pending",
         priority: "high"
       }
     ]
   })
   ```

3. **Update todolist** as you complete each step

## Workflow

1. Determine target sprint (ask if unclear).
2. Organize tasks by user story with dependencies and parallel markers.
3. Include exact file paths for each task.
4. Add INPUT -> OUTPUT -> VERIFY for each task where applicable.
5. Keep links to requirements and sprint docs in Obsidian format (`[[...]]`).

## Rules

- Do not implement code.
- Tasks must be independently testable by story.
- Use the question tool for sprint selection or options.
- If required inputs are missing, STOP and direct the user to run `/specify` then `/plan` first.
- **This command is the ONLY one that creates TASKS.md** - do not create TASKS.md in other commands.

## STOP POINT

After writing tasks, use the question tool:

```javascript
question({
  questions: [{
    header: "Tasks Complete",
    question: "Tasks recorded in docs/sprint/Sprint-XX/TASKS.md. What would you like to do next?",
    options: [
      { label: "Proceed to /impl", description: "Start implementation with task execution" },
      { label: "Review tasks first", description: "Let me check the task breakdown" },
      { label: "Adjust tasks", description: "Make changes to the tasks" }
    ]
  }]
})
```
