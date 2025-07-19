package apresentacao

import (
	"testing"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo11/dominio"
)

func TestDescontoPorUsuario_AplicarComMapa(t *testing.T) {
	regra := NewDescontoPorUsuario()
	baseTotal := 100.0
	p, _ := dominio.NewProduto("Teste", 10.0)
	produtos := []*dominio.Produto{p}

	// Cenário 1: Usuário Premium (10% de desconto)
	usuarioPremium, _ := dominio.NewUsuario("Premium", dominio.TipoUsuarioPremium)
	pedidoPremium := dominio.NewPedido(produtos, usuarioPremium, "")
	totalPremium := regra.Aplicar(pedidoPremium, baseTotal)
	if totalPremium != 90.0 {
		t.Errorf("Desconto Premium incorreto. Esperado: 90.0, Obtido: %f", totalPremium)
	}

	// Cenário 2: Usuário Plus (20% de desconto)
	usuarioPlus, _ := dominio.NewUsuario("Plus", dominio.TipoUsuarioPlus)
	pedidoPlus := dominio.NewPedido(produtos, usuarioPlus, "")
	totalPlus := regra.Aplicar(pedidoPlus, baseTotal)
	if totalPlus != 80.0 {
		t.Errorf("Desconto Plus incorreto. Esperado: 80.0, Obtido: %f", totalPlus)
	}

	// Cenário 3: Usuário Normal (sem desconto)
	usuarioNormal, _ := dominio.NewUsuario("Normal", dominio.TipoUsuarioPadrao)
	pedidoNormal := dominio.NewPedido(produtos, usuarioNormal, "")
	totalNormal := regra.Aplicar(pedidoNormal, baseTotal)
	if totalNormal != baseTotal {
		t.Errorf("Nenhum desconto esperado. Esperado: %f, Obtido: %f", baseTotal, totalNormal)
	}
}

func TestDescontoPorCupom_AplicarComMapa(t *testing.T) {
	regra := NewDescontoPorCupom()
	baseTotal := 100.0
	p, _ := dominio.NewProduto("Teste", 10.0)
	usuario, _ := dominio.NewUsuario("Teste", dominio.TipoUsuarioPadrao)
	produtos := []*dominio.Produto{p}

	// Cenário 1: Cupom 10OFF
	pedido10OFF := dominio.NewPedido(produtos, usuario, dominio.Cupom10OFF)
	total10OFF := regra.Aplicar(pedido10OFF, baseTotal)
	if total10OFF != 90.0 {
		t.Errorf("Desconto 10OFF incorreto. Esperado: 90.0, Obtido: %f", total10OFF)
	}

	// Cenário 2: Cupom 5OFF
	pedido5OFF := dominio.NewPedido(produtos, usuario, dominio.Cupom5OFF)
	total5OFF := regra.Aplicar(pedido5OFF, baseTotal)
	if total5OFF != 95.0 {
		t.Errorf("Desconto 5OFF incorreto. Esperado: 95.0, Obtido: %f", total5OFF)
	}
}

func TestRegraDeTaxaDeEnvio_Aplicar(t *testing.T) {
	regra := &RegraDeTaxaDeEnvio{}
	p, _ := dominio.NewProduto("Teste", 10.0)
	usuario, _ := dominio.NewUsuario("Teste", dominio.TipoUsuarioPadrao)
	produtos := []*dominio.Produto{p}
	pedido := dominio.NewPedido(produtos, usuario, "")

	// Cenário 1: Total baixo, aplica taxa
	totalBaixo := 80.0
	totalComTaxa := regra.Aplicar(pedido, totalBaixo)
	if totalComTaxa != 84.0 {
		t.Errorf("Taxa de envio incorreta. Esperado: 84.0, Obtido: %f", totalComTaxa)
	}

	// Cenário 2: Total alto, não aplica taxa
	totalAlto := 120.0
	totalSemTaxa := regra.Aplicar(pedido, totalAlto)
	if totalSemTaxa != totalAlto {
		t.Errorf("Nenhuma taxa esperada. Esperado: %f, Obtido: %f", totalAlto, totalSemTaxa)
	}
}
