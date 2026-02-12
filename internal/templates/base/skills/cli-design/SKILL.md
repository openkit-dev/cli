---
name: cli-design
description: CLI architecture and command UX patterns. Use for command design, flags, output contracts, and interactive/non-interactive flows.
allowed-tools: Read, Write, Edit, Bash
---

# CLI Design Patterns

## Core Principles

- Commands are predictable, composable, and script-friendly.
- Non-interactive mode must always be available.
- Output contracts are stable (`text` for humans, `json` for machines).
- Errors are actionable (what failed, why, how to fix).

## Command Shape

- Use `tool <resource> <verb>` for grouped commands.
- Keep top-level verbs small (`init`, `sync`, `check`, `upgrade`, `context`).
- Prefer explicit flags over hidden behavior.

## Flag Rules

- `--yes` for non-interactive confirmation bypass.
- `--type` for forcing detection when needed.
- `--overlays` for explicit feature augmentation.
- Boolean flags default to safe behavior.

## Input and Output

- Inputs are validated early with clear errors.
- Support machine output format when command is automatable.
- Print compact success summaries and deterministic paths.

## Safety

- Never execute destructive actions implicitly.
- Require confirmation for risky operations unless `--yes` is present.
- Detect ambiguity and ask for explicit selection.

## Testing Checklist

- Unit tests for parser/flag behavior.
- Integration tests for interactive and non-interactive runs.
- Snapshot checks for output stability on major commands.
