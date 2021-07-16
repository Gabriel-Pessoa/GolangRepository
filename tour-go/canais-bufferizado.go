package main

import "fmt"

func main() {
	ch := make(chan int, 2) //Fornecendo o tamanho do buffer como o segundo argumento para make para inicializar um canal bufferizado
	ch <- 1
	ch <- 2
	//ch <- 3 ERRO! Canal foi bufferizado apenas para dois "buffers/entradas"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) // Deadlock, não existe mais nenhuma saída além das duas acessadas anteriormente
}
