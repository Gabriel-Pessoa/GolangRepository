package main

import "fmt"

func main() {
	i, j := 42, 2701 // definindo variáveis por inferência

	p := &i         // definindo ponteiro por inferência, referenciando variável 'i'
	fmt.Println(*p) // acessando valor do ponteiro
	fmt.Println(i)  // acessando variáveis para qual o ponteiro faz referência

	p = &j       // mundando referência do ponteiro
	*p = *p / 37 // mudando valor da variável referenciada pelo ponteiro (desreferenciamento ou indirecionamento)

	fmt.Println(j)
}
