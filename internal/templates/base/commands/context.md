---
description: Create a contextual docs pack for this project.
---

# /context - Context Preflight

$ARGUMENTS

## Overview

Generate a verified context pack for the current project, documenting risks, drift, and gaps with file-cited evidence. Outputs are contextual (no frontend/backend docs unless applicable).

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

Criar apenas quando o projeto tiver esse contexto:

- Frontend -> `docs/FRONTEND.md`
- Backend/API -> `docs/BACKEND.md` e/ou `docs/API.md`
- Database -> `docs/DATABASE.md`

## Templates (REQUIRED)

Use these templates as a starting point (adjust content, keep file names):

- `.opencode/templates/DOCS-CONTEXT.md` -> `docs/CONTEXT.md`
- `.opencode/templates/DOCS-QUALITY_GATES.md` -> `docs/QUALITY_GATES.md`
- `.opencode/templates/DOCS-SECURITY.md` -> `docs/SECURITY.md`
- `.opencode/templates/DOCS-ACTION_ITEMS.md` -> `docs/ACTION_ITEMS.md`
- `.opencode/templates/DOCS-FRONTEND.md` -> `docs/FRONTEND.md` (if applicable)
- `.opencode/templates/DOCS-BACKEND.md` -> `docs/BACKEND.md` (if applicable)
- `.opencode/templates/DOCS-API.md` -> `docs/API.md` (if applicable)
- `.opencode/templates/DOCS-DATABASE.md` -> `docs/DATABASE.md` (if applicable)
