package dominio

/*
Passo 8: Estrutura em Camadas (Clean Architecture).
A camada de domínio contém as entidades e regras de negócio essenciais.
Estas são as estruturas de dados e a lógica que representam o coração
do sistema, independentes de qualquer detalhe de implementação externa.
*/

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

func (p *Pedido) Subtotal() float64 {
	var subtotal float64
	for _, produto := range p.Produtos {
		subtotal += produto.Preco
	}
	return subtotal
}