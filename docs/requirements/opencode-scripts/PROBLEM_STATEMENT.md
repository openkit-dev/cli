# Problem Statement: Port NPX scripts into CLI templates

## Problem

The CLI embedded templates (`cli/internal/templates/base/**`) reference `.opencode/scripts/*` in multiple places (rules and commands), but the CLI does not currently ship these scripts in the embedded base pack.

Examples of referenced-but-missing scripts:

- `.opencode/scripts/checklist.py`
- `.opencode/scripts/verify_all.py`
- `.opencode/scripts/auto_preview.py`
- `.opencode/scripts/session_manager.py`

This creates broken workflows after `openkit opencode sync` (documentation and commands instruct users to run scripts that are not present).

## Goal

Port the generic, standard-library Python scripts from the legacy NPX package into the Go CLI embedded templates so they are installed as part of the OpenCode target.

## Non-goals

- Do not add Python dependencies.
- Do not make the scripts agent-specific beyond their `.opencode/` location.
- Do not change CLI sync semantics.

## Related

- [[requirements/opencode-scripts/HUB-OPENCODE-SCRIPTS.md]]
- [[requirements/opencode-scripts/ACCEPTANCE_CRITERIA.md]]
- [[requirements/opencode-scripts/PLAN.md]]
- [[sprint/Sprint-03/HUB-SPRINT-03.md]]
