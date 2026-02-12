# OpenKit Integration: Flag --memory

**Created**: 2026-02-09
**Status**: Draft

## Objetivo

Adicionar flag `--memory` aos comandos `openkit init` e `openkit sync` para habilitar instalacao automatica do plugin de memoria semantica local.

---

## 1. Especificacao da Flag

### 1.1 Comando `openkit init`

```bash
# Uso basico
openkit init my-app --memory

# Combinado com outras flags
openkit init my-app --ai opencode --memory
openkit init --here --memory
```

**Comportamento:**
1. Executa init normal (templates, git, etc.)
2. Cria estrutura `.opencode/plugins/semantic-memory/`
3. Copia plugin de memoria do embedded templates
4. Cria `.opencode/memory/` com config padrao
5. Adiciona `memory.enabled: true` ao opencode.json

### 1.2 Comando `openkit sync`

```bash
# Habilitar memoria em projeto existente
openkit sync opencode --memory

# Desabilitar memoria
openkit sync opencode --no-memory
```

**Comportamento:**
1. Se `--memory`: instala plugin se nao existir
2. Se `--no-memory`: remove plugin e config (preserva dados em `.opencode/memory/`)
3. Atualiza opencode.json conforme flag

---

## 2. Estrutura de Arquivos Instalada

```
.opencode/
├── plugins/
│   └── semantic-memory/
│       ├── index.ts                 # Plugin principal
│       ├── package.json             # Dependencias (lancedb, onnx)
│       └── lib/
│           ├── memory.ts            # Core memory logic
│           ├── embeddings.ts        # ONNX embedding
│           └── storage.ts           # LanceDB wrapper
├── memory/
│   ├── config.json                  # Configuracao do plugin
│   └── .gitignore                   # Ignorar index.lance/
└── opencode.json                    # Atualizado com memory config
```

### 2.1 Arquivo config.json Padrao

```json
{
  "version": "1.0.0",
  "embedding": {
    "model": "nomic-embed-text",
    "runtime": "onnx"
  },
  "retrieval": {
    "max_results": 10,
    "min_similarity": 0.7,
    "token_budget": 4000
  },
  "curation": {
    "ttl_days": 90,
    "max_per_project": 500,
    "prune_unused_after_days": 30
  },
  "extraction": {
    "on_session_idle": true,
    "patterns": ["decision", "architecture", "pattern", "fix", "solution"]
  },
  "debug": {
    "verbose": false,
    "show_injection_indicator": true
  }
}
```

### 2.2 Atualizacao do opencode.json

```json
{
  "$schema": "https://opencode.ai/config.json",
  "instructions": [".opencode/rules/MASTER.md"],
  "plugin": ["./plugins/semantic-memory"],
  "memory": {
    "enabled": true
  }
}
```

---

## 3. Implementacao em Go

### 3.1 Adicionar Flag ao init.go

```go
// internal/cli/init.go

var (
	flagAgent  string
	flagHere   bool
	flagForce  bool
	flagNoGit  bool
	flagMemory bool  // NOVO
)

func init() {
	initCmd.Flags().StringVar(&flagAgent, "ai", "", "AI agent to configure")
	initCmd.Flags().BoolVar(&flagHere, "here", false, "Initialize in current directory")
	initCmd.Flags().BoolVar(&flagForce, "force", false, "Overwrite existing files")
	initCmd.Flags().BoolVar(&flagNoGit, "no-git", false, "Skip git initialization")
	initCmd.Flags().BoolVar(&flagMemory, "memory", false, "Enable semantic memory plugin")  // NOVO
}

func runInit(args []string) {
	// ... existing code ...

	// Extract templates
	printInfo("Extracting templates...")
	if err := templates.Extract(projectDir, agent); err != nil {
		exitWithError(fmt.Sprintf("Failed to extract templates: %v", err))
	}

	// NOVO: Install memory plugin if requested
	if flagMemory {
		printInfo("Installing semantic memory plugin...")
		if err := installMemoryPlugin(projectDir); err != nil {
			printWarning(fmt.Sprintf("Memory plugin installation failed: %v", err))
		} else {
			printInfo("Semantic memory enabled")
		}
	}

	// ... rest of function ...
}

// NOVO
func installMemoryPlugin(projectDir string) error {
	// 1. Create directories
	pluginDir := filepath.Join(projectDir, ".opencode", "plugins", "semantic-memory")
	memoryDir := filepath.Join(projectDir, ".opencode", "memory")
	
	if err := os.MkdirAll(pluginDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(memoryDir, 0755); err != nil {
		return err
	}

	// 2. Extract memory plugin from embedded templates
	if err := templates.ExtractMemoryPlugin(pluginDir); err != nil {
		return err
	}

	// 3. Create default config.json
	if err := createMemoryConfig(memoryDir); err != nil {
		return err
	}

	// 4. Create .gitignore for memory data
	gitignore := filepath.Join(memoryDir, ".gitignore")
	if err := os.WriteFile(gitignore, []byte("index.lance/\nmodel/\n"), 0644); err != nil {
		return err
	}

	// 5. Update opencode.json
	if err := updateOpencodeJsonMemory(projectDir, true); err != nil {
		return err
	}

	return nil
}
```

### 3.2 Adicionar Flag ao sync (agent_targets.go)

```go
// internal/cli/agent_targets.go

var (
	flagSyncMemory   bool
	flagSyncNoMemory bool
)

func init() {
	// ... existing flags ...
	agentTargetsCmd.Flags().BoolVar(&flagSyncMemory, "memory", false, "Enable semantic memory plugin")
	agentTargetsCmd.Flags().BoolVar(&flagSyncNoMemory, "no-memory", false, "Disable semantic memory plugin")
}

func runAgentTargets(cmd *cobra.Command, args []string) {
	// ... existing sync logic ...

	// Handle memory flag
	if flagSyncMemory {
		if err := installMemoryPlugin(projectRoot); err != nil {
			printWarning(fmt.Sprintf("Memory plugin installation failed: %v", err))
		}
	} else if flagSyncNoMemory {
		if err := disableMemoryPlugin(projectRoot); err != nil {
			printWarning(fmt.Sprintf("Memory plugin removal failed: %v", err))
		}
	}
}
```

### 3.3 Templates Embarcados

```go
// internal/templates/memory/embed.go

package memory

import "embed"

//go:embed plugin/*
var PluginFS embed.FS

//go:embed config.json
var DefaultConfig []byte
```

---

## 4. Plugin de Memoria (TypeScript)

### 4.1 index.ts (Entry Point)

```typescript
// .opencode/plugins/semantic-memory/index.ts

import type { Plugin } from "@opencode-ai/plugin"
import { SemanticMemory } from "./lib/memory"

export const MemoryPlugin: Plugin = async (ctx) => {
  const configPath = `${ctx.worktree}/.opencode/memory/config.json`
  const dbPath = `${ctx.worktree}/.opencode/memory/index.lance`
  
  const memory = new SemanticMemory({ configPath, dbPath })
  
  try {
    await memory.initialize()
  } catch (err) {
    console.error("[semantic-memory] Failed to initialize:", err)
    return {}
  }

  return {
    // Carregar contexto relevante no inicio
    "session.created": async (input, output) => {
      const relevant = await memory.getRelevantContext(input.sessionId, 10)
      memory.setSessionCache(relevant)
    },

    // Injetar memoria no compaction
    "experimental.session.compacting": async (input, output) => {
      const memories = memory.getSessionCache()
      if (memories.length > 0) {
        const formatted = memories.map(m => `- [${m.type}] ${m.title}: ${m.content}`).join('\n')
        output.context.push(`## Project Memory\n${formatted}`)
      }
    },

    // Extrair conhecimento ao fim da sessao
    "session.idle": async (input) => {
      await memory.extractFromSession(input.sessionId, ctx.client)
    },

    // Tool customizada para queries manuais
    tool: {
      memory_query: {
        description: "Query project memory for relevant context",
        args: {
          query: { type: "string", description: "Search query" },
          limit: { type: "number", description: "Max results (default: 5)" }
        },
        async execute(args) {
          const results = await memory.search(args.query, args.limit || 5)
          return results.map(r => `[${r.type}] ${r.title}: ${r.content}`).join('\n')
        }
      }
    }
  }
}
```

### 4.2 package.json

```json
{
  "name": "semantic-memory",
  "version": "0.1.0",
  "type": "module",
  "dependencies": {
    "lancedb": "^0.4.0",
    "onnxruntime-node": "^1.17.0"
  }
}
```

---

## 5. Acceptance Criteria

### AC-001: Flag --memory no init

**Dado** usuario executando `openkit init my-app --memory`
**Quando** init completar
**Entao**:
- [ ] Diretorio `.opencode/plugins/semantic-memory/` existe
- [ ] Diretorio `.opencode/memory/` existe
- [ ] Arquivo `.opencode/memory/config.json` existe com valores padrao
- [ ] Arquivo `.opencode/memory/.gitignore` ignora `index.lance/` e `model/`
- [ ] opencode.json contem `"memory": { "enabled": true }`
- [ ] opencode.json contem plugin path em `"plugin": ["./plugins/semantic-memory"]`

### AC-002: Flag --memory no sync

**Dado** projeto existente sem memoria
**Quando** usuario executar `openkit sync opencode --memory`
**Entao**:
- [ ] Plugin de memoria e instalado
- [ ] Config e criado
- [ ] opencode.json e atualizado

### AC-003: Flag --no-memory no sync

**Dado** projeto com memoria habilitada
**Quando** usuario executar `openkit sync opencode --no-memory`
**Entao**:
- [ ] Plugin de memoria e removido
- [ ] `memory.enabled` e `false` no opencode.json
- [ ] Dados em `.opencode/memory/index.lance/` sao preservados

### AC-004: Instalacao de Dependencias

**Dado** plugin de memoria instalado
**Quando** OpenCode iniciar
**Entao**:
- [ ] Bun instala dependencias automaticamente
- [ ] LanceDB inicializa sem erros
- [ ] ONNX model e baixado na primeira execucao

---

## 6. Riscos e Mitigacoes

| Risco | Mitigacao |
|-------|-----------|
| ONNX model grande (~500MB) | Download assincrono, cache em `~/.cache/opencode/models/` |
| LanceDB incompativel com plataforma | Fallback para JSON simples |
| Bun nao instalado | Erro claro com instrucoes |
| OpenCode < 1.0 sem suporte a plugins | Verificar versao, warning se incompativel |

---

## 7. Proximos Passos

1. [ ] Implementar flag `--memory` no `internal/cli/init.go`
2. [ ] Implementar flag `--memory` no `internal/cli/agent_targets.go`
3. [ ] Criar templates embarcados em `internal/templates/memory/`
4. [ ] Implementar `installMemoryPlugin()` e `disableMemoryPlugin()`
5. [ ] Criar plugin TypeScript basico
6. [ ] Testar instalacao e inicializacao
7. [ ] Documentar uso no README

## Related

- [[requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[requirements/semantic-memory/ARCHITECTURE_ANALYSIS.md]]
- [[requirements/semantic-memory/RECOMMENDATIONS.md]]
- [[sprint/Sprint-05/TASKS.md]]
