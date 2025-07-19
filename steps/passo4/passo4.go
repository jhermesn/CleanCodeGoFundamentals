package passo4

import "fmt"

/*
Passo 4: Introduzir Tipos de Domínio.
Neste passo, substituímos os tipos de dados genéricos (map, slices de float)
por estruturas de domínio (structs) como Pedido, Usuario e Produto.
Isso torna o código mais expressivo, seguro e alinhado aos conceitos
de negócio, evitando a manipulação de dados primitivos e sem semântica.
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

func CalcularPrecoTotal(pedido *Pedido) {
	var total float64
	for _, produto := range pedido.Produtos {
		total += produto.Preco
	}

	if pedido.Usuario.Tipo == TipoUsuarioPremium {
		total -= total * DescontoPremium
	} else if pedido.Usuario.Tipo == TipoUsuarioPlus {
		total -= total * DescontoPlus
	}

	if pedido.Cupom == Cupom10OFF {
		total -= ValorDesconto10
	} else if pedido.Cupom == Cupom5OFF {
		total -= ValorDesconto5
	}

	if total < ValorMinimoParaTaxa {
		total += total * TaxaDeEnvio
	}

	fmt.Println("Preço total:", total)
}