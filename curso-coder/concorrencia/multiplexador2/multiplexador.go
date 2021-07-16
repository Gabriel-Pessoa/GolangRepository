package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()

	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string { // recebo duas entrada (origens) e transformo em um
	c := make(chan string)
	go func() {
		// for infinito que irá executar enquanto a goroutine estiver ativa
		for {
			// multiplexando para o mesmo canal
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c // retorna canal multiplexado
}

func main() {
	c := juntar(falar("João"), falar("Maria"))

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
