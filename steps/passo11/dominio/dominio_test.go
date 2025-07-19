package dominio

import (
	"testing"
)

func TestNewProduto_Sucesso(t *testing.T) {
	produto, err := NewProduto("Produto A", 50.50)
	if err != nil {
		t.Fatalf("Criação de produto falhou: %v", err)
	}
	if produto.Nome != "Produto A" || produto.Preco != 50.50 {
		t.Errorf("Dados do produto incorretos. Esperado: 'Produto A', 50.50. Obtido: '%s', %f", produto.Nome, produto.Preco)
	}
}

func TestNewProduto_ErroPrecoNegativo(t *testing.T) {
	_, err := NewProduto("Produto Inválido", -10.0)
	if err == nil {
		t.Error("Esperado um erro de preço inválido, mas nenhum erro foi retornado")
	}
	if err != ErrPrecoInvalido {
		t.Errorf("Tipo de erro incorreto. Esperado: %v, Obtido: %v", ErrPrecoInvalido, err)
	}
}

func TestNewUsuario_Sucesso(t *testing.T) {
	usuario, err := NewUsuario("Cliente Premium", TipoUsuarioPremium)
	if err != nil {
		t.Fatalf("Criação de usuário falhou: %v", err)
	}
	if usuario.Nome != "Cliente Premium" || usuario.Tipo != TipoUsuarioPremium {
		t.Errorf("Dados do usuário incorretos.")
	}
}

func TestNewUsuario_ErroTipoInvalido(t *testing.T) {
	_, err := NewUsuario("Cliente Inválido", "invalido")
	if err == nil {
		t.Error("Esperado um erro de tipo de usuário inválido, mas nenhum erro foi retornado")
	}
}

func TestPedido_Subtotal(t *testing.T) {
	p1, _ := NewProduto("Produto A", 50.50)
	p2, _ := NewProduto("Produto B", 10.00)
	p3, _ := NewProduto("Produto C", 39.50)
	usuario, _ := NewUsuario("Cliente", TipoUsuarioPadrao)

	pedido := NewPedido([]*Produto{p1, p2, p3}, usuario, "")

	esperado := 100.0
	resultado := pedido.Subtotal()

	if resultado != esperado {
		t.Errorf("Subtotal incorreto. Esperado: %f, Obtido: %f", esperado, resultado)
	}
}
