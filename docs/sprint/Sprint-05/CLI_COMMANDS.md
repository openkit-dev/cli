# CLI Memory Commands - Sprint 05

**Implemented**: 2026-02-09 14:15 (UTC-3)
**Task**: T008
**Status**: Complete

---

## Overview

Implemented 7 CLI subcommands for managing semantic memory via `openkit memory`.

## Commands

### 1. `openkit memory list`

List all stored memories with optional filtering.

```bash
# List all memories (default: 20)
openkit memory list

# Filter by type
openkit memory list --type decision
openkit memory list --type pattern
openkit memory list --type error
openkit memory list --type spec
openkit memory list --type context

# Limit results
openkit memory list --limit 5
openkit memory list -n 10
```

**Output:**
```
Memories (5 total)
==================================================

  [DECISION] Use PostgreSQL for database
    We decided to use PostgreSQL instead of MySQL because it has better JSON support...
    Salience: 0.85 | Accessed: 5x | Last: 2024-02-09 11:46
    Files: src/db/schema.ts, docker-compose.yml

  [PATTERN] Error handling pattern for API routes
    All API routes should wrap their handlers in a try-catch block...
    Salience: 0.78 | Accessed: 3x | Last: 2024-02-09 09:00
    Files: src/routes/api.ts, src/middleware/errorHandler.ts
```

---

### 2. `openkit memory search <query>`

Search memories by text content.

```bash
# Search for keywords
openkit memory search authentication
openkit memory search "error handling"
openkit memory search websocket

# Limit results
openkit memory search database --limit 3
```

**Output:**
```
Search Results for 'authentication' (1 found)
==================================================

  [SPEC] Authentication requirements
    Users must be able to sign in with email/password or OAuth...
    Salience: 0.80 | Tokens: ~42
```

**Note:** This is text-based search. For semantic search, use the `memory_query` tool in OpenCode.

---

### 3. `openkit memory stats`

Show comprehensive memory statistics.

```bash
openkit memory stats
```

**Output:**
```
Semantic Memory Statistics
==================================================

Storage
-------
  Total Memories: 5
  Total Tokens:   ~230
  Total Accesses: 20
  Avg Access:     4.0 per memory

By Type
-------
  decision:    1
  pattern:     1
  error:       1
  context:     1
  spec:        1

Session History
---------------
  Sessions Tracked:     4
  Compaction Events:    4 (100%)
  Total Memories Injected: 14
  Total Tokens Injected:   ~2340
  Avg Tokens per Session:  ~585

  Estimated Savings: ~7415 tokens/session (93%)

Configuration
-------------
  Token Budget:    4000
  Max Results:     10
  Min Similarity:  0.70
  TTL:             90 days
  Max per Project: 500
  Verbose:         true

Health
------
  [--] LanceDB not initialized (will be created on first memory)
  [OK] Configuration loaded
  [OK] 4 sessions tracked
```

---

### 4. `openkit memory prune`

Clean up old and unused memories.

```bash
# Preview what would be deleted
openkit memory prune --dry-run

# Run with confirmation prompt
openkit memory prune

# Skip confirmation
openkit memory prune --force
```

**Output:**
```
Memory Garbage Collection
==================================================

  Expired (TTL > 90 days):    2 memories
  Unused (> 30 days, <2 access): 5 memories
  Over Cap (> 500):            0 memories

  [DRY RUN] Would delete 7 memories
```

---

### 5. `openkit memory export [file]`

Export memories to JSON file.

```bash
# Export to default file
openkit memory export

# Export to specific file
openkit memory export backup.json
openkit memory export ~/backups/memories-2024-02-09.json
```

**Output:**
```
Export Memories
==================================================

  Exported 5 memories to backup.json
```

---

### 6. `openkit memory config`

Show or modify memory configuration.

```bash
# Show current configuration
openkit memory config

# Toggle verbose mode
openkit memory config --verbose
```

**Output (show):**
```
Memory Configuration
==================================================

  Embedding:
    model:   nomic-embed-text
    runtime: onnx

  Retrieval:
    max_results:    10
    min_similarity: 0.70
    token_budget:   4000

  Curation:
    ttl_days:              90
    max_per_project:       500
    prune_unused_after:    30 days

  Extraction:
    on_session_idle: true
    patterns:        decision, architecture, pattern, fix, solution

  Debug:
    verbose:              false
    show_injection_indicator: true

  Config file: /path/to/.opencode/memory/config.json
```

**Output (toggle verbose):**
```
  Verbose mode enabled
```

---

### 7. `openkit memory debug`

Debug memory system status.

```bash
openkit memory debug
```

**Output:**
```
Memory System Debug
==================================================

Installation Check
------------------
  [OK] Plugin directory
  [OK] Plugin index.ts
  [OK] Memory directory
  [OK] Config file
  [--] LanceDB directory (will be created on first use)
  [OK] Metrics file

Configuration
-------------
  [OK] Config loaded successfully
       Version: 1.0.0
       Verbose: true

Session Metrics
---------------
  [OK] 4 sessions tracked

  Recent sessions:
    sess_def... | Loaded: 5 | Injected: 4 | Tokens: ~680
    sess_ghi... | Loaded: 4 | Injected: 2 | Tokens: ~320
    sess_jkl... | Loaded: 5 | Injected: 5 | Tokens: ~890

Recommendations
---------------
  - Run more OpenCode sessions to accumulate metrics
```

---

## Files Created

| File | Lines | Description |
|------|-------|-------------|
| `internal/cli/memory.go` | 750+ | All memory CLI commands |

## Data Files

The CLI reads from these files in `.opencode/memory/`:

| File | Description |
|------|-------------|
| `config.json` | Memory configuration |
| `metrics.json` | Session metrics history |
| `memories.json` | Memory export (JSON fallback) |
| `index.lance/` | LanceDB vector database |

## Limitations

1. **Text Search Only**: CLI provides text-based search. For semantic (vector) search, use the `memory_query` tool in OpenCode.

2. **LanceDB Read**: CLI cannot directly read from LanceDB (requires Bun/Node runtime). Use `openkit memory export` in OpenCode to create a JSON snapshot.

3. **Prune Execution**: Full pruning requires the OpenCode plugin. CLI shows what would be deleted but cannot modify LanceDB directly.

## Usage Flow

```
1. Initialize project with memory:
   openkit init my-project --memory

2. Run OpenCode sessions (memories accumulate automatically)

3. Check status:
   openkit memory stats
   openkit memory debug

4. Search memories:
   openkit memory search "authentication"

5. List by type:
   openkit memory list --type decision

6. Export for backup:
   openkit memory export backup.json

7. Enable verbose logging:
   openkit memory config --verbose

8. Clean up old memories:
   openkit memory prune --dry-run
   openkit memory prune
```

## Integration with OpenCode Tools

| CLI Command | OpenCode Tool | Difference |
|-------------|---------------|------------|
| `memory list` | `memory_stats` | CLI shows list, tool shows stats |
| `memory search` | `memory_query` | CLI is text-based, tool is semantic |
| `memory stats` | `memory_stats` | Similar output |
| `memory config` | Edit config.json | CLI can toggle verbose |
| `memory debug` | `memory_debug` | Similar functionality |
| `memory export` | N/A | CLI only |
| `memory prune` | Auto on init | CLI shows preview |

## Related

- [[sprint/Sprint-05/HUB-SPRINT-05.md]]
- [[sprint/Sprint-05/IMPLEMENTATION_STATUS.md]]
- [[sprint/Sprint-05/TASKS.md]]
- [[CONTEXT.md]]
