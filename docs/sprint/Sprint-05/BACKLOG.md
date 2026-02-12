# Sprint Backlog - Sprint 05

**Created**: 2026-02-09 13:05 (UTC-3)
**Sprint Goal**: Implement semantic memory plugin for OpenCode/OpenKit

## Story Status Legend
- `[ ]` Not Started
- `[>]` In Progress
- `[x]` Complete
- `[~]` Blocked

---

## P0 - Foundation (Critical Path)

### US-001: CLI Flag for Memory Installation
**Status**: `[ ]` | **Points**: 5 | **Owner**: TBD

**User Story**: As a developer, I want to use `--memory` flag with `openkit init` to automatically install the memory plugin, so that I can enable persistent context without manual setup.

**Acceptance Criteria**:
- [ ] `openkit init my-app --memory` creates `.opencode/plugins/semantic-memory/`
- [ ] Default config.json created in `.opencode/memory/`
- [ ] opencode.json updated with `memory.enabled: true`
- [ ] .gitignore created to exclude vector database files

**Tasks**:
- [ ] Add `--memory` flag to `internal/cli/init.go`
- [ ] Implement `installMemoryPlugin()` function
- [ ] Create embedded templates in `internal/templates/memory/`
- [ ] Update opencode.json programmatically

**Dependencies**: None

---

### US-002: LanceDB Storage Integration
**Status**: `[ ]` | **Points**: 8 | **Owner**: TBD

**User Story**: As the memory plugin, I want to store embeddings in LanceDB, so that I can perform fast vector similarity searches locally.

**Acceptance Criteria**:
- [ ] LanceDB initializes without errors
- [ ] Can write memories with vectors (Float32[768])
- [ ] Can query by similarity with threshold
- [ ] Can filter by project, type, date
- [ ] Database persists in `.opencode/memory/index.lance/`

**Tasks**:
- [ ] Create `lib/storage.ts` with LanceDB wrapper
- [ ] Implement Memory schema with metadata
- [ ] Add createMemory(), searchMemories(), deleteMemory() methods
- [ ] Test storage CRUD operations

**Dependencies**: None

---

### US-003: ONNX Embeddings Integration
**Status**: `[ ]` | **Points**: 8 | **Owner**: TBD

**User Story**: As the memory plugin, I want to generate embeddings locally using ONNX, so that I don't depend on external APIs.

**Acceptance Criteria**:
- [ ] nomic-embed-text model loads successfully
- [ ] Can generate embeddings for text (768 dimensions)
- [ ] Model cached in `~/.cache/opencode/models/`
- [ ] Batch embedding for efficiency
- [ ] Embedding generation < 500ms per memory

**Tasks**:
- [ ] Create `lib/embeddings.ts` with ONNX Runtime
- [ ] Implement model download and caching
- [ ] Add embed() method for single text
- [ ] Add embedBatch() for multiple texts
- [ ] Handle errors gracefully

**Dependencies**: None

---

## P1 - Core Functionality

### US-004: Session Idle Hook for Extraction
**Status**: `[ ]` | **Points**: 8 | **Owner**: TBD

**User Story**: As the memory plugin, I want to extract decisions from agent responses when the session ends, so that important knowledge is persisted.

**Acceptance Criteria**:
- [ ] Hook `session.idle` fires correctly
- [ ] Extracts decisions via pattern matching
- [ ] Generates embeddings for extracted memories
- [ ] Stores in LanceDB with metadata
- [ ] No duplicate memories created

**Tasks**:
- [ ] Implement `session.idle` hook in index.ts
- [ ] Create extraction logic in `lib/memory.ts`
- [ ] Add decision detection heuristics (regex patterns)
- [ ] Generate embeddings and store
- [ ] Test with sample session data

**Dependencies**: US-002, US-003

---

### US-005: Session Compacting Hook for Injection
**Status**: `[ ]` | **Points**: 5 | **Owner**: TBD

**User Story**: As the memory plugin, I want to inject relevant memories into the compaction prompt, so that the agent has access to past context.

**Acceptance Criteria**:
- [ ] Hook `experimental.session.compacting` fires correctly
- [ ] Searches for relevant memories by semantic similarity
- [ ] Respects token budget (max 4K tokens)
- [ ] Only injects memories above threshold (0.7)
- [ ] Formatted as Markdown list

**Tasks**:
- [ ] Implement `experimental.session.compacting` hook
- [ ] Add memory search with similarity threshold
- [ ] Format memories for injection
- [ ] Test context injection
- [ ] Verify token limits respected

**Dependencies**: US-002, US-004

---

### US-006: CLI Memory Commands
**Status**: `[ ]` | **Points**: 5 | **Owner**: TBD

**User Story**: As a developer, I want to inspect and manage memories via CLI, so that I can debug and curate the memory database.

**Acceptance Criteria**:
- [ ] `/memory status` shows count, size, config
- [ ] `/memory inspect` lists recent 10 memories
- [ ] `/memory search <query>` performs semantic search
- [ ] `/memory forget <id>` removes specific memory
- [ ] `/memory clear` clears all memories (with confirmation)

**Tasks**:
- [ ] Add custom tool `memory_query` to plugin
- [ ] Implement status, inspect, search, forget, clear operations
- [ ] Add confirmation prompts for destructive operations
- [ ] Test CLI commands

**Dependencies**: US-002, US-003

---

## P2 - Configuration & Polish

### US-007: Configuration Management
**Status**: `[ ]` | **Points**: 3 | **Owner**: TBD

**User Story**: As a developer, I want to configure memory behavior via config.json, so that I can customize TTL, limits, and retrieval settings.

**Acceptance Criteria**:
- [ ] Config file validates on load
- [ ] TTL settings respected (default 90 days)
- [ ] Max memories per project enforced (default 500)
- [ ] Similarity threshold configurable
- [ ] Token budget configurable

**Tasks**:
- [ ] Define JSON schema for config
- [ ] Implement config loader with validation
- [ ] Apply config to memory operations
- [ ] Test config overrides

**Dependencies**: US-002

---

### US-008: Garbage Collection
**Status**: `[ ]` | **Points**: 5 | **Owner**: TBD

**User Story**: As the memory plugin, I want to automatically remove expired and unused memories, so that the database doesn't grow indefinitely.

**Acceptance Criteria**:
- [ ] Memories older than TTL are deleted
- [ ] Memories never accessed after 30 days are deleted
- [ ] Hard cap of 500 memories enforced
- [ ] GC runs on plugin initialization
- [ ] GC respects "never expire" types (context)

**Tasks**:
- [ ] Implement pruneMemories() function
- [ ] Add TTL check logic
- [ ] Add access count tracking
- [ ] Enforce hard cap with LRU eviction
- [ ] Test GC behavior

**Dependencies**: US-002, US-007

---

### US-009: Sync Command Flag
**Status**: `[ ]` | **Points**: 3 | **Owner**: TBD

**User Story**: As a developer, I want to enable/disable memory on existing projects via `openkit sync --memory`, so that I can add memory to projects without re-initializing.

**Acceptance Criteria**:
- [ ] `openkit sync opencode --memory` installs plugin if missing
- [ ] `openkit sync opencode --no-memory` disables plugin
- [ ] Sync preserves existing memory data
- [ ] opencode.json updated correctly

**Tasks**:
- [ ] Add `--memory` flag to `internal/cli/agent_targets.go`
- [ ] Implement sync-time plugin installation
- [ ] Test sync with existing projects

**Dependencies**: US-001

---

### US-010: Documentation
**Status**: `[ ]` | **Points**: 3 | **Owner**: TBD

**User Story**: As a developer, I want comprehensive documentation on using the memory plugin, so that I can understand how it works and troubleshoot issues.

**Acceptance Criteria**:
- [ ] README updated with `--memory` flag usage
- [ ] Configuration guide in docs/
- [ ] Troubleshooting section with common issues
- [ ] Architecture diagram of plugin components

**Tasks**:
- [ ] Update main README with memory section
- [ ] Create docs/memory-plugin.md
- [ ] Add troubleshooting guide
- [ ] Create architecture diagram

**Dependencies**: US-001 through US-009

---

## Burndown

| Day | Stories Completed | Stories Remaining | Points Completed | Points Remaining |
|-----|-------------------|-------------------|------------------|------------------|
| 0   | 0                 | 10                | 0                | 50               |

---

## Notes

- Focus on P0 and P1 first (core functionality)
- P2 can be deferred if timeline is tight
- Integration testing is critical given OpenCode plugin complexity
- Consider early spike to validate LanceDB + ONNX integration before full implementation

## Related

- [[docs/sprint/Sprint-05/HUB-SPRINT-05.md]]
- [[docs/sprint/Sprint-05/SPRINT_GOAL.md]]
- [[docs/sprint/Sprint-05/TASKS.md]]
- [[docs/sprint/Sprint-05/RISK_REGISTER.md]]
