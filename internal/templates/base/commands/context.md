---
description: Create a contextual docs pack for this project (MANDATORY before /specify).
---

# /context - Context Preflight (MANDATORY)

$ARGUMENTS

## Overview

Generate a verified context pack for the current project, documenting risks, drift, and gaps with file-cited evidence. Outputs are contextual (no frontend/backend docs unless applicable).

**IMPORTANT:** This command is MANDATORY before starting the SDD workflow.
- `/context` generates technical context (CONTEXT.md, SECURITY.md, etc.)
- `/brainstorm` is OPTIONAL and used for exploring options when scope is unclear

**Discovery Gate Flow:**
```
/context (MANDATORY) → /specify → /clarify → /plan → /tasks → /impl
    ↑
/brainstorm (OPTIONAL, when scope unclear)
```

## Workflow

1. Discover structure/configs with `rg` and directory listings; record evidence with file paths and short snippets.
2. Map FE routing/data fetching and BE endpoints/models/migrations; mark missing items as "not found."
3. Diagnose risks (security, config, drift, missing tests, missing configuration files, logger/correlation-id).
4. Produce/update required `docs/` files with tables and concise summaries.
5. If external/network tools are needed (pip-audit, npm audit, shadcn llms.txt), note blockers in `docs/ACTION_ITEMS.md`.

## Output Requirements

- Create or update the full context pack (see list below).
- Cite evidence in each doc with explicit file paths.
- Put cross-repo impacts in `docs/ACTION_ITEMS.md` with severity and owner.
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

Sempre criar/atualizar:

- `docs/CONTEXT.md` (executive summary + overview + evidencias)
- `docs/QUALITY_GATES.md` (linters, tests, CI, checks)
- `docs/SECURITY.md` (threats, controls, gaps, prioritized actions)
- `docs/ACTION_ITEMS.md` (backlog priorizado por impacto x effort)
- `docs/README.md` (documentation hub)
- `docs/GLOSSARY.md` (shared terminology)
- `docs/requirements/README.md` (requirements hub)
- `docs/sprint/README.md` (sprint hub)
- `docs/MIGRATION_CHECKLIST.md` (legacy docs migration, when needed)

Criar apenas quando o projeto tiver esse contexto:

- Frontend -> `docs/FRONTEND.md`
- Backend/API -> `docs/BACKEND.md` e/ou `docs/API.md`
- Database -> `docs/DATABASE.md`

## Templates (REQUIRED)

Use these templates as a starting point (adjust content, keep file names):

- `.opencode/templates/DOCS-CONTEXT.md` -> `docs/CONTEXT.md`
- `.opencode/templates/DOCS-README.md` -> `docs/README.md`
- `.opencode/templates/DOCS-GLOSSARY.md` -> `docs/GLOSSARY.md`
- `.opencode/templates/DOCS-MIGRATION_CHECKLIST.md` -> `docs/MIGRATION_CHECKLIST.md` (if needed)
- `.opencode/templates/DOCS-QUALITY_GATES.md` -> `docs/QUALITY_GATES.md`
- `.opencode/templates/DOCS-SECURITY.md` -> `docs/SECURITY.md`
- `.opencode/templates/DOCS-ACTION_ITEMS.md` -> `docs/ACTION_ITEMS.md`
- `.opencode/templates/DOCS-FRONTEND.md` -> `docs/FRONTEND.md` (if applicable)
- `.opencode/templates/DOCS-BACKEND.md` -> `docs/BACKEND.md` (if applicable)
- `.opencode/templates/DOCS-API.md` -> `docs/API.md` (if applicable)
- `.opencode/templates/DOCS-DATABASE.md` -> `docs/DATABASE.md` (if applicable)
