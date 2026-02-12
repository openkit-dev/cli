# OpenKit Content Protocol (Canonical)

This protocol defines how agent content is packaged, identified, versioned, and rendered into agent-specific layouts. It describes a content pack format, not a repository layout.

## Canonical Artifact Types

Artifacts are addressed by a stable `id` (string) and a `type` (enum). A pack MAY ship multiple variants that render to different targets.

Canonical `type` values:
- `rule` (policy/instructions that constrain behavior)
- `command` (a named action; may be Markdown/TOML depending on target)
- `prompt` (top-level agent/system prompt)
- `skill` (a reusable capability bundle; usually a directory rooted at `SKILL.md`)
- `template` (scaffold assets; opaque)
- `agent` (agent/subagent definition; if a target runtime supports it)
- `workflow` (multi-step procedure documentation; opaque to sync)

## Canonical Pack Layout

Within the compressed pack (`.tar.gz`), the canonical layout is:

```
<pack-root>/
  manifest.json
  checksums.sha256
  artifacts/
    rules/
    commands/
    prompts/
    skills/
      <skill-name>/
        SKILL.md
        ...
    templates/
    agents/
    workflows/
```

Notes:
- `artifacts/**` contains the raw bytes. Packs MUST NOT rely on external downloads.
- `checksums.sha256` covers `manifest.json` and all files under `artifacts/`.

## Naming Rules and IDs

- `id` MUST be stable across versions.
- `id` SHOULD be filename-safe: `[a-z0-9][a-z0-9._-]*`.
- IDs SHOULD be namespaced when reused across ecosystems, e.g. `opencode.core.master-rules`.
- Skills use the directory name as the canonical skill name (`skills/<name>/...`).

## Versioning Semantics

- Packs use Semantic Versioning: `MAJOR.MINOR.PATCH`.
- `MAJOR`: breaking changes to canonical IDs, manifest schema, or sync behavior expectations.
- `MINOR`: additive artifacts or additive targets; backward compatible.
- `PATCH`: fixes to artifact content or metadata without changing meaning.

## Cross-Agent Rendering Model

The same canonical artifact can be mapped to multiple agents via per-agent renderer targets:
- `copy`: exact bytes copied to a target path
- `render`: transform a canonical artifact into a target format/layout (e.g., render rules into a single entrypoint file)
- `template`: treat artifact as a template with a limited variable model (no arbitrary code execution)

Rendering MUST be deterministic given:
- pack bytes
- manifest
- explicit render inputs (agent target, project root, and explicit variables)

## Pack Publishing Format

Required:
- Distribution as a single `tar.gz` file
- `manifest.json`
- `checksums.sha256` with `sha256` for all required files

Optional (recommended for supply-chain integrity):
- Detached signature (e.g. `manifest.sig`)
- Provenance metadata (signing key ID, build timestamp)

## Related

- [[content-protocol/README.md]]
- [[content-protocol/MANIFEST_SCHEMA.md]]
- [[content-protocol/MANAGED_STATE_SCHEMA.md]]
- [[content-protocol/SYNC_SEMANTICS.md]]
