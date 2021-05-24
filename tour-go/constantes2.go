package main

import "fmt"

const (
	//crie um número enorme deslocando 1 bit para a esquerda 100 casas
	// O número 0 seguido por cem zeros
	Big   = 1 << 100
	Small = Big >> 99
)

func retornaInt(x int) int { return x*10 + 1 }

func retornaFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(retornaInt(Small))
	fmt.Println(retornaFloat(Small))
	fmt.Println(retornaFloat(Big))
}
