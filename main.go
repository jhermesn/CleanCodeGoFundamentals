package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo1"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo2"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo3"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo4"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo5"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo6"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo7"
	passo8apresentacao "github.com/jhermesn/CleanCodeGoFundamentals/steps/passo8/apresentacao"
	passo9apresentacao "github.com/jhermesn/CleanCodeGoFundamentals/steps/passo9/apresentacao"
	passo10apresentacao "github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/apresentacao"
	passo11apresentacao "github.com/jhermesn/CleanCodeGoFundamentals/steps/passo11/apresentacao"
)

func main() {
	if len(os.Args) > 1 {
		passo, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argumento inválido. Forneça um número de passo válido.")
			os.Exit(1)
		}
		fmt.Printf("\n--- Executando Passo %d (modo não interativo) ---\n", passo)
		executarPasso(passo)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Demonstração de Evolução de Código com Clean Code e Clean Architecture ---")
		fmt.Println("Escolha o passo que deseja executar (1-11) ou digite 'sair':")
		fmt.Println("1.  Código Inicial (Monolítico e com 'Code Smells')")
		fmt.Println("2.  Nomes Significativos")
		fmt.Println("3.  Remoção de Números Mágicos")
		fmt.Println("4.  Introdução de Tipos de Domínio (Structs)")
		fmt.Println("5.  Funções com Responsabilidade Única (SRP)")
		fmt.Println("6.  Coesão com Métodos e Serviços")
		fmt.Println("7.  Princípio Aberto/Fechado (OCP)")
		fmt.Println("8.  Estrutura em Camadas (Clean Architecture)")
		fmt.Println("9.  Tratamento Robusto de Erros")
		fmt.Println("10. Injeção de Dependência e Testes")
		fmt.Println("11. Refatoração de Regras Semânticas")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "sair" {
			fmt.Println("Saindo...")
			break
		}

		passo, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número de 1 a 11.")
			continue
		}

		fmt.Printf("\n--- Executando Passo %d ---\n", passo)
		executarPasso(passo)
	}
}

func executarPasso(passo int) {
	switch passo {
	case 1:
		precos := []float64{50.0, 80.0}
		usuarioMap := map[string]string{"t": "pl"} // "pl" para 20% de desconto
		cupom := "10OFF"
		passo1.Cp(precos, usuarioMap, cupom)
	case 2:
		precos := []float64{50.0, 80.0}
		usuarioMap := map[string]string{"tipo": "plus"} // "plus" para 20% de desconto
		cupom := "10OFF"
		passo2.CalcularPrecoTotal(precos, usuarioMap, cupom)
	case 3:
		precos := []float64{50.0, 80.0}
		usuarioMap := map[string]string{"tipo": "plus"}
		cupom := "10OFF"
		passo3.CalcularPrecoTotal(precos, usuarioMap, cupom)
	case 4:
		pedido := &passo4.Pedido{
			Produtos: []*passo4.Produto{{Nome: "Produto A", Preco: 50.0}, {Nome: "Produto B", Preco: 80.0}},
			Usuario:  &passo4.Usuario{Nome: "Cliente Plus", Tipo: "plus"},
			Cupom:    "10OFF",
		}
		passo4.CalcularPrecoTotal(pedido)
	case 5:
		pedido := &passo5.Pedido{
			Produtos: []*passo5.Produto{{Nome: "Produto A", Preco: 50.0}, {Nome: "Produto B", Preco: 80.0}},
			Usuario:  &passo5.Usuario{Nome: "Cliente Plus", Tipo: "plus"},
			Cupom:    "10OFF",
		}
		passo5.CalcularPrecoTotal(pedido)
	case 6:
		pedido := &passo6.Pedido{
			Produtos: []*passo6.Produto{{Nome: "Produto A", Preco: 50.0}, {Nome: "Produto B", Preco: 80.0}},
			Usuario:  &passo6.Usuario{Nome: "Cliente Plus", Tipo: "plus"},
			Cupom:    "10OFF",
		}
		passo6.CalcularPrecoTotal(pedido)
	case 7:
		pedido := &passo7.Pedido{
			Produtos: []*passo7.Produto{{Nome: "Produto A", Preco: 50.0}, {Nome: "Produto B", Preco: 80.0}},
			Usuario:  &passo7.Usuario{Nome: "Cliente Plus", Tipo: "plus"},
			Cupom:    "10OFF",
		}
		passo7.CalcularPrecoTotal(pedido)
	case 8:
		passo8apresentacao.CalcularPrecoTotal()
	case 9:
		passo9apresentacao.CalcularPrecoTotal()
	case 10:
		passo10apresentacao.CalcularPrecoTotal()
	case 11:
		passo11apresentacao.CalcularPrecoTotal()
	default:
		fmt.Println("Passo inválido. Por favor, escolha um número de 1 a 11.")
	}
}