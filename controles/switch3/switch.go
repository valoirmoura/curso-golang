package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "string"
	case func():
		return "Função"
	default:
		return "Não Sei"
	}
}

func main() {
	fmt.Println(tipo(2))
	fmt.Println(tipo(2.4))
	fmt.Println(tipo("Valoir Moura"))
	fmt.Println(tipo(func() {}))

}
