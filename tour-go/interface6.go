package main

import "fmt"

func main() {
	// interface que especifíca zero métodos
	// usada para valores do tipo desconhecido
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
