package main

import "fmt"

func main() {
	var a int     // Tipo int tem valor zero inicial
	var b float64 // Tipo float tem valor zero inicial
	var c bool    // Tipo bool tem valor false inicial
	var d string  // Tipo string tem valor inicial vazio
	var e *int    // Tipo ponteiro tem valor nulo inicial

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)

}
