package rotas


var rotasUsuarios = []Rota {
	{
		URI: "/usuarios"
		Method: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao : false,
	},
	{
		URI: "/usuarios"
		Method: http.MethodGet,
		Funcao: controller.BuscarUsuarios,
		RequerAutenticacao : true,
	},
	{
		URI: "/usuarios/{usuarioId}"
		Method: http.MethodGet,
		Funcao: controller.BuscarUsuario,
		RequerAutenticacao : true,
	},
	{
		URI: "/usuarios/{usuarioId}"
		Method: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao : true,
	},
	{
		URI: "/usuarios/{usuarioId}"
		Method: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao : true,
	}
}	