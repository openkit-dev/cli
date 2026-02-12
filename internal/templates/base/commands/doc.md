---
description: Write or update technical docs in docs/ for this project.
---

# /doc - Docs Writer

$ARGUMENTS

## Overview

Create clear, actionable documentation based on code changes and context pack inputs.

## Workflow

1. Update only `docs/` files; reference code paths without editing them.
2. Ensure API indexes, frontend routes, and data fetching docs are updated when relevant.
3. Add Action Items for inconsistencies or cross-repo impacts.
4. Use Obsidian-compatible links for internal docs references (`[[docs/...]]`).
5. Add a `## Related` section with at least two useful internal links.
6. Follow canonical docs filenames from `.opencode/rules/DOCS_FILE_GLOSSARY.md`.

## Output Requirements

- Provide docs updates and docstring insertion notes as needed.
- Use tables and examples where needed for clarity.
- When presenting user choices, use a `question` tool with proper structure:

```javascript
question({
  questions: [{
      question: "What documentation type?",
      header: "Type",
      options: [
        { label: "API Docs", description: "Endpoints and schemas" },
        { label: "User Guide", description: "Usage and examples" }
      ]
    }]
})
```

## Key Documentation Types

- **README/How-to**: installation, envs, build, run, migrations, workers.
- **API**: update `docs/API.md` (method, route, schemas, auth) and ensure OpenAPI is current.
- **Front-end**: document routes (TanStack), hooks, complex components, UI patterns (shadcn) and accessibility.
- **Security**: usage warnings (auth/roles, uploads, limits, privacy).
- **ADRs**: `docs/adr/ADR_0001_<TITLE>.md` (decision, options, trade-offs).
- **Runbooks**: `docs/runbooks/<SERVICE>.md` (alarm -> diagnosis -> action).
- **Glossary**: update `docs/GLOSSARY.md` with domain-specific terms introduced in the feature.
- **Migration**: use `docs/MIGRATION_CHECKLIST.md` when upgrading legacy docs to Obsidian format.

## Glossary Updates

When documenting a new feature:

1. **Check for new terms:** Does this feature introduce domain-specific terminology?
2. **Update glossary:** Add new terms to `docs/GLOSSARY.md`
3. **Link references:** In feature docs, link to `docs/GLOSSARY.md` for standard terms
4. **Template sections:** All SDD templates include a Glossary section for feature-specific terms

**Standard Reference:**
> For standard terminology definitions, see `docs/GLOSSARY.md`
