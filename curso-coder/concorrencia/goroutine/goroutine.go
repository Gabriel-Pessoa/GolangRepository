package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	/* Código síncrono, aguarda a execução para depois executar outro*/
	//fale("Maria", "Por que não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc!", 1)

	//thread: fio ou linha, pedaço do processo capaz de executar determinadas tarefas como se fosse um subsistema, subconjunto.
	// Podemos definir thread do contexto de microprocessadores: capacidade de encadeamento de instruções que podem ser executadas simultaneamente

	/* Código assíncrono, executa os dois simultaneamente, porém a função acaba antes da primeira execução*/
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	/*Código assíncrono, que executa durante as repetições da segunda função*/
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns", 5)
}
