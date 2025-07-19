package dominio

import "errors"

/*
Passo 10: Injeção de Dependência e Testes Automatizados.
A camada de domínio permanece inalterada em sua lógica, mas agora é
validada por um conjunto de testes unitários que garantem a corretude
do cálculo de subtotal e do tratamento de erros, provando sua robustez.
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
