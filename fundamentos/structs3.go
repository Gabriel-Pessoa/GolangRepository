package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // tipo Vetex
	v2 = Vertex{X: 1}  // Y: 0 é implícito
	v3 = Vertex{}      // X:0 e Y:0
	p  = &Vertex{1, 2} // constroi um ponteiro para uma struct literal
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
