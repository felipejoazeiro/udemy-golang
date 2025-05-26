package main
import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool{
	media := (n1 + n2)/2

	if media >6{
		return true 
	}else if media <6{
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main(){
	alunoEstaAprovado(5,7)
}