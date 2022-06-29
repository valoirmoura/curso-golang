package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // cria um array com 5 posições, porém, é o compilador que contará

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { //Ignorando o Índice
		fmt.Println(num)
	}
}
