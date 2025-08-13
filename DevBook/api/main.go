package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	config.Carregar()
	fmt.Println(config.StringConexaoBanco)

	fmt.Println(config.SecretKey)

	fmt.Println("Rodando a API")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}