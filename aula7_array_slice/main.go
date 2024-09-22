package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello");

	var array1 [5]string
	array1[0]= "posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1","Posicao 2","Posicao 3","Posicao 4","Posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	slice := []int{10,11,12,13,14,15,16,17}
	fmt.Println(slice)

	slice = append(slice, 21)

	slice2 :=array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao Alterada"
	fmt.Println(slice2)

	slice3 := make([]float32,10,15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	usuario := map[string]string{
		"email":"Pedro",
		"sobrenome":"Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro":"Joao",
			"ultimo":"Pedro",
		},
		"curso":{
			"nome":"Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2,"nome")
	fmt.Println(usuario2)

	usuario2["signo"]=map[string]string{
		"nome":"Gemeos",
	}

	fmt.Println(usuario2)

}












































