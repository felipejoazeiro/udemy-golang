package main

import "fmt"

type usuario struct{
	nome 		strign
	idade 		uint8
	endereco 	endereco
}

type endereco struct{
	logradouro 	string
	numero 		uint8
}

func main(){
	fmt.Println("Arquivo Structs")
	enderecoExemplo := endereco{"Rua dos bobos", 0}

	var u usuario 
	u.nome ="Felipe"
	u.idade = 21
	u.endereco = enderecoExemplo
	fmt.Println(u)


	u2 := usuario{"Gabriel", 36}
	fmt.Prinln(u2)

	usuario3 := usuario{idade:25}
	fmt.Println(usuario3)
}