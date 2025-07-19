package aplicacao

import (
	"errors"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo11/dominio"
)

/*
Passo 11: Refatoração de Regras Semânticas.
A camada de aplicação foi purificada. As constantes de valores de desconto,
que são detalhes de implementação, foram removidas desta camada e movidas
para junto das regras concretas na camada de apresentação.
*/

var (
	ErrPedidoInvalido = errors.New("pedido não pode ser nulo")
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

	total := pedido.Subtotal()

	for _, regra := range s.regras {
		total = regra.Aplicar(pedido, total)
	}

	return total, nil
}