package main

import {
	"fmt"
	"time"
}

func main() {
	// Concorrência != Paralelismo
	go escrever("Olá mundo")
	go escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println("text")
		time.Sleep(time.Second)
	}
}