/*
O Channel, ou canal, é comumente usado com goroutine. Ele (o canal) é um mecanismo de sincronismo,
 ele espera os dados chegarem de funções que estão rodando de forma assincrona - função marcada com goroutine.
 O canal para o código até o valor chegar.

 OBS: CUIDADO!!! Para não cair num cenário de DEADLOCK: o canal trava o código aguardando o valor que nunca virá. A goroutine já acabou a sua execução,
 mas o canal está no fluxo principal aguardando um valor de uma goroutine(função marcada com go) que já finalizou a sua execução
*/

package main

import "fmt"

func main() {
	// Criamos um canal de forma parecida com um slice. Canal não é uma estrutura à parte da linguagem. Ele é um tipo: assim como um int é um tipo, o float.
	// Dentro dele pode trafegar dados de um determinado tipo

	// Dentro desse canal de comunicação irá trafegar (envio e recebimento) números inteiros.
	// Segundo parâmetro é o buffer
	ch := make(chan int, 1)

	// Como enviar dados para um canal
	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2 // enviando novo valor para o canal (escrita)
	fmt.Println(<-ch)
}
