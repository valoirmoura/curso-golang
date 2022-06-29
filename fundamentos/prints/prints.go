package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha!")

	fmt.Println(" Nova")
	fmt.Println("Linha!")

	x := 3.141516

	//Neste caso foi convertido o valor da variavel x para String assim é possivel utilizar o operador de concatenhação "+"
	xs := fmt.Sprint(x)
	fmt.Println("O valor de X é" + xs)

	//Em GO, não se faz necessário utilizar o operador de concatenação, basta colocar uma vírgula e a variável!
	fmt.Println("O valor de X em GOLANG é", x)

	/** É possivel trabalhar com dados formatados em GO
	 * .2f significa que queremos utilizar duas casas decimais.
	 */
	fmt.Printf("O Valor de X formato em Golang é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	/**
	 * "\n" é para quebrar linha
	 * "%d" é para valores inteiros
	 * "%f" é para valores flutuantes
	 * "%t" é para valores booleanos
	 * "%s" é para valores Strings
	 * "%v" é um genérico que serve para qualquer tipo
	 */
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
