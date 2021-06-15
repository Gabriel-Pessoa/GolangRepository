package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p1
	primeiro = p2
	return // retorno limpo: como já decidi a ordem no tipo da saída esperada na função,
	// eu não preciso retornar esses valores explicitamente com return
}

func main() {
	// não sou obrigado a invocar a função com os mesmo nomes do retorno nomeado,
	// até pq eu só poderia utilizar esses retornos apenas uma vez dentro de um escopo específico
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
