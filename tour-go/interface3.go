package main

import (
	"fmt"
)

type Calc interface {
	Sum10() int
}

type Numbers struct {
	num1, num2 int
}

func (n *Numbers) Sum10() int {
	return n.num1 + n.num2 + 10
}

type F int

func (f F) Sum10() int {
	result := int(f + 10)
	fmt.Println(result)
	return result
}

func main() {
	var c Calc

	/*Um valor de uma interface contém um valor de um tipo concreto subjacente específico.
	Chamando um método em um valor de interface executa o método do mesmo nome no seu tipo subjacente.*/
	c = &Numbers{3, 2}
	fmt.Println(c.Sum10())
	describe(c)

	c = F(10)
	describe(c)
	c.Sum10()

	var empty *Numbers
	c = empty
	describe(c)
	c.Sum10()
}

func describe(c Calc) {
	fmt.Printf("(%v, %T) \n", c, c)
}
