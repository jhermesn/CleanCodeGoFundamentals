package dominio

import (
	"testing"
)

func TestPedido_Subtotal_Sucesso(t *testing.T) {
	pedido := &Pedido{
		Produtos: []*Produto{
			{Nome: "Produto A", Preco: 50.50},
			{Nome: "Produto B", Preco: 10.00},
			{Nome: "Produto C", Preco: 39.50},
		},
	}
	esperado := 100.0

	resultado, err := pedido.Subtotal()

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado != esperado {
		t.Errorf("Subtotal incorreto. Esperado: %f, Obtido: %f", esperado, resultado)
	}
}

func TestPedido_Subtotal_ErroPrecoNegativo(t *testing.T) {
	pedido := &Pedido{
		Produtos: []*Produto{
			{Nome: "Produto A", Preco: 50.0},
			{Nome: "Produto Inválido", Preco: -10.0},
		},
	}

	_, err := pedido.Subtotal()

	if err == nil {
		t.Errorf("Esperado um erro de preço inválido, mas nenhum erro foi retornado")
	}
	if err != ErrPrecoInvalido {
		t.Errorf("Tipo de erro incorreto. Esperado: %v, Obtido: %v", ErrPrecoInvalido, err)
	}
}