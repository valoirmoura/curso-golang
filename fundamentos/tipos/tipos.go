package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numericos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal Inteiro é", reflect.TypeOf(320000))

	//Somente valores positivos (uint8(equivalente byte), uint16(equivalente ao short), uint32(equivalente ao Int), uint64(equivalente ao Long) o "u" sigifinica "unsign" ou seja sem sinal
	var b byte = 255 //O byte é apenas um alias para uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	//valores com sinal, ou seja aceita positivo e negativo (uint8, uint16, uint32, uint64)
	i1 := math.MaxInt64
	fmt.Println("O Valor de máximo de i1 é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	//representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O Valor de i2 é", i2)

	//Números Reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O Tipo da Variável X é", reflect.TypeOf(x))      //O tipo aqui é float32 pois foi informado explicitamente na declaração da variável
	fmt.Println("O Tipo do Valor 49.99 é", reflect.TypeOf(49.99)) //O tipo aqui é float64, ele foi inserido implicitamente, então o GO utiliza a arquitetura da maquina para a declaração.

	//boolean
	bo := true
	fmt.Println("O tipo da variável bo é", reflect.TypeOf(bo))
	fmt.Println("O valor da variável", !bo)

	//String : delimititado por aspas duplas
	s1 := "Olá meu Nome é Valoir"
	fmt.Println(s1 + " é um Lindo!!")
	fmt.Println("O tamanho da string é", len(s1))

	//String com múltiplas linhas
	s2 := `
		Olá
		Meu Nome é
		Valoir Leite Moura!!!
	`
	fmt.Println(s2)
	fmt.Println(len(s2))
}
