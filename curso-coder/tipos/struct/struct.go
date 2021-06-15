package main

import "fmt"

type produto struct { // é um agrupador de dados, funciona mais ou menos como atributos de um classe
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor), para que eu possa realizar ações como: p.precoComDesconto()
// reparar que o receiver é um valor, uma cópia; não uma referência
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())
	// não teve alteração no preço inicial do produto 1, porque o método recebe uma cópia do produto 1 faz a operação
	// e retornar um novo valor, sem alterar o original
	fmt.Println(produto1.preco)

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
}
