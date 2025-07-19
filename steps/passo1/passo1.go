package passo1

import "fmt"

/* package passo1 demonstra o ponto de partida do código.
A função monolítica abaixo calcula o preço total de um pedido e possui diversos problemas:
- Nomes de funções e variáveis curtos e não descritivos.
- "Números Mágicos" sem explicação.
- Múltiplas responsabilidades (cálculo de subtotal, descontos, taxas e impressão do resultado).
- Uso de tipos de dados genéricos em vez de estruturas de domínio.
- Lógica complexa e aninhada.
*/

func Cp(p []float64, u map[string]string, c string) {
	var t float64
	for _, pc := range p {
		t += pc
	}

	if u["t"] == "pr" {
		t *= 0.9
	} else if u["t"] == "pl" {
		t *= 0.8
	}

	if c == "10OFF" {
		t -= 10
	} else if c == "5OFF" {
		t -= 5
	}

	if t < 100 {
		t *= 1.05
	}

	fmt.Println(t)
}