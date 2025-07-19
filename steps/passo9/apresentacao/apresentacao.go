package apresentacao

import (
	"fmt"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo9/aplicacao"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo9/dominio"
)

/*
Passo 9: Tratamento Robusto de Erros.
A camada de apresentação agora verifica os erros retornados pela camada de
aplicação. É responsabilidade desta camada decidir como tratar e apresentar
os erros ao usuário final (neste caso, imprimindo no console).
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
	fmt.Println("--- Cenário de Sucesso ---")
	pedidoSucesso := &dominio.Pedido{
		Produtos: []*dominio.Produto{
			{Nome: "Produto A", Preco: 50.0},
			{Nome: "Produto B", Preco: 80.0},
		},
		Usuario: &dominio.Usuario{Nome: "Cliente Premium", Tipo: dominio.TipoUsuarioPremium},
		Cupom:   dominio.Cupom10OFF,
	}
	executarCalculo(pedidoSucesso)

	fmt.Println("\n--- Cenário de Erro (Produto com Preço Negativo) ---")
	pedidoErro := &dominio.Pedido{
		Produtos: []*dominio.Produto{
			{Nome: "Produto C", Preco: -10.0},
		},
		Usuario: &dominio.Usuario{Nome: "Cliente Plus", Tipo: dominio.TipoUsuarioPlus},
	}
	executarCalculo(pedidoErro)
}

func executarCalculo(pedido *dominio.Pedido) {
	servico := &aplicacao.ServicoDeCalculoDePreco{
		Regras: []aplicacao.RegraDeCalculo{
			&DescontoPorUsuario{},
			&DescontoPorCupom{},
			&RegraDeTaxaDeEnvio{},
		},
	}

	precoFinal, err := servico.Calcular(pedido)
	if err != nil {
		fmt.Println("Ocorreu um erro ao calcular o preço:", err)
		return
	}

	fmt.Println("Preço total:", precoFinal)
}