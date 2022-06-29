package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3} // array possuí tamanho fixo em GO, não sendo possível aumentar o seu tamanho.
	var slice1 []int          // slice é um array dinâmico, podendo ser criado com ou sem tamanho definido.

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // podemos inicializar um array com a função make, ou seja no exemplo o slice será do tipo int e terá tamanho 2 inicial, como o valor inicial do tipo int é 0 ambos os valores serão 0.
	copy(slice2, slice1)
	fmt.Println(slice2)

}
