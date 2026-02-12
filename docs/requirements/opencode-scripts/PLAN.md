# Implementation Plan: Port NPX scripts into CLI templates

## Steps

1. Copy the scripts from `npx/.opencode/scripts/` into `cli/internal/templates/base/scripts/`:
   - `checklist.py`
   - `verify_all.py`
   - `auto_preview.py`
   - `session_manager.py`

2. Make small portability fixes while keeping behavior intact:
   - Use `sys.executable` instead of hardcoding `python`.
   - Remove non-ASCII status markers from output.

3. Update `cli/README.md` to document the primary workflows (condensed).

## Verification

- `go test ./...` in `cli/`.

## Related

- [[docs/requirements/opencode-scripts/HUB-OPENCODE-SCRIPTS.md]]
- [[docs/requirements/opencode-scripts/PROBLEM_STATEMENT.md]]
- [[docs/content-protocol/PROTOCOL.md]]
- [[docs/sprint/Sprint-03/TASKS.md]]
