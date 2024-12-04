package controller

import (
	"fmt"
	"io"
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
		1,
	}

	for k := range notas {
		div := math.Floor(float64(valor) / float64(notas[k]))

		if div > 0 {
			valor -= int(div) * notas[k]
			res := int(div)
			fmt.Printf("Você sacou %d nota(s) de %d Real(Reais)\n", res, notas[k])
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
		1,
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
