---
name: library-patterns
description: Library and SDK design patterns. Use for public API stability, semantic versioning, and migration guidance.
allowed-tools: Read, Write, Edit
---

# Library Patterns

## API Surface

- Keep public API small and intentional.
- Separate public contracts from internal helpers.
- Avoid leaking implementation details in exported types.

## Versioning

- Follow semantic versioning strictly.
- Breaking changes require MAJOR version bumps.
- Deprecate first, remove later (document timeline).

## Compatibility

- Provide migration notes for every breaking change.
- Keep defaults backwards-compatible when possible.
- Add compatibility tests for critical public functions.

## Documentation

- Maintain `PUBLIC_API.md` with examples.
- Maintain `VERSIONING.md` and `BREAKING_CHANGES.md`.
- Include minimal but complete quick-start snippets.

## Error Contracts

- Return stable, typed errors when possible.
- Keep error messages deterministic for automation.
- Do not change error shapes in patch releases.
