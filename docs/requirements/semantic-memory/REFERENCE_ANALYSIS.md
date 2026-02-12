# Reference Analysis: claude-mem

**Created**: 2026-02-09
**Status**: Draft
**Reference**: https://github.com/thedotmack/claude-mem

## Overview

claude-mem e um plugin de memoria para Claude Code com 26K+ stars que implementa persistencia de contexto entre sessoes. Esta analise identifica licoes aprendidas e pontos de melhoria para nosso plugin de memoria semantica no OpenCode/OpenKit.

---

## 1. Arquitetura do claude-mem

### 1.1 Componentes Principais

| Componente | Descricao |
|------------|-----------|
| **6 Lifecycle Hooks** | SessionStart, UserPromptSubmit, PostToolUse, Stop, SessionEnd + pre-hook |
| **Worker Service** | Express.js HTTP na porta 37777 com 10 endpoints de busca |
| **Database** | SQLite + FTS5 (Full-Text Search) + ChromaDB (vetorial) |
| **MCP Tools** | 4 ferramentas: search, timeline, get_observations, __IMPORTANT |
| **Web Viewer** | React UI para visualizacao de memorias em tempo real |
| **Claude Agent SDK** | Compressao de memorias via AI |

### 1.2 Fluxo de Dados

```
Hook (stdin) → SQLite → Worker Service → Claude SDK (compressao) → SQLite → Proxima Sessao
```

### 1.3 Schema do Banco

```sql
-- Sessoes
sdk_sessions (id, sdk_session_id, project, status, created_at, ...)

-- Observacoes (memorias individuais)
observations (id, session_id, tool_name, title, narrative, facts, concepts, type, files_read, files_modified, ...)

-- Resumos de sessao
session_summaries (id, sdk_session_id, request, investigated, learned, completed, next_steps, ...)

-- Prompts do usuario
user_prompts (id, sdk_session_id, prompt_text, ...)

-- FTS5 Virtual Tables
observations_fts, session_summaries_fts, user_prompts_fts
```

---

## 2. Pontos Fortes do claude-mem

### 2.1 Progressive Disclosure (3-Layer Workflow)

O padrao de busca em 3 camadas e elegante:

1. **search** - Retorna indice compacto com IDs (~50-100 tokens/resultado)
2. **timeline** - Contexto cronologico ao redor de um ponto
3. **get_observations** - Detalhes completos apenas para IDs filtrados (~500-1000 tokens)

**Economia de tokens: ~10x** comparado a buscar tudo de uma vez.

### 2.2 FTS5 + Busca Hibrida

- SQLite FTS5 para full-text search (sub-10ms)
- ChromaDB para busca semantica vetorial
- Triggers automaticos para sincronizacao FTS5

### 2.3 Estrutura de Observacao

Campos hierarquicos bem definidos:
- `title`, `subtitle` - Resumo curto
- `narrative` - Descricao completa
- `text` - Conteudo bruto
- `facts` - Fatos extraidos
- `concepts` - Tags semanticas
- `type` - Classificacao (decision, bugfix, feature, refactor, discovery, change)
- `files_read`, `files_modified` - Rastreamento de arquivos

### 2.4 Web Viewer

UI React com:
- Stream em tempo real via SSE
- Scroll infinito com deduplicacao
- Filtro por projeto
- Persistencia de settings

### 2.5 Compressao via AI

Usa Claude Agent SDK para comprimir observacoes, extraindo:
- Titulo e subtitulo
- Narrativa estruturada
- Fatos chave
- Conceitos/tags

---

## 3. Pontos Fracos / Oportunidades de Melhoria

### 3.1 Dependencia de Servico Externo (Claude API)

**Problema:** Compressao de memorias depende de chamadas a API Claude.
- Custo por compressao
- Latencia de rede
- Dependencia de disponibilidade

**Nossa abordagem:** Usar modelo local (nomic-embed-text via ONNX) para embeddings. Compressao opcional via LLM local ou API.

### 3.2 Complexidade Excessiva

**Problema:** 6 hooks + worker HTTP + MCP server + ChromaDB + React UI = muitas partes moveis.

**Nossa abordagem:** MVP minimalista:
- 2-3 hooks essenciais (session.idle, session.compacting)
- Storage unico (LanceDB)
- Sem UI separada inicialmente (comandos CLI)

### 3.3 Instalacao e Manutencao

**Problema:** Requer Bun, uv, ChromaDB - multiplas dependencias externas.

**Nossa abordagem:** 
- Unica dependencia: LanceDB (embedded)
- ONNX runtime para embeddings (bundled)
- Instalacao via flag simples: `openkit init --memory`

### 3.4 Token Cost no Startup

**Problema:** MCP tools adicionam ~2500 tokens ao contexto inicial.

**Nossa abordagem:** Injecao minima via `session.compacting`, nao como tools.

### 3.5 Sem Garbage Collection Automatico

**Problema:** claude-mem nao tem politica clara de expiracao/limpeza.

**Nossa abordagem:** TTL por tipo + hard cap + GC automatico semanal.

---

## 4. Licoes Aprendidas

### 4.1 Adotar: Progressive Disclosure

O padrao de 3 camadas (index → context → details) e excelente. Devemos adaptar:

```
Query → Semantic Search (LanceDB) → Return top-N IDs + scores
            ↓
    Optional: Expand context around ID
            ↓
    Inject only high-relevance content
```

### 4.2 Adotar: Estrutura de Observacao

Schema hierarquico com campos definidos:

```typescript
interface Memory {
  id: string
  type: 'decision' | 'pattern' | 'error' | 'spec' | 'context'
  title: string
  content: string
  facts: string[]
  concepts: string[]
  files: string[]
  salience: number
  created_at: number
  accessed_at: number
  access_count: number
}
```

### 4.3 Adotar: FTS5 como Fallback

FTS5 e excelente para busca exata de keywords tecnicas. Considerar busca hibrida:
- LanceDB para semantic search
- SQLite FTS5 para keyword match
- Merge de resultados

### 4.4 Adaptar: Tipos de Observacao

claude-mem usa: decision, bugfix, feature, refactor, discovery, change

Nossa proposta: decision, pattern, error, spec, context

### 4.5 NAO Adotar: Worker Service HTTP

Overhead desnecessario para nosso caso. Plugin OpenCode pode fazer tudo inline.

### 4.6 NAO Adotar: Compressao via API

Custo e latencia. Preferir embeddings locais sem compressao LLM.

### 4.7 NAO Adotar: UI Web Separada

Complexidade desnecessaria para MVP. CLI commands suficientes.

---

## 5. Comparativo de Abordagens

| Aspecto | claude-mem | Nossa Proposta |
|---------|------------|----------------|
| **Arquitetura** | Hooks + Worker HTTP + MCP | Plugin OpenCode unico |
| **Storage** | SQLite + ChromaDB | LanceDB (unico) |
| **Embeddings** | ChromaDB (vetorial) | nomic-embed-text (ONNX local) |
| **Compressao** | Claude API (AI) | Sem compressao (opcional) |
| **Busca** | FTS5 + ChromaDB hibrida | LanceDB vetorial + filtros |
| **Injecao** | MCP Tools (4) | Hook session.compacting |
| **UI** | React Web Viewer | CLI commands |
| **Dependencias** | Bun, uv, ChromaDB, Express | LanceDB, ONNX Runtime |
| **Custo** | API calls para compressao | Zero (local-only) |
| **Complexidade** | Alta (6 hooks, 3 servicos) | Baixa (1 plugin, 2-3 hooks) |

---

## 6. Impacto na Arquitetura OpenKit

### 6.1 Flag de Instalacao

```bash
# Init com memoria habilitada
openkit init --memory

# Sync com memoria habilitada
openkit sync opencode --memory
```

### 6.2 Estrutura de Arquivos

```
.opencode/
├── plugins/
│   └── semantic-memory/           # Plugin de memoria
│       ├── index.ts               # Plugin principal (hooks)
│       └── lib/
│           ├── memory.ts          # Core memory logic
│           ├── embeddings.ts      # ONNX embedding
│           └── storage.ts         # LanceDB wrapper
├── memory/
│   ├── index.lance/               # LanceDB database
│   ├── config.json                # Configuracao
│   └── model/                     # ONNX model cache
└── tools/
    └── memory-query.ts            # Tool opcional para queries manuais
```

### 6.3 Configuracao em opencode.json

```json
{
  "memory": {
    "enabled": true,
    "model": "nomic-embed-text",
    "retrieval": {
      "max_results": 10,
      "min_similarity": 0.7,
      "token_budget": 4000
    },
    "curation": {
      "ttl_days": 90,
      "max_per_project": 500
    }
  }
}
```

---

## 7. Recomendacoes Atualizadas

### 7.1 Simplificar Drasticamente

1. **Plugin unico** (nao worker HTTP separado)
2. **Storage unico** (LanceDB, nao SQLite + ChromaDB)
3. **2-3 hooks** (session.idle, session.compacting, opcional message.updated)
4. **CLI commands** (nao UI web)

### 7.2 Adotar Progressive Disclosure

Mesmo sem MCP tools, o principio se aplica:
1. Buscar top-10 por similaridade
2. Filtrar por threshold (0.7+)
3. Injetar apenas memorias relevantes (max 4K tokens)

### 7.3 Estrutura de Memoria Inspirada

```typescript
interface Memory {
  // Identificacao
  id: string
  project: string
  
  // Conteudo
  type: 'decision' | 'pattern' | 'error' | 'spec' | 'context'
  title: string
  content: string
  facts: string[]
  concepts: string[]
  files: string[]
  
  // Embedding
  vector: Float32Array
  
  // Lifecycle
  salience: number
  created_at: number
  accessed_at: number
  access_count: number
  expires_at: number | null
}
```

### 7.4 Extracao Simplificada

Em vez de compressao AI complexa:
1. Detectar padroes de decisao via regex/heuristica
2. Extrair titulo do primeiro paragrafo
3. Gerar embedding do conteudo completo
4. Extrair conceitos via keyword extraction (TF-IDF ou similar)

### 7.5 Flag --memory no OpenKit

```go
// cmd/openkit/init.go
func initCmd() *cobra.Command {
  var enableMemory bool
  
  cmd := &cobra.Command{
    Use: "init",
    Run: func(cmd *cobra.Command, args []string) {
      // ... existing init logic ...
      
      if enableMemory {
        installMemoryPlugin(projectRoot)
      }
    },
  }
  
  cmd.Flags().BoolVar(&enableMemory, "memory", false, "Enable semantic memory plugin")
  return cmd
}
```

---

## 8. Proximos Passos

1. [ ] Implementar flag `--memory` no OpenKit CLI
2. [ ] Criar estrutura basica do plugin em `.opencode/plugins/semantic-memory/`
3. [ ] Integrar LanceDB como storage
4. [ ] Implementar embeddings ONNX locais
5. [ ] Criar hooks session.idle e session.compacting
6. [ ] Testar extracao de decisoes
7. [ ] Implementar comandos CLI `/memory status`, `/memory inspect`

---

## 9. Conclusao

O claude-mem oferece uma implementacao madura de memoria para agentes, mas com complexidade excessiva para nosso caso de uso. Devemos:

**Adotar:**
- Progressive disclosure (3-layer workflow conceitual)
- Estrutura de observacao hierarquica
- Tipos de memoria bem definidos
- FTS5 como fallback para keywords

**Simplificar:**
- Unico plugin (sem worker HTTP)
- Unico storage (LanceDB)
- Embeddings locais (ONNX)
- CLI commands (sem UI web)

**Adicionar:**
- Flag `--memory` no OpenKit init/sync
- TTL e garbage collection automatico
- Configuracao via opencode.json

## Related

- [[requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[requirements/semantic-memory/ARCHITECTURE_ANALYSIS.md]]
- [[requirements/semantic-memory/ALTERNATIVES_DISCARDED.md]]
- [[requirements/semantic-memory/RECOMMENDATIONS.md]]
