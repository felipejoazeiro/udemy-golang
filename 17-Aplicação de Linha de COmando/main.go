package main

import (
	"fmt"
	"linha-de-comando/app" // Importando o pacote app
)

func main(){
	fmt.Println("Hello, World!")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}