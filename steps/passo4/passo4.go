package passo4

// package passo4 aplica o Princípio da Responsabilidade Única (SRP).
//
// A função de cálculo de preço foi dividida em funções menores e coesas,
// cada uma com uma única responsabilidade. A função principal agora
// apenas orquestra as chamadas, tornando o fluxo mais claro e testável.

const (
	FatorDescontoPremiumPasso4       = 0.90
	FatorDescontoPlusPasso4          = 0.95
	CupomDesconto10Passo4            = "DESC10"
	ValorCupom10Passo4               = 10.0
	CupomDesconto5Passo4             = "DESC5"
	ValorCupom5Passo4                = 5.0
	ValorMinimoParaEnvioGratisPasso4 = 100.0
	TaxaDeEnvioPasso4                = 5.0
	TipoUsuarioPremiumPasso4         = "premium"
	TipoUsuarioPlusPasso4            = "plus"
)

type UsuarioPasso4 struct {
	Tipo string
}

type ProdutoPasso4 struct {
	Preco float64
}

type PedidoPasso4 struct {
	Produtos []ProdutoPasso4
	Cupom    string
}

// CalcularPrecoPedidoRefatoradoPasso4 orquestra as etapas do cálculo de preço.
func CalcularPrecoPedidoRefatoradoPasso4(pedido PedidoPasso4, usuario UsuarioPasso4) float64 {
	total := calcularTotalDosProdutosPasso4(pedido.Produtos)
	total = aplicarDescontoPorTipoUsuarioPasso4(total, usuario)
	total = aplicarDescontoDeCupomPasso4(total, pedido.Cupom)
	total = adicionarTaxaDeEnvioPasso4(total)
	return total
}

// calcularTotalDosProdutos soma os preços de todos os produtos do pedido.
func calcularTotalDosProdutosPasso4(produtos []ProdutoPasso4) float64 {
	var total float64
	for _, produto := range produtos {
		total += produto.Preco
	}
	return total
}

// aplicarDescontoPorTipoUsuario aplica o desconto com base no tipo de usuário.
func aplicarDescontoPorTipoUsuarioPasso4(total float64, usuario UsuarioPasso4) float64 {
	if usuario.Tipo == TipoUsuarioPremiumPasso4 {
		return total * FatorDescontoPremiumPasso4
	}
	if usuario.Tipo == TipoUsuarioPlusPasso4 {
		return total * FatorDescontoPlusPasso4
	}
	return total
}

// aplicarDescontoDeCupom aplica o desconto de um cupom de valor fixo.
func aplicarDescontoDeCupomPasso4(total float64, cupom string) float64 {
	if cupom == CupomDesconto10Passo4 {
		return total - ValorCupom10Passo4
	}
	if cupom == CupomDesconto5Passo4 {
		return total - ValorCupom5Passo4
	}
	return total
}

// adicionarTaxaDeEnvio adiciona a taxa se o total for abaixo do mínimo.
func adicionarTaxaDeEnvioPasso4(total float64) float64 {
	if total < ValorMinimoParaEnvioGratisPasso4 {
		return total + TaxaDeEnvioPasso4
	}
	return total
}
