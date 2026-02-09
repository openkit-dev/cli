---
trigger: always_on
agent: opencode
---

# Semantic Memory Protocol (OpenCode Only)

> Rules for using the semantic memory plugin to optimize context across sessions.

**Note:** This protocol applies ONLY to OpenCode. Other agents (Claude, Cursor, Gemini) do not have access to these memory tools.

---

## CRITICAL: Direct Tool Usage (MANDATORY)

**NEVER check if the memory plugin exists before using memory tools.**

The memory tools (`memory_query`, `memory_stats`, `memory_save`, `memory_debug`) are built-in system tools. Use them DIRECTLY without:
- Running `ls` to check plugin directories
- Asking for permission to use them
- Verifying plugin existence

**WRONG:**
```
# First let me check if the plugin exists
$ ls -d .opencode/plugins/semantic-memory/
```

**CORRECT:**
```
[Use memory_stats directly]
```

If a memory tool fails, the error message will indicate the issue. Do not pre-check.

---

## Memory Tools Available

The semantic memory plugin provides these tools:

| Tool | Purpose | When to Use |
|------|---------|-------------|
| `memory_context` | Get optimized context for current task | At the START of complex tasks |
| `memory_save` | Save important decisions/patterns | After making significant decisions |
| `memory_query` | Search for specific memories | When you need specific past context |
| `memory_stats` | View memory statistics | To check system health |
| `memory_debug` | Debug memory system | To verify plugin is working |

---

## Mandatory Memory Usage

### 1. At Session Start (Automatic)

The plugin automatically loads relevant context via the `session.created` hook. No action required.

### 2. Before Complex Tasks (RECOMMENDED)

**When starting a complex task, use `memory_context` to retrieve relevant context:**

```
Before implementing authentication, let me check what decisions we've made before.

[Use memory_context with task="implement user authentication"]
```

**Triggers for memory_context:**
- Starting implementation of a feature mentioned before
- Working on code that was discussed in previous sessions
- Making architectural decisions
- Debugging issues that may have been fixed before

### 3. After Important Decisions (MANDATORY)

**When you make a significant decision, save it with `memory_save`:**

```
We decided to use PostgreSQL instead of MySQL for better JSON support.

[Use memory_save with:
  type="decision"
  title="Use PostgreSQL for database"
  content="Chose PostgreSQL over MySQL because: 1) Better JSON support with jsonb, 2) Native array types for tags, 3) Better performance for complex queries"
  files="src/db/schema.ts, docker-compose.yml"
]
```

**Types of decisions to save:**
- Architecture decisions (database, framework, patterns)
- Error fixes (what caused it, how it was fixed)
- Design patterns adopted
- Configuration choices
- Security decisions

### 4. Decision Classification

Use these types when saving memories:

| Type | When to Use | Example |
|------|-------------|---------|
| `decision` | Architecture, tech choices | "Use JWT for auth" |
| `pattern` | Code patterns, conventions | "Error handling pattern" |
| `error` | Bug fixes, issue resolutions | "Fixed memory leak in WS" |
| `spec` | Requirements, acceptance criteria | "Auth requirements" |
| `context` | General project knowledge | "Project uses TanStack Query" |

---

## Memory Usage Patterns

### Pattern 1: Feature Implementation

```
User: "Implement user profile editing"

1. [memory_context with task="user profile editing"]
   -> Retrieves past decisions about user model, API patterns, UI conventions

2. Implement the feature using retrieved context

3. [memory_save with type="pattern" if new patterns were established]
```

### Pattern 2: Bug Investigation

```
User: "The WebSocket keeps disconnecting"

1. [memory_context with task="websocket disconnection issues"]
   -> May retrieve past fixes for similar issues

2. If no relevant memories, investigate and fix

3. [memory_save with type="error" documenting the fix]
```

### Pattern 3: Architecture Discussion

```
User: "Should we use Redis or PostgreSQL for caching?"

1. [memory_context with task="caching strategy"]
   -> Retrieves past decisions about performance, infrastructure

2. Make recommendation based on context + new analysis

3. [memory_save with type="decision" once user confirms choice]
```

---

## What NOT to Save

Do NOT use `memory_save` for:
- Trivial changes (typo fixes, formatting)
- Temporary solutions (TODOs, workarounds marked for removal)
- User-specific preferences (unless they affect the whole project)
- Sensitive data (credentials, tokens, personal info)

---

## Memory Health Checks

### Periodically Check Stats

When working on a project for extended periods:

```
[memory_stats]
```

Look for:
- Total memories (should grow over time)
- Access patterns (frequently accessed = valuable)
- Token budget usage (should be < 4000)

### Debug If Issues Suspected

If context seems missing or irrelevant:

```
[memory_debug with action="status"]
```

Check:
- Plugin status: ACTIVE
- Compaction triggered: Should be true for long sessions
- Memories loaded: Should be > 0 if project has history

---

## Integration with Orchestrator

### During Orchestration Mode

When the orchestrator coordinates multiple agents:

1. **Before Planning Phase:**
   ```
   [memory_context with task="<mission description>"]
   ```
   Retrieve relevant past decisions before creating new plans.

2. **After Implementation Phase:**
   ```
   [memory_save with type="decision" for each major architectural choice]
   ```
   Capture decisions made during the mission.

3. **In Verification Phase:**
   Memory is automatically preserved for future sessions.

### Passing Context to Sub-Agents

When invoking specialist agents via Task tool, include relevant memories in the prompt:

```
Task(
  subagent_type: "backend-specialist",
  prompt: """
  Implement user authentication API.
  
  Relevant context from project memory:
  - We use JWT tokens (not sessions)
  - PostgreSQL for user storage
  - bcrypt for password hashing
  
  [Full task description...]
  """
)
```

---

## Automatic Behaviors

The plugin handles these automatically:

| Behavior | Trigger | Action |
|----------|---------|--------|
| Context loading | Session start | Loads relevant memories into cache |
| Context injection | Session compaction | Injects memories into compressed context |
| Knowledge extraction | Session idle | Extracts decisions from conversation |
| Garbage collection | Plugin init | Removes expired/unused memories |
| Access tracking | Memory retrieval | Updates access counts for LRU |

---

## Configuration Reference

Memory behavior can be adjusted in `.opencode/memory/config.json`:

```json
{
  "retrieval": {
    "max_results": 10,        // Max memories to inject
    "min_similarity": 0.7,    // Similarity threshold (0-1)
    "token_budget": 4000      // Max tokens for context
  },
  "extraction": {
    "on_session_idle": true,  // Auto-extract decisions
    "patterns": ["decision", "architecture", "pattern", "fix", "solution"]
  },
  "debug": {
    "verbose": false          // Enable detailed logs
  }
}
```

To enable verbose logging for debugging:
```bash
openkit memory config --verbose
```

---

## CLI Commands (Outside OpenCode)

These commands are available in the terminal:

```bash
# View all memories
openkit memory list

# Search memories
openkit memory search "authentication"

# View statistics
openkit memory stats

# Debug system
openkit memory debug

# Export for backup
openkit memory export backup.json
```

---

## Summary

1. **Use `memory_context`** before complex tasks to retrieve relevant context
2. **Use `memory_save`** after important decisions to preserve knowledge
3. **Let automatic hooks** handle session start/end context management
4. **Check `memory_stats`** periodically to monitor system health
5. **Include memory context** when delegating to sub-agents
