package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		// select: Aguarda até que um de seus cases possam executar, então ele executa esse case.
		// Permite uma espera na goroutine sobre as operações de comunicação múltiplas.
		// Escolhe um ao acaso se vários estiverem prontos.
		select {
		case c <- x:
			time.Sleep(time.Second)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return // importante para finalizar o loop
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
