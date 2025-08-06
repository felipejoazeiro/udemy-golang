package main

type usuario struct {
	nome string
	idade uint18
}

func (u usuario) salvar(){
	fmt.Println("Salvando os dados do Usuário %s no banco de dados", u.nome)
}

func (u usuario) maiorDeIdade() bool{
	return u.idade >= 18
}

func (u *usuario) fazerAniversario(){
	u.idade++
}

func main(){
	escrever()
}

func escrever() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.maiorDeIdade()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}