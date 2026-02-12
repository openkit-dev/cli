# Problem Statement: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft
**Owner**: TBD

## Contexto

O ecossistema OpenCode/OpenKit carece de mecanismo de persistencia de conhecimento entre sessoes. Cada nova sessao de agente inicia com contexto zerado, resultando em:

1. **Context drift** - Decisoes anteriores sao esquecidas ou contraditas
2. **Repeticao de erros** - Mesmos problemas sao resolvidos multiplas vezes
3. **Perda de padroes** - Convencoes do projeto nao sao mantidas
4. **Custo de tokens** - Re-descoberta de contexto a cada sessao

## Problema

Como persistir conhecimento tecnico relevante (decisoes, padroes, erros, specs) entre sessoes de forma seletiva, local-first, e transparente ao desenvolvedor?

## Restricoes

- **Local-first**: Sem dependencia de servicos externos
- **CLI-compatible**: Deve funcionar em ambiente terminal
- **Low-overhead**: Nao degradar performance do OpenCode
- **Privacy-aware**: Nao persistir dados sensiveis
- **Non-invasive**: Transparente no uso diario

## Nao-Objetivos

- Memoria conversacional completa (historico de chat)
- Context window infinita
- Memoria global entre projetos distintos
- Sincronizacao com servicos cloud

## Metricas de Sucesso

1. Reducao de context drift mensuravel em projetos longos
2. Zero latencia perceptivel na inicializacao de sessao
3. Precisao de recuperacao >= 70% em queries de teste
4. Adocao opt-in sem friccao

## Stakeholders

- Desenvolvedores usando OpenCode diariamente
- Equipe OpenKit/OpenCode (manutencao)
- Comunidade open-source (contribuicoes)

## Referencias

- [OpenCode Plugin System](https://opencode.ai/docs/plugins/)
- [OpenCode Hooks Documentation](https://opencode.ai/docs/plugins/#events)
- [LanceDB Documentation](https://lancedb.com/)

## Related

- [[docs/requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md]]
- [[docs/requirements/semantic-memory/USER_STORIES.md]]
- [[docs/requirements/semantic-memory/ACCEPTANCE_CRITERIA.md]]
- [[docs/requirements/semantic-memory/RISKS.md]]
