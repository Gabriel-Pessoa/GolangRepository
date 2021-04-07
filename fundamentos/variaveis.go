package main

import "fmt"

// a declaração var pode estar num pacote ou num escopo de função

var c, python, java bool //lista de variáveis, o tipo é o último passado

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
