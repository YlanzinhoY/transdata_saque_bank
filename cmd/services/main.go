package services

import (
	"fmt"
	"math"
)

func Saque(valor int) string {
	notas := []int{
		100,
		50,
		20,
		10,
		5,
		2,
		1,
	}

	for k := range notas {

		div := math.Floor(float64(valor) / float64(notas[k]))

		if div > 0 {
			valor -= int(div) * notas[k]
			res := int(div)
			valorFinal := fmt.Sprintf("Você sacou %d nota(s) de %d\n", res, notas[k])
			return valorFinal
		}

	}

	return ""
}
