# User Stories: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft

## Epic: Memoria Semantica Entre Sessoes

### US-001: Persistencia Automatica de Decisoes

**Como** desenvolvedor usando OpenCode
**Quero** que decisoes arquiteturais importantes sejam automaticamente salvas
**Para** que nao precise repetir o mesmo contexto em sessoes futuras

**Criterios de Aceitacao:**
- [ ] Decisoes sao extraidas automaticamente ao fim da sessao
- [ ] Extracao identifica palavras-chave como "decidimos", "vamos usar", "escolhemos"
- [ ] Memorias sao persistidas em `.opencode/memory/`
- [ ] Sem intervencao manual do usuario

**Prioridade:** P0

---

### US-002: Recuperacao de Contexto de Projeto

**Como** desenvolvedor iniciando uma nova sessao
**Quero** que o agente tenha acesso a decisoes e padroes anteriores
**Para** manter consistencia entre sessoes

**Criterios de Aceitacao:**
- [ ] Contexto relevante e carregado no inicio da sessao
- [ ] Memorias sao injetadas no prompt de forma transparente
- [ ] Limite de tokens respeitado (max 4K tokens de memoria)
- [ ] Apenas memorias com alta relevancia sao injetadas

**Prioridade:** P0

---

### US-003: Inspecao de Memorias

**Como** desenvolvedor
**Quero** poder visualizar quais memorias estao armazenadas
**Para** entender o que o agente "lembra" sobre o projeto

**Criterios de Aceitacao:**
- [ ] Comando `/memory inspect` lista memorias recentes
- [ ] Exibe: conteudo, data de criacao, tipo, frequencia de uso
- [ ] Ordenacao por relevancia ou recencia
- [ ] Paginacao para listas longas

**Prioridade:** P1

---

### US-004: Remocao de Memorias Incorretas

**Como** desenvolvedor
**Quero** poder remover memorias especificas
**Para** corrigir informacoes desatualizadas ou incorretas

**Criterios de Aceitacao:**
- [ ] Comando `/memory forget <id>` remove memoria
- [ ] Confirmacao antes de remover
- [ ] Comando `/memory clear` limpa todas as memorias (com confirmacao)
- [ ] Operacao e reversivel via backup

**Prioridade:** P1

---

### US-005: Busca em Memorias

**Como** desenvolvedor
**Quero** buscar memorias por texto ou conceito
**Para** encontrar decisoes ou padroes especificos

**Criterios de Aceitacao:**
- [ ] Comando `/memory search <query>` retorna memorias relevantes
- [ ] Busca e semantica (nao apenas keyword)
- [ ] Resultados ordenados por relevancia
- [ ] Exibe score de similaridade

**Prioridade:** P2

---

### US-006: Expiracao Automatica de Memorias

**Como** desenvolvedor
**Quero** que memorias antigas expirem automaticamente
**Para** evitar acumulo de informacao desatualizada

**Criterios de Aceitacao:**
- [ ] TTL configuravel por tipo de memoria
- [ ] Padroes razoaveis (90 dias para decisoes, 30 dias para erros)
- [ ] Memorias nunca acessadas expiram mais rapido
- [ ] GC executa periodicamente (semanal)

**Prioridade:** P2

---

### US-007: Indicador de Uso de Memoria

**Como** desenvolvedor
**Quero** saber quando o agente usou memorias na resposta
**Para** entender de onde veio certo contexto

**Criterios de Aceitacao:**
- [ ] Indicador sutil quando memoria e injetada
- [ ] Formato: "[Memory: 3 items used]"
- [ ] Desabilitavel via config
- [ ] Modo debug mostra quais memorias foram usadas

**Prioridade:** P2

---

### US-008: Configuracao de Memoria

**Como** desenvolvedor
**Quero** configurar comportamento da memoria via `opencode.json`
**Para** adaptar o plugin ao meu fluxo de trabalho

**Criterios de Aceitacao:**
- [ ] Opcao para desabilitar memoria completamente
- [ ] Configuracao de TTL por tipo
- [ ] Limite de memorias por projeto
- [ ] Threshold de similaridade configuravel

**Prioridade:** P1

---

### US-009: Export de Memorias

**Como** desenvolvedor
**Quero** exportar memorias para arquivo JSON
**Para** backup ou migracao entre projetos

**Criterios de Aceitacao:**
- [ ] Comando `/memory export` gera JSON
- [ ] Inclui conteudo, metadata, embeddings (opcional)
- [ ] Formato compativel com import
- [ ] Exporta para stdout ou arquivo

**Prioridade:** P3

---

### US-010: Import de Memorias

**Como** desenvolvedor
**Quero** importar memorias de arquivo JSON
**Para** restaurar backup ou migrar conhecimento

**Criterios de Aceitacao:**
- [ ] Comando `/memory import <file>` carrega memorias
- [ ] Valida formato do arquivo
- [ ] Detecta duplicatas (merge ou skip)
- [ ] Re-gera embeddings se modelo mudou

**Prioridade:** P3

---

## Matriz de Prioridade

| Prioridade | User Stories | Sprint Target |
|------------|--------------|---------------|
| P0 | US-001, US-002 | v0.1 (MVP) |
| P1 | US-003, US-004, US-008 | v0.2 |
| P2 | US-005, US-006, US-007 | v0.3 |
| P3 | US-009, US-010 | v1.0 |

## Related

- [[requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[requirements/semantic-memory/PROBLEM_STATEMENT.md]]
- [[requirements/semantic-memory/ACCEPTANCE_CRITERIA.md]]
- [[requirements/semantic-memory/RISKS.md]]
