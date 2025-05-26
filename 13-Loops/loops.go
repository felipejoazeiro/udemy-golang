package main
func main(){
	i:=0

	for i < 10{
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	for z:= 0; z < 10; z+=2{
		fmt.Println("Incrementando z", z)
		time.Sleep(time.Second)
	}
	fmt.Println("Término")

	nomes := [3]string("João", "Davi", "Lucas")

	for indice, nome := range nomes {
		fmt.Println(indice,nome)
	}

	for nome := range nomes{
		fmt.Println(nome)
	}

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	usuario := map[string]string{
		"nome": "Leonardo",
		"sobrenome": "Silva"
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct{
		nome string,
		sobrenome string,
	}

	usuario2 := usuarioStruct{"Zé", "Júnior"}

	for chave, valor := range usuario2 {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

	
}