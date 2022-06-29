package main

import "fmt"

func main() {
	i := 1

	//GO não tem aritmética de ponteiros
	var p *int = nil

	p = &i //atribuindo o endereço de memória da variável i para p
	*p++
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(p)
}
