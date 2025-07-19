package apresentacao

import (
	"fmt"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/aplicacao"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/dominio"
)

/*
Passo 10: Injeção de Dependência e Testes Automatizados.
A camada de apresentação utiliza um "Contêiner de DI" para compor a
aplicação. O resultado é um sistema totalmente desacoplado e agora
comprovadamente testável, com as camadas de domínio e aplicação
cobertas por testes unitários, representando o estado da arte das
boas práticas de engenharia de software.
*/

type DescontoPorUsuario struct{}

func (d *DescontoPorUsuario) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	if pedido.Usuario.Tipo == dominio.TipoUsuarioPremium {
		return total - (total * aplicacao.DescontoPremium)
	}
	if pedido.Usuario.Tipo == dominio.TipoUsuarioPlus {
		return total - (total * aplicacao.DescontoPlus)
	}
	return total
}

type DescontoPorCupom struct{}

func (d *DescontoPorCupom) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	if pedido.Cupom == dominio.Cupom10OFF {
		return total - aplicacao.ValorDesconto10
	}
	if pedido.Cupom == dominio.Cupom5OFF {
		return total - aplicacao.ValorDesconto5
	}
	return total
}

type RegraDeTaxaDeEnvio struct{}

func (t *RegraDeTaxaDeEnvio) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	if total < aplicacao.ValorMinimoParaTaxa {
		return total + (total * aplicacao.TaxaDeEnvio)
	}
	return total
}

type container struct {
	servicoDeCalculo *aplicacao.ServicoDeCalculoDePreco
}

func newContainer() *container {
	regras := []aplicacao.RegraDeCalculo{
		&DescontoPorUsuario{},
		&DescontoPorCupom{},
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
	pedidoSucesso := &dominio.Pedido{
		Produtos: []*dominio.Produto{
			{Nome: "Produto A", Preco: 50.0},
			{Nome: "Produto B", Preco: 80.0},
		},
		Usuario: &dominio.Usuario{Nome: "Cliente Premium", Tipo: dominio.TipoUsuarioPremium},
		Cupom:   dominio.Cupom10OFF,
	}
	executarCalculo(di.servicoDeCalculo, pedidoSucesso)

	fmt.Println("\n--- Cenário de Erro (Produto com Preço Negativo) ---")
	pedidoErro := &dominio.Pedido{
		Produtos: []*dominio.Produto{
			{Nome: "Produto C", Preco: -10.0},
		},
		Usuario: &dominio.Usuario{Nome: "Cliente Plus", Tipo: dominio.TipoUsuarioPlus},
	}
	executarCalculo(di.servicoDeCalculo, pedidoErro)
}

func executarCalculo(servico *aplicacao.ServicoDeCalculoDePreco, pedido *dominio.Pedido) {
	precoFinal, err := servico.Calcular(pedido)
	if err != nil {
		fmt.Println("Ocorreu um erro ao calcular o preço:", err)
		return
	}

	fmt.Println("Preço total:", precoFinal)
}
