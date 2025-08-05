package main

type cachorro struct {
	Nome 	string 	`json:"nome"`
	Raca 	string 	`json:"raca"`
	Idade 	uint 	`json:"idade"`
}

func main() {
	cachorroEmJson := `{"name": "Rex","raca": "Dálmata","cursos": 3}`

	c := cachorro{}
	
	fmt.Println(c)

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJson := `{"nome": "Toby", "raca": "Poodle", "idade": 2}`

	c2 := make(map[string] string)

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}