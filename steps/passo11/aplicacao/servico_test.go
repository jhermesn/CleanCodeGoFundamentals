package aplicacao

import (
	"errors"
	"testing"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo11/dominio"
)

type mockRegraCalculo struct {
	valorAplicado float64
}

func (m *mockRegraCalculo) Aplicar(pedido *dominio.Pedido, total float64) float64 {
	return total - m.valorAplicado
}

func TestServicoDeCalculoDePreco_Calcular_Sucesso(t *testing.T) {
	produto, _ := dominio.NewProduto("Produto Teste", 100)
	usuario, _ := dominio.NewUsuario("Usuário Teste", dominio.TipoUsuarioPadrao)
	pedido := dominio.NewPedido([]*dominio.Produto{produto}, usuario, "")

	regrasMock := []RegraDeCalculo{
		&mockRegraCalculo{valorAplicado: 20}, // Simula um desconto de 20
		&mockRegraCalculo{valorAplicado: 5},  // Simula um desconto de 5
	}
	servico := NovoServicoDeCalculoDePreco(regrasMock)

	esperado := 75.0 // 100 - 20 - 5

	resultado, err := servico.Calcular(pedido)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado != esperado {
		t.Errorf("Cálculo incorreto. Esperado: %f, Obtido: %f", esperado, resultado)
	}
}

func TestServicoDeCalculoDePreco_Calcular_ErroPedidoNulo(t *testing.T) {
	servico := NovoServicoDeCalculoDePreco(nil)

	_, err := servico.Calcular(nil)

	if err == nil {
		t.Error("Esperado um erro de pedido inválido, mas nenhum foi retornado")
	}
	if !errors.Is(err, ErrPedidoInvalido) {
		t.Errorf("Tipo de erro incorreto. Esperado: %v, Obtido: %v", ErrPedidoInvalido, err)
	}
}
