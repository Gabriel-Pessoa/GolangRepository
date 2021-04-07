package main

import "fmt"

//Observando o pattern: as propriedades são definidas em letra Maiúscula
type Vertex struct {
	X int
	Y int
}

func main() {
	// var v = Vertex(type Vertex)
	v := Vertex{1, 2}
	v.X = 4 // alterando valor da propriedade "X" no objeto "V"
	fmt.Println(v.X)
}
