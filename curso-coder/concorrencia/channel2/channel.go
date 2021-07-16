package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre as goroutines
// O channel é tipado

// Segundo parâmetro é um canal que trafega dados inteiros
func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second) // espera 1 segundo
	c <- 2 * base           // enviando dados para esse canal. O código aqui para até que a informação seja lida pelo o canal

	time.Sleep(time.Second) // espera 1 segundo
	c <- 3 * base

	time.Sleep(3 * time.Second) // espera 3 segundos
	c <- 4 * base

}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	fmt.Println("Executando 1...")

	a, b := <-c, <-c // os dois primeiros valores do canal são atribuídos. Recebendo dados do canal. Só sairá dessa linha quando os dois valores chegarem.
	fmt.Println(a, b)
	fmt.Println("Executando 2...")

	fmt.Println(<-c) // o último dos três valores recebidos pelo canal
	fmt.Println("Executando 3...")

	//fmt.Println(<-c) // deadlock, não existe mais nehuma goroutine
}
