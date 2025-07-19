package passo5

import "fmt"

/*
Passo 5: Funções com Responsabilidade Única (SRP).
Neste passo, a função monolítica `CalcularPrecoTotal` foi dividida em
funções menores e com responsabilidades únicas, seguindo o SRP.
Temos funções para calcular o subtotal, aplicar descontos e taxas.
Isso torna o código mais modular, legível e fácil de testar.
*/

const (
	TipoUsuarioPremium = "premium"
	TipoUsuarioPlus    = "plus"
	Cupom10OFF         = "10OFF"
	Cupom5OFF          = "5OFF"
)

const (
	DescontoPremium     = 0.10
	DescontoPlus        = 0.20
	ValorDesconto10     = 10.0
	ValorDesconto5      = 5.0
	ValorMinimoParaTaxa = 100.0
	TaxaDeEnvio         = 0.05
)

type Produto struct {
	Nome  string
	Preco float64
}

type Usuario struct {
	Nome string
	Tipo string
}

type Pedido struct {
	Produtos []*Produto
	Usuario  *Usuario
	Cupom    string
}

func calcularSubtotal(produtos []*Produto) float64 {
	var subtotal float64
	for _, produto := range produtos {
		subtotal += produto.Preco
	}
	return subtotal
}

func aplicarDescontoPorUsuario(total float64, usuario *Usuario) float64 {
	if usuario.Tipo == TipoUsuarioPremium {
		return total - (total * DescontoPremium)
	}
	if usuario.Tipo == TipoUsuarioPlus {
		return total - (total * DescontoPlus)
	}
	return total
}

func aplicarDescontoPorCupom(total float64, cupom string) float64 {
	if cupom == Cupom10OFF {
		return total - ValorDesconto10
	}
	if cupom == Cupom5OFF {
		return total - ValorDesconto5
	}
	return total
}

func aplicarTaxaDeEnvio(total float64) float64 {
	if total < ValorMinimoParaTaxa {
		return total + (total * TaxaDeEnvio)
	}
	return total
}

func CalcularPrecoTotal(pedido *Pedido) {
	total := calcularSubtotal(pedido.Produtos)
	total = aplicarDescontoPorUsuario(total, pedido.Usuario)
	total = aplicarDescontoPorCupom(total, pedido.Cupom)
	total = aplicarTaxaDeEnvio(total)

	fmt.Println("Preço total:", total)
}