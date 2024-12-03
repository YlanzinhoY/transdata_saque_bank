package controller

import (
	"fmt"
	"math"
)

func Controller(valor int) {

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
			fmt.Printf("Você sacou %d notas de %d\n", res, notas[k])
		}

	}

}
