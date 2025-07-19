package apresentacao

import (
	"testing"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/aplicacao"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo10/dominio"
)

func TestDescontoPorUsuario_Aplicar(t *testing.T) {
	regra := &DescontoPorUsuario{}
	baseTotal := 100.0

	// Cenário 1: Usuário Premium
	pedidoPremium := &dominio.Pedido{Usuario: &dominio.Usuario{Tipo: dominio.TipoUsuarioPremium}}
	totalPremium := regra.Aplicar(pedidoPremium, baseTotal)
	esperadoPremium := baseTotal - (baseTotal * aplicacao.DescontoPremium)
	if totalPremium != esperadoPremium {
		t.Errorf("Desconto Premium incorreto. Esperado: %f, Obtido: %f", esperadoPremium, totalPremium)
	}

	// Cenário 2: Usuário Plus
	pedidoPlus := &dominio.Pedido{Usuario: &dominio.Usuario{Tipo: dominio.TipoUsuarioPlus}}
	totalPlus := regra.Aplicar(pedidoPlus, baseTotal)
	esperadoPlus := baseTotal - (baseTotal * 0.20)
	if totalPlus != esperadoPlus {
		t.Errorf("Desconto Plus incorreto. Esperado: %f, Obtido: %f", esperadoPlus, totalPlus)
	}

	// Cenário 3: Usuário Normal (sem desconto)
	pedidoNormal := &dominio.Pedido{Usuario: &dominio.Usuario{Tipo: "normal"}}
	totalNormal := regra.Aplicar(pedidoNormal, baseTotal)
	if totalNormal != baseTotal {
		t.Errorf("Nenhum desconto esperado. Esperado: %f, Obtido: %f", baseTotal, totalNormal)
	}
}

func TestDescontoPorCupom_Aplicar(t *testing.T) {
	regra := &DescontoPorCupom{}
	baseTotal := 100.0

	// Cenário 1: Cupom 10OFF
	pedido10OFF := &dominio.Pedido{Cupom: dominio.Cupom10OFF}
	total10OFF := regra.Aplicar(pedido10OFF, baseTotal)
	if total10OFF != 90.0 {
		t.Errorf("Desconto 10OFF incorreto. Esperado: 90.0, Obtido: %f", total10OFF)
	}

	// Cenário 2: Cupom 5OFF
	pedido5OFF := &dominio.Pedido{Cupom: dominio.Cupom5OFF}
	total5OFF := regra.Aplicar(pedido5OFF, baseTotal)
	if total5OFF != 95.0 {
		t.Errorf("Desconto 5OFF incorreto. Esperado: 95.0, Obtido: %f", total5OFF)
	}
}

func TestRegraDeTaxaDeEnvio_Aplicar(t *testing.T) {
	regra := &RegraDeTaxaDeEnvio{}

	// Cenário 1: Total baixo, aplica taxa
	totalBaixo := 80.0
	totalComTaxa := regra.Aplicar(&dominio.Pedido{}, totalBaixo)
	esperadoComTaxa := totalBaixo + (totalBaixo * aplicacao.TaxaDeEnvio)
	if totalComTaxa != esperadoComTaxa {
		t.Errorf("Taxa de envio incorreta. Esperado: %f, Obtido: %f", esperadoComTaxa, totalComTaxa)
	}

	// Cenário 2: Total alto, não aplica taxa
	totalAlto := 120.0
	totalSemTaxa := regra.Aplicar(&dominio.Pedido{}, totalAlto)
	if totalSemTaxa != totalAlto {
		t.Errorf("Nenhuma taxa esperada. Esperado: %f, Obtido: %f", totalAlto, totalSemTaxa)
	}
}