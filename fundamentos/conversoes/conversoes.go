package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) //converter o tipo para float64, realizamos um parse

	nota := 6.9
	notaFinal := int(nota) //a conversão sempre é realizada para baixo
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste " + string(97)) //converte o int para string, o valor 97 é o valor da tabela ASCII

	//inteiro para String
	fmt.Println("Teste " + strconv.Itoa(123)) //converte para String

	//string para int
	num, _ := strconv.Atoi("123") //Em GO é possível retornar dois valores, o primeiro é o valor convertido e o segundo é um erro, atenção ao _, em GO não é possivel declarar uma variável e não utiliza-la, com o _ é possivel ignorar esta variável.
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") //converte para bool
	if b {
		fmt.Println("Verdadeiro")
	}

	//Converter Inteiro para Float é necessário realizar o Cast ou seja float64(numero inteiro)
}
