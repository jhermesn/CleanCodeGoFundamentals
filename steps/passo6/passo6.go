package passo6

// package passo6 consolida a aplicação dos princípios SOLID, usando o padrão Strategy
// para criar um sistema de cálculo de preços flexível e extensível.

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

// --- Funções Auxiliares ---
func CalcularTotalDosProdutos(produtos []Produto) float64 {
	var total float64
	for _, produto := range produtos {
		total += produto.Preco
	}
	return total
}

func AdicionarTaxaDeEnvio(total float64) float64 {
	if total < ValorMinimoParaEnvioGratis {
		return total + TaxaDeEnvio
	}
	return total
}

// --- Cálculo Final ---

// CalcularPrecoFinal itera sobre as regras de desconto e as aplica ao total.
func CalcularPrecoFinal(totalInicial float64, descontos []ICalculadorDeDesconto) float64 {
	total := totalInicial
	for _, desconto := range descontos {
		total = desconto.Aplicar(total)
	}
	return AdicionarTaxaDeEnvio(total)
}
