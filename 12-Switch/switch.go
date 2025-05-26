package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
		case 1:
			return "Domingo"
		case 2:
			return "Segunda-Feira"
		case 3:
			return "Terça-Feira"
		case 4:
			return "Quarta-Feira"
		default:
			return "Número Inválido"
	}
}

func diaDaSemana2(numero int)string{
	switch{
	case numero == 1:
		return "Domingo"
	default:
		return "Número Inválido"
	}
}

func main(){
	fmtPrintln("Switch")
	dia:= diaDaSemana(1)
	fmt.Println(dia)

	dia2:=diaDaSemana2(1)
	fmt.Println(dia2)
}