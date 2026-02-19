---
description: Analyze project context and generate documentation (MANDATORY before /specify)
---

# /discover - Project Discovery

$ARGUMENTS

## Overview

Generate a verified context pack for the current project, documenting risks, drift, and gaps with file-cited evidence. Outputs are contextual (no frontend/backend docs unless applicable).

**IMPORTANT:** This command is MANDATORY before starting any feature work.

**Discovery Gate Flow:**
```
/discover (MANDATORY) → /specify → /create → /verify → /deploy
```

## Workflow

1. Discover structure/configs with `rg` and directory listings; record evidence with file paths and short snippets.
2. Map FE routing/data fetching and BE endpoints/models/migrations; mark missing items as "not found."
3. Diagnose risks (security, config, drift, missing tests, missing configuration files, logger/correlation-id).
4. Produce/update required `memory/` files with tables and concise summaries.
5. If external/network tools are needed (pip-audit, npm audit), note blockers in `memory/ACTION_ITEMS.md`.

## Output Requirements

- Create or update the full context pack (see list below).
- Cite evidence in each doc with explicit file paths.
- Put cross-repo impacts in `memory/ACTION_ITEMS.md` with severity and owner.
- Use Obsidian-compatible links (`[[...]]`) for all internal documentation references.
- Add `## Related` section in each generated docs artifact where applicable.
- Enforce canonical docs filenames from `.opencode/rules/DOCS_FILE_GLOSSARY.md`.
- When presenting user choices, use a `question` tool with proper structure:

```javascript
question({
  questions: [{
      question: "What context scope?",
      header: "Scope",
      options: [
        { label: "Full", description: "Complete context pack" },
        { label: "Backend Only", description: "API and database only" },
        { label: "Frontend Only", description: "Routes and components only" }
      ]
    }]
})
```

## Context Pack Files (Contextual)

Always create/update:

- `memory/CONTEXT.md` (executive summary + overview + evidence)
- `memory/QUALITY_GATES.md` (linters, tests, CI, checks)
- `memory/SECURITY.md` (threats, controls, gaps, prioritized actions)
- `memory/ACTION_ITEMS.md` (backlog prioritized by impact x effort)
- `memory/HUB-DOCS.md` (documentation hub)
- `memory/GLOSSARY.md` (shared terminology)
- `memory/requirements/HUB-REQUIREMENTS.md` (requirements hub)
- `memory/sprint/HUB-SPRINTS.md` (sprint hub)

Create only when project has this context:

- Frontend → `memory/FRONTEND.md`
- Backend/API → `memory/BACKEND.md` and/or `memory/API.md`
- Database → `memory/DATABASE.md`

## Templates (REQUIRED)

Use these templates:

- `.opencode/templates/DOCS-CONTEXT.md` → `memory/CONTEXT.md`
- `.opencode/templates/DOCS-README.md` → `memory/HUB-DOCS.md`
- `.opencode/templates/DOCS-GLOSSARY.md` → `memory/GLOSSARY.md`
- `.opencode/templates/DOCS-QUALITY_GATES.md` → `memory/QUALITY_GATES.md`
- `.opencode/templates/DOCS-SECURITY.md` → `memory/SECURITY.md`
- `.opencode/templates/DOCS-ACTION_ITEMS.md` → `memory/ACTION_ITEMS.md`
- `.opencode/templates/DOCS-FRONTEND.md` → `memory/FRONTEND.md` (if applicable)
- `.opencode/templates/DOCS-BACKEND.md` → `memory/BACKEND.md` (if applicable)
- `.opencode/templates/DOCS-API.md` → `memory/API.md` (if applicable)
- `.opencode/templates/DOCS-DATABASE.md` → `memory/DATABASE.md` (if applicable)

---

## STOP Point

After discovery is complete, use the question tool:

```javascript
question({
  questions: [{
      header: "Discovery Complete",
      question: "Context generated. Proceed to specification (/specify)?",
      options: [
        { label: "Yes, proceed to /specify", description: "Start specifying the feature" },
        { label: "Review context first", description: "Check generated docs" }
      ]
    }]
})
```
