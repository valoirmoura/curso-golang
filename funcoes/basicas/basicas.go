package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

//passamos dois parametros, inserimos o nome do parametro seguido do seu tipo
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

//passamos dois parâmetros e apenas um tipo, isso significa que ambos os parâmetros serão do mesmo tipo
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {

	f1()
	f2("Parametro 1", "Parametro 2")

	r3, r4 := f3(), f4("Parametro 1", "Parametro 2")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	//_, r52 = f5() //Neste Caso é ignorado o primeiro retorno
	fmt.Println(r51, r52)

}
