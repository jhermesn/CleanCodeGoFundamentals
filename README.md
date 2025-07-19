# Estudo de Caso: Refatoração com Clean Code e Clean Architecture em Go

Este repositório é um estudo de caso prático que demonstra a aplicação progressiva de princípios de **Boas Práticas**, **Clean Code**, **SOLID** e **Clean Architecture** para refatorar uma aplicação em Go. O objetivo é mostrar como um código inicialmente monolítico e problemático pode ser transformado em uma solução robusta, modular, testável e de fácil manutenção.

O projeto foi apresentado como avaliação final da disciplina de Fundamentos de Engenharia de Software em Julho de 2025, ministrada pelo professor Anderson Costa na UEPA. A análise completa dos princípios teóricos que guiaram esta refatoração pode ser encontrada nos documentos na pasta [`docs/`](./docs/).

## O Cenário

O ponto de partida é uma única função responsável por calcular o preço final de um pedido. Essa função precisa considerar:
- A soma dos preços dos produtos.
- Descontos baseados no tipo de cliente (Premium, Plus).
- Descontos de cupons.
- A aplicação de uma taxa de envio para pedidos abaixo de um certo valor.

---

## A Jornada de Refatoração

A evolução do código é dividida em 11 passos, cada um focado em resolver um ou mais "code smells" específicos, aplicando gradualmente os princípios de engenharia de software.

### Passo 1: O Ponto de Partida
- **Arquivo:** [`steps/passo1/passo1.go`](./steps/passo1/passo1.go)
- **Descrição:** O código inicial é uma função monolítica com múltiplos problemas, servindo como um exemplo clássico de débito técnico.
- **Violações Principais:**
  - **Nomes Não Significativos:** Nomes de função e variáveis curtos e crípticos (`Cp`, `p`, `u`, `c`).
  - **Múltiplas Responsabilidades (Violação do SRP):** A função calcula subtotal, descontos e taxas em um único bloco de código.
  - **Números Mágicos:** Valores como `0.9`, `100` e `1.05` estão "hardcoded", sem explicação de seu propósito.
  - **Tipos de Dados Genéricos:** Uso de `map[string]string` para representar um usuário, uma modelagem de domínio fraca.

### Passo 2: Nomes Significativos
- **Arquivo:** [`steps/passo2/passo2.go`](./steps/passo2/passo2.go)
- **Mudanças:** A função e as variáveis foram renomeadas para nomes claros e autoexplicativos (`CalcularPrecoTotal`, `precos`, `usuario`, `cupom`).
- **Princípio Aplicado:** **Clean Code (Nomes Significativos)**. A legibilidade é drasticamente melhorada, tornando a intenção do código compreensível.

### Passo 3: Eliminar Números Mágicos
- **Arquivo:** [`steps/passo3/passo3.go`](./steps/passo3/passo3.go)
- **Mudanças:** "Números mágicos" e "strings mágicas" foram extraídos para constantes nomeadas (`DescontoPremium`, `ValorMinimoParaTaxa`, `TipoUsuarioPremium`).
- **Princípio Aplicado:** **Clean Code (Código Autoexplicativo)** e **DRY (Don't Repeat Yourself)**. O código se torna mais fácil de manter, pois os valores de configuração estão centralizados e seu significado é explícito.

### Passo 4: Introduzir Tipos de Domínio
- **Arquivo:** [`steps/passo4/passo4.go`](./steps/passo4/passo4.go)
- **Mudanças:** Os tipos de dados genéricos são substituídos por `structs` que modelam o negócio (`Pedido`, `Usuario`, `Produto`).
- **Princípio Aplicado:** **Clean Architecture (Entidades)**. O código começa a refletir o domínio real, tornando-se mais expressivo e seguro contra erros de tipo.

### Passo 5: Funções com Responsabilidade Única
- **Arquivo:** [`steps/passo5/passo5.go`](./steps/passo5/passo5.go)
- **Mudanças:** A função monolítica é decomposta em funções menores e focadas, cada uma com uma única responsabilidade (`calcularSubtotal`, `aplicarDescontoPorUsuario`, etc.).
- **Princípio Aplicado:** **SOLID (Single Responsibility Principle - SRP)**. A função principal torna-se um orquestrador, melhorando a modularidade, legibilidade e, crucialmente, a testabilidade do código.

### Passo 6: Melhorar a Coesão com Métodos e Serviços
- **Arquivo:** [`steps/passo6/passo6.go`](./steps/passo6/passo6.go)
- **Mudanças:**
  - `calcularSubtotal` torna-se um método do `Pedido` (`pedido.Subtotal()`), aumentando a coesão.
  - As funções de cálculo são agrupadas em um `ServicoDeCalculoDePreco`, introduzindo o conceito de um serviço de aplicação.
- **Princípio Aplicado:** **Alta Coesão** e **Clean Architecture (Camada de Casos de Uso)**. O design é refinado, separando a lógica da entidade da lógica do caso de uso.

### Passo 7: Princípio Aberto/Fechado (OCP)
- **Arquivo:** [`steps/passo7/passo7.go`](./steps/passo7/passo7.go)
- **Mudanças:**
  - É criada a interface `RegraDeCalculo`.
  - As lógicas de desconto e taxa são extraídas para `structs` separadas que implementam essa interface.
  - O serviço de cálculo passa a receber uma lista de regras, iterando sobre elas.
- **Princípio Aplicado:** **SOLID (Open-Closed Principle - OCP)**. O sistema torna-se extensível. Novas regras de cálculo podem ser adicionadas sem modificar o código do serviço, apenas injetando uma nova implementação da regra.

### Passo 8: Estrutura em Camadas (Clean Architecture)
- **Arquivos:**
  - [`steps/passo8/dominio/dominio.go`](./steps/passo8/dominio/dominio.go)
  - [`steps/passo8/aplicacao/servico.go`](./steps/passo8/aplicacao/servico.go)
  - [`steps/passo8/apresentacao/apresentacao.go`](./steps/passo8/apresentacao/apresentacao.go)
- **Mudanças:** O código é fisicamente separado em três pacotes (`dominio`, `aplicacao`, `apresentacao`), implementando a **Arquitetura em Camadas**.
- **Princípio Aplicado:** **Clean Architecture (Separação de Responsabilidades e Regra da Dependência)**. As dependências apontam para dentro (`apresentacao` -> `aplicacao` -> `dominio`), isolando o núcleo de negócio de detalhes externos.

### Passo 9: Tratamento Robusto de Erros
- **Arquivos:**
  - [`steps/passo9/dominio/dominio.go`](./steps/passo9/dominio/dominio.go)
  - [`steps/passo9/aplicacao/servico.go`](./steps/passo9/aplicacao/servico.go)
  - [`steps/passo9/apresentacao/apresentacao.go`](./steps/passo9/apresentacao/apresentacao.go)
- **Mudanças:** As funções passam a retornar erros. A camada de domínio valida suas invariantes (ex: preço não pode ser negativo), e as camadas superiores propagam e tratam esses erros.
- **Princípio Aplicado:** **Clean Code (Tratamento Robusto de Erros)**. O sistema torna-se mais resiliente e previsível, lidando de forma explícita com estados de falha.

### Passo 10: Testes Automatizados
- **Arquivos:**
  - [`steps/passo10/dominio/dominio_test.go`](./steps/passo10/dominio/dominio_test.go)
  - [`steps/passo10/aplicacao/servico_test.go`](./steps/passo10/aplicacao/servico_test.go)
  - [`steps/passo10/apresentacao/apresentacao_test.go`](./steps/passo10/apresentacao/apresentacao_test.go)
- **Mudanças:** São adicionados testes unitários para cada camada, validando a funcionalidade e a arquitetura.
- **Princípio Aplicado:** **Clean Architecture (Testabilidade)** e **Clean Code (Testes Limpos - FIRST)**.
  - **Testes de Domínio:** Validam a lógica de negócio pura.
  - **Testes de Aplicação:** Usam **mocks/stubs** para testar a lógica do caso de uso em isolamento.
  - **Testes de Apresentação:** Validam as implementações concretas das regras.

### Passo 11: Refatoração Semântica e Construtores
- **Arquivos:**
  - [`steps/passo11/dominio/dominio.go`](./steps/passo11/dominio/dominio.go)
  - [`steps/passo11/aplicacao/servico.go`](./steps/passo11/aplicacao/servico.go)
  - [`steps/passo11/apresentacao/apresentacao.go`](./steps/passo11/apresentacao/apresentacao.go)
- **Mudanças:**
  - **Domínio:** Introduz construtores (ex: `NewProduto`) que garantem que os objetos só podem ser criados em estado válido, movendo a validação para o momento da criação.
  - **Aplicação:** As constantes de valores de desconto são removidas para purificar a camada.
  - **Apresentação:** As constantes são movidas para junto das regras que as usam. As regras são refatoradas para usar mapas, eliminando as cadeias de `if/else`.
- **Princípio Aplicado:** **Domain-Driven Design (Factories e Agregados)** e **Alta Coesão**. O código torna-se ainda mais seguro e semanticamente expressivo.

---
## Licença

Distribuído sob a licença MIT. Veja a [LICENSE](LICENSE) para mais informações.