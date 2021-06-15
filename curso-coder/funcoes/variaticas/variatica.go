package main

import "fmt"

// função variáticas: recebe de 1 a n parâmetros
func media(numero ...float64) float64 {
	var soma float64
	for _, valor := range numero {
		soma += valor
	}
	qtd := len(numero)
	return soma / float64(qtd)
}

func main() {
	resultado := media(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 15, 27, 35, 356)
	fmt.Printf("Média: %.2f", resultado)
}
