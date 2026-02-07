---
description: Generate executable tasks from the approved plan.
subtask: true
---

# /tasks - Task Breakdown

$ARGUMENTS

## Purpose

Convert the plan into a task list that can be executed by `/impl`.

## Inputs

Read from (MANDATORY):

- `docs/requirements/<feature>/PLAN.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/DATA_CONTRACTS.md` (required if data is involved)
- `docs/requirements/<feature>/contracts/` (required if APIs/events are involved)

If data is involved and `DATA_CONTRACTS.md` is missing, STOP and direct the user to update `/plan`.
If APIs/events are involved and `contracts/` is missing, STOP and direct the user to create it using `.opencode/templates/SDD-Contracts.md`.

## Output

- `docs/sprint/Sprint-XX/TASKS.md`

## Template

- `.opencode/templates/SDD-Tasks.md`

## Required Paths

- Always write tasks to `docs/sprint/Sprint-XX/TASKS.md`
- Include feature reference in the header and link to `docs/requirements/<feature>/PLAN.md`

## Workflow

1. Determine target sprint (ask if unclear).
2. Organize tasks by user story with dependencies and parallel markers.
3. Include exact file paths for each task.
4. Add INPUT -> OUTPUT -> VERIFY for each task where applicable.

## Rules

- Do not implement code.
- Tasks must be independently testable by story.
- Use the question tool for sprint selection or options.
- If required inputs are missing, STOP and direct the user to run `/plan` first.

## STOP POINT

After writing tasks:

> Use the question tool to ask "Tasks recorded in docs/sprint/. Proceed to /impl?"
