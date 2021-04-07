package main

/*
Um closure é uma função valor que referencia variáveis de fora de seu corpo.
A função pode acessar e atribuir nas variáveis referenciadas; nesse sentido a função é "limitada" às variáveis.
*/
import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() // variáveis recebem como atribuição outra função que recebe x como parâmetro, e ambos referenciam sum
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
