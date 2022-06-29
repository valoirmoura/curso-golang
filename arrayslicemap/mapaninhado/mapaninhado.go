package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela silva": 15456.78,
			"Guga Pereira":   84567.78,
		},
		"J": {
			"Jose João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
