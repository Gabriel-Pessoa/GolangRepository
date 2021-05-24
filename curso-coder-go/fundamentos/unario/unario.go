package main

import "fmt"

func main() {
	x := 1
	y := 2

	// go só suporte operadores pós fixados
	x++
	y--
	fmt.Println(x)
	fmt.Println(y)

	//fmt.Println(x == y--) isso não é permitido, nem recomendado em outras linguagens
}
