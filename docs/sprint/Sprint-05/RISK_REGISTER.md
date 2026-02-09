# Risk Register - Sprint 05

**Created**: 2026-02-09 13:05 (UTC-3)
**Sprint Goal**: Implement semantic memory plugin for OpenCode/OpenKit

## Risk Matrix

| ID | Risk | Probability | Impact | Score | Mitigation | Owner | Status |
|----|------|-------------|--------|-------|------------|-------|--------|
| R1 | ONNX model size (~500MB) causes slow download | High | Medium | 6 | Lazy download, show progress, cache in ~/.cache | TBD | Open |
| R2 | LanceDB incompatible with platform | Medium | High | 6 | Test on Linux/macOS/Windows, fallback to JSON | TBD | Open |
| R3 | Hook integration breaks OpenCode | Medium | High | 6 | Minimal hooks, error handling, graceful degradation | TBD | Open |
| R4 | Embedding generation too slow | Medium | Medium | 4 | Use lightweight model, batch processing, async | TBD | Open |
| R5 | Memory database grows too large | Medium | Low | 3 | Hard cap, GC, TTL enforcement | TBD | Open |
| R6 | Plugin conflicts with other plugins | Low | High | 4 | Test with common plugins, namespace isolation | TBD | Open |
| R7 | Config validation fails silently | Low | Medium | 2 | Strict schema validation, error logging | TBD | Open |
| R8 | Token budget exceeded on injection | High | Medium | 6 | Strict budget enforcement, truncation | TBD | Open |

---

## Detailed Risk Analysis

### R1: ONNX Model Size

**Description**: nomic-embed-text model is ~500MB, causing slow initial download.

**Impact**: 
- Poor first-run experience
- Users on slow connections may timeout
- Storage space concerns

**Mitigation Strategy**:
1. **Lazy Download**: Only download when first memory is created, not on installation
2. **Progress Indicator**: Show download progress with ETA
3. **Cache Location**: Use `~/.cache/opencode/models/` (shared across projects)
4. **Compression**: Use compressed model format if available
5. **Fallback**: Option to use smaller model (all-MiniLM-L6-v2, 80MB)

**Contingency**: If model download fails, plugin gracefully degrades to no-op mode

**Status**: Open | **Next Review**: After T005 implementation

---

### R2: LanceDB Platform Incompatibility

**Description**: LanceDB may not work on all platforms (especially older Linux distros or Windows).

**Impact**:
- Plugin fails to initialize
- Users cannot use memory feature
- Negative user experience

**Mitigation Strategy**:
1. **Cross-Platform Testing**: Test on Ubuntu 20.04+, macOS 12+, Windows 10+
2. **Fallback Storage**: Implement JSON-based storage if LanceDB fails
3. **Clear Error Messages**: Inform user of platform issue with workaround
4. **Dependencies Check**: Verify system requirements (glibc version, etc.)

**Contingency**: JSON fallback storage with linear search (acceptable for < 500 memories)

**Status**: Open | **Next Review**: After T004 implementation

---

### R3: Hook Integration Breaks OpenCode

**Description**: Plugin hooks may interfere with OpenCode's normal operation.

**Impact**:
- OpenCode crashes or hangs
- Sessions fail to start
- User loses trust in plugin

**Mitigation Strategy**:
1. **Minimal Hooks**: Start with only `session.idle` and `session.compacting`
2. **Error Handling**: Wrap all hook logic in try-catch
3. **Timeout Guards**: Set max execution time for hook operations
4. **Graceful Degradation**: If hook fails, log error but don't crash
5. **Kill Switch**: Config option to disable hooks entirely

**Contingency**: If hooks cause issues, release hotfix to disable plugin

**Status**: Open | **Next Review**: After T006, T007 implementation

---

### R4: Embedding Generation Too Slow

**Description**: ONNX inference may be slow on CPU-only machines.

**Impact**:
- Perception of plugin "hanging"
- Increased latency on session end
- User frustration

**Mitigation Strategy**:
1. **Async Processing**: Run embedding generation in background
2. **Batch Processing**: Generate embeddings for multiple memories at once
3. **Lightweight Model**: Use nomic-embed-text (optimized for CPU)
4. **Caching**: Don't re-embed identical content
5. **Progress Feedback**: Show "Processing memories..." message

**Acceptance Threshold**: < 500ms per memory on typical dev machine

**Contingency**: If too slow, switch to API-based embedding (optional config)

**Status**: Open | **Next Review**: After T005 implementation

---

### R5: Memory Database Grows Too Large

**Description**: Without proper curation, database could grow to GBs.

**Impact**:
- Storage space issues
- Slower search queries
- Degraded performance

**Mitigation Strategy**:
1. **Hard Cap**: Enforce 500 memories per project
2. **TTL Enforcement**: Auto-delete memories older than 90 days
3. **GC on Startup**: Run pruning on plugin initialization
4. **LRU Eviction**: Remove least-recently-used memories when cap reached
5. **User Visibility**: `/memory status` shows size and cap

**Acceptance Threshold**: < 100MB per project

**Contingency**: Implement aggressive GC if size exceeds threshold

**Status**: Open | **Next Review**: After T010 implementation

---

### R6: Plugin Conflicts with Other Plugins

**Description**: Memory plugin may conflict with other OpenCode plugins.

**Impact**:
- Errors or crashes
- Unexpected behavior
- Users blame our plugin

**Mitigation Strategy**:
1. **Namespace Isolation**: Use unique prefixes for all exports
2. **Dependency Isolation**: Bundle dependencies in plugin directory
3. **Testing**: Test with popular plugins (if available)
4. **Documentation**: List known incompatibilities

**Contingency**: If conflict found, document workaround or fix priority

**Status**: Open | **Next Review**: During T013 integration testing

---

### R7: Config Validation Fails Silently

**Description**: Invalid config.json may be loaded without warning.

**Impact**:
- Unpredictable behavior
- Hard-to-debug issues
- User confusion

**Mitigation Strategy**:
1. **Strict Schema**: Use JSON schema validation
2. **Early Validation**: Validate on plugin load
3. **Error Logging**: Log validation errors clearly
4. **Fallback to Defaults**: Use defaults for invalid values
5. **User Notification**: Show warning in OpenCode if config invalid

**Contingency**: Always fallback to defaults on validation failure

**Status**: Open | **Next Review**: After T009 implementation

---

### R8: Token Budget Exceeded on Injection

**Description**: Plugin may inject too many memories, exceeding token budget.

**Impact**:
- Context window overflow
- LLM errors
- Degraded agent performance

**Mitigation Strategy**:
1. **Strict Budget**: Hard limit of 4K tokens for memory injection
2. **Progressive Disclosure**: Inject only top-N most relevant
3. **Truncation**: Cut off memories if budget reached
4. **Token Estimation**: Estimate tokens before injection
5. **User Control**: Config option to adjust budget

**Acceptance Threshold**: Never exceed 4K tokens

**Contingency**: If exceeded, truncate memories aggressively

**Status**: Open | **Next Review**: After T007 implementation

---

## Risk Tracking

### Open Risks: 8
### Mitigated Risks: 0
### Closed Risks: 0

### High Priority Risks (Score >= 6): 5
- R1, R2, R3, R6, R8

### Actions Needed:
1. Schedule cross-platform testing (R2)
2. Prepare fallback storage implementation (R2)
3. Define hook timeout limits (R3)
4. Benchmark ONNX on target machines (R4)
5. Implement token estimation (R8)

---

## Sprint Retrospective Risk Review

*To be completed at end of sprint*

### Risks Realized:
- TBD

### Risks Mitigated:
- TBD

### New Risks Identified:
- TBD

### Lessons Learned:
- TBD
