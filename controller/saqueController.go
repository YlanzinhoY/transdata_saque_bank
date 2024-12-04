package controller

import (
	"fmt"
	"io"
	"math"
)

func SaqueController(valor int) {
	notas := []int{100, 50, 20, 10, 5, 2}
	resultado := make([]int, len(notas))
	melhorResultado := make([]int, len(notas))
	menorQuantidadeDeNotas := math.MaxInt

	var busca func(restante, index, qtdNotas int)
	busca = func(restante, index, qtdNotas int) {
		if restante == 0 {
			if qtdNotas < menorQuantidadeDeNotas {
				menorQuantidadeDeNotas = qtdNotas
				copy(melhorResultado, resultado)
			}
			return
		}

		if restante < 0 || index >= len(notas) {
			return
		}

		busca(restante, index+1, qtdNotas)

		for i := 1; restante >= i*notas[index]; i++ {
			resultado[index] = i
			busca(restante-i*notas[index], index+1, qtdNotas+i)
			resultado[index] = 0
		}
	}

	busca(valor, 0, 0)

	if menorQuantidadeDeNotas == math.MaxInt {
		fmt.Println("Não foi possível sacar o valor com as notas disponíveis.")
		return
	}

	fmt.Printf("Valor: %d\n", valor)
	for i, qtd := range melhorResultado {
		if qtd > 0 {
			fmt.Printf("Você sacou %d nota(s) de %d Real(Reais)\n", qtd, notas[i])
		}
	}
}

func SaqueControllerComIo(valor int, w io.Writer) {
	notas := []int{
		100,
		50,
		20,
		10,
		5,
		2,
	}

	for k := range notas {
		div := math.Floor(float64(valor) / float64(notas[k]))

		if div > 0 {
			valor -= int(div) * notas[k]
			res := int(div)
			_, err := fmt.Fprintf(w, "Você sacou %d nota(s) de %d\n", res, notas[k])
			if err != nil {
				return
			}
		}
	}
}
