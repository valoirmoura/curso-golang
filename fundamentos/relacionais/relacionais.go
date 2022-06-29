package main

import (
	"fmt"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana") //Operador de Igualdade ==
	fmt.Println("!=", 3 != 2)                     //Operador de Diferença !=
	fmt.Println("<", 3 < 2)                       //Operador de Menor que <
	fmt.Println(">", 3 > 2)                       //Operador de Maior que >
	fmt.Println("<=", 3 <= 2)                     //Operador de Menor ou Igual <=
	fmt.Println("<=", 2 <= 3)                     //Operador de Maior ou Igual >=
	fmt.Println(">=", 3 >= 3)                     //Operador de Maior ou Igual >=
	fmt.Println(">=", 2 >= 3)                     //Operador de Menor ou Igual <=

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2)
}
