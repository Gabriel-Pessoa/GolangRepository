// Go tem uma única estrutura de laço, o for
// a expressão não precisa ser cercada de ( ) mas os chaves { } são obrigatórios.
package main

import "fmt"

func main() {
	sum := 0

	// variável i, visivel apenas no escopo de declaração for
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	sum = 1

	// O for é tbm o while do Go
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

}
