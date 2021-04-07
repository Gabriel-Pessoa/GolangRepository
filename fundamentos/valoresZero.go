package main

import "fmt"

// São valores iniciais para as variáveis sem atribuições, dependendo do seu tipo

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
