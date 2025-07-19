package aplicacao

import "github.com/jhermesn/CleanCodeGoFundamentals/steps/passo8/dominio"

/*
Passo 8: Estrutura em Camadas (Clean Architecture).
A camada de aplicação contém a lógica específica dos casos de uso do sistema.
Ela orquestra as entidades do domínio para executar as regras de negócio
e depende do domínio, mas não de camadas externas como a apresentação.
*/

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

func (s *ServicoDeCalculoDePreco) Calcular(pedido *dominio.Pedido) float64 {
	total := pedido.Subtotal()
	for _, regra := range s.Regras {
		total = regra.Aplicar(pedido, total)
	}
	return total
}