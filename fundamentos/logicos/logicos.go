package main

import "fmt"

/**
 * Funções em Golang possibilitam o multiplos retornos
 * Na função compras temos dois parametros do tipo bool e tres retornos do tipo bool, isso faz com que condicionais tipo if não sejam necessárias
 */
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 //ou Exlusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
