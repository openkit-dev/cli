# Migration Guide: Context-Aware OpenKit

## Overview

OpenKit now adapts to project type instead of assuming `backend/` + `frontend/` layout for all repositories.

## What Changed

- Added detection engine in `internal/detection/`.
- Refactored `context` command in `internal/cli/context.go`.
- Added project type definitions in `.opencode/project-types/`.
- Added overlays in `.opencode/overlays/`.
- Added new base skills in `internal/templates/base/skills/`:
  - `cli-design`
  - `library-patterns`
  - `desktop-patterns`
  - `serverless-patterns`
  - `iac-patterns`

## Behavior Changes in `openkit context`

- Auto-detects project type with evidence.
- Supports non-interactive mode via flags:
  - `--yes`
  - `--type`
  - `--overlays`
- Generates only relevant docs for detected type.

## Backward Compatibility

- Web projects still generate `docs/BACKEND.md` and `docs/FRONTEND.md`.
- Existing workflows remain valid.

## Recommended Upgrade Steps

1. Pull latest OpenKit changes.
2. Run `go test ./...`.
3. Run `openkit context --yes` in a target project.
4. Validate generated docs for project relevance.

## Rollback

If needed, pin to previous OpenKit release and restore previous generated docs from git history.
