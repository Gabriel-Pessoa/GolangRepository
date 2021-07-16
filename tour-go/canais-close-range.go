package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Apenas o remetente deve fechar um canal, nunca o receptor. O envio em um canal fechado irá causar um pânico.
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		//v, ok := <-ch //ok é false se não há mais valores a receber e o canal está fechado.
		fmt.Println(i)
	}
}
