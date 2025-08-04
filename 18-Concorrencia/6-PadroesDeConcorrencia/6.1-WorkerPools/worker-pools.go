package main

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)

	for i:=0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i<45; i++ {
		resultado := <-resultado
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, result chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}