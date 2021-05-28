package main

import "fmt"

func main() {
	i := 1 // não pode criar uma variável do tipo int recebendo um nil

	var p *int = nil // aloquei um espaço de memória do tipo ponteiro,
	//dentro dele temos um endereço de memória que aponta para uma variável(nesse primeiro momento para lugar nenhum, nil)

	p = &i // pegue o endereço da variável i e atribua ao ponteiro p

	*p++ // desreferenciando, ou seja, pegando o valor associado a esse ponteiro
	i++

	// Go não tem aritmética de ponteiros, ou seja, não tem operações com ponteiros
	// p++ -->  isso não funciona

	fmt.Println(p, *p, i, &i)
	//p: ponteiro, o seu valor literal é o endereço
	//*p: referência do ponteiro, o valor da variável que ele aponta
	//i: variável do tipo inteiro
	//&i: acessando o referência da variável i
}
