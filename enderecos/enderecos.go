package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem um tipo válido e o retona
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	endrecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalabraDoEndereco := strings.Split(endrecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalabraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalabraDoEndereco)
	}

	return "Tipo Inválido"
}