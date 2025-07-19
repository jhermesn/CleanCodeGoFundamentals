package aplicacao

import (
	"errors"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/dominio"
)

/*
Passo 10: Injeção de Dependência e Testes Automatizados.
A camada de aplicação, que recebe suas dependências via construtor,
agora tem sua lógica de orquestração validada por testes unitários.
O uso de mocks comprova que a camada é testável em total isolamento,
um dos principais benefícios desta arquitetura.
*/

var (
	ErrPedidoInvalido = errors.New("pedido não pode ser nulo")
)

const (
	DescontoPremium     = 0.10
	DescontoPlus        = 0.20
	ValorDesconto10     = 10.0
	ValorDesconto5      = 5.0
	ValorMinimoParaTaxa = 100.0
	TaxaDeEnvio         = 0.05
)

type RegraDeCalculo interface {
	Aplicar(pedido *dominio.Pedido, total float64) float64
}

type ServicoDeCalculoDePreco struct {
	regras []RegraDeCalculo
}

func NovoServicoDeCalculoDePreco(regras []RegraDeCalculo) *ServicoDeCalculoDePreco {
	return &ServicoDeCalculoDePreco{regras: regras}
}

func (s *ServicoDeCalculoDePreco) Calcular(pedido *dominio.Pedido) (float64, error) {
	if pedido == nil {
		return 0, ErrPedidoInvalido
	}

	total, err := pedido.Subtotal()
	if err != nil {
		return 0, err
	}

	for _, regra := range s.regras {
		total = regra.Aplicar(pedido, total)
	}

	return total, nil
}
