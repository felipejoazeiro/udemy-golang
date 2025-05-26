package main

import "fmt"

func soma(numeros ...int){
	fmt.Println(numeros)
}

func escrever(texto string, numero ...int) string {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main(){
	soma(1,2,3,4,5,6,200,102,12,13)
}