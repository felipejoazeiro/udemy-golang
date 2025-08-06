package main
import "fmt"

func funcao1(){
	fmt.Prinln("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 floar32) bool {
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2)/2

	if media >= 6{
		fmt.Println("Média calculada. Resultado será retornado")
		return true
	}
	fmt.Println("Média calculada. Resultado será retornado")
	return false
}

func main(){
	defer funcao1()
	//DEFER  = ADIAR

	funcao2()

	fmt.Println(alunoEstaAprovado(7,8))
}