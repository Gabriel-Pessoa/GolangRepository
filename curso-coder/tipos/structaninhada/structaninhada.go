package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			// o uso da atribuição curta em sequência, não é mt indico, pode gerar confusão do que cada propriedade representa
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			// o uso da atribuição com as propriedades explícitas é mais indicado
			item{produtoID: 11, qtde: 100, preco: 3.13},
		},
	}
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
