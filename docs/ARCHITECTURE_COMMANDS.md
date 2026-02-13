# OpenKit Command Architecture

> Nova estrutura de comandos para o sistema OpenKit
> Versão: 2.0

---

## Overview

O sistema OpenKit utiliza **7 comandos principais** que cobrem todo o fluxo de desenvolvimento de software. Cada comando tem responsabilidade única e bem definida.

### Princípios

1. **Minimalismo** - Apenas comandos essenciais
2. **Unificação** - Comandos antigos unificados em novos
3. **Fluxo linear** - Discovery → Specify → Create → Verify
4. **Orquestração** - Para casos complexos, usar orchestrator

---

## Os 7 Comandos

### 1. /discover

**Propósito:** Análise inicial e contexto do projeto

**Responsabilidade:**
- Gerar contexto técnico do projeto
- Identificar stack, estrutura, riscos
- Documentar estado atual

**Quando usar:**
- Sempre no início de qualquer novo trabalho
- Obrigatório antes de `/specify`

**Artefatos gerados:**
- `docs/CONTEXT.md`
- `docs/SECURITY.md`
- `docs/QUALITY_GATES.md`

**Comandos antigos substituídos:**
- `/context`

---

### 2. /specify

**Propósito:** Especificação completa + Planejamento + Tasks

**Responsabilidade:**
- Criar especificação da feature
- Definir user stories e acceptance criteria
- Planejar tarefas com prioridades
- Criar breakdown de implementação

**Quando usar:**
- Para qualquer nova feature
- Para tarefas complexas que precisam de planejamento

**Artefatos gerados:**
- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/RISKS.md`
- `docs/requirements/<feature>/PLAN.md`
- `docs/sprint/Sprint-XX/TASKS.md`

**Comandos antigos substituídos:**
- `/specify`
- `/clarify`
- `/plan`
- `/tasks`

---

### 3. /create

**Propósito:** Implementação de código

**Responsabilidade:**
- Executar tarefas do plano
- Criar/modificar código
- Coordenar múltiplos domínios (se necessário)

**Quando usar:**
- Após `/specify` estar completo
- Para implementar features
- Para fazer alterações no código

**Fluxo interno:**
- P0: Foundation (DB + Security)
- P1: Core Backend
- P2: UI/UX
- P3: Polish (Tests + Perf)

**Comandos antigos substituídos:**
- `/impl`

---

### 4. /verify

**Propósito:** Verificação e validação

**Responsabilidade:**
- Executar testes
- Rodar linters
- Verificar segurança
- Checar performance
- Validar accessibility

**Quando usar:**
- Após `/create` completar
- Para verificar qualidade do código
- Antes de deploy

**Scripts executados:**
- Security scan
- Lint + Types
- Tests
- UX audit
- Lighthouse (se aplicável)

**Comandos antigos substituídos:**
- `/test`
- `/checklist`
- `/analyze`

---

### 5. /orchestrate

**Propósito:** Orquestração universal

**Responsabilidade:**
- Detectar complexidade da tarefa
- Coordenar múltiplos agentes
- Executar fluxo completo automaticamente

**Quando usar:**
- Tarefas complexas que envolvem múltiplos domínios
- Quando não se encaixa em um comando específico
- Para executar fluxo completo automaticamente

**Modo Router:**
- Detecta palavras-chave e redireciona para comando adequado
- "test..." → `/verify`
- "debug..." → `/debug`

**Modo Orchestrator:**
- Análise de complexidade
- Execução P0→P1→P2→P3
- Verificação automática

**Comandos antigos substituídos:**
- `/engineer`

---

### 6. /debug

**Propósito:** Debugging sistemático

**Responsabilidade:**
- Investigar bugs e erros
- Análise de root cause
- Propor e implementar soluções

**Quando usar:**
- Quando algo não está funcionando
- Para investigar erros
- Para debugging de produção

**Fases:**
1. Symptom Analysis
2. Information Gathering
3. Hypothesis Testing
4. Resolution

**Comandos antigos:**
- Mantém `/debug` (já era direto)

---

### 7. /deploy

**Propósito:** Deploy seguro

**Responsabilidade:**
- Preparar ambiente de deploy
- Executar deploy
- Verificar pós-deploy

**Quando usar:**
- Para enviar para staging
- Para enviar para produção
- Para gerenciar preview servers

**Comandos antigos:**
- Mantém `/deploy` (já era direto)
- Absorve `/preview`

---

## Fluxo de Trabalho

### Fluxo Padrão (Linear)

```
/discover     → /specify     → /create     → /verify     → /deploy
   (análise)     (design)      (build)       (check)        (ship)
```

### Fluxo Orquestrado

```
/orchestrate  →  (detecta complexidade)
                  ↓
              Execução automática
                  ↓
              /discover + /specify + /create + /verify
                  ↓
              Resultado
```

### Fluxo de Debug

```
/debug        →  Investigação
                  ↓
              Solução
                  ↓
              /verify (opcional)
```

---

## Mapeamento de Comandos Antigos

| Comando Antigo | Comando Novo | Status |
|----------------|--------------|--------|
| `/context` | `/discover` | Substituído |
| `/specify` | `/specify` | Mantido (unificado) |
| `/clarify` | `/specify` | Absorvido |
| `/plan` | `/specify` | Absorvido |
| `/tasks` | `/specify` | Absorvido |
| `/impl` | `/create` | Substituído |
| `/test` | `/verify` | Substituído |
| `/checklist` | `/verify` | Absorvido |
| `/analyze` | `/verify` | Absorvido |
| `/engineer` | `/orchestrate` | Substituído |
| `/debug` | `/debug` | Mantido |
| `/deploy` | `/deploy` | Mantido |
| `/ui-ux` | `/orchestrate` | Absorvido |
| `/doc` | `/orchestrate` | Absorvido |
| `/status` | `/verify` | Absorvido |
| `/preview` | `/deploy` | Absorvido |
| `/create` | `/orchestrate` | Absorvido |
| `/brainstorm` | `/discover` | Absorvido |

---

## Regras de Uso

### Obrigatoriedades

1. `/discover` é **sempre obrigatório** antes de `/specify`
2. `/specify` deve estar completo antes de `/create`
3. `/verify` deve passar antes de `/deploy`

### STOP Points

| Após | Pergunta |
|------|----------|
| `/discover` | "Contexto gerado. Prosseguir para especificação?" |
| `/specify` | "Plano registrado. Aprovar e implementar?" |
| `/create` (P0) | "Foundation completo. Prosseguir para P1?" |
| `/create` (P1) | "Backend completo. Prosseguir para P2?" |
| `/create` (P2) | "UI/UX completo. Prosseguir para P3?" |
| `/create` (P3) | "Implementação completa. Executar verificação?" |
| `/verify` | "Verificação completa. Prosseguir para deploy?" |
| `/deploy` | "Deploy completo. Confirmar sucesso?" |

---

## Integração com Agentes

### Agentes utilizados por comando

| Comando | Agentes Envolvidos |
|---------|-------------------|
| `/discover` | explorer-agent |
| `/specify` | project-planner, product-owner |
| `/create` | database-architect, security-auditor, backend-specialist, frontend-specialist, test-engineer |
| `/verify` | test-engineer, security-auditor, performance-optimizer |
| `/orchestrate` |Todos (coordenados) |
| `/debug` | debugger |
| `/deploy` | devops-engineer |

---

## Exemplo de Uso

### Exemplo 1: Nova Feature

```bash
# 1. Descobrir contexto
/discover

# 2. Especificar e planejar
/specify add user authentication

# 3. Implementar
/create from docs/sprint/Sprint-XX/TASKS.md

# 4. Verificar
/verify

# 5. Deploy
/deploy production
```

### Exemplo 2: Bug Fix

```bash
# Debugging
/debug login not working after update

# Verificar se corrigiu
/verify

# Deploy se necessário
/deploy staging
```

### Exemplo 3: Tarefa Complexa

```bash
# Orquestrador detecta complexidade e executa tudo
/orchestrate build e-commerce with Stripe checkout
```

---

## Resumo

| Comando | Fase | Entrada | Saída |
|---------|------|---------|-------|
| `/discover` | Discovery | - | CONTEXT.md, SECURITY.md |
| `/specify` | Specification | - | REQUIREMENTS.md, PLAN.md, TASKS.md |
| `/create` | Implementation | TASKS.md | Código |
| `/verify` | Verification | Código | Relatório de verificações |
| `/orchestrate` | All | - | Tudo (modo automático) |
| `/debug` | Investigation | Bug | Solução |
| `/deploy` | Deployment | Código | App no ar |

---

## Referências

- [[MASTER.md]] - Regras principais do sistema
- [[AGENTS.md]] - Lista de agentes disponíveis
- [[SKILLS.md]] - Skills e seus propósitos
