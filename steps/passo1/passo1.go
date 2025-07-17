package passo1

// package passo1 demonstra o ponto de partida do código.
// A função monolítica abaixo calcula o preço total de um pedido e possui diversos problemas:
// - Nomes de variáveis curtos e não descritivos.
// - "Números Mágicos" sem explicação.
// - Múltiplas responsabilidades (cálculo de subtotal, descontos, taxas).
// - Uso de tipos de dados genéricos em vez de estruturas de domínio.
// - Lógica complexa e aninhada.

// CalcularPrecoPasso1 calcula o preço final com base nos preços (p), usuário (u) e cupom (c).
func CalcularPrecoPasso1(p []float64, u map[string]string, c string) float64 {
	var t float64
	// Calcula o subtotal dos produtos.
	for _, preco := range p {
		t += preco
	}

	// Aplica desconto por tipo de usuário.
	if u["tipo"] == "premium" {
		t *= 0.9 // 10% de desconto.
	} else if u["tipo"] == "plus" {
		t *= 0.95 // 5% de desconto.
	}

	// Aplica cupom de desconto.
	if c == "DESC10" {
		t -= 10
	} else if c == "DESC5" {
		t -= 5
	}

	// Adiciona taxa de envio para pedidos abaixo de 100.
	if t < 100 {
		t += 5.0
	}

	return t
}
