package passo6

import "fmt"

/*
Passo 6: Melhorar a Coesão com Métodos e Serviços.
Neste passo, aumentamos a coesão do código. A função para calcular o subtotal
foi transformada em um método do `Pedido`, pois está diretamente ligada aos seus
dados. As lógicas de cálculo de descontos e taxas foram agrupadas em um
`ServicoDeCalculoDePreco`, centralizando as regras de negócio relacionadas
ao cálculo do preço final em uma única estrutura.
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

type ServicoDeCalculoDePreco struct{}

func (s *ServicoDeCalculoDePreco) Calcular(pedido *Pedido) float64 {
	total := pedido.Subtotal()
	total = s.aplicarDescontoPorUsuario(total, pedido.Usuario)
	total = s.aplicarDescontoPorCupom(total, pedido.Cupom)
	total = s.aplicarTaxaDeEnvio(total)
	return total
}

func (s *ServicoDeCalculoDePreco) aplicarDescontoPorUsuario(total float64, usuario *Usuario) float64 {
	if usuario.Tipo == TipoUsuarioPremium {
		return total - (total * DescontoPremium)
	}
	if usuario.Tipo == TipoUsuarioPlus {
		return total - (total * DescontoPlus)
	}
	return total
}

func (s *ServicoDeCalculoDePreco) aplicarDescontoPorCupom(total float64, cupom string) float64 {
	if cupom == Cupom10OFF {
		return total - ValorDesconto10
	}
	if cupom == Cupom5OFF {
		return total - ValorDesconto5
	}
	return total
}

func (s *ServicoDeCalculoDePreco) aplicarTaxaDeEnvio(total float64) float64 {
	if total < ValorMinimoParaTaxa {
		return total + (total * TaxaDeEnvio)
	}
	return total
}

func CalcularPrecoTotal(pedido *Pedido) {
	servico := &ServicoDeCalculoDePreco{}
	precoFinal := servico.Calcular(pedido)
	fmt.Println("Preço total:", precoFinal)
}