package main

import "fmt"

// sum recebe dois parâmetros do tipo int. Observe que o tipo vem após o nome da variável (a int, b int).

// func sum(a int, b int) int {
// 	return a + b
// }

//quando dois ou mais parêmtro compartilham o mesmo tipo, podemos omitir o tipo de todos menos o último
func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(4, 5))
}
