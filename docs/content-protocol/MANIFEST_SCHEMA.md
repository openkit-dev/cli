# Content Pack Manifest Schema

This document defines a concrete JSON manifest for an OpenKit content pack. The manifest is stored at `manifest.json` at the pack root.

## Field-Level Definition (Normative)

Top-level fields:
- `manifest_version` (string): schema version for this manifest document (e.g. `"1"`).
- `pack` (object): identity and version.
- `compat` (object): supported agents and minimum versions.
- `artifacts` (array): canonical artifact catalog.
- `targets` (array): per-agent rendering instructions.
- `migrations` (object, optional): renames/deprecations and notes.

### `pack`

- `id` (string): stable pack identifier, e.g. `"openkit-core"`.
- `version` (string): SemVer for the pack (e.g. `"1.4.0"`).
- `description` (string, optional)
- `license` (string, optional)
- `homepage` (string, optional)

### `compat`

- `agents` (array): supported agent targets.
  - `name` (enum): `"opencode" | "claude" | "gemini" | "codex"`
  - `min_agent_version` (string, optional): minimum agent runtime version
  - `min_cli_version` (string, optional): minimum OpenKit CLI version required to sync
  - `notes` (string, optional)

### `artifacts[]`

Each artifact entry:
- `id` (string): stable canonical ID.
- `type` (enum): `rule | command | prompt | skill | template | agent | workflow`.
- `source` (string): path within pack root (must be under `artifacts/`).
- `sha256` (string): hex digest of the artifact bytes.
- `metadata` (object, optional):
  - `description` (string, optional)
  - `tags` (array of strings, optional)
  - `language` (string, optional): e.g. `"en"`, `"pt-BR"`
  - `sensitive` (boolean, optional): indicates higher review needs

### `targets[]`

Defines how artifacts render into each agent layout.

Target fields:
- `agent` (enum): `opencode | claude | gemini | codex`
- `artifact_id` (string): references `artifacts[].id`
- `output_path` (string): relative path from project root (no `..`)
- `mode` (enum): `copy | render | template`
- `constraints` (object, optional): agent-specific constraints
  - `max_bytes` (integer, optional)
  - `requires_trust` (boolean, optional)
  - `format` (string, optional): e.g. `"md"`, `"toml"`, `"json"`
- `render` (object, optional): required when `mode=render` or `mode=template`
  - `engine` (string): named renderer (e.g. `"md-section-append"`, `"toml-command"`)
  - `inputs` (object, optional): limited variable map

### `migrations` (optional)

- `renames` (array, optional):
  - `from_output_path` (string)
  - `to_output_path` (string)
  - `since` (string): pack version introducing rename
- `deprecated` (array, optional):
  - `output_path` (string)
  - `since` (string)
  - `remove_after` (string, optional)

## JSON Schema (Draft 2020-12)

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://openkit.dev/schemas/content-pack-manifest.v1.json",
  "type": "object",
  "required": ["manifest_version", "pack", "compat", "artifacts", "targets"],
  "additionalProperties": false,
  "properties": {
    "manifest_version": {"type": "string", "const": "1"},
    "pack": {
      "type": "object",
      "required": ["id", "version"],
      "additionalProperties": false,
      "properties": {
        "id": {"type": "string", "minLength": 1},
        "version": {"type": "string", "minLength": 1},
        "description": {"type": "string"},
        "license": {"type": "string"},
        "homepage": {"type": "string"}
      }
    },
    "compat": {
      "type": "object",
      "required": ["agents"],
      "additionalProperties": false,
      "properties": {
        "agents": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "required": ["name"],
            "additionalProperties": false,
            "properties": {
              "name": {"enum": ["opencode", "claude", "gemini", "codex"]},
              "min_agent_version": {"type": "string"},
              "min_cli_version": {"type": "string"},
              "notes": {"type": "string"}
            }
          }
        }
      }
    },
    "artifacts": {
      "type": "array",
      "minItems": 1,
      "items": {
        "type": "object",
        "required": ["id", "type", "source", "sha256"],
        "additionalProperties": false,
        "properties": {
          "id": {"type": "string", "minLength": 1},
          "type": {"enum": ["rule", "command", "prompt", "skill", "template", "agent", "workflow"]},
          "source": {
            "type": "string",
            "pattern": "^artifacts/"
          },
          "sha256": {
            "type": "string",
            "pattern": "^[a-f0-9]{64}$"
          },
          "metadata": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
              "description": {"type": "string"},
              "tags": {"type": "array", "items": {"type": "string"}},
              "language": {"type": "string"},
              "sensitive": {"type": "boolean"}
            }
          }
        }
      }
    },
    "targets": {
      "type": "array",
      "minItems": 1,
      "items": {
        "type": "object",
        "required": ["agent", "artifact_id", "output_path", "mode"],
        "additionalProperties": false,
        "properties": {
          "agent": {"enum": ["opencode", "claude", "gemini", "codex"]},
          "artifact_id": {"type": "string", "minLength": 1},
          "output_path": {
            "type": "string",
            "minLength": 1,
            "not": {"pattern": "\\.\\."}
          },
          "mode": {"enum": ["copy", "render", "template"]},
          "constraints": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
              "max_bytes": {"type": "integer", "minimum": 0},
              "requires_trust": {"type": "boolean"},
              "format": {"type": "string"}
            }
          },
          "render": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
              "engine": {"type": "string"},
              "inputs": {"type": "object"}
            }
          }
        }
      }
    },
    "migrations": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "renames": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["from_output_path", "to_output_path", "since"],
            "additionalProperties": false,
            "properties": {
              "from_output_path": {"type": "string"},
              "to_output_path": {"type": "string"},
              "since": {"type": "string"}
            }
          }
        },
        "deprecated": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["output_path", "since"],
            "additionalProperties": false,
            "properties": {
              "output_path": {"type": "string"},
              "since": {"type": "string"},
              "remove_after": {"type": "string"}
            }
          }
        }
      }
    }
  }
}
```

## Minimal Example

```json
{
  "manifest_version": "1",
  "pack": {"id": "openkit-core", "version": "1.0.0"},
  "compat": {"agents": [{"name": "opencode", "min_cli_version": "0.9.0"}]},
  "artifacts": [
    {
      "id": "rules.master",
      "type": "rule",
      "source": "artifacts/rules/master.md",
      "sha256": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
      "metadata": {"description": "Baseline safety rules", "sensitive": true}
    }
  ],
  "targets": [
    {
      "agent": "opencode",
      "artifact_id": "rules.master",
      "output_path": ".opencode/rules/MASTER.md",
      "mode": "copy",
      "constraints": {"format": "md"}
    }
  ]
}
```

## Related

- [[content-protocol/README.md]]
- [[content-protocol/PROTOCOL.md]]
- [[content-protocol/SYNC_SEMANTICS.md]]
