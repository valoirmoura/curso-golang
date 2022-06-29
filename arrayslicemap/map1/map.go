package main

import "fmt"

func main() {
	//Map devem ser inicializados
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[321654987] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 123456789)
	fmt.Println(aprovados)
}
