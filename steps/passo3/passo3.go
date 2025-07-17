package passo3

// package passo3 substitui "Números Mágicos" por constantes e introduz structs
// para modelar o domínio, melhorando a manutenibilidade e a segurança de tipos.

// Constantes nomeadas para evitar "Números Mágicos".
const (
	FatorDescontoPremiumPasso3       = 0.90
	FatorDescontoPlusPasso3          = 0.95
	CupomDesconto10Passo3            = "DESC10"
	ValorCupom10Passo3               = 10.0
	CupomDesconto5Passo3             = "DESC5"
	ValorCupom5Passo3                = 5.0
	ValorMinimoParaEnvioGratisPasso3 = 100.0
	TaxaDeEnvioPasso3                = 5.0
	TipoUsuarioPremiumPasso3         = "premium"
	TipoUsuarioPlusPasso3            = "plus"
)

// Usuario representa um usuário do sistema.
type UsuarioPasso3 struct {
	Tipo string
}

// Produto representa um item do pedido.
type ProdutoPasso3 struct {
	Preco float64
}

// Pedido agrupa os produtos e o cupom de um pedido.
type PedidoPasso3 struct {
	Produtos []ProdutoPasso3
	Cupom    string
}

// CalcularPrecoPedidoComEstruturas usa constantes e structs para maior clareza.
func CalcularPrecoPedidoComEstruturasPasso3(pedido PedidoPasso3, usuario UsuarioPasso3) float64 {
	var total float64
	// Calcula o subtotal dos produtos.
	for _, produto := range pedido.Produtos {
		total += produto.Preco
	}

	// Aplica desconto por tipo de usuário.
	if usuario.Tipo == TipoUsuarioPremiumPasso3 {
		total *= FatorDescontoPremiumPasso3
	} else if usuario.Tipo == TipoUsuarioPlusPasso3 {
		total *= FatorDescontoPlusPasso3
	}

	// Aplica cupom de desconto.
	if pedido.Cupom == CupomDesconto10Passo3 {
		total -= ValorCupom10Passo3
	} else if pedido.Cupom == CupomDesconto5Passo3 {
		total -= ValorCupom5Passo3
	}

	// Adiciona taxa de envio.
	if total < ValorMinimoParaEnvioGratisPasso3 {
		total += TaxaDeEnvioPasso3
	}

	return total
}
