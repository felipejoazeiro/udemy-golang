package enderecos_test

import (
	"testing"
	. "introducao-testes/enderecos"
)

type cenarioDeTeste struct {
	endereceInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada Qualquer", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"", "Tipo Inválido"},
		{"Estrada de Testes", "Estrada"},
		{"Rua Ana", "Rua"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.endereceInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado,)
		}
	}

}


func TestQualquer(t *testing.T) {
	t.Parallel()
	if > 2 {
		t.Errorf("Teste quebrou!")
	}
}