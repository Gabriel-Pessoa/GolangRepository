package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

// função exec recebe como 1ºparâmetro: uma função com dois parâmetro do tipo int que retorna um inteiro.
// função exec recebe como 2º e 3º parâmetro: p1 e p2 do tipo int
// função exec retorna um int
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
