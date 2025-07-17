package passo2

// Neste passo, aplicamos o princípio de "Nomenclatura Clara e Significativa".
//
// Mudanças realizadas:
// - A função `CalcularPreco` foi renomeada para `CalcularPrecoPedido` para descrever melhor sua finalidade.
// - O parâmetro `p` foi renomeado para `precosDosProdutos`.
// - O parâmetro `u` foi renomeado para `informacoesDoUsuario`.
// - O parâmetro `c` foi renomeado para `cupomDeDesconto`.
// - A variável local `t` foi renomeada para `total`.
//
// Benefícios:
// - O código se torna mais autoexplicativo.
// - Reduz a necessidade de comentários para explicar o que cada variável faz.
// - Facilita a leitura e o entendimento por outros desenvolvedores (ou por você mesmo no futuro).

func CalcularPrecoPedidoPasso2(precosDosProdutos []float64, informacoesDoUsuario map[string]string, cupomDeDesconto string) float64 {
	var total float64
	// 1. Calcula o total dos produtos
	for _, preco := range precosDosProdutos {
		total += preco
	}

	// 2. Aplicar desconto com base no tipo de usuário
	if informacoesDoUsuario["tipo"] == "premium" {
		total = total * 0.9 // 10% de desconto para premium
	} else if informacoesDoUsuario["tipo"] == "plus" {
		total = total * 0.95 // 5% de desconto para plus
	}

	// 3. Aplicar cupom de desconto
	if cupomDeDesconto == "DESC10" {
		total = total - 10
	} else if cupomDeDesconto == "DESC5" {
		total = total - 5
	}

	// 4. Adicionar taxa de envio
	if total < 100 {
		total = total + 5.0 // Taxa de R$5 para pedidos abaixo de R$100
	}

	return total
}
