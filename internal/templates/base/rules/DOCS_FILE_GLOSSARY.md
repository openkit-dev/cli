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

- `memory/HUB-DOCS.md`
- `memory/GLOSSARY.md`
- `memory/CONTEXT.md`
- `memory/SECURITY.md`
- `memory/QUALITY_GATES.md`
- `memory/ACTION_ITEMS.md`
- `memory/ARCHITECTURE.md`
- `memory/COMMANDS.md`
- `memory/SKILLS.md`
- `memory/WORKFLOW.md`

## Contextual Docs (create only when applicable)

- `memory/FRONTEND.md`
- `memory/BACKEND.md`
- `memory/API.md`
- `memory/DATABASE.md`
- `memory/CHANGELOG.md`
- `memory/MIGRATION_CHECKLIST.md`

## Requirements Docs

- `memory/requirements/HUB-REQUIREMENTS.md`
- `memory/requirements/<feature>/HUB-<FEATURE>.md`
- `memory/requirements/<feature>/PROBLEM_STATEMENT.md`
- `memory/requirements/<feature>/USER_STORIES.md`
- `memory/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `memory/requirements/<feature>/DATA_CONTRACTS.md`
- `memory/requirements/<feature>/RISKS.md`
- `memory/requirements/<feature>/PLAN.md`
- `memory/requirements/<feature>/RESEARCH.md` (optional)
- `memory/requirements/<feature>/QUICKSTART.md` (optional)
- `memory/requirements/<feature>/ANALYSIS.md` (optional)
- `memory/requirements/<feature>/TECH_STACK.md` (optional)
- `memory/requirements/<feature>/CHECKLIST.md` (optional)

## Sprint Docs

- `memory/sprint/HUB-SPRINTS.md`
- `memory/sprint/Sprint-XX/HUB-SPRINT-XX.md`
- `memory/sprint/Sprint-XX/SPRINT_GOAL.md`
- `memory/sprint/Sprint-XX/BACKLOG.md`
- `memory/sprint/Sprint-XX/TASKS.md`
- `memory/sprint/Sprint-XX/RISK_REGISTER.md`

## ADR and Runbooks

- `memory/adr/HUB-ADR.md`
- `memory/adr/ADR_0001_<TITLE>.md`
- `memory/runbooks/HUB-RUNBOOKS.md`
- `memory/runbooks/<SERVICE>.md`

## Rules

- If a command, agent, or skill references a documentation filename, it MUST use this glossary.
- If a new canonical docs file is introduced, update this file in the same change.
