package main

import (
	"fmt"
)

func calculos(n1,n2 float64) (float64,float64){
	soma:=n1+n2
	subtracao:=n1+n2
	return soma,subtracao
}

func  main(){
	soma, sub := calculos(10,20)
	fmt.Println("Hello GO")
	fmt.Println(soma,sub)
	fmt.Printf("A soma foi de %v e a sub foi de %.2f ",soma,sub)

	

}