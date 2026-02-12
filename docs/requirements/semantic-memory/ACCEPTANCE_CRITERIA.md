# Acceptance Criteria: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft

## MVP (v0.1) - Criterios de Aceitacao

### AC-001: Plugin Inicializa Corretamente

**Dado** um projeto com `.opencode/` configurado
**Quando** o OpenCode iniciar
**Entao** o plugin de memoria deve:
- [ ] Inicializar sem erros
- [ ] Criar `.opencode/memory/` se nao existir
- [ ] Carregar configuracao de `config.json` se existir
- [ ] Usar defaults se config nao existir
- [ ] Nao bloquear inicializacao do OpenCode (< 100ms)

---

### AC-002: Extracao de Memorias ao Fim da Sessao

**Dado** uma sessao de OpenCode com respostas do LLM
**Quando** a sessao ficar idle (hook `session.idle`)
**Entao** o plugin deve:
- [ ] Analisar respostas da sessao
- [ ] Identificar decisoes, padroes, e solucoes
- [ ] Gerar embeddings para conteudo relevante
- [ ] Persistir em LanceDB
- [ ] Nao duplicar memorias existentes (similaridade > 0.95)
- [ ] Logar extracao em modo debug

---

### AC-003: Injecao de Contexto no Compaction

**Dado** memorias existentes em `.opencode/memory/`
**Quando** o OpenCode compactar a sessao (hook `session.compacting`)
**Entao** o plugin deve:
- [ ] Buscar memorias relevantes ao contexto atual
- [ ] Selecionar top N memorias (max 10)
- [ ] Respeitar budget de tokens (max 4K)
- [ ] Injetar em `output.context` no formato Markdown
- [ ] Adicionar indicador "[Memory: N items]" se configurado

---

### AC-004: Comando /memory status

**Dado** memorias existentes no projeto
**Quando** usuario executar `/memory status`
**Entao** deve exibir:
- [ ] Total de memorias armazenadas
- [ ] Espaco em disco usado
- [ ] Distribuicao por tipo (se tipos separados)
- [ ] Data da ultima extracao
- [ ] Configuracao ativa

---

### AC-005: Comando /memory inspect

**Dado** memorias existentes no projeto
**Quando** usuario executar `/memory inspect`
**Entao** deve exibir:
- [ ] Lista das 10 memorias mais recentes
- [ ] ID, conteudo truncado (100 chars), data
- [ ] Tipo de memoria (se aplicavel)
- [ ] Opcao de paginacao (`/memory inspect --page 2`)

---

### AC-006: Privacidade e Seguranca

**Dado** respostas do LLM com potencial conteudo sensivel
**Quando** extracao de memorias ocorrer
**Entao** o plugin deve:
- [ ] Filtrar padroes de API keys (regex: `[A-Za-z0-9]{32,}`)
- [ ] Ignorar conteudo de paths blacklistados (.env, *.pem, credentials.*)
- [ ] Nao persistir tokens ou senhas explicitas
- [ ] Logar warning se conteudo sensivel detectado

---

### AC-007: Performance

**Dado** uso normal do plugin
**Quando** operacoes de memoria ocorrerem
**Entao**:
- [ ] Inicializacao do plugin: < 100ms
- [ ] Geracao de embedding (por memoria): < 500ms
- [ ] Busca de memorias relevantes: < 200ms
- [ ] Injecao no compaction: < 50ms
- [ ] Nao degradar UX do OpenCode perceptivelmente

---

### AC-008: Desabilitacao do Plugin

**Dado** configuracao `memory.enabled: false` em `opencode.json`
**Quando** OpenCode iniciar
**Entao** o plugin deve:
- [ ] Nao registrar hooks
- [ ] Nao criar arquivos em `.opencode/memory/`
- [ ] Nao consumir recursos
- [ ] Comandos `/memory *` retornam "Memory plugin disabled"

---

## v0.2 - Criterios Adicionais

### AC-009: Busca Semantica

**Dado** memorias existentes
**Quando** usuario executar `/memory search "authentication flow"`
**Entao** deve:
- [ ] Gerar embedding da query
- [ ] Buscar memorias similares (cosine similarity)
- [ ] Retornar top 10 resultados ordenados por score
- [ ] Exibir score de similaridade (0.0 - 1.0)

---

### AC-010: Remocao de Memorias

**Dado** memoria com ID `mem_123abc`
**Quando** usuario executar `/memory forget mem_123abc`
**Entao** deve:
- [ ] Solicitar confirmacao
- [ ] Remover do LanceDB
- [ ] Atualizar indices
- [ ] Confirmar remocao

---

### AC-011: Configuracao via opencode.json

**Dado** configuracao customizada em `opencode.json`
**Quando** plugin inicializar
**Entao** deve:
- [ ] Ler e aplicar `memory.ttl_days`
- [ ] Ler e aplicar `memory.max_per_project`
- [ ] Ler e aplicar `memory.min_similarity`
- [ ] Validar configuracao e logar erros

---

## Definicao de Pronto (DoD)

Para uma User Story ser considerada completa:

1. [ ] Todos os Acceptance Criteria verificados
2. [ ] Testes unitarios passando (>= 80% coverage)
3. [ ] Testes de integracao com OpenCode
4. [ ] Documentacao atualizada
5. [ ] Code review aprovado
6. [ ] Sem regressoes de performance
7. [ ] Sem warnings de linter

## Related

- [[docs/requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[docs/requirements/semantic-memory/PROBLEM_STATEMENT.md]]
- [[docs/requirements/semantic-memory/USER_STORIES.md]]
- [[docs/sprint/Sprint-05/TASKS.md]]
