# Sprint Goal - Sprint 05

**Created**: 2026-02-09 13:05 (UTC-3)
**Duration**: 2 weeks
**Status**: Planning

## Goal Statement

Implement semantic memory plugin for OpenCode/OpenKit that enables persistent context across agent sessions through local-first vector storage and automatic knowledge extraction.

## Success Criteria

1. [ ] `--memory` flag functional in `openkit init` and `openkit sync` commands
2. [ ] Plugin structure installed in `.opencode/plugins/semantic-memory/`
3. [ ] LanceDB storage operational with embeddings
4. [ ] Session hooks (session.idle, session.compacting) extracting and injecting context
5. [ ] CLI commands `/memory status` and `/memory inspect` working
6. [ ] Zero external dependencies (fully local-first)
7. [ ] Documentation complete with usage examples

## Scope

### In Scope
- Go CLI modifications for `--memory` flag
- TypeScript plugin with LanceDB integration
- ONNX embeddings (nomic-embed-text) integration
- Basic extraction heuristics (decisions, patterns, errors)
- Progressive disclosure for context injection
- TTL and garbage collection logic
- CLI memory management commands

### Out of Scope
- Web UI (deferred to v0.2)
- Re-ranking (deferred to v0.3)
- Global memory across projects (deferred to v0.3)
- Compression via LLM API
- FTS5 hybrid search (deferred to v0.2)

## Key Deliverables

1. **Go CLI Changes**
   - `internal/cli/init.go` - add `--memory` flag
   - `internal/cli/agent_targets.go` - add `--memory` flag to sync
   - `internal/templates/memory/` - embedded plugin templates

2. **TypeScript Plugin**
   - `.opencode/plugins/semantic-memory/index.ts` - main plugin
   - `.opencode/plugins/semantic-memory/lib/memory.ts` - core logic
   - `.opencode/plugins/semantic-memory/lib/embeddings.ts` - ONNX integration
   - `.opencode/plugins/semantic-memory/lib/storage.ts` - LanceDB wrapper

3. **Configuration**
   - `.opencode/memory/config.json` - default settings
   - `opencode.json` updates - enable memory

4. **Documentation**
   - README updates with `--memory` usage
   - Plugin configuration guide
   - Troubleshooting section

## Risks

| Risk | Mitigation |
|------|------------|
| ONNX model size (~500MB) | Lazy download, cache in ~/.cache/opencode/models/ |
| LanceDB platform compatibility | Test on Linux, macOS, Windows; fallback to JSON |
| Hook integration complexity | Start with minimal hooks (session.idle only) |
| Performance impact on OpenCode startup | Async initialization, timeout guards |

## Definition of Done

- All acceptance criteria met
- Unit tests passing (>= 80% coverage where applicable)
- Integration tests with OpenCode successful
- Documentation reviewed and approved
- No critical bugs
- Code review complete
- Sprint retrospective conducted

## Related

- [[docs/sprint/Sprint-05/HUB-SPRINT-05.md]]
- [[docs/sprint/Sprint-05/BACKLOG.md]]
- [[docs/sprint/Sprint-05/TASKS.md]]
- [[docs/requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
