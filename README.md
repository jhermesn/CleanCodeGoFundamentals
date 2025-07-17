# Estudo de Caso: Refatoração e Clean Code em Go

Este repositório é um estudo de caso prático que demonstra a aplicação de princípios de **Clean Code** e **SOLID** para refatorar um código em Go. O objetivo é mostrar, de forma progressiva, como um código inicialmente simples, mas problemático, pode ser transformado em uma solução robusta, testável e de fácil manutenção. Projeto apresentado como avaliação final da disciplina de Fundamentos de Engenharia de Software em Julho de 2025, ministrada pelo professor Anderson Costa na UEPA.

## O Cenário

O ponto de partida é uma única função responsável por calcular o preço final de um pedido. Essa função precisa considerar:
- A soma dos preços dos produtos.
- Descontos baseados no tipo de cliente (Premium, Plus).
- Descontos de cupons.
- A aplicação de uma taxa de envio para pedidos abaixo de um certo valor.

## A Evolução do Código

A refatoração foi dividida em etapas, e cada etapa está em seu próprio pacote para facilitar a execução e comparação.

### Passo 1: `steps/passo1/passo1.go`
O código original. É uma função monolítica com vários problemas clássicos:
- **Nomes ruins**: Variáveis como `p`, `u`, `c` e `t` não são descritivas.
- **Números Mágicos**: Valores como `0.9`, `100.0` e `5.0` estão espalhados pelo código sem explicação.
- **Múltiplas Responsabilidades**: A mesma função faz tudo, tornando-a difícil de ler, testar e modificar.

### Passo 2: `steps/passo2/passo2.go`
A primeira e mais simples melhoria: renomear a função e suas variáveis para nomes claros e significativos. O código começa a se autodocumentar.

### Passo 3: `steps/passo3/passo3.go`
Aqui, atacamos dois problemas:
1.  Os "números mágicos" são substituídos por **constantes nomeadas**, tornando o código mais legível e fácil de manter.
2.  Tipos genéricos como `map[string]string` são substituídos por **`structs`** (`Usuario`, `Pedido`), modelando melhor o domínio do problema.

### Passo 4: `steps/passo4/passo4.go`
Aplicação do **Princípio da Responsabilidade Única (SRP)**. A função original é quebrada em funções menores, cada uma com um único propósito. A função principal passa a orquestrar as chamadas, e seu fluxo se torna muito mais claro.

### Passo 5: `steps/passo5/passo5.go`
A evolução final do design, aplicando os princípios **Aberto/Fechado (OCP)** e de **Inversão de Dependência (DIP)**.
- Uma interface `ICalculadorDeDesconto` é criada para abstrair as regras de desconto.
- A lógica de cálculo passa a depender dessa abstração, e não de implementações concretas.
- Isso permite que **novas regras de desconto sejam adicionadas sem alterar o código existente**, tornando o sistema extensível.

### Passo 6: `steps/passo6/passo6.go` e `steps/passo6/passo6_test.go`
O resultado final da refatoração: um código altamente testável. Este pacote contém o código final e seus testes unitários, demonstrando como a nova arquitetura facilita a validação e garante a qualidade do software.

## Como Executar

### Executando cada passo
Você pode executar a versão de cada passo da refatoração usando o `main.go`. Passe o nome do passo como um argumento de linha de comando.

```bash
# Exemplo para executar a lógica do Passo 1
go run main.go passo1

# Exemplo para executar a lógica do Passo 5
go run main.go passo5
```

### Executando os Testes
Para validar a solução final, o `passo6` foi configurado para rodar os testes de unidade diretamente.

```bash
# Executa os testes unitários do Passo 6
go run main.go passo6
```
Isso demonstrará como a nova arquitetura facilita a validação e garante a qualidade do software.

## Como Executar com Docker

Com o Docker e o Docker Compose instalados, você pode construir e executar a aplicação em um container.

### Construindo a Imagem
Execute o comando abaixo na raiz do projeto para construir a imagem Docker:
```bash
docker-compose build
```

### Executando uma Etapa Específica
Você pode executar qualquer passo da refatoração usando o `docker-compose run`. O serviço removerá o container após a execução (`--rm`).

```bash
# Exemplo para executar a lógica do Passo 1
docker-compose run --rm app passo1

# Exemplo para executar a lógica do Passo 6 (versão final)
docker-compose run --rm app passo6
```

### Executando os Testes no Container
Para rodar os testes dentro do ambiente Docker, você pode sobrescrever o comando de entrada do container:
```bash
docker-compose run --rm --entrypoint "go test -v ./..." app
```
[LICENSE](LICENSE).