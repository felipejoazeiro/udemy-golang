package main 

func main() {
	canal := multiplexar(escrever("Olá mundo!"), escrever("Programndo em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}


func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- string {
	canalDeSaida := nake(chan string)
	go func(){
		for {
			select {
			case mensagem := <-canalsDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintg("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}