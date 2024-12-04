package controller_test

import (
	"fmt"
	"testing"

	"github.com/ylanzinhoy/transdata_saque_bank/cmd/services"
)

func ControllerTest(t *testing.T) {
	t.Run("525", func(t *testing.T) {
		res := services.Saque(525)
		fmt.Println(res)
	})
}
