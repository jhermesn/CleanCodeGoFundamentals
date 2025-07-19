package aplicacao

import (
	"errors"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo9/dominio"
)

/*
Passo 9: Tratamento Robusto de Erros.
A camada de aplicação agora propaga os erros vindos do domínio e pode
adicionar seus próprios erros de lógica de aplicação. O caso de uso
de cálculo de preço agora retorna um erro, permitindo que a camada
de apresentação decida como lidar com falhas.
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
	Regras []RegraDeCalculo
}

func (s *ServicoDeCalculoDePreco) Calcular(pedido *dominio.Pedido) (float64, error) {
	if pedido == nil {
		return 0, ErrPedidoInvalido
	}

	total, err := pedido.Subtotal()
	if err != nil {
		return 0, err
	}

	for _, regra := range s.Regras {
		total = regra.Aplicar(pedido, total)
	}

	return total, nil
}