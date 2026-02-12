---
name: obsidian-docs
description: Obsidian-compatible documentation patterns for wikilinks, references, and connected docs graph in OpenKit.
allowed-tools: Read, Glob, Grep
---

# Obsidian Docs

Use this skill when creating or updating documentation in `docs/`.

## Objective

Produce documentation that behaves like an Obsidian knowledge graph:
- easy to navigate with wikilinks
- stable for long-term references
- optimized for retrieval and memory workflows

## Core Rules

1. Use wikilinks for internal docs references.
   - `[[docs/README.md]]`
   - `[[docs/requirements/<feature>/PLAN.md]]`
2. Use heading links for precise references.
   - `[[docs/SECURITY.md#Threats]]`
3. Keep external URLs as standard Markdown links.
   - `[Obsidian Help](https://help.obsidian.md/)`
4. Add `## Related` with meaningful connections.
5. Keep naming and paths stable to avoid link churn.

## Related Section Template

```markdown
## Related

- [[docs/README.md]]
- [[docs/GLOSSARY.md]]
```

## Feature Artifact Related Template

```markdown
## Related

- [[docs/requirements/<feature>/PROBLEM_STATEMENT.md]]
- [[docs/requirements/<feature>/USER_STORIES.md]]
- [[docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md]]
- [[docs/requirements/<feature>/PLAN.md]]
- [[docs/sprint/Sprint-XX/TASKS.md]]
```

## Notes

- Prefer explicit full paths inside the vault.
- Avoid dead links and stale headings.
- Block references are optional and should be used only when line-level traceability is needed.
