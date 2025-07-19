package aplicacao

import (
	"errors"
	"testing"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/dominio"
)

type mockRegraCalculo struct {
	valorAplicado float64
}

func (m *mockRegraCalculo) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	return total + m.valorAplicado
}

func TestServicoDeCalculoDePreco_Calcular_Sucesso(t *testing.T) {
	pedido := &dominio.Pedido{
		Produtos: []*dominio.Produto{{Preco: 100}},
	}
	regrasMock := []RegraDeCalculo{
		&mockRegraCalculo{valorAplicado: -10},
		&mockRegraCalculo{valorAplicado: 5},
	}
	servico := NovoServicoDeCalculoDePreco(regrasMock)

	esperado := 95.0

	resultado, err := servico.Calcular(pedido)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado != esperado {
		t.Errorf("Cálculo incorreto. Esperado: %f, Obtido: %f", esperado, resultado)
	}
}

func TestServicoDeCalculoDePreco_Calcular_ErroNoSubtotal(t *testing.T) {
	pedidoComErro := &dominio.Pedido{
		Produtos: []*dominio.Produto{{Preco: -10}},
	}
	servico := NovoServicoDeCalculoDePreco(nil)

	_, err := servico.Calcular(pedidoComErro)

	if err == nil {
		t.Errorf("Esperado um erro, mas nenhum foi retornado")
	}
	if !errors.Is(err, dominio.ErrPrecoInvalido) {
		t.Errorf("Tipo de erro incorreto. Esperado: %v, Obtido: %v", dominio.ErrPrecoInvalido, err)
	}
}