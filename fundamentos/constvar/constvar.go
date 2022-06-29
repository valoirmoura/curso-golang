package main

import (
	"fmt"
	"math"
)

func main() {

	/**
	 * Constante em Golang é declarada com a palavra-chave const
	 * Variável em Golang é declarada com a palavra-chave var
	 */

	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64 é inferido pelo GO

	/**
	 * Em Go é possível criar uma variável e já inicializá-la
	 * Toda Variável criada em GO deve ser utilizada, caso contrário o GO irá gerar um Erro de compilação
	 */
	area := PI * math.Pow(raio, 2)
	fmt.Println("Area da Circunferência é", area)

	//Declarando um bloco de constantes
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//Em uma unica linha é possivel declarar varias variáveis e inclusive inicializa-las
	var e, f = true, false //neste exemplo a variavel e é true e a f é false.
	fmt.Println(e, f)

	/**
	 * Não confundir GOLANG é uma linguagem fortemente tipada, nestes casos o tipo é inferido assim que a variável é inicializada.
	 * Valores Atribuidos: g = 2; h = false; i = "epa.
	 */
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)
}
