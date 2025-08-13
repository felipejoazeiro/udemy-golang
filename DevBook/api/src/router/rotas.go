package rotas 

type Rota struct {
	Uri 					string
	Metodo 					string
	Funcao 					func(http.ReponseWriter, *http.Request)
	RequerAutenticacao 		bool	
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios 
	rotas := append(rotas,rotaLogin)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao)),).Methods(rota.Metodo)
		}else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}
	}

	return r
}