package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array : possui tamanho fixo
	s1 := []int{1, 2, 3}  //slice: possui tamanho variável

	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice não é um array! Slice define um pedaço de um array.
	s2 := a2[1:3]       //pega do indice 1 até o indice 3 porem não inclui o indice 3
	fmt.Println(a2, s2) //Obs. O Slice não criou um novo array ele apontou para os indice do array a2

	s3 := a2[:2] //vai do ínicio do Array a2 até o índice 2 não incluido o índice 2
	fmt.Println(a2, s3)

	//Slice possui tamanho e um ponteiro.
	s4 := s2[:1] //pega do ínicio do array até o indice 1 e não pegando o indice 1
	fmt.Println(s4)

}
