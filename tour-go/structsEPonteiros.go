package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v   // definindo ponteiro por inferência; referenciando Vertex (obj)
	p.X = 1e9 // sintax sugar com ponteiro: sem o operador "*" para alterar valores da variável referenciada

	fmt.Println(v)
}
