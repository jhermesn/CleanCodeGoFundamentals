package dominio

import "errors"

/*
Passo 9: Tratamento Robusto de Erros.
A camada de domínio agora inclui validações. As entidades garantem a
consistência dos dados (ex: preço não pode ser negativo). As funções
retornam erros para sinalizar estados inválidos, tornando o sistema
mais robusto e previsível.
*/

var (
	ErrPrecoInvalido = errors.New("o preço do produto não pode ser negativo")
)

const (
	TipoUsuarioPremium = "premium"
	TipoUsuarioPlus    = "plus"
	Cupom10OFF         = "10OFF"
	Cupom5OFF          = "5OFF"
)

type Produto struct {
	Nome  string
	Preco float64
}

type Usuario struct {
	Nome string
	Tipo string
}

type Pedido struct {
	Produtos []*Produto
	Usuario  *Usuario
	Cupom    string
}

func (p *Pedido) Subtotal() (float64, error) {
	var subtotal float64
	for _, produto := range p.Produtos {
		if produto.Preco < 0 {
			return 0, ErrPrecoInvalido
		}
		subtotal += produto.Preco
	}
	return subtotal, nil
}