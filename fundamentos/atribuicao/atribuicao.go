package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inferência de Tipo
	i += 3 //i = i + 3
	i -= 3 //i = i - 3
	i /= 3 // i = i / 3
	i *= 3 //i = i * 3
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2 //Variavel x = 1, Variável y = 2
	fmt.Println(x, y)
}
