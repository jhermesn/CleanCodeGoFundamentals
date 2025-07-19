package passo2

import "fmt"

/*
Passo 2: Nomes Significativos.
Neste passo, a função e as variáveis foram renomeadas para refletir suas
responsabilidades e tornar o código mais legível e autoexplicativo,
sem alterar a lógica original.
*/

func CalcularPrecoTotal(precos []float64, usuario map[string]string, cupom string) {
	var total float64
	for _, preco := range precos {
		total += preco
	}

	if usuario["tipo"] == "premium" {
		total *= 0.9
	} else if usuario["tipo"] == "plus" {
		total *= 0.8
	}

	if cupom == "10OFF" {
		total -= 10
	} else if cupom == "5OFF" {
		total -= 5
	}

	if total < 100 {
		total *= 1.05
	}

	fmt.Println("Preço total:", total)
}