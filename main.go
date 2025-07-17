package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo1"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo2"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo3"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo4"
	"github.com/jhermesn/CleanCodeGoFundamentals/steps/passo5"
)

const (
	Passo1 = "passo1"
	Passo2 = "passo2"
	Passo3 = "passo3"
	Passo4 = "passo4"
	Passo5 = "passo5"
	Passo6 = "passo6"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, especifique a versão a ser executada (ex: passo1, passo2, ...)")
		return
	}
	versao := os.Args[1]

	switch versao {
	case Passo1:
		executarPasso1()
	case Passo2:
		executarPasso2()
	case Passo3:
		executarPasso3()
	case Passo4:
		executarPasso4()
	case Passo5:
		executarPasso5()
	case Passo6:
		executarPasso6()
	default:
		log.Fatalf("Versão desconhecida: %s", versao)
	}
}

func executarPasso1() {
	titulo := "Executando Passo 1: Código Inicial"
	precos := []float64{80.0, 40.0}
	usuario := map[string]string{"tipo": "premium"}
	cupom := "DESC10"

	precoFinal := passo1.CalcularPrecoPasso1(precos, usuario, cupom)
	imprimirResultado(titulo, "Produtos (R$120), Usuário Premium, Cupom DESC10", precoFinal)
}

func executarPasso2() {
	titulo := "Executando Passo 2: Nomenclatura Clara"
	precos := []float64{80.0, 40.0}
	usuario := map[string]string{"tipo": "premium"}
	cupom := "DESC10"

	precoFinal := passo2.CalcularPrecoPedidoPasso2(precos, usuario, cupom)
	imprimirResultado(titulo, "Produtos (R$120), Usuário Premium, Cupom DESC10", precoFinal)
}

func executarPasso3() {
	titulo := "Executando Passo 3: Sem Números Mágicos e com Estruturas"
	usuario := passo3.UsuarioPasso3{Tipo: passo3.TipoUsuarioPremiumPasso3}
	pedido := passo3.PedidoPasso3{
		Produtos: []passo3.ProdutoPasso3{{Preco: 80.0}, {Preco: 40.0}},
		Cupom:    passo3.CupomDesconto10Passo3,
	}

	precoFinal := passo3.CalcularPrecoPedidoComEstruturasPasso3(pedido, usuario)
	imprimirResultado(titulo, "Produtos (R$120), Usuário Premium, Cupom DESC10", precoFinal)
}

func executarPasso4() {
	titulo := "Executando Passo 4: Funções Pequenas e SRP"
	usuario := passo4.UsuarioPasso4{Tipo: passo4.TipoUsuarioPremiumPasso4}
	pedido := passo4.PedidoPasso4{
		Produtos: []passo4.ProdutoPasso4{{Preco: 80.0}, {Preco: 40.0}},
		Cupom:    passo4.CupomDesconto10Passo4,
	}

	precoFinal := passo4.CalcularPrecoPedidoRefatoradoPasso4(pedido, usuario)
	imprimirResultado(titulo, "Produtos (R$120), Usuário Premium, Cupom DESC10", precoFinal)
}

func executarPasso5() {
	titulo := "Executando Passo 5: Lógica de Montagem no Main"
	usuario := passo5.Usuario{Tipo: passo5.TipoUsuarioPremium}
	pedido := passo5.Pedido{
		Produtos: []passo5.Produto{{Preco: 80.0}, {Preco: 40.0}},
		Cupom:    passo5.CupomDesconto10,
	}
	subtotal := passo5.CalcularTotalDosProdutos(pedido.Produtos)
	var descontosAplicaveis []passo5.ICalculadorDeDesconto
	if usuario.Tipo == passo5.TipoUsuarioPremium {
		descontosAplicaveis = append(descontosAplicaveis, passo5.DescontoUsuario{Fator: passo5.FatorDescontoPremium})
	} else if usuario.Tipo == passo5.TipoUsuarioPlus {
		descontosAplicaveis = append(descontosAplicaveis, passo5.DescontoUsuario{Fator: passo5.FatorDescontoPlus})
	}
	if pedido.Cupom == passo5.CupomDesconto10 {
		descontosAplicaveis = append(descontosAplicaveis, passo5.DescontoCupom{Valor: passo5.ValorCupom10})
	} else if pedido.Cupom == passo5.CupomDesconto5 {
		descontosAplicaveis = append(descontosAplicaveis, passo5.DescontoCupom{Valor: passo5.ValorCupom5})
	}
	precoFinal := passo5.CalcularPrecoFinal(subtotal, descontosAplicaveis)
	imprimirResultado(titulo, "Produtos (R$120), Usuário Premium, Cupom DESC10", precoFinal)
}

func executarPasso6() {
	fmt.Println("--- Executando os testes de unidade para o Passo 6... ---")
	cmd := exec.Command("go", "test", "-v", "./steps/passo6/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("A execução dos testes falhou: %v", err)
	}
}

func imprimirResultado(titulo, cenario string, precoFinal float64) {
	fmt.Printf("--- %s ---\n", titulo)
	fmt.Printf("Cenário: %s\n", cenario)
	fmt.Printf("Preço Final Calculado: R$%.2f\n", precoFinal)
	fmt.Println(strings.Repeat("-", len(titulo)+6))
}
