package main

func main(){
	soma := 1+2
	subtracao := 1-2
	divisao := 10/4
	multiplicacao := 10*5
	restoDivisao := 10%2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	soma:=numero1 + numero2 // Não funciona devido tipos de dados diferentes

	// FIm dos aritméticos

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"

	//Fim dos Operadores de Atribuição

	//Operadores Relacionais
	fmt.Println(1>2) //false
	fmt.Println(1>=2) //false
	fmt.Println(1==2) //false
	fmt.Println(1<=2)
	fmt.Println(1>2)
	fmt.Println(1<2)
	fmt.Println(1!=2)

	//FIM DOS RELACIONAIS

	//Operadores Lógicos
	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Fim dos Operadores Lógicos

	//Operadores Uniários
	numero:=10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	//FIM DOS Operadores Uniários

	var texto string
	if numero > 5{ 
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	fmt.Println(texto)
}