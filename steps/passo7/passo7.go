package passo7

import "fmt"

/*
Passo 7: Princípio Aberto/Fechado (OCP).
Neste passo, aplicamos o Princípio Aberto/Fechado. Criamos a interface
`RegraDeCalculo` para representar qualquer regra que possa modificar o preço.
Agora, para adicionar um novo desconto ou taxa, basta criar uma nova struct
que implemente essa interface, sem precisar alterar o `ServicoDeCalculoDePreco`.
O sistema está aberto para extensão (novas regras), mas fechado para modificação.
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

func (p *Pedido) Subtotal() float64 {
	var subtotal float64
	for _, produto := range p.Produtos {
		subtotal += produto.Preco
	}
	return subtotal
}

type RegraDeCalculo interface {
	Aplicar(pedido *Pedido, total float64) float64
}

type DescontoPorUsuario struct{}

func (d *DescontoPorUsuario) Aplicar(pedido *Pedido, total float64) float64 {
	if pedido.Usuario.Tipo == TipoUsuarioPremium {
		return total - (total * DescontoPremium)
	}
	if pedido.Usuario.Tipo == TipoUsuarioPlus {
		return total - (total * DescontoPlus)
	}
	return total
}

type DescontoPorCupom struct{}

func (d *DescontoPorCupom) Aplicar(pedido *Pedido, total float64) float64 {
	if pedido.Cupom == Cupom10OFF {
		return total - ValorDesconto10
	}
	if pedido.Cupom == Cupom5OFF {
		return total - ValorDesconto5
	}
	return total
}

type RegraDeTaxaDeEnvio struct{}

func (t *RegraDeTaxaDeEnvio) Aplicar(pedido *Pedido, total float64) float64 {
	if total < ValorMinimoParaTaxa {
		return total + (total * TaxaDeEnvio)
	}
	return total
}

type ServicoDeCalculoDePreco struct {
	Regras []RegraDeCalculo
}

func (s *ServicoDeCalculoDePreco) Calcular(pedido *Pedido) float64 {
	total := pedido.Subtotal()
	for _, regra := range s.Regras {
		total = regra.Aplicar(pedido, total)
	}
	return total
}

func CalcularPrecoTotal(pedido *Pedido) {
	servico := &ServicoDeCalculoDePreco{
		Regras: []RegraDeCalculo{
			&DescontoPorUsuario{},
			&DescontoPorCupom{},
			&RegraDeTaxaDeEnvio{},
		},
	}
	precoFinal := servico.Calcular(pedido)
	fmt.Println("Preço total:", precoFinal)
}