// Go tem ponteiros. Um ponteiro guarda na memória o endereço de uma variável
package main

import "fmt"

func main() {
	var p *int // definindo um ponteiro

	i := 42
	p = &i // operador "&" faz "p" apontar para "i"

	// sem o operador "*" não consegue-se acessar valor do ponteiro
	fmt.Println(*p) // Lê i através do ponteiro p
	*p = 21         // definindo i através do seu ponteiro p (desreferenciamento ou indirecionamento)
	fmt.Println(i)
}
