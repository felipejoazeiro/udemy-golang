package controllers

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return	
	}

	/* if erro = usuario.Preparar(); erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	} */

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return	
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil || usuarioSalvoNoBanco.ID == 0 {
		respostas.Error(w, http.StatusUnauthorized, errors.New("usuário ou senha inválidos"))
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Error(w, http.StatusUnauthorized, errors.New("usuário ou senha inválidos"))
		return
	}

	token, _ = autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	fmt.Println(token)

	w.Write([]byte(token))

	respostas.JSON(w, http.StatusOK, usuarioSalvoNoBanco)
}