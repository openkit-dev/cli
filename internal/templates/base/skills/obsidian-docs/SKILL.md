---
name: obsidian-docs
description: Obsidian-compatible documentation patterns for wikilinks, references, and connected docs graph in OpenKit.
allowed-tools: Read, Glob, Grep
---

# Obsidian Docs

Use this skill when creating or updating documentation in `openkit-memory/`.

## Objective

Produce documentation that behaves like an Obsidian knowledge graph:
- easy to navigate with wikilinks
- stable for long-term references
- optimized for retrieval and memory workflows

## Core Rules

1. Use wikilinks for internal docs references.
   - `[[HUB-DOCS.md]]`
   - `[[requirements/<feature>/PLAN.md]]`
2. Use heading links for precise references.
   - `[[SECURITY.md#Threats]]`
3. Keep external URLs as standard Markdown links.
   - `[Obsidian Help](https://help.obsidian.md/)`
4. Add `## Related` with meaningful connections.
5. Keep naming and paths stable to avoid link churn.

## Hub Sections (Mandatory)

For every hub note (`HUB-*.md`), include both sections:

1. `## Context`
   - Briefly explain what the hub represents and when to use it.
2. `## Status Overview`
   - Provide a concise operational snapshot (phase, completion, blockers, readiness, or risk level).

Suggested minimal status format:

```markdown
## Status Overview

- Phase: [Planning | Execution | Verification | Done]
- Completion: [0-100]%
- Blockers: [None or short note]
```

## Related Section Template

```markdown
## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
```

## Feature Artifact Related Template

```markdown
## Related

- [[requirements/<feature>/PROBLEM_STATEMENT.md]]
- [[requirements/<feature>/USER_STORIES.md]]
- [[requirements/<feature>/ACCEPTANCE_CRITERIA.md]]
- [[requirements/<feature>/PLAN.md]]
- [[sprint/Sprint-XX/TASKS.md]]
```

## Notes

- Prefer explicit full paths inside the vault.
- Avoid dead links and stale headings.
- Block references are optional and should be used only when line-level traceability is needed.
