package dominio

import "errors"

/*
Passo 11: Refatoração de Regras Semânticas.
Esta camada não sofreu alterações, demonstrando que as mudanças na
implementação das regras de negócio foram contidas em outra camada,
respeitando os princípios da arquitetura limpa.
*/

var (
	ErrPrecoInvalido = errors.New("o preço do produto não pode ser negativo")
)

type TipoUsuario string

const (
	TipoUsuarioPadrao  TipoUsuario = "padrao"
	TipoUsuarioPremium TipoUsuario = "premium"
	TipoUsuarioPlus    TipoUsuario = "plus"
	Cupom10OFF                     = "10OFF"
	Cupom5OFF                      = "5OFF"
)

func (t TipoUsuario) IsValid() bool {
	switch t {
	case TipoUsuarioPadrao, TipoUsuarioPremium, TipoUsuarioPlus:
		return true
	}
	return false
}

type Produto struct {
	Nome  string
	Preco float64
}

func NewProduto(nome string, preco float64) (*Produto, error) {
	if preco < 0 {
		return nil, ErrPrecoInvalido
	}
	return &Produto{Nome: nome, Preco: preco}, nil
}

type Usuario struct {
	Nome string
	Tipo TipoUsuario
}

func NewUsuario(nome string, tipo TipoUsuario) (*Usuario, error) {
	if !tipo.IsValid() {
		return nil, errors.New("tipo de usuário inválido")
	}
	return &Usuario{Nome: nome, Tipo: tipo}, nil
}

type Pedido struct {
	Produtos []*Produto
	Usuario  *Usuario
	Cupom    string
}

func NewPedido(produtos []*Produto, usuario *Usuario, cupom string) *Pedido {
	return &Pedido{
		Produtos: produtos,
		Usuario:  usuario,
		Cupom:    cupom,
	}
}

func (p *Pedido) Subtotal() float64 {
	var subtotal float64
	for _, produto := range p.Produtos {
		subtotal += produto.Preco
	}
	return subtotal
}