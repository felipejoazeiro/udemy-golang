package modelos

type Usuario struct {
	ID 			uint64 			`json: "id,omitempty"`
	Nome		string			`json:"nome,omitempty"`
	Nick 		string			`json:"nick,omitempty"`
	Email		string			`json:"email,omitempty"`
	Senha		string			`json:"senha,omitempty"`
	CriadoEm	time.time		`json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return erros.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return erros.New("O Nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return erros.New("O Email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return erros.New("O Email é inválido")
	}

	if usuario.Senha == "" {
		return erros.New("O Senha é obrigatório e não pode estar em branco")
	}

	return nil
}


func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}