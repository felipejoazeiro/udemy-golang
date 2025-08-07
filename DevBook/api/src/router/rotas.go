package rotas 

type Rota struct {
	Uri 					string
	Metodo 					string
	Funcao 					func(http.ReponseWriter, *http.Request)
	RequerAutenticacao 		bool	
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios 

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}