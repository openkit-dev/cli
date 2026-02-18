# Docs Graph Policy

## Context

This policy standardizes how OpenKit documentation contributes to the Obsidian graph and long-term memory retrieval.

It operationalizes the linking rules referenced in [[requirements/memory-kernel-rust-cli/MEMORY_KERNEL_V1.md]] and enforced by runtime checks in `openkit-rs memory doctor`.

## Rules

1. Inline linking is mandatory on first relevant mention of internal artifacts.
2. `## Related` is mandatory for canonical requirements/sprint/docs artifacts.
3. Canonical filenames and hubs must be preserved for stable graph references.
4. Links should prefer project-relative wikilinks and avoid duplicate aliases.

## Checklist

- [ ] Inline wikilinks exist for key artifact references in each updated file.
- [ ] `## Related` exists and includes principal upstream/downstream artifacts.
- [ ] New docs are discoverable from at least one hub file.
- [ ] `obsidian-link-lint` returns zero broken links.

## Verification

- Runtime: `openkit-rs memory doctor --json --write`
- Graph lint: `obsidian-link-lint` over `docs/`

## Related

- [[requirements/memory-kernel-rust-cli/MEMORY_KERNEL_V1.md]]
- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[sprint/Sprint-07/TASKS.md]]
- [[HUB-DOCS.md]]
