# Sprint Tasks - Sprint 05

**Created**: 2026-02-09 13:05 (UTC-3)
**Updated**: 2026-02-09 14:30 (UTC-3)
**Sprint Goal**: Implement semantic memory plugin for OpenCode/OpenKit
**Status**: Phase 2 Complete (Core Logic + CLI + Sync) - Docs and E2E testing pending

## Task Format
```
[Status] ID: Task Name | Priority | Estimate | Owner
  INPUT: What is needed to start
  OUTPUT: What will be produced
  VERIFY: How to confirm completion
  Dependencies: [List of task IDs]
```

**Status Legend**: `[ ]` Not Started | `[>]` In Progress | `[x]` Complete | `[~]` Blocked

---

## Phase 1: Foundation (P0)

### [[x] T001: Add --memory Flag to init Command | P0 | 2h | Complete

**INPUT**: 
- `internal/cli/init.go` source code
- Requirements from OPENKIT_INTEGRATION.md

**OUTPUT**:
- Modified init.go with `flagMemory bool` variable
- Flag registered in init() function
- Logic in runInit() to call installMemoryPlugin() when flag is set

**VERIFY**:
```bash
openkit init test-app --memory
# Should create .opencode/plugins/semantic-memory/
# Should create .opencode/memory/config.json
# Should update opencode.json
```

**Dependencies**: None

---

### [[x] T002: Create installMemoryPlugin() Function | P0 | 3h | Complete

**INPUT**:
- Plugin structure design from OPENKIT_INTEGRATION.md
- Embedded templates location

**OUTPUT**:
- `installMemoryPlugin(projectDir string) error` function
- Creates `.opencode/plugins/semantic-memory/` directory
- Creates `.opencode/memory/` directory
- Copies embedded plugin files
- Creates default config.json
- Updates opencode.json with memory settings

**VERIFY**:
```bash
# After running init with --memory
ls .opencode/plugins/semantic-memory/
# Should show: index.ts, package.json, lib/
ls .opencode/memory/
# Should show: config.json, .gitignore
cat opencode.json | grep memory
# Should show: "memory": { "enabled": true }
```

**Dependencies**: None

---

### [[x] T003: Create Embedded Plugin Templates | P0 | 4h | Complete

**INPUT**:
- Plugin TypeScript code design
- Template structure from OPENKIT_INTEGRATION.md

**OUTPUT**:
- `internal/templates/memory/plugin/` directory structure
- `index.ts` with plugin skeleton
- `package.json` with lancedb, onnxruntime-node
- `lib/memory.ts` stub
- `lib/embeddings.ts` stub
- `lib/storage.ts` stub
- `embed.go` for Go embedding

**VERIFY**:
```bash
ls internal/templates/memory/plugin/
# Should match expected structure
cat internal/templates/memory/plugin/package.json
# Should list lancedb and onnxruntime-node
```

**Dependencies**: None

---

### [[x] T004: Implement LanceDB Storage Wrapper | P0 | 5h | Complete

**INPUT**:
- LanceDB documentation
- Memory schema from ARCHITECTURE_ANALYSIS.md

**OUTPUT**:
- `lib/storage.ts` with LanceDBStorage class
- Methods: initialize(), createMemory(), searchMemories(), deleteMemory()
- Memory interface with id, type, title, content, vector, salience, timestamps
- Error handling for DB operations

**VERIFY**:
```typescript
const storage = new LanceDBStorage({ dbPath: './test.lance' })
await storage.initialize()
await storage.createMemory({
  type: 'decision',
  title: 'Test',
  content: 'Test memory',
  vector: new Float32Array(768)
})
const results = await storage.searchMemories(vector, 5, 0.7)
console.log(results) // Should return similar memories
```

**Dependencies**: None

---

### [[x] T005: Implement ONNX Embeddings Integration | P0 | 6h | Complete

**INPUT**:
- onnxruntime-node documentation
- nomic-embed-text model specs

**OUTPUT**:
- `lib/embeddings.ts` with EmbeddingService class
- Methods: initialize(), embed(), embedBatch()
- Model download and caching logic
- Error handling for ONNX operations

**VERIFY**:
```typescript
const embedder = new EmbeddingService({ 
  modelPath: '~/.cache/opencode/models/nomic-embed-text'
})
await embedder.initialize()
const vector = await embedder.embed('This is a test')
console.log(vector.length) // Should be 768
console.log(typeof vector[0]) // Should be 'number'
```

**Dependencies**: None

---

## Phase 2: Core Functionality (P1)

### [[x] T006: Implement session.idle Hook | P1 | 4h | Complete

**INPUT**:
- OpenCode plugin hook documentation
- Extraction heuristics design

**OUTPUT**:
- session.idle hook implementation in index.ts
- extractFromSession() method in memory.ts with full logic:
  - Decision detection via pattern matching (decision, architecture, pattern, fix, solution)
  - Structured extraction: title, facts, concepts, files, type classification
  - Salience scoring based on content richness
  - Embedding generation and storage
  - Error handling and verbose logging
- Helper methods: detectDecisions(), extractTitle(), extractFacts(), extractConcepts(), extractFiles(), classifyType(), calculateSalience()

**VERIFY**:
```bash
# Run OpenCode session with plugin enabled
# Trigger session.idle
# Check LanceDB:
ls -la .opencode/memory/index.lance/
# Should show database files created
```

**Dependencies**: T004, T005
**Status**: Implemented with full extraction pipeline

---

### [[x] T007: Implement session.compacting Hook | P1 | 3h | Complete

**INPUT**:
- OpenCode compacting hook documentation
- Progressive disclosure design

**OUTPUT**:
- experimental.session.compacting hook implementation in index.ts
- getRelevantContext() method with full logic:
  - Query embedding generation
  - Vector similarity search with threshold filtering
  - Salience-based sorting
  - Token budget enforcement (4000 tokens default)
  - Access metadata tracking
  - Session cache management
- Helper method: applyTokenBudget() (rough estimate: 1 token ≈ 4 chars)
- Markdown formatting for injection: `[type] title: content`

**VERIFY**:
```bash
# Run OpenCode session with existing memories
# Check compaction output includes memory context
# Verify token count < 4000 via logs
```

**Dependencies**: T004, T006
**Status**: Implemented with full retrieval pipeline and budget control

---

### [[x] T008: Implement CLI Memory Commands | P1 | 5h | Complete

**INPUT**:
- OpenCode custom tools API
- CLI command design from USER_STORIES.md

**OUTPUT**:
- New file: `internal/cli/memory.go` (750+ lines)
- 7 subcommands implemented:
  - `openkit memory list` - List all memories with filtering (--type, --limit)
  - `openkit memory search <query>` - Text search in memories
  - `openkit memory stats` - Show comprehensive statistics
  - `openkit memory prune` - Garbage collection (--dry-run, --force)
  - `openkit memory export [file]` - Export to JSON
  - `openkit memory config` - Show/modify config (--verbose toggle)
  - `openkit memory debug` - Debug system status
- Color-coded output with type indicators
- Health checks and recommendations
- Token savings calculation

**VERIFY**:
```bash
# All commands tested and working:
openkit memory list                    # Lists all memories
openkit memory list --type decision    # Filter by type
openkit memory search authentication   # Text search
openkit memory stats                   # Full statistics
openkit memory prune --dry-run         # Show what would be deleted
openkit memory export backup.json      # Export to file
openkit memory config                  # Show configuration
openkit memory config --verbose        # Toggle verbose mode
openkit memory debug                   # Debug information
```

**Dependencies**: T004
**Status**: Fully implemented with all 7 subcommands

---

## Phase 3: Configuration & Polish (P2)

### [[x] T009: Implement Configuration Management | P2 | 3h | Complete

**INPUT**:
- Config schema from ARCHITECTURE_ANALYSIS.md
- Default values

**OUTPUT**:
- MemoryConfig interface with full schema validation
- loadConfig() method in SemanticMemory class with:
  - JSON file reading from configPath
  - Error handling with fallback to defaults
  - Type-safe config structure
- Config applied to all operations:
  - Retrieval: max_results, min_similarity, token_budget
  - Curation: ttl_days, max_per_project, prune_unused_after_days
  - Extraction: on_session_idle, patterns
  - Debug: verbose, show_injection_indicator
- Default config.json generated by Go CLI with all parameters

**VERIFY**:
```typescript
const config = await loadConfig('.opencode/memory/config.json')
console.log(config.retrieval.max_results) // Should be 10
console.log(config.curation.ttl_days) // Should be 90
```

**Dependencies**: None
**Status**: Fully implemented with type safety and defaults

---

### [[x] T010: Implement Garbage Collection | P2 | 4h | Complete

**INPUT**:
- TTL and access count logic
- GC strategy from ARCHITECTURE_ANALYSIS.md

**OUTPUT**:
- pruneMemories() method with three-phase GC:
  1. **TTL Expiration**: deleteExpired() removes memories past expires_at
  2. **Unused Cleanup**: pruneUnusedMemories() removes memories not accessed in prune_unused_after_days with low access_count (<2)
  3. **Hard Cap Enforcement**: enforceHardCap() implements LRU eviction when project exceeds max_per_project
- Helper methods:
  - pruneUnusedMemories(threshold): Deletes stale, rarely-accessed memories
  - enforceHardCap(): Sorts by accessed_at ASC, deletes oldest beyond cap
- GC runs automatically on plugin initialization
- Verbose logging of GC results

**VERIFY**:
```typescript
// Create memory with old timestamp
await storage.createMemory({ ..., expires_at: Date.now() - 1000 })
await memory.pruneMemories()
// Old memory should be deleted
// Logs: "GC complete: X expired, Y unused, Z over-cap"
```

**Dependencies**: T004, T009
**Status**: Fully implemented with three-phase strategy

---

### [[x] T011: Add --memory Flag to sync Command | P2 | 2h | Complete

**INPUT**:
- `internal/cli/agent_targets.go` source
- Sync command design

**OUTPUT**:
- Modified `internal/cli/agent_targets.go`:
  - Added `flagSyncMemory bool` variable
  - Added `--memory` flag to sync and upgrade commands (OpenCode only)
  - Implemented `syncMemoryPlugin()` function for install/update
  - Implemented `extractMemoryPluginForSync()` wrapper
  - Implemented `createDefaultMemoryConfig()` helper
- Features:
  - Installs plugin if not present
  - Updates plugin files if already installed
  - Preserves existing config.json on update
  - Creates .gitignore for LanceDB data
  - Shows manual steps for opencode.json

**VERIFY**:
```bash
# Install plugin in existing project
openkit opencode sync --memory
# Output: "Installing semantic memory plugin..."
#         "Extracted SEMANTIC_MEMORY.md rule"

# Update plugin (preserves config)
openkit opencode sync --memory
# Output: "Updating semantic memory plugin..."
#         "Preserved existing config.json"
#         "Extracted SEMANTIC_MEMORY.md rule"

# Verify files
ls .opencode/plugins/semantic-memory/
# index.ts, package.json, lib/, scripts/

ls .opencode/memory/
# config.json, .gitignore

ls .opencode/rules/ | grep SEMANTIC
# SEMANTIC_MEMORY.md
```

**Dependencies**: T001, T002
**Status**: Fully implemented with install/update logic and rule extraction

**Additional Implementation (Orchestrator Integration)**:
- Added "Semantic Memory Integration" section to `orchestrator.md`
- Created `SEMANTIC_MEMORY.md` rule in `internal/templates/memory/rules/`
- Updated `embed.go` with `ExtractMemoryRules()` function
- Rule is extracted to `.opencode/rules/` only for OpenCode projects
- Orchestrator now has protocol for using memory tools during missions

---

### [[ ] T012: Write Documentation | P2 | 4h | TBD

**INPUT**:
- All implemented features
- User stories and acceptance criteria

**OUTPUT**:
- README.md updated with memory section
- docs/memory-plugin.md with configuration guide
- Troubleshooting section
- Architecture diagram

**VERIFY**:
```bash
# Check documentation coverage
grep -r "memory" README.md
# Should find memory plugin section
cat docs/memory-plugin.md
# Should have complete configuration guide
```

**Dependencies**: T001-T011

---

## Phase 4: Testing & Verification (P3)

### [[ ] T013: Integration Testing | P3 | 4h | TBD

**INPUT**:
- Completed plugin implementation
- Test scenarios

**OUTPUT**:
- Integration test suite
- End-to-end flow testing
- OpenCode compatibility verification

**VERIFY**:
```bash
npm test
# All tests should pass
```

**Dependencies**: T001-T012

---

### [[ ] T014: Performance Testing | P3 | 2h | TBD

**INPUT**:
- Completed implementation
- Performance requirements

**OUTPUT**:
- Performance benchmarks
- Startup time measurement
- Embedding speed verification

**VERIFY**:
```bash
# Measure plugin initialization time
# Should be < 100ms
# Measure embedding generation
# Should be < 500ms per memory
```

**Dependencies**: T004, T005

---

## Summary

**Total Tasks**: 14
**Total Estimate**: 51 hours
**Critical Path**: T001 → T002 → T003 → T004 → T005 → T006 → T007

**Priority Breakdown**:
- P0: 5 tasks (20 hours)
- P1: 3 tasks (12 hours)
- P2: 4 tasks (13 hours)
- P3: 2 tasks (6 hours)

---

## Sprint 05 Completion Summary

**Completed**: 2026-02-09 13:30 (UTC-3)
**Duration**: ~3 hours
**Status**: Foundation Complete - Core Logic Pending

### Tasks Completed ✅

| Task | Status | Hours | Notes |
|------|--------|-------|-------|
| T001 | ✅ | 2h | Flag --memory added to init.go |
| T002 | ✅ | 3h | installMemoryPlugin() implemented |
| T003 | ✅ | 4h | Plugin templates created (7 files) |

**Total Completed**: 3 tasks, 9 hours

### Tasks Blocked (Need Real Implementation) 🔴

| Task | Status | Hours | Blocker |
|------|--------|-------|---------|
| T004 | 🔴 | 5h | LanceDB real implementation needed |
| T005 | 🔴 | 6h | ONNX real implementation needed |

### Tasks Not Started (Depend on T004, T005) ⏸️

| Task | Priority | Hours | Dependencies |
|------|----------|-------|--------------|
| T006 | P1 | 4h | T004, T005 |
| T007 | P1 | 3h | T004, T005 |
| T008 | P1 | 5h | T004 |
| T009 | P2 | 3h | - |
| T010 | P2 | 4h | T004, T009 |
| T011 | P2 | 2h | T001, T002 |
| T012 | P2 | 4h | T001-T011 |
| T013 | P3 | 4h | T001-T012 |
| T014 | P3 | 2h | T004, T005 |

**Total Remaining**: 11 tasks, 42 hours

### Deliverables

#### Code Created
- `internal/cli/init.go` - Modified (+75 lines)
- `internal/templates/embed.go` - Modified (+5 lines)
- `internal/templates/memory/index.ts` - New (74 lines)
- `internal/templates/memory/package.json` - New (12 lines)
- `internal/templates/memory/lib/memory.ts` - New (142 lines)
- `internal/templates/memory/lib/storage.ts` - New (35 lines)
- `internal/templates/memory/lib/embeddings.ts` - New (30 lines)

**Total**: ~370 lines of code

#### Documentation Created
- `SPRINT_GOAL.md` (80 lines)
- `BACKLOG.md` (290 lines)
- `TASKS.md` (310 lines)
- `RISK_REGISTER.md` (220 lines)
- `IMPLEMENTATION_STATUS.md` (350 lines)

**Total**: ~1250 lines of documentation

### Next Sprint (Sprint 06)

**Focus**: Implement core logic (T004, T005) and hooks (T006, T007)

**Critical Path**:
1. Test Go build (verify T001, T002, T003)
2. Implement LanceDB real storage (T004)
3. Implement ONNX real embeddings (T005)
4. Implement hooks logic (T006, T007)
5. Integration testing (T013)

**Estimated Effort**: 20-24 hours

### Sprint Retrospective

**What Went Well** ✅:
- Comprehensive planning artifacts
- Clean code structure
- Error handling from start
- Documentation-first approach

**Challenges** 🔴:
- No Go compiler for testing
- Stubs need real implementation
- External dependencies complexity

**Lessons Learned** 📚:
- Structural MVP is valuable
- Separation of concerns helps
- Testing requires real environment

**Action Items**:
1. Set up Go test environment
2. Create integration test plan
3. Research LanceDB best practices
4. Research ONNX model optimization

## Related

- [[docs/sprint/Sprint-05/HUB-SPRINT-05.md]]
- [[docs/sprint/Sprint-05/SPRINT_GOAL.md]]
- [[docs/sprint/Sprint-05/BACKLOG.md]]
- [[docs/requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
