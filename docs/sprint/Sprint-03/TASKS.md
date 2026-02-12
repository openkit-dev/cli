# Sprint 03 Tasks

## Task 1: Port scripts into embedded templates

INPUT: `npx/.opencode/scripts/*`
OUTPUT: `cli/internal/templates/base/scripts/*` (embedded)
VERIFY: scripts contain only standard library and no non-ASCII output

## Task 2: Update CLI README workflows

INPUT: existing `cli/README.md` + NPX workflows
OUTPUT: `cli/README.md` includes concise workflow documentation and verification scripts
VERIFY: README instructions match installed file paths

## Task 3: Verification

INPUT: code and template changes
OUTPUT: passing tests
VERIFY: `go test ./...` in `cli/`

## Related

- [[sprint/Sprint-03/HUB-SPRINT-03.md]]
- [[sprint/Sprint-03/SPRINT_GOAL.md]]
- [[sprint/Sprint-03/BACKLOG.md]]
- [[requirements/opencode-scripts/PLAN.md]]
