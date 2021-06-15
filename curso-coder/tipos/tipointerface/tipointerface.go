package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} // interface é o tipo mais genérico, podendo assumir qualquer valor
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang"}
	fmt.Println(coisa2)
}
