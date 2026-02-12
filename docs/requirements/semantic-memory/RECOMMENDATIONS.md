# Technical Recommendations: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft

## Stack Recomendada (MVP)

| Componente | Escolha | Justificativa |
|------------|---------|---------------|
| **Arquitetura** | Plugin OpenCode | Integracao nativa, hooks disponiveis, distribuicao npm |
| **Storage** | LanceDB | Performance, local-first, versionavel, zero-dependency |
| **Embeddings** | nomic-embed-text (ONNX) | Qualidade, local, sem custo, 768 dims |
| **Busca** | Hibrida (semantic + filtros) | Precisao sem complexidade excessiva |
| **Curadoria** | TTL + hard cap | Simples, previsivel, configuravel |
| **DX** | Invisivel + comandos opcionais | Nao intrusivo, debug quando necessario |

---

## Decisoes para MVP

| # | Decisao | Escolha | Alternativa Descartada |
|---|---------|---------|------------------------|
| 1 | Escopo de memoria | Por projeto apenas | Global entre projetos |
| 2 | Tipos de memoria | Namespace unico | Tipos separados (decisions, patterns, etc) |
| 3 | TTL padrao | 90 dias | Sem expiracao |
| 4 | Cap de memorias | 500 por projeto | Sem limite |
| 5 | Re-ranking | Nao | Cross-encoder ou LLM |
| 6 | Resumo de memorias | Nao (injecao literal) | Resumo automatico |
| 7 | Ponto de extracao | `session.idle` apenas | Tempo real por mensagem |
| 8 | Ponto de injecao | `session.compacting` | Modificacao de prompt |

---

## Roadmap Sugerido

### v0.1 (MVP)
- [ ] Storage LanceDB funcional
- [ ] Embedding local (nomic-embed-text via ONNX)
- [ ] Hooks: `session.idle`, `session.compacting`
- [ ] Extracao basica de decisoes
- [ ] CLI: `/memory status`, `/memory inspect`

### v0.2
- [ ] Tipos de memoria separados (decisions, patterns, errors, specs)
- [ ] Busca hibrida com filtros por tipo e recencia
- [ ] Comandos: `/memory search`, `/memory forget`
- [ ] Configuracao via `opencode.json`

### v0.3
- [ ] Decaimento de salience
- [ ] Garbage collection automatico periodico
- [ ] Export/import de memorias (JSON)
- [ ] Memoria global opcional (`~/.config/opencode/memory/`)

### v1.0
- [ ] Re-ranking opcional (cross-encoder)
- [ ] Resumo para memorias longas (> 500 tokens)
- [ ] Dashboard visual (TUI ou web)
- [ ] Metricas de uso e precisao

---

## Configuracao Default Sugerida

```json
{
  "$schema": "https://opencode.ai/config.json",
  "memory": {
    "enabled": true,
    "storage": {
      "type": "lancedb",
      "path": ".opencode/memory/index.lance"
    },
    "embedding": {
      "model": "nomic-embed-text",
      "runtime": "onnx",
      "fallback": "text-embedding-3-small"
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
}
```

---

## Principios Aplicados

| Principio | Como Foi Aplicado |
|-----------|-------------------|
| Simplicidade > sofisticacao prematura | MVP sem re-ranking, resumo, ou tipos separados |
| Contexto seletivo > contexto grande | Limite de 2-8K tokens para memoria, threshold de similaridade |
| Persistencia curada > memoria bruta | Extracao apenas de decisoes, padroes, specs |
| Local-first > infra externa | LanceDB + ONNX, zero servicos externos |
| Spec-driven > heuristica magica | Tipos de memoria definidos, configuracao explicita |

---

## Proximos Passos Imediatos

1. **Validar stack:** Prototipo minimo com LanceDB + nomic-embed
2. **Testar hooks:** Verificar comportamento de `session.idle` e `session.compacting`
3. **Definir schema:** Estrutura de metadata para memorias
4. **Implementar extracao:** Logica para identificar decisoes em respostas do LLM
5. **Criar testes:** Integracao com OpenCode, performance de embedding

---

## Metricas de Sucesso

| Metrica | Target MVP | Metodo de Medicao |
|---------|------------|-------------------|
| Precisao de recuperacao | >= 70% | Teste manual com queries conhecidas |
| Latencia de inicializacao | < 100ms | Profiling de `session.created` |
| Latencia de embedding | < 500ms por memoria | Profiling de extracao |
| Reducao de context drift | Qualitativo | Feedback de usuarios |
| Taxa de adocao | > 30% dos usuarios OpenCode | Analytics opt-in |

## Related

- [[requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[requirements/semantic-memory/ARCHITECTURE_ANALYSIS.md]]
- [[requirements/semantic-memory/REFERENCE_ANALYSIS.md]]
- [[requirements/semantic-memory/ALTERNATIVES_DISCARDED.md]]
