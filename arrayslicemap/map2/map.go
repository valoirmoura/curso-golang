package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1340.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
