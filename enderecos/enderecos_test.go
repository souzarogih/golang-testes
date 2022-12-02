package enderecos

import "testing"

// TESTE DE UNIDADE

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := [] cenarioDeTeste {
		{ "Rua ABC", "Rua"},
		{ "Avenida Paulista", "Avenida"},
		{ "Rodovia dos Imigrantes", "Rodovia"},
		{ "Praça das Rosas", "Tipo Inválido"},
		{ "Estrada Qualquer", "Estrada"},
		{ "RUA OS BOBOS", "Rua"},
		{ "AVENIDA REBOUÇAS", "Avenida"},
		{ "", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido [%s] é diferente do esperado [%s]",
			tipoDeEnderecoRecebido,
		cenario.retornoEsperado,)
		}
	}
}

func TestTipoDeEndereco2(t *testing.T) {
	t.Parallel()
	// Forma simples de fazer teste
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava [%s] e recebeu [%s]",
	tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}