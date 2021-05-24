package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // pega o nano segundo da atual e passa como fonte para gerar um número aleatório
	r := rand.New(s)                           // a partir da fonte configurada na linha acima, sempre irá gerar um número novo
	return r.Intn(10)                          // gera um número
}

func main() {
	// variável i fica acessível apenas dentro desse bloco
	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou!")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu!")
		fmt.Println(i)
	}

	// fmt.Println(i) i não fica acessível fora do escopo da condição

}
