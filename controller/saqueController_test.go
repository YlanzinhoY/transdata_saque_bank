package controller_test

import (
	"bytes"
	"github.com/ylanzinhoy/transdata_saque_bank/controller"
	"testing"
)

func TestController(t *testing.T) {
	tests := []struct {
		name       string
		valor      int
		wantOutput string
	}{
		{
			name:       "Teste com 186",
			valor:      186,
			wantOutput: "Você sacou 1 nota(s) de 100\nVocê sacou 1 nota(s) de 50\nVocê sacou 1 nota(s) de 20\nVocê sacou 1 nota(s) de 10\nVocê sacou 1 nota(s) de 5\nVocê sacou 1 nota(s) de 1\n",
		},
		{
			name:       "Teste com 555",
			valor:      555,
			wantOutput: "Você sacou 5 nota(s) de 100\nVocê sacou 1 nota(s) de 50\nVocê sacou 1 nota(s) de 5\n",
		},
		{
			name:       "Teste com 5",
			valor:      5,
			wantOutput: "Você sacou 1 nota(s) de 5\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			controller.SaqueControllerComIo(tt.valor, &buf)
			got := buf.String()
			if got != tt.wantOutput {
				t.Errorf("Controller() = %q, want %q", got, tt.wantOutput)
			}
		})
	}
}
