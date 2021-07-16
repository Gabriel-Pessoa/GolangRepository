// O código do exemplo soma os números em uma slice, distribuindo o trabalho entre duas goroutines.

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // envia sum para c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // como maps e slices, os canais devem ser criados antes de se usar

	// Por padrão, enviam e recebem bloco até o outro lado estar pronto.
	// Isso permite que goroutines sincronizem sem bloqueios explícitos ou variáveis ​​de condição.
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// Uma vez que ambas as goroutines completaram seu processamento, calcula o resultado final.

	x, y := <-c, <-c // recebe de c

	fmt.Println(x, y, x+y)
}
