package enderecos

func tipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}


	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavradoEndereco := strings.Split(endereco, " ")[0]

	enderecoTemUmTipoValido := false 
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco{
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo Inválido"
}

