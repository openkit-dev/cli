# Risks: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft

## Risk Register

| ID | Risco | Probabilidade | Impacto | Score | Mitigacao |
|----|-------|---------------|---------|-------|-----------|
| R1 | Memoria irrelevante injetada no contexto | Alta | Medio | 6 | Threshold de similaridade alto (0.7+), limite de tokens |
| R2 | Performance de embedding lenta | Media | Alto | 6 | Batch processing, cache, modelo ONNX leve |
| R3 | Storage cresce indefinidamente | Media | Baixo | 3 | Hard cap por projeto, GC automatico |
| R4 | Conflito com updates do OpenCode | Baixa | Alto | 4 | API estavel, fallback gracioso, testes de integracao |
| R5 | Dados sensiveis persistidos | Media | Alto | 6 | Filtro de .env, tokens, secrets, regex patterns |
| R6 | Embeddings divergem entre modelos | Baixa | Medio | 2 | Hash de modelo, re-embedding quando muda |
| R7 | Custo de tokens na injecao | Media | Medio | 4 | Budget fixo (2-8K tokens), truncamento |
| R8 | Developer experience ruim | Media | Alto | 6 | Invisibilidade por padrao, comandos de debug claros |

---

## Detalhamento de Riscos Criticos

### R1: Memoria Irrelevante Injetada

**Descricao:** Busca semantica retorna memorias que parecem relevantes mas nao sao, poluindo o contexto do agente.

**Cenarios:**
- Termos ambiguos (ex: "auth" pode ser autenticacao ou autorizacao)
- Memorias desatualizadas ainda presentes
- Falsos positivos por similaridade de embedding

**Mitigacoes:**
1. Threshold de similaridade minimo (0.7 cosine)
2. Filtro por recencia (preferir memorias recentes)
3. Limite de memorias injetadas (max 10)
4. Feedback loop: permitir usuario marcar memoria como irrelevante

---

### R2: Performance de Embedding Lenta

**Descricao:** Geracao de embeddings locais consome tempo excessivo, degradando experiencia.

**Cenarios:**
- Primeira execucao (cold start do modelo ONNX)
- Batch grande de memorias para indexar
- Hardware limitado (sem GPU)

**Mitigacoes:**
1. Modelo leve (nomic-embed-text: ~500MB)
2. Cache de embeddings por hash de conteudo
3. Background processing (nao bloquear sessao)
4. Fallback para API em caso de timeout

---

### R5: Dados Sensiveis Persistidos

**Descricao:** Memorias contem tokens, senhas, ou dados confidenciais.

**Cenarios:**
- LLM menciona variavel de ambiente em resposta
- Codigo copiado com credenciais hardcoded
- Stack traces com dados sensiveis

**Mitigacoes:**
1. Regex filter para padroes conhecidos (API keys, tokens, passwords)
2. Blacklist de paths (.env, credentials.*, *.pem)
3. Revisao manual opcional antes de persistir
4. Encriptacao at-rest (opcional)

---

### R8: Developer Experience Ruim

**Descricao:** Plugin causa friccao ou confusao no uso diario.

**Cenarios:**
- Memorias incorretas sao usadas, gerando respostas erradas
- Usuario nao entende de onde veio certo contexto
- Plugin interfere com fluxo de trabalho

**Mitigacoes:**
1. Invisibilidade por padrao (zero output em uso normal)
2. Indicador claro quando memoria e usada ("[Memory: 3 items]")
3. Comandos de debug acessiveis (`/memory inspect`)
4. Opt-out simples via config

---

## Contingencias

| Risco | Se Ocorrer | Acao |
|-------|------------|------|
| R1 | Usuario reporta contexto poluido | Aumentar threshold, adicionar comando `/memory clear session` |
| R2 | Timeout em embedding | Fallback para API, log warning |
| R4 | Breaking change no OpenCode | Desabilitar plugin, notificar usuario, aguardar fix |
| R5 | Vazamento de dados | Purge imediato, revisao de filtros, notificacao |

## Related

- [[docs/requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[docs/requirements/semantic-memory/PROBLEM_STATEMENT.md]]
- [[docs/requirements/semantic-memory/ARCHITECTURE_ANALYSIS.md]]
- [[docs/sprint/Sprint-05/RISK_REGISTER.md]]
