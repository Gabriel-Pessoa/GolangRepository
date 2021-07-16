package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)

	// enviando o valor 1 para o canal
	c <- 1 // operação bloqueante, ou seja, bloqueia o código até "alguém" do outro lado receba esse dado (leia)
	fmt.Println("Só depois que for lido")
}

func main() {
	c := make(chan int) // canal sem buffer. Buffer: delimita a quant. de dados que não geram bloqueio no canal. Ou seja, delimita a quant. de dados que irão acumular -sem interromper o tráfego- antes de gerar
	// o bloqueio chegando ao limite definido no buffer, que só é liberado mediante a leitura das informações.

	go rotina(c) // função marcada com go de goroutine, forma de trabalha com código assíncrono em go

	fmt.Println(<-c) // operação também bloqueante: channel forma de gerenciar goroutines em go, ou seja, definir um síncronismo no código. Nesse caso, trava o código aguardando a leitura do valor
	fmt.Println("Foi lido")

	fmt.Println(<-c)   // deadlock
	fmt.Println("Fim") // nunca será executado por conta do deadlock acima

}
