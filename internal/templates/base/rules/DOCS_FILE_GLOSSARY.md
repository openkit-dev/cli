---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# DOCUMENTATION FILE GLOSSARY

Canonical documentation filenames for all OpenKit projects.

## Naming Convention

- Documentation files MUST use uppercase snake case: `TUDO_MAIUSCULO.md`.
- Example: `PROBLEM_STATEMENT.md`, `TASKS.md`, `TECH_STACK.md`.

## Core Docs (always available)

- `docs/README.md`
- `docs/GLOSSARY.md`
- `docs/CONTEXT.md`
- `docs/SECURITY.md`
- `docs/QUALITY_GATES.md`
- `docs/ACTION_ITEMS.md`
- `docs/ARCHITECTURE.md`
- `docs/COMMANDS.md`
- `docs/SKILLS.md`
- `docs/WORKFLOW.md`

## Contextual Docs (create only when applicable)

- `docs/FRONTEND.md`
- `docs/BACKEND.md`
- `docs/API.md`
- `docs/DATABASE.md`
- `docs/CHANGELOG.md`
- `docs/MIGRATION_CHECKLIST.md`

## Requirements Docs

- `docs/requirements/README.md`
- `docs/requirements/<feature>/README.md`
- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/DATA_CONTRACTS.md`
- `docs/requirements/<feature>/RISKS.md`
- `docs/requirements/<feature>/PLAN.md`
- `docs/requirements/<feature>/RESEARCH.md` (optional)
- `docs/requirements/<feature>/QUICKSTART.md` (optional)
- `docs/requirements/<feature>/ANALYSIS.md` (optional)
- `docs/requirements/<feature>/TECH_STACK.md` (optional)
- `docs/requirements/<feature>/CHECKLIST.md` (optional)

## Sprint Docs

- `docs/sprint/README.md`
- `docs/sprint/Sprint-XX/README.md`
- `docs/sprint/Sprint-XX/SPRINT_GOAL.md`
- `docs/sprint/Sprint-XX/BACKLOG.md`
- `docs/sprint/Sprint-XX/TASKS.md`
- `docs/sprint/Sprint-XX/RISK_REGISTER.md`

## ADR and Runbooks

- `docs/adr/README.md`
- `docs/adr/ADR_0001_<TITLE>.md`
- `docs/runbooks/README.md`
- `docs/runbooks/<SERVICE>.md`

## Rules

- If a command, agent, or skill references a documentation filename, it MUST use this glossary.
- If a new canonical docs file is introduced, update this file in the same change.
