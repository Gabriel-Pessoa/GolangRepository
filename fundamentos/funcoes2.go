package main

import "fmt"

// O retorno de funções podem ser dois ou mais tipos
func swap(x, y string) (string, string) {
	return x, y
}

// parêmtros retornados nomeados
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// atribuindo retorno da função para variáveis
	a, b := swap("Gabriel", "Pessoa")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
