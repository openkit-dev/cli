# Architecture Analysis: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft

## 1. Arquitetura Geral

### 1.1 Opcoes Arquiteturais

| Arquitetura | Descricao | Pros | Contras |
|-------------|-----------|------|---------|
| **A: Plugin OpenCode** | Plugin JS/TS usando sistema de hooks nativo | Integracao natural, acesso a eventos, facil distribuicao via npm | Dependencia do runtime OpenCode, limitado a eventos disponiveis |
| **B: MCP Server Local** | Servidor MCP local que expoe tools de memoria | Padrao aberto, compativel com outros agentes, isolamento | Overhead de processo separado, latencia adicional |
| **C: Hibrido (Plugin + Storage Go)** | Plugin para hooks + modulo Go no OpenKit para storage | Performance nativa, integracao profunda com OpenKit | Complexidade de desenvolvimento, dois runtimes |
| **D: Tool Customizada** | Custom tool em `.opencode/tools/` | Simples, sem dependencias externas | Sem acesso a eventos, apenas acionada pelo LLM |

### 1.2 Recomendacao: Arquitetura A (Plugin OpenCode)

**Justificativa:**
- O sistema de plugins do OpenCode oferece hooks robustos (`session.idle`, `session.created`, `tool.execute.after`, `experimental.session.compacting`)
- Acesso direto ao SDK client para interacao com o agente
- Distribuicao simplificada via npm ou `.opencode/plugins/`
- Alinhamento com principio local-first (plugin roda no mesmo processo)

### 1.3 Estrutura de Memoria no Filesystem

```
.opencode/
├── memory/
│   ├── index.lance/           # LanceDB database
│   ├── config.json            # Configuracao do plugin
│   └── snapshots/             # Backups periodicos
├── plugins/
│   └── semantic-memory.ts     # Plugin principal
└── tools/
    └── memory-query.ts        # Tool opcional para queries manuais
```

### 1.4 Memoria por Projeto vs Global

| Abordagem | Quando Usar |
|-----------|-------------|
| **Por projeto** (`.opencode/memory/`) | Contexto especifico do projeto, isolamento de dados |
| **Global** (`~/.config/opencode/memory/`) | Padroes gerais, preferencias do desenvolvedor |
| **Hibrido** | Memoria de projeto + referencia a padroes globais |

**Recomendacao MVP:** Por projeto apenas, para evitar vazamento de contexto entre projetos.

### 1.5 Tipos de Memoria

| Tipo | Descricao | Exemplo |
|------|-----------|---------|
| `decisions` | Decisoes arquiteturais e de design | "Usar Tanstack Query para data fetching" |
| `patterns` | Padroes de codigo observados | "Componentes seguem estrutura feature-first" |
| `errors` | Erros resolvidos e solucoes | "Fix: CORS issue com proxy reverso" |
| `specs` | Especificacoes e requisitos | "API deve retornar paginacao cursor-based" |
| `context` | Contexto de projeto (stack, deps) | "Next.js 14, App Router, Prisma" |

**Recomendacao:** Iniciar com namespace unico (`memories`) e evoluir para tipos separados conforme necessidade validada.

---

## 2. Persistencia e Armazenamento

### 2.1 Comparativo de Opcoes

| Opcao | Simplicidade | Performance | Versionamento | CLI-friendly | Acoplamento |
|-------|--------------|-------------|---------------|--------------|-------------|
| **LanceDB** | Alta | Excelente | Git-friendly (arquivos) | Otimo | Baixo |
| **ChromaDB** | Media | Boa | Requer backup manual | Requer server | Medio |
| **SQLite + FTS5** | Alta | Boa (text) | Git-friendly | Otimo | Baixo |
| **FAISS** | Baixa | Excelente | Arquivos binarios | Medio | Medio |
| **Qdrant** | Media | Excelente | Requer server | Requer server | Alto |
| **JSON + Embeddings file** | Muito Alta | Baixa | Git-friendly | Otimo | Nenhum |

### 2.2 LanceDB (Recomendado)

```
Pros:
- Zero-dependency (embedded, escrito em Rust)
- Formato colunar (eficiente para vetores)
- Arquivos locais (`.lance/`) - versionavel
- Query hibrida nativa (vetor + filtros)
- Suporte a TypeScript/Python

Contras:
- Biblioteca relativamente nova
- Documentacao menos madura que alternativas
```

### 2.3 Arquitetura de Storage Hibrida

```
┌─────────────────────────────────────────────┐
│              Storage Layer                   │
├─────────────────────────────────────────────┤
│  LanceDB                                     │
│  ├── vectors: Float32[768]                   │
│  ├── metadata: { type, salience, created }   │
│  └── content: string                         │
├─────────────────────────────────────────────┤
│  Metadata Index (SQLite ou JSON)             │
│  ├── memory_id                               │
│  ├── project_context                         │
│  ├── last_accessed                           │
│  └── access_count                            │
└─────────────────────────────────────────────┘
```

---

## 3. Estrategias de Embeddings

### 3.1 Modelo de Embedding

| Modelo | Dimensao | Performance | Custo | Disponibilidade |
|--------|----------|-------------|-------|-----------------|
| **all-MiniLM-L6-v2** | 384 | Rapido | Gratis (local) | ONNX |
| **text-embedding-3-small** | 1536 | API | $0.02/1M tokens | OpenAI |
| **nomic-embed-text** | 768 | Rapido | Gratis (local) | ONNX |
| **bge-small-en-v1.5** | 384 | Rapido | Gratis (local) | ONNX |

**Recomendacao MVP:** `nomic-embed-text` via ONNX (768 dims, bom trade-off qualidade/performance, roda local)

### 3.2 Chunking Strategy

| Tipo de Conteudo | Tamanho Chunk | Overlap |
|------------------|---------------|---------|
| Decisoes (curtas) | 256-512 tokens | 0% |
| Padroes (medios) | 512-1024 tokens | 10% |
| Specs (longos) | 1024-2048 tokens | 15% |
| Erros (variaveis) | Dinamico por secao | 10% |

### 3.3 Quando NAO Gerar Embedding

- Conteudo efemero (logs de debug, outputs brutos)
- Conteudo duplicado (ja existe embedding similar)
- Conteudo muito curto (< 50 tokens)
- Conteudo sensivel (tokens, senhas, .env)

### 3.4 Evitar Re-embedding

```typescript
// Hash do conteudo para detectar duplicatas
const contentHash = sha256(content)
const existing = await db.query({ hash: contentHash })
if (existing) return existing.id

// Similaridade para near-duplicates
const embedding = await embed(content)
const similar = await db.search(embedding, { threshold: 0.95 })
if (similar.length > 0) {
  return similar[0].id // Merge ou skip
}
```

---

## 4. Recuperacao de Contexto (Read Path)

### 4.1 Estrategias de Busca

| Estrategia | Quando Usar | Trade-offs |
|------------|-------------|------------|
| **Semantica pura** | Queries abertas | Pode trazer conteudo irrelevante |
| **Hibrida (semantic + filtros)** | Queries especificas | Melhor precisao, requer metadata |
| **Keyword + Semantic** | Termos tecnicos especificos | Combina exatidao e contexto |

**Recomendacao:** Busca hibrida com filtros por tipo e recencia

### 4.2 Re-ranking

| Abordagem | Complexidade | Beneficio |
|-----------|--------------|-----------|
| **Sem re-ranking** | Baixa | Suficiente para MVP |
| **Cross-encoder local** | Media | +10-15% precisao |
| **LLM re-ranking** | Alta | Melhor qualidade, alto custo |

**Recomendacao MVP:** Sem re-ranking.

### 4.3 Limites de Tokens na Injecao

```
Contexto total do OpenCode: ~128K-200K tokens (varia por modelo)
Budget para memoria: 2K-8K tokens (1-4% do contexto)

Estrategia de alocacao:
├── Decisoes recentes: 1K tokens (alta prioridade)
├── Padroes do projeto: 1K tokens (media prioridade)
├── Erros relevantes: 0.5K tokens (baixa prioridade)
└── Specs relacionadas: 1.5K tokens (variavel)
```

### 4.4 Resumo vs Injecao Literal

| Abordagem | Pros | Contras |
|-----------|------|---------|
| **Literal** | Precisao maxima, sem perda de info | Consome mais tokens |
| **Resumo** | Eficiente em tokens | Pode perder detalhes criticos |
| **Hibrido** | Resumo + snippets literais para codigo | Complexidade de implementacao |

**Recomendacao:** Injecao literal para MVP (simplicidade).

---

## 5. Curadoria e Politica de Esquecimento

### 5.1 Estrategias de Esquecimento

| Estrategia | Descricao | Implementacao MVP |
|------------|-----------|-------------------|
| **TTL por tipo** | Memorias expiram apos N dias | Sim |
| **Decaimento de salience** | Score decresce com tempo | Opcional |
| **Hard cap por projeto** | Max N memorias por projeto | Sim |
| **Remocao de lixo semantico** | Remove memorias nunca acessadas | Sim |

### 5.2 Configuracao Sugerida (MVP)

```json
{
  "memory": {
    "ttl": {
      "decisions": "90d",
      "patterns": "180d",
      "errors": "30d",
      "specs": "365d",
      "context": "never"
    },
    "limits": {
      "max_per_project": 500,
      "max_per_type": 100,
      "min_access_count": 0,
      "prune_threshold": 0.3
    },
    "decay": {
      "enabled": false,
      "half_life_days": 30
    }
  }
}
```

### 5.3 Garbage Collection

```typescript
async function pruneMemories() {
  const stale = await db.query({
    last_accessed: { $lt: Date.now() - TTL[type] },
    access_count: { $eq: 0 }
  })
  
  for (const memory of stale) {
    await db.delete(memory.id)
  }
}
```

---

## 6. Integracao com OpenCode

### 6.1 Hooks Disponiveis e Uso

| Hook | Uso para Memoria | Prioridade |
|------|------------------|------------|
| `session.created` | Carregar contexto inicial | P0 |
| `session.idle` | Extrair memorias da sessao | P0 |
| `session.compacted` | Capturar resumo de sessao | P1 |
| `tool.execute.after` | Capturar resultados de tools | P2 |
| `message.updated` | Monitorar respostas do LLM | P2 |
| `experimental.session.compacting` | Injetar contexto no compaction | P0 |

### 6.2 Exemplo de Plugin

```typescript
import type { Plugin } from "@opencode-ai/plugin"
import { SemanticMemory } from "./semantic-memory"

export const MemoryPlugin: Plugin = async (ctx) => {
  const memory = new SemanticMemory({
    dbPath: `${ctx.worktree}/.opencode/memory/index.lance`,
    config: `${ctx.worktree}/.opencode/memory/config.json`
  })

  await memory.initialize()

  return {
    "session.created": async (input, output) => {
      const relevant = await memory.query({
        project: ctx.worktree,
        limit: 10
      })
      memory.setSessionContext(relevant)
    },

    "experimental.session.compacting": async (input, output) => {
      const context = memory.getSessionContext()
      if (context.length > 0) {
        output.context.push(`## Project Memory\n${context.map(m => `- ${m.content}`).join('\n')}`)
      }
    },

    "session.idle": async (input) => {
      await memory.extractFromSession(input.sessionId, ctx.client)
    }
  }
}
```

### 6.3 O que NAO Modificar

- Nao modificar prompts do agente diretamente
- Nao interceptar tools de forma invasiva
- Nao alterar comportamento de permissions

---

## 7. Experiencia do Desenvolvedor (DX)

### 7.1 Matriz de Interacao

| Cenario | Visibilidade | Interacao |
|---------|--------------|-----------|
| **Uso normal** | Invisivel | Nenhuma |
| **Debug/troubleshooting** | Explicita | Comando `/memory inspect` |
| **Curadoria manual** | Explicita | Comando `/memory forget <id>` |
| **Configuracao** | Explicita | Arquivo `config.json` |

### 7.2 Comandos Sugeridos

```
/memory status       # Estatisticas de uso
/memory inspect      # Listar memorias recentes
/memory search <q>   # Buscar memorias
/memory forget <id>  # Remover memoria especifica
/memory clear        # Limpar todas as memorias
/memory export       # Exportar para JSON
```

### 7.3 Feedback ao Desenvolvedor

| Tipo | Quando | Como |
|------|--------|------|
| **Silencioso** | Operacao normal | Nenhum |
| **Toast sutil** | Memoria extraida | "Captured decision about auth flow" |
| **Log verboso** | Debug mode | Detalhes de operacoes |
