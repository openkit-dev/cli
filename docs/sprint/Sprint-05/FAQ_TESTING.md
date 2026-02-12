# FAQ: Testando a Otimizacao de Contexto

**Criado**: 2026-02-09 13:45 (UTC-3)
**Sprint**: Sprint-05
**Topico**: Verificacao e Debug do Plugin de Memoria Semantica

---

## Pergunta 1: Como verificar se o OpenCode usa o contexto otimizado?

### Problema

O hook `experimental.session.compacting` injeta contexto, mas precisamos confirmar que:
1. O hook esta sendo chamado
2. O contexto injetado chega ao modelo
3. O agente esta usando as memorias injetadas

### Solucao Implementada

#### 1. Marcador Visual no Contexto

Quando `debug.show_injection_indicator: true` (padrao), o plugin injeta um comentario HTML:

```markdown
<!-- [SEMANTIC-MEMORY] Injected 5 memories (~1200 tokens) -->
## Project Memory (Optimized Context)

> This context was automatically retrieved based on semantic relevance.
> Total: 5 memories | ~1200 tokens

- [decision] Use LanceDB for storage: ...
- [pattern] Error handling pattern: ...
```

**Como verificar**: Peca ao agente para mostrar seu contexto atual ou procure por `[SEMANTIC-MEMORY]` nas respostas.

#### 2. Tool de Debug: `memory_debug`

```
Use memory_debug with action='status'
```

Retorna:
```
Plugin Status: ACTIVE
Session Cache: 5 memories loaded
Compaction Hook: TRIGGERED (or NOT YET TRIGGERED)
Extraction Hook: TRIGGERED (or NOT YET TRIGGERED)

Current Session Metrics:
  - Memories loaded at start: 5
  - Memories injected in compaction: 5
  - Tokens injected: ~1200
```

#### 3. Arquivo de Metricas

O plugin salva metricas em `.opencode/memory/metrics.json`:

```json
[
  {
    "sessionId": "abc123",
    "startTime": 1707494700000,
    "memoriesLoaded": 5,
    "memoriesInjected": 5,
    "tokensInjected": 1200,
    "compactionTriggered": true,
    "extractionTriggered": true
  }
]
```

#### 4. Logs Verbose

Com `debug.verbose: true` no config.json:

```
[semantic-memory] Session abc123 started
[semantic-memory] Query: "implement authentication..."
[semantic-memory] Loaded 5 relevant memories
[semantic-memory] Memory types: decision, pattern, error
[semantic-memory] Compaction triggered
[semantic-memory] Injected 5 memories (~1200 tokens)
```

---

## Pergunta 2: Como verificar reducao de tokens e drift?

### Metricas de Tokens

#### Script de Verificacao

Criamos `scripts/verify_optimization.ts`:

```bash
cd .opencode/plugins/semantic-memory
bun run scripts/verify_optimization.ts --verbose
```

Saida:
```
============================================================
SEMANTIC MEMORY VERIFICATION REPORT
============================================================

[PASS] Plugin Installation
   Plugin files are installed correctly

[PASS] Configuration
   Config loaded: token_budget=4000, max_results=10

[PASS] Database
   LanceDB database directory exists

[PASS] Session Metrics
   25 sessions tracked, 20 with memory injection

[PASS] Token Savings
   Estimated 75% token reduction (~6000 tokens saved per session)

[PASS] Context Relevance
   80% of recent sessions received relevant context

============================================================
SUMMARY: 6 passed, 0 warnings, 0 failed
============================================================
```

#### Tool: `memory_stats`

```
Use memory_stats
```

Retorna:
```
Semantic Memory Statistics:

Total Memories: 42
Total Tokens Stored: ~15000
Total Access Count: 156
Average Access per Memory: 3.7

By Type:
  - decision: 18
  - pattern: 12
  - error: 8
  - context: 4

Session History:
  - Sessions tracked: 25
  - Avg tokens injected per session: ~1500

Configuration:
  - Token budget: 4000
  - Max results: 10
  - Min similarity: 0.7
  - TTL: 90 days
  - Max per project: 500
```

### Calculo de Economia

| Metrica | Valor Estimado |
|---------|----------------|
| Contexto completo (sem otimizacao) | ~8000 tokens |
| Contexto otimizado (com memoria) | ~1500 tokens |
| Economia por sessao | ~6500 tokens (81%) |
| Economia por 100 sessoes | ~650,000 tokens |

**Nota**: O contexto "completo" assume que o OpenCode incluiria todo o historico da sessao. Com memoria semantica, apenas as partes relevantes sao injetadas.

### Deteccao de Drift

**Drift** = quando o contexto injetado nao e relevante para a tarefa atual.

#### Indicadores de Drift

1. **Access Count Baixo**: Memorias com `access_count < 2` apos 30 dias sao provavelmente irrelevantes
2. **Taxa de Injecao**: Se `memoriesInjected` e 0 em muitas sessoes, as memorias nao estao correspondendo
3. **Relevancia por Sessao**: Verificar se `memoriesLoaded > 0` mas `memoriesInjected == 0` (threshold muito alto)

#### Ajustes para Reduzir Drift

1. **Aumentar `min_similarity`** (0.7 -> 0.8): Memorias menos relevantes nao sao injetadas
2. **Reduzir `token_budget`** (4000 -> 2000): Menos memorias, mais focadas
3. **Reduzir `max_results`** (10 -> 5): Apenas as mais relevantes

---

## Pergunta 3: Precisamos de uma tool customizada?

### Resposta: SIM, e ja implementamos!

O plugin agora oferece **6 tools**:

| Tool | Proposito | Quando Usar |
|------|-----------|-------------|
| `memory_query` | Busca semantica | Quando precisa de contexto especifico |
| `memory_context` | **Injecao sob demanda** | Quando precisa de contexto otimizado para uma tarefa |
| `memory_save` | Salvamento manual | Quando toma uma decisao importante |
| `memory_stats` | Estatisticas | Para verificar saude do sistema |
| `memory_debug` | Debug | Para verificar se o plugin funciona |
| (hook automatico) | Injecao automatica | Durante compaction (passivo) |

### Tool Principal: `memory_context`

Esta e a tool mais importante para uso ativo:

```
Use memory_context with task="implement user authentication"
```

Retorna:
```markdown
## Optimized Context for: "implement user authentication"

**Retrieved:** 5 memories (~1200 tokens)
**Query matched:** 12 total memories

---

### [DECISION] Use JWT for session management
We decided to use JWT tokens instead of server-side sessions because...
> Salience: 0.85 | Accessed: 5x | Files: src/auth/jwt.ts

---

### [PATTERN] Error handling in auth routes
Always wrap auth operations in try-catch and return...
> Salience: 0.78 | Accessed: 3x | Files: src/routes/auth.ts

---

*This context was retrieved from semantic memory based on relevance to your task.*
```

### Fluxo de Uso Recomendado

```
Usuario: "Implemente autenticacao com OAuth"
           |
           v
Agente: [Automaticamente] usa memory_context
           |
           v
        Recebe contexto relevante:
        - Decisoes anteriores sobre auth
        - Padroes de codigo estabelecidos
        - Erros anteriores e solucoes
           |
           v
        Implementa com contexto otimizado
           |
           v
        [Ao final] usa memory_save para registrar novas decisoes
```

### Comparacao: Hook vs Tool

| Aspecto | Hook (Automatico) | Tool (Sob Demanda) |
|---------|-------------------|-------------------|
| **Controle** | Nenhum | Total |
| **Timing** | Durante compaction | A qualquer momento |
| **Visibilidade** | Transparente | Resultado visivel |
| **Customizacao** | Config global | Por chamada |
| **Garantia** | Pode nao acionar | Sempre funciona |

**Recomendacao**: Use ambos!
- Hook para injecao passiva durante compaction
- Tool para busca ativa quando precisa de contexto especifico

---

## Checklist de Verificacao

### Antes de Usar em Producao

- [ ] Plugin instalado (`openkit init --memory`)
- [ ] Config.json ajustado para seu caso de uso
- [ ] `debug.verbose: true` para primeiros testes
- [ ] Algumas memorias criadas (via sessoes ou `memory_save`)

### Durante Uso

- [ ] Verificar `memory_stats` periodicamente
- [ ] Monitorar `metrics.json` para trends
- [ ] Ajustar `min_similarity` se drift alto
- [ ] Ajustar `token_budget` se custo alto

### Verificacao de Funcionamento

```bash
# 1. Verificar instalacao
ls -la .opencode/plugins/semantic-memory/

# 2. Verificar config
cat .opencode/memory/config.json

# 3. Verificar database
ls -la .opencode/memory/index.lance/

# 4. Verificar metricas
cat .opencode/memory/metrics.json

# 5. Rodar script de verificacao
cd .opencode/plugins/semantic-memory
bun run scripts/verify_optimization.ts
```

### Dentro do OpenCode

```
# Verificar status
Use memory_debug with action='status'

# Ver estatisticas
Use memory_stats

# Testar busca
Use memory_query with query="authentication"

# Buscar contexto otimizado
Use memory_context with task="implement feature X"

# Salvar decisao importante
Use memory_save with type='decision', title='Use PostgreSQL', content='...'
```

---

## Proximos Passos

1. **Testar em ambiente real**: Rodar OpenCode com plugin ativo
2. **Acumular memorias**: Usar por algumas sessoes
3. **Verificar metricas**: Analisar `metrics.json`
4. **Ajustar config**: Otimizar thresholds baseado em dados reais
5. **Documentar patterns**: Criar guia de uso para o time

---

## Arquivos Relevantes

| Arquivo | Proposito |
|---------|-----------|
| `.opencode/plugins/semantic-memory/index.ts` | Plugin principal (469 linhas) |
| `.opencode/plugins/semantic-memory/lib/memory.ts` | Logica de memoria (~400 linhas) |
| `.opencode/plugins/semantic-memory/lib/storage.ts` | LanceDB wrapper (220 linhas) |
| `.opencode/plugins/semantic-memory/lib/embeddings.ts` | ONNX embeddings (185 linhas) |
| `.opencode/plugins/semantic-memory/scripts/verify_optimization.ts` | Script de verificacao |
| `.opencode/memory/config.json` | Configuracao |
| `.opencode/memory/metrics.json` | Metricas de sessao |
| `.opencode/memory/index.lance/` | Database LanceDB |

## Related

- [[docs/sprint/Sprint-05/HUB-SPRINT-05.md]]
- [[docs/sprint/Sprint-05/TASKS.md]]
- [[docs/sprint/Sprint-05/BUILD_VERIFICATION.md]]
- [[docs/QUALITY_GATES.md]]
