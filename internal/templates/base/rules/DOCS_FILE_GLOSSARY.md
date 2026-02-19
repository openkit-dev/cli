---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# DOCUMENTATION FILE GLOSSARY

Canonical documentation filenames for all OpenKit projects.

## Naming Convention

- Documentation files MUST use canonical uppercase naming.
- Artifact files use uppercase snake case (for example `PROBLEM_STATEMENT.md`, `TASKS.md`, `TECH_STACK.md`).
- Hub files use the `HUB-<RESOURCE>.md` convention (for example `HUB-DOCS.md`, `HUB-SPRINT-XX.md`).

## Core Docs (always available)

- `openkit-memory/HUB-DOCS.md`
- `openkit-memory/GLOSSARY.md`
- `openkit-memory/CONTEXT.md`
- `openkit-memory/SECURITY.md`
- `openkit-memory/QUALITY_GATES.md`
- `openkit-memory/ACTION_ITEMS.md`
- `openkit-memory/ARCHITECTURE.md`
- `openkit-memory/COMMANDS.md`
- `openkit-memory/SKILLS.md`
- `openkit-memory/WORKFLOW.md`

## Contextual Docs (create only when applicable)

- `openkit-memory/FRONTEND.md`
- `openkit-memory/BACKEND.md`
- `openkit-memory/API.md`
- `openkit-memory/DATABASE.md`
- `openkit-memory/CHANGELOG.md`
- `openkit-memory/MIGRATION_CHECKLIST.md`

## Requirements Docs

- `openkit-memory/requirements/HUB-REQUIREMENTS.md`
- `openkit-memory/requirements/<feature>/HUB-<FEATURE>.md`
- `openkit-memory/requirements/<feature>/PROBLEM_STATEMENT.md`
- `openkit-memory/requirements/<feature>/USER_STORIES.md`
- `openkit-memory/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `openkit-memory/requirements/<feature>/DATA_CONTRACTS.md`
- `openkit-memory/requirements/<feature>/RISKS.md`
- `openkit-memory/requirements/<feature>/PLAN.md`
- `openkit-memory/requirements/<feature>/RESEARCH.md` (optional)
- `openkit-memory/requirements/<feature>/QUICKSTART.md` (optional)
- `openkit-memory/requirements/<feature>/ANALYSIS.md` (optional)
- `openkit-memory/requirements/<feature>/TECH_STACK.md` (optional)
- `openkit-memory/requirements/<feature>/CHECKLIST.md` (optional)

## Sprint Docs

- `openkit-memory/sprint/HUB-SPRINTS.md`
- `openkit-memory/sprint/Sprint-XX/HUB-SPRINT-XX.md`
- `openkit-memory/sprint/Sprint-XX/SPRINT_GOAL.md`
- `openkit-memory/sprint/Sprint-XX/BACKLOG.md`
- `openkit-memory/sprint/Sprint-XX/TASKS.md`
- `openkit-memory/sprint/Sprint-XX/RISK_REGISTER.md`

## ADR and Runbooks

- `openkit-memory/adr/HUB-ADR.md`
- `openkit-memory/adr/ADR_0001_<TITLE>.md`
- `openkit-memory/runbooks/HUB-RUNBOOKS.md`
- `openkit-memory/runbooks/<SERVICE>.md`

## Rules

- If a command, agent, or skill references a documentation filename, it MUST use this glossary.
- If a new canonical docs file is introduced, update this file in the same change.
