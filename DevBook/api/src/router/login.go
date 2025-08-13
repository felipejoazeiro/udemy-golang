package rotas

var rotaLogin = Rota{
	Uri: 					"/login",
	Metodo: 				http.MethodPost,
	Funcao: 				controllers.Login,
	RequerAutenticacao: 	false,
}