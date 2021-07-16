package main

import (
	"fmt"
	"time"
)

func rotinas(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou após 3ªposição!")
	ch <- 4
	fmt.Println("Executou após 4ªposição!") // a execução é finalizada antes de chegar nesse trecho, que fica desocupado dps da liberação da primeira posição
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // definindo canal com buffer de até 3 posições
	go rotinas(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch) // lendo o primeiro dado enviado, e também desocupando 1 posição no buffer de 3 posições
}
