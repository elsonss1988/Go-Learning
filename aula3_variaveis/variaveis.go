package main

import (
	"errors"
	"fmt"
)

func main(){
	 var greeting string = "hello GO"
	 helloBR := "Ola Mundo"

	 fmt.Println(greeting)
	 fmt.Println(helloBR)

	 var(
		nome string = "Elson"
		sobrenome string = "West"
   )
   
   fmt.Println(nome,sobrenome)

   pais,estado := "Vera Cruz", "Sao Pedro"

   fmt.Println(pais,estado)

   const nascimento string = "Distrito 101"

   fmt.Println(nascimento)

   pais, estado = estado, pais

   fmt.Println(
	"Informacao trocada "+pais+"outra info trocada "+estado)

   var erro error
   fmt.Println(erro)

    erro = errors.New("Erro interno")
   fmt.Println(erro)
}


