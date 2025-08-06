package main

func main(){
	fmt.Println(somar(2,2))

	var f = func(txt string){
		fmt.Prinln(txt)
	}

	f("Texto da função 1")

	resultadoSoma, resultadoSUbtração := calculosMatematicos(10,20)

	resultadoSoma2, _ := calculosMatematicos(5, 10)
}

fun calculosMatematicos(n1,n2 int8) (int8, int8){
	soma:=n1+n2
	subtracao := n1 - n2
	return soma, subtracao
}

func somar(n1 int, n2 int){
	return n1 + n2
}