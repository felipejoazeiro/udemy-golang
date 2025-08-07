package rotas


var rotasUsuarios = []Rota {
	{
		URI: "/usuarios"
		Metdod: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao : false,
	},
	{
		URI: "/usuarios"
		Metdod: http.MethodGet,
		Funcao: controller.BuscarUsuarios,
		RequerAutenticacao : false,
	},
	{
		URI: "/usuarios/{usuarioId}"
		Metdod: http.MethodGet,
		Funcao: controller.BuscarUsuario,
		RequerAutenticacao : false,
	},
	{
		URI: "/usuarios/{usuarioId}"
		Metdod: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao : false,
	},
	{
		URI: "/usuarios/{usuarioId}"
		Metdod: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao : false,
	}
}	