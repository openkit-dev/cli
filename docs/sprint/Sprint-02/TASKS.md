# Sprint 02 Tasks

## Task 1: TOML command generator

INPUT: embedded `base/commands/*.md`
OUTPUT: `.gemini/commands/openkit/*.toml` desired files
VERIFY: unit test for frontmatter parsing + prompt escaping

STATUS: done
EVIDENCE: `cli/internal/targets/gemini_commands.go`, `cli/internal/targets/gemini_commands_test.go`

## Task 2: Gemini target file mapping

INPUT: embedded `base/rules/*`, `base/skills/**`, generated TOMLs
OUTPUT: `openkit gemini sync` installs commands, rules, skills, and updates `GEMINI.md`
VERIFY: `openkit gemini sync --dry-run` shows deterministic plan and correct paths

STATUS: done
EVIDENCE: `cli/internal/targets/targets.go`

## Task 3: Documentation alignment

INPUT: new Gemini target behavior
OUTPUT: updated `cli/docs/agent-compat/agents/gemini.md`
VERIFY: doc matches actual sync paths

STATUS: done
EVIDENCE: `cli/docs/agent-compat/agents/gemini.md`

## Task 4: Verification

INPUT: code changes
OUTPUT: passing Go tests
VERIFY: `go test ./...` in `cli/`

STATUS: done
EVIDENCE: `go test ./...` (package `github.com/openkit-devtools/openkit/internal/targets` is passing)

## Related

- [[sprint/Sprint-02/HUB-SPRINT-02.md]]
- [[sprint/Sprint-02/SPRINT_GOAL.md]]
- [[sprint/Sprint-02/BACKLOG.md]]
- [[requirements/gemini-sync/PLAN.md]]
