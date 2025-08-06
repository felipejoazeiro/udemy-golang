package main

import "fmt"

func main(){
	usuario := map[int]string{
		1: "Pedro",
		2: "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": 	"Pedro",
		}
		"curso": {
			"nome": 	"Engenharia",
			"campus": 	"Campus 1",
		}
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
}