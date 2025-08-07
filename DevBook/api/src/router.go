package router

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}