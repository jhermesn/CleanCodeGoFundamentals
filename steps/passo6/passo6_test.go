package passo6

import "testing"

// package passo6 contém os testes de unidade para a lógica de cálculo de preço.
// Os testes validam as funções auxiliares, as estratégias de desconto e a orquestração final,
// garantindo que a refatoração preservou e isolou corretamente as regras de negócio.

// --- Testes para Funções Auxiliares ---

func TestCalcularTotalDosProdutos(t *testing.T) {
	produtos := []Produto{{Preco: 50.5}, {Preco: 49.5}}
	totalEsperado := 100.0

	totalCalculado := CalcularTotalDosProdutos(produtos)

	if totalCalculado != totalEsperado {
		t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", totalEsperado, totalCalculado)
	}
}

func TestAdicionarTaxaDeEnvio(t *testing.T) {
	t.Run("Deve adicionar taxa para total abaixo de 100", func(t *testing.T) {
		totalInicial := 99.99
		totalEsperado := 104.99 // 99.99 + 5.00
		totalCalculado := AdicionarTaxaDeEnvio(totalInicial)
		if totalCalculado != totalEsperado {
			t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", totalEsperado, totalCalculado)
		}
	})

	t.Run("NÃO deve adicionar taxa para total igual ou acima de 100", func(t *testing.T) {
		totalInicial := 100.0
		totalEsperado := 100.0
		totalCalculado := AdicionarTaxaDeEnvio(totalInicial)
		if totalCalculado != totalEsperado {
			t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", totalEsperado, totalCalculado)
		}
	})
}

// --- Testes para as Estratégias de Desconto ---

func TestDescontoUsuario(t *testing.T) {
	desconto := DescontoUsuario{Fator: 0.8} // 20% de desconto
	totalInicial := 200.0
	totalEsperado := 160.0

	totalCalculado := desconto.Aplicar(totalInicial)

	if totalCalculado != totalEsperado {
		t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", totalEsperado, totalCalculado)
	}
}

func TestDescontoCupom(t *testing.T) {
	desconto := DescontoCupom{Valor: 15.0}
	totalInicial := 100.0
	totalEsperado := 85.0

	totalCalculado := desconto.Aplicar(totalInicial)

	if totalCalculado != totalEsperado {
		t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", totalEsperado, totalCalculado)
	}
}

func TestDescontoCupomNaoDeveNegativar(t *testing.T) {
	desconto := DescontoCupom{Valor: 20.0}
	totalInicial := 10.0
	totalEsperado := 0.0 // O total não pode ser negativo.

	totalCalculado := desconto.Aplicar(totalInicial)

	if totalCalculado != totalEsperado {
		t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", totalEsperado, totalCalculado)
	}
}

// --- Teste para a Função de Orquestração ---

func TestCalcularPrecoFinal(t *testing.T) {
	subtotal := 120.0
	// Aplica desconto de usuário (10%) e depois cupom (R$10).
	descontos := []ICalculadorDeDesconto{
		DescontoUsuario{Fator: FatorDescontoPremium}, // 120 * 0.9 = 108
		DescontoCupom{Valor: ValorCupom10},           // 108 - 10 = 98
	}

	// O total (98.0) é menor que 100, então a taxa de envio é adicionada.
	precoFinalEsperado := 98.0 + TaxaDeEnvio // 103.0

	precoFinalCalculado := CalcularPrecoFinal(subtotal, descontos)

	if precoFinalCalculado != precoFinalEsperado {
		t.Errorf("Esperado R$%.2f, mas foi calculado R$%.2f", precoFinalEsperado, precoFinalCalculado)
	}
}

// --- Teste de Integração ---

func TestCenarioCompleto(t *testing.T) {
	// Valida o cenário de ponta a ponta, simulando como a aplicação montaria os componentes.

	// 1. Dados de entrada do pedido.
	usuario := Usuario{Tipo: TipoUsuarioPremium}
	pedido := Pedido{
		Produtos: []Produto{{Preco: 80.0}, {Preco: 40.0}},
		Cupom:    CupomDesconto10,
	}

	// 2. Calcula o subtotal.
	subtotal := CalcularTotalDosProdutos(pedido.Produtos)

	// 3. Constrói a lista de descontos aplicáveis.
	var descontosAplicaveis []ICalculadorDeDesconto
	if usuario.Tipo == TipoUsuarioPremium {
		descontosAplicaveis = append(descontosAplicaveis, DescontoUsuario{Fator: FatorDescontoPremium})
	}
	if pedido.Cupom == CupomDesconto10 {
		descontosAplicaveis = append(descontosAplicaveis, DescontoCupom{Valor: ValorCupom10})
	}

	// 4. Executa o cálculo final.
	precoFinalCalculado := CalcularPrecoFinal(subtotal, descontosAplicaveis)

	// 5. Valida o resultado esperado.
	// Subtotal (120.0) -> Desconto Premium (108.0) -> Cupom (98.0) -> Taxa de Envio (103.0)
	precoFinalEsperado := 103.0

	if precoFinalCalculado != precoFinalEsperado {
		t.Errorf("Cenário completo falhou. Esperado R$%.2f, mas foi calculado R$%.2f", precoFinalEsperado, precoFinalCalculado)
	}
}
