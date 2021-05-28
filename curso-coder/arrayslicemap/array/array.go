package main

import "fmt"

func main() {
	// array são estruturas homogênas (mesmo tipo) e estático (tamanho fixo)
	var notas [3]float64
	fmt.Println(notas)

	//estrutura indexada, acesso ao elemento através do seu índice
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	// notas[3] = 10 // esse índice está fora do range do array

	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f", media)
}
