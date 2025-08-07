package main

func main() {
	config.Carregar()
	fmt.Println(config.StringConexaoBanco)


	fmt.Println("Rodando a API")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}