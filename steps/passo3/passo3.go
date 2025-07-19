package passo3

import "fmt"

/*
Passo 3: Eliminar Números Mágicos.
Neste passo, os "números mágicos" e "strings mágicas" foram substituídos
por constantes nomeadas. Isso torna o código mais legível e facilita a
manutenção, pois os valores ficam centralizados e com um significado claro.
*/

const (
	TipoUsuarioPremium = "premium"
	TipoUsuarioPlus    = "plus"

	DescontoPremium = 0.10
	DescontoPlus    = 0.20

	Cupom10OFF = "10OFF"
	Cupom5OFF  = "5OFF"

	ValorDesconto10 = 10.0
	ValorDesconto5  = 5.0

	ValorMinimoParaTaxa = 100.0
	TaxaDeEnvio         = 0.05
)

func CalcularPrecoTotal(precos []float64, usuario map[string]string, cupom string) {
	var total float64
	for _, preco := range precos {
		total += preco
	}

	if usuario["tipo"] == TipoUsuarioPremium {
		total -= total * DescontoPremium
	} else if usuario["tipo"] == TipoUsuarioPlus {
		total -= total * DescontoPlus
	}

	if cupom == Cupom10OFF {
		total -= ValorDesconto10
	} else if cupom == Cupom5OFF {
		total -= ValorDesconto5
	}

	if total < ValorMinimoParaTaxa {
		total += total * TaxaDeEnvio
	}

	fmt.Println("Preço total:", total)
}