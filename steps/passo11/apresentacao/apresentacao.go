package apresentacao

import (
	"fmt"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo11/aplicacao"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo11/dominio"
)

/*
Passo 11: Refatoração de Regras Semânticas.
Neste passo, os valores de desconto foram movidos para esta camada, junto
às regras que os utilizam, aumentando a coesão. As regras agora são
totalmente autocontidas, e a camada de aplicação não tem mais conhecimento
sobre detalhes de implementação.
*/

const (
	descontoPremium     = 0.10
	descontoPlus        = 0.20
	valorDesconto10     = 10.0
	valorDesconto5      = 5.0
	valorMinimoParaTaxa = 100.0
	taxaDeEnvio         = 0.05
)

type DescontoPorUsuario struct {
	descontos map[dominio.TipoUsuario]float64
}

func NewDescontoPorUsuario() *DescontoPorUsuario {
	return &DescontoPorUsuario{
		descontos: map[dominio.TipoUsuario]float64{
			dominio.TipoUsuarioPremium: descontoPremium,
			dominio.TipoUsuarioPlus:    descontoPlus,
		},
	}
}

func (d *DescontoPorUsuario) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	if desconto, encontrado := d.descontos[pedido.Usuario.Tipo]; encontrado {
		return total - (total * desconto)
	}
	return total
}

type DescontoPorCupom struct {
	descontos map[string]float64
}

func NewDescontoPorCupom() *DescontoPorCupom {
	return &DescontoPorCupom{
		descontos: map[string]float64{
			dominio.Cupom10OFF: valorDesconto10,
			dominio.Cupom5OFF:  valorDesconto5,
		},
	}
}

func (d *DescontoPorCupom) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	if desconto, encontrado := d.descontos[pedido.Cupom]; encontrado {
		return total - desconto
	}
	return total
}

type RegraDeTaxaDeEnvio struct{}

func (t *RegraDeTaxaDeEnvio) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	if total < valorMinimoParaTaxa {
		return total + (total * taxaDeEnvio)
	}
	return total
}

type container struct {
	servicoDeCalculo *aplicacao.ServicoDeCalculoDePreco
}

func newContainer() *container {
	regras := []aplicacao.RegraDeCalculo{
		NewDescontoPorUsuario(),
		NewDescontoPorCupom(),
		&RegraDeTaxaDeEnvio{},
	}
	servico := aplicacao.NovoServicoDeCalculoDePreco(regras)

	return &container{
		servicoDeCalculo: servico,
	}
}

func CalcularPrecoTotal() {
	di := newContainer()

	fmt.Println("--- Cenário de Sucesso ---")
	p1, err := dominio.NewProduto("Produto A", 50.0)
	if err != nil {
		fmt.Println("Erro ao criar produto:", err)
		return
	}
	p2, err := dominio.NewProduto("Produto B", 80.0)
	if err != nil {
		fmt.Println("Erro ao criar produto:", err)
		return
	}
	usuario, err := dominio.NewUsuario("Cliente Premium", dominio.TipoUsuarioPremium)
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return
	}
	pedidoSucesso := dominio.NewPedido(
		[]*dominio.Produto{p1, p2},
		usuario,
		dominio.Cupom5OFF,
	)
	executarCalculo(di.servicoDeCalculo, pedidoSucesso)

	fmt.Println("\n--- Cenário de Erro (Produto com Preço Negativo) ---")
	produtoInvalido, err := dominio.NewProduto("Produto Inválido", -10.0)
	if err != nil {
		fmt.Println("Ocorreu um erro ao criar o produto:", err)
	} else {
		usuarioErro, _ := dominio.NewUsuario("Cliente Plus", dominio.TipoUsuarioPlus)
		pedidoErro := dominio.NewPedido([]*dominio.Produto{produtoInvalido}, usuarioErro, "")
		executarCalculo(di.servicoDeCalculo, pedidoErro)
	}
}

func executarCalculo(servico *aplicacao.ServicoDeCalculoDePreco, pedido *dominio.Pedido) {
	precoFinal, err := servico.Calcular(pedido)
	if err != nil {
		fmt.Println("Ocorreu um erro ao calcular o preço:", err)
		return
	}

	fmt.Println("Preço total:", precoFinal)
}