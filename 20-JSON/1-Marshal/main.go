package main
type cachorro struct {
	Nome 	string 	`json:"nome"`
	Raca 	string 	`json:"raca"`
	Idade 	uint 	`json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Prinln(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c := map[string]string {
		"name": "Toby",
		"raca": "Poodle"
	}

	cachorro2EmJson, erro := json.Marshal(c2)

	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJson)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))

}