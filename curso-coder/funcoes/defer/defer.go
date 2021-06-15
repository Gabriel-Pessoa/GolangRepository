// O principal uso do defer é a limpeza de recursos, como: arquivos abertos, conexões de rede e conexões de banco de dados.
// O defer faz com que nosso programa fique mais limpo e menos suscetível a erros, mantendo as chamadas para fechar
// o arquivo/recurso próximas da chamada aberta.

package main

import "fmt"

func obterValorAprovado(numero int) int {
	// defer: adia a execução de uma função até o momento antes do retorno da função que está inserido.
	// Os argumentos da chamada adiada são avaliados imediatamente, mas a chamada da função não é executada até que a função envolvente retorne.
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto!")
		return 5000
	}

	fmt.Println("Valor baixo!")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
