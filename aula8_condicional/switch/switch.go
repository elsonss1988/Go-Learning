package main

import (
	"fmt"
)

func diaDaSemana(numero int) string{
	switch numero{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terca"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "invalido"
	}
}

func weekDay(numero int) string{
	var day string
	switch {
	case numero == 1:
		day= "Domingo"
	case numero == 2:
		day= "Segunda"
	case numero == 3:
		day= "Terca"
	case numero == 4:
		day= "Quarta"
	case numero == 5:
		day= "Quinta"
	case numero == 6:
		day= "Sexta"
	case numero == 7:
		day= "Sabado"
		fallthrough
	default:
		day = day + " Mande outro dia entre 1 a 7"
	}
	return day
}


func main(){
	fmt.Println()
	dia := diaDaSemana(200)
	fmt.Println(dia)
	day:= weekDay(7)
	fmt.Println(day)
}