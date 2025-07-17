package passo5

// package passo5 aplica os princípios Aberto/Fechado (OCP) e Inversão de Dependência (DIP)
// usando o padrão de projeto Strategy. O sistema agora é extensível a novas regras
// de desconto sem modificar o código de cálculo existente.

// --- Constantes e Estruturas Base (mantidas dos passos anteriores) ---
const (
	FatorDescontoPremium       = 0.90
	FatorDescontoPlus          = 0.95
	CupomDesconto10            = "DESC10"
	ValorCupom10               = 10.0
	CupomDesconto5             = "DESC5"
	ValorCupom5                = 5.0
	ValorMinimoParaEnvioGratis = 100.0
	TaxaDeEnvio                = 5.0
	TipoUsuarioPremium         = "premium"
	TipoUsuarioPlus            = "plus"
)

type Usuario struct {
	Tipo string
}

type Produto struct {
	Preco float64
}

type Pedido struct {
	Produtos []Produto
	Cupom    string
}

// --- Abstração e Implementações ---

// ICalculadorDeDesconto define o contrato para qualquer regra de desconto.
type ICalculadorDeDesconto interface {
	Aplicar(total float64) float64
}

// DescontoUsuario implementa a regra de desconto por tipo de usuário.
type DescontoUsuario struct {
	Fator float64
}

func (d DescontoUsuario) Aplicar(total float64) float64 {
	return total * d.Fator
}

// DescontoCupom implementa a regra de desconto por cupom.
type DescontoCupom struct {
	Valor float64
}

func (d DescontoCupom) Aplicar(total float64) float64 {
	novoTotal := total - d.Valor
	if novoTotal < 0 {
		return 0 // Evita total negativo.
	}
	return novoTotal
}

// CalcularTotalDosProdutos soma os preços de todos os produtos do pedido.
func CalcularTotalDosProdutos(produtos []Produto) float64 {
	var total float64
	for _, produto := range produtos {
		total += produto.Preco
	}
	return total
}

func adicionarTaxaDeEnvio(total float64) float64 {
	if total < ValorMinimoParaEnvioGratis {
		return total + TaxaDeEnvio
	}
	return total
}

// --- Cálculo Final ---

// CalcularPrecoFinal itera sobre uma lista de descontos e os aplica ao total.
// A função depende da abstração ICalculadorDeDesconto, não de implementações concretas.
func CalcularPrecoFinal(totalInicial float64, descontos []ICalculadorDeDesconto) float64 {
	total := totalInicial
	for _, desconto := range descontos {
		total = desconto.Aplicar(total)
	}
	return adicionarTaxaDeEnvio(total)
}
