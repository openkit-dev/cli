# Alternatives Discarded: Plugin de Memoria Semantica

**Created**: 2026-02-09
**Status**: Draft

## 1. Historico Completo de Chat

**Descricao:** Persistir todas as mensagens de todas as sessoes.

**Por que descartado:**
- Volume massivo de dados (90%+ e ruido)
- Context drift significativo
- Custo de storage e embedding proibitivo
- Viola principio "contexto seletivo > contexto grande"
- Degradacao de performance em queries

---

## 2. Context Window Infinita

**Descricao:** Simular contexto ilimitado via concatenacao de historico.

**Por que descartado:**
- Fisicamente impossivel (limites de modelo: 128K-200K tokens)
- Custo proibitivo de tokens para cada requisicao
- Degradacao de performance em contextos longos (lost in the middle)
- Solucao preguicosa que ignora o problema real de selecao

---

## 3. Memoria Global Entre Projetos

**Descricao:** Compartilhar memorias entre todos os projetos do usuario.

**Por que descartado:**
- Risco de vazamento de contexto (decisao de projeto A influenciando projeto B)
- Contaminacao de padroes incompativeis
- Viola principio "local-first"
- Complexidade de permissoes e isolamento
- Potencial leak de informacao sensivel

---

## 4. Dependencia de Servicos Externos

**Descricao:** Usar Pinecone, Weaviate, ou outros servicos cloud para storage.

**Por que descartado:**
- Latencia de rede inaceitavel para uso CLI
- Custo recorrente (nao alinhado com open-source)
- Dependencia de disponibilidade de terceiros
- Privacidade e seguranca (dados enviados para cloud)
- Viola principio "local-first"

---

## 5. Memoria Baseada Apenas em Timestamps

**Descricao:** Priorizar memorias recentes por data de criacao.

**Por que descartado:**
- Recencia != relevancia (decisao de 6 meses atras pode ser mais importante que de ontem)
- Perde contexto importante mas antigo
- Nao captura relacoes semanticas entre memorias
- Ignora frequencia de acesso como sinal de importancia

---

## 6. Embedding via API Apenas

**Descricao:** Usar apenas OpenAI/Cohere para embeddings.

**Por que descartado:**
- Custo por uso (embora baixo)
- Latencia de rede
- Dependencia de conectividade
- Viola principio "local-first"

**Alternativa adotada:** ONNX models locais (nomic-embed-text) com fallback para API.

---

## 7. Re-ranking LLM em Todas as Queries

**Descricao:** Usar LLM para re-ordenar resultados de busca.

**Por que descartado:**
- Custo alto por query
- Latencia significativa (100-500ms adicional)
- Overengineering para MVP
- Beneficio marginal para volumes pequenos de memoria

**Alternativa adotada:** Sem re-ranking no MVP. Cross-encoder local se necessario.

---

## 8. Schema Dinamico de Tipos

**Descricao:** Inferir tipos de memoria automaticamente via LLM.

**Por que descartado:**
- Inconsistencia de classificacao
- Custo de classificacao por memoria
- Complexidade de schema evolution
- Viola principio "spec-driven > heuristica magica"

**Alternativa adotada:** Tipos fixos definidos em spec, com namespace unico no MVP.

---

## 9. Indexacao em Tempo Real

**Descricao:** Indexar cada mensagem do LLM instantaneamente.

**Por que descartado:**
- Overhead de CPU/memoria durante sessao
- Ruido de mensagens intermediarias
- Duplicacao de conteudo (mesma info em multiplas mensagens)

**Alternativa adotada:** Extracao apenas em `session.idle` (fim da sessao).

---

## 10. UI Grafica para Gerenciamento

**Descricao:** Dashboard web para visualizar/gerenciar memorias.

**Por que descartado para MVP:**
- Complexidade de desenvolvimento
- Viola principio "CLI-compatible"
- Overhead de manutencao

**Alternativa adotada:** Comandos CLI (`/memory status`, `/memory inspect`). UI opcional para v1.0+.
