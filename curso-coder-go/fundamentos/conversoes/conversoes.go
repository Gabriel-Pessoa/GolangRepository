package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // por padrão é float64
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) // cuidado

	//cuidado...
	fmt.Println("Teste " + string(123)) // cuidado, converte para o elemento da tabela ASCII

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 123)

	b, _ := strconv.ParseBool("1") // considera verdadeiro apenas true e 1
	if b {
		fmt.Println("Verdadeiro")
	}

}
