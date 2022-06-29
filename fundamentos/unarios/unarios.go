package main

import "fmt"

func main() {
	x := 1
	y := 2

	//apenas postFix: GO somente existe pós fixada, préfixada não existe
	x++ //o pré fixada que seria --x não existe
	y++

	fmt.Println(x, y)
}
