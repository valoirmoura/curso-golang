package main

import "fmt"

//Não tem operador Ternario, alternativa é utiliza if

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {

	nota := 5.0

	fmt.Println(obterResultado(nota))

}
