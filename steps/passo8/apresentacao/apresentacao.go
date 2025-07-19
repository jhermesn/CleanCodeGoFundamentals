package apresentacao

import (
	"fmt"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo8/aplicacao"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo8/dominio"
)

/*
Passo 8: Estrutura em Camadas (Clean Architecture).
A camada de apresentação é a mais externa. Ela é responsável por
conectar o sistema ao mundo exterior (neste caso, o terminal) e por
compor a aplicação, injetando as dependências concretas.
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

func CalcularPrecoTotal() {
	pedido := &dominio.Pedido{
		Produtos: []*dominio.Produto{
			{Nome: "Produto A", Preco: 50.0},
			{Nome: "Produto B", Preco: 80.0},
		},
		Usuario: &dominio.Usuario{Nome: "Cliente Premium", Tipo: dominio.TipoUsuarioPremium},
		Cupom:   dominio.Cupom10OFF,
	}

	servico := &aplicacao.ServicoDeCalculoDePreco{
		Regras: []aplicacao.RegraDeCalculo{
			&DescontoPorUsuario{},
			&DescontoPorCupom{},
			&RegraDeTaxaDeEnvio{},
		},
	}

	precoFinal := servico.Calcular(pedido)
	fmt.Println("Preço total:", precoFinal)
}