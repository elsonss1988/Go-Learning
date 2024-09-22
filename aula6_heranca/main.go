package main

import (
	"fmt"
)

type people struct {
	nome  string
	idade int
}

type employer struct {
	people
	titulo  string
	salario int
}

func main() {
	cliente1 := employer{
		people: people{
			nome:  "João",
			idade: 27,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	cliente2 := people{
		nome:  "Joana",
		idade: 33,
	}

	fmt.Printf("Hello World \n\n")
	fmt.Println(cliente1.people.nome)
	fmt.Println(cliente1.nome)
	fmt.Println(cliente1.people.idade)
	fmt.Println(cliente1.idade)
	fmt.Println(cliente1.titulo)
	fmt.Println(cliente1.salario)
	fmt.Println(cliente2.nome)
	fmt.Println(cliente2.idade)

	pessoa := people{"Mauricio", 40}
	cliente3 := employer{people{"Vanderlei", 50}, "Politico", 10000}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
	fmt.Println(cliente3)
	fmt.Println(pessoa)
}