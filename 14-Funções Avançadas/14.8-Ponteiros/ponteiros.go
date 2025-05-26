package main

import "fmt"

func inverSinal(numero int) int {
	return numero * -1
}

func invertSinalComPonteiro(numero *int){
	*numero = *numero * -1
}

func main(){
	numero := 20

	numeroInvertido := inverSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero = := 40
	fmt.Println(novoNumero)
	invertSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

