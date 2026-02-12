# BACKEND / INTERNAL ARCHITECTURE

**Created**: 2026-02-08
**Last Updated**: 2026-02-10

## Architecture

Monolithic Go application organized by functional packages in `internal/`. The CLI tool follows a clean separation of concerns with packages for each major subsystem.

## Package Overview

| Package | Responsibility | Key Files |
|---|---|---|
| `internal/agents` | Agent registry and agent metadata management. | `registry.go` |
| `internal/cli` | Cobra command definitions, flags, and CLI entry points. | `root.go`, `init.go`, `agent_targets.go`, `memory.go`, `upgrade.go`, `uninstall.go`, `check.go`, `version.go` |
| `internal/managedstate` | State persistence for tracking managed files, conflict detection, drift detection. | `managedstate.go` |
| `internal/platform` | OS-specific abstractions for paths and environment. | `paths.go` |
| `internal/selfupdate` | Binary auto-update mechanism with GitHub API integration. | `upgrade.go`, `checker.go`, `github_latest.go`, `state.go` |
| `internal/syncer` | Synchronization logic for deploying agent configurations. | `syncer.go` |
| `internal/targets` | Agent-specific content generators for each supported agent. | `targets.go`, `claude_commands.go`, `cursor_content.go`, `codex_content.go`, `gemini_commands.go` |
| `internal/templates` | Embedded project scaffolding templates. | `embed.go`, base templates, memory templates, root templates |
| `internal/ui` | Console output helpers (Success, Error, Info, Warning). | `ui.go` |

## Data Models

### Managed State (`internal/managedstate`)

**Schema**: File-based JSON stored in `.openkit/managed.json`

```go
type ManagedState struct {
    SchemaVersion int                    `json:"schema_version"`
    Agents        map[string]AgentState  `json:"agents"`
}

type AgentState struct {
    Pack struct {
        ID      string `json:"id"`
        Version string `json:"version"`
    } `json:"pack"`
    Files map[string]FileState `json:"files"`
}

type FileState struct {
    InstalledSHA256 string `json:"installed_sha256"`
    Mode            string `json:"mode"` // "copy" or "template"
}
```

**Evidence**: `internal/managedstate/managedstate.go`

### Configuration (`opencode.json`)

**Schema**: JSON schema defined at `https://opencode.ai/config.json`

```json
{
  "$schema": "https://opencode.ai/config.json",
  "instructions": ["path/to/rules"],
  "default_agent": "orchestrator",
  "agent": {
    "<agent_name>": {
      "description": "string",
      "mode": "primary|subagent",
      "prompt": "path/to/prompt",
      "tools": {
        "<tool_name>": true|false
      },
      "permission": {
        "<tool_name>": "allow|ask|deny"
      }
    }
  }
}
```

**Evidence**: `opencode.json`

### Sync Plan (`internal/syncer`)

**Schema**: Planning structure for file synchronization

```go
type Plan struct {
    Entries    []PlanEntry
    Create     []string
    Update     []string
    Overwrite  []string
    Skip       []string
    Conflicts  []string
    Delete     []string
    Orphaned   []string
}

type PlanEntry struct {
    Action     Action     // "create", "update", "overwrite", "skip", "conflict", "delete"
    Path       string
    Reason     string
    ArtifactID string
}
```

**Evidence**: `internal/syncer/syncer.go`

## Command Flow

### Root Command Flow

```
openkit <command>
  ↓
cmd/openkit/main.go
  ↓
internal/cli/root.go.Execute()
  ↓
internal/cli/root.Command.PersistentPreRunE()
  ↓ (maybeNotifyUpdate)
internal/selfupdate.Checker.Check()
  ↓
Command execution (init, check, version, <agent> subcommands)
```

### Agent Sync Flow

```
openkit <agent> sync
  ↓
internal/cli/agent_targets.go.NewSyncCmd()
  ↓
internal/syncer.GeneratePlan()
  ↓ (from embedded templates)
internal/templates/embed.go
  ↓ (compare with current state)
internal/managedstate/ManagedState.Load()
  ↓ (create plan)
internal/syncer.Plan.Apply()
  ↓ (write files, update state)
internal/managedstate/ManagedState.Save()
```

### Self-Update Flow

```
openkit upgrade
  ↓
internal/cli/upgrade.go
  ↓
internal/selfupdate/upgrade.Upgrade()
  ↓ (check latest release)
internal/selfupdate/github_latest.GetLatestRelease()
  ↓ (download binary)
internal/selfupdate/github_latest.DownloadRelease()
  ↓ (verify checksum)
internal/selfupdate/checksum.Verify()
  ↓ (replace binary)
internal/selfupdate/upgrade.ReplaceBinary()
```

## Key Algorithms

### Conflict Detection

**Location**: `internal/syncer/syncer.go`

**Logic**:
1. Generate desired file list from embedded templates
2. Load current state from `.openkit/managed.json`
3. For each desired file:
   - If file exists and is managed: Compare SHA256
     - Different: Mark as `update`
     - Same: Mark as `skip`
   - If file exists and is unmanaged: Mark as `conflict`
   - If file does not exist: Mark as `create`
4. For each managed file not in desired list: Mark as `orphaned`

### Drift Detection

**Location**: `internal/syncer/syncer.go`

**Logic**:
1. Calculate SHA256 of current file on disk
2. Compare with SHA256 stored in managed state
3. If different: File has drifted (manually modified)
4. Action: Warn user, require `--overwrite` flag to proceed

### Version Injection

**Location**: `Makefile`

**Logic**:
```makefile
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"
```

Injected into binary at build time and accessible via `main.version`, `main.commit`, `main.date`.

## Entry Points

### CLI Entry Point

**File**: `cmd/openkit/main.go`

**Responsibilities**:
- Parse command-line arguments
- Initialize CLI root command
- Execute Cobra command tree

### Agent-Specific Subcommands

**File**: `internal/cli/agent_targets.go`

**Functions**:
- `NewAgentCmd()` - Creates agent subcommands (sync, upgrade, doctor, uninstall)
- `NewSyncCmd()` - Sync agent configuration
- `NewUpgradeCmd()` - Upgrade agent configuration
- `NewDoctorCmd()` - Check agent health
- `NewUninstallCmd()` - Remove agent configuration

### Memory Commands (OpenCode)

**File**: `internal/cli/memory.go`

**Functions**:
- `NewMemoryCmd()` - Memory management root command
- `NewMemoryListCmd()` - List stored memories
- `NewMemorySearchCmd()` - Search memories
- `NewMemoryStatsCmd()` - Show statistics
- `NewMemoryExportCmd()` - Export memories to JSON
- `NewMemoryPruneCmd()` - Clean up old memories
- `NewMemoryConfigCmd()` - Show/modify configuration
- `NewMemoryDebugCmd()` - Debug system status

## External Dependencies

| Package | Version | Purpose | Usage |
|---------|---------|---------|-------|
| `github.com/spf13/cobra` | v1.10.2 | CLI framework | Command definition, flags, help |
| `github.com/fatih/color` | v1.18.0 | Colored output | Success, error, info, warning messages |
| `golang.org/x/mod` | v0.32.0 | Go module versioning | Version parsing and comparison |

## Configuration Files

| File | Purpose | Format |
|------|---------|--------|
| `go.mod` | Go module definition | Go module format |
| `go.sum` | Dependency checksums | Hash list |
| `.golangci.yml` | Linter configuration | YAML |
| `.goreleaser.yaml` | Release automation | YAML |
| `opencode.json` | Agent system config | JSON with schema |
| `.openkit/managed.json` | Managed file state | JSON |
| `.openkit/package.json` | Node.js deps for scripts | NPM format |

## Error Handling

**Strategy**: Return errors from functions, use `exitWithError()` helper for CLI errors.

**Evidence**: `internal/cli/root.go`

```go
func exitWithError(msg string) {
    red := color.New(color.FgRed, color.Bold)
    red.Fprintf(os.Stderr, "Error: %s\n", msg)
    os.Exit(1)
}
```

## Logging

**Strategy**: No structured logging currently. Output goes directly to stdout/stderr with color formatting.

**Gap**: Consider adding structured logging (e.g., `slog` in Go 1.21+) for audit trails.

## Testing Strategy

**Unit Tests**: Test individual packages in isolation
- `internal/ui/ui_test.go`
- `internal/cli/upgrade_test.go`
- `internal/cli/agent_targets_integration_test.go`
- `internal/targets/*_test.go`
- `internal/selfupdate/upgrade_test.go`

**Integration Tests**: Test command flows end-to-end
- `internal/cli/agent_targets_integration_test.go`

**Gap**: No end-to-end tests for complete user workflows.

## Performance Considerations

**Optimizations**:
- Embedded templates: No file system overhead at runtime
- SHA256 caching: Compute once and store in managed state
- HTTP timeout: 2 second timeout for update checks

**Potential Issues**:
- Large template filesystem increases binary size
- No caching for GitHub API calls (update checks)

## Security Considerations

**Security Boundaries**:
- File system access: Full access to project directory
- Network access: Limited to GitHub API for updates
- Command execution: Via `bash` tool (controlled by agent permissions)

**Gaps**:
- No sandbox or containerization
- No input validation for template parameters
- No audit logging for sensitive operations

See [[SECURITY.md]] for detailed security analysis.

## Related

- [[HUB-DOCS.md]]
- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
