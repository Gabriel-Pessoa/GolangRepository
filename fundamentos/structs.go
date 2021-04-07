package main

import "fmt"

// Coleção de campos (objeto)
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
