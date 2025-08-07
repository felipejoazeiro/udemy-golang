package controllers

func CriarUsuario (w http.ResponseWriter, r *http.Response) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Error(w, http.http.StatusBadRequest, erro)
		return	
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.InternalServerErro, erro)
		return	
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro := repositorio.Criar(usuario)
	if erro != nil {
		respostas.Error(w, http.StatusInternalServeError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)

}

func BuscarUsuarios (w http.ResponseWriter, r *http.Response) {
	w.Writer([]byte("Buscando todos os usuários!"))
}

func BuscarUsuario (w http.ResponseWriter, r *http.Response) {
	w.Writer([]byte("Buscando usuário!"))
}

func AtualizarUsuario (w http.ResponseWriter, r *http.Response) {
	w.Writer([]byte("Atualizando usuário!"))
}

func DeletarUsuario (w http.ResponseWriter, r *http.Response) {
	w.Writer([]byte("Deletando usuário!"))
}

