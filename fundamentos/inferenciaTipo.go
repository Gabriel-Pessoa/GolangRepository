package main

import "fmt"

func main() {
	// A linguagem infere o seu tipo pelo valor atribuido
	v := 42
	f := 3.142
	g := 0.867 + 0.5i

	fmt.Printf("v é do tipo %T\n", v)
	fmt.Printf("f é do tipo %T\n", f)
	fmt.Printf("g é do tipo %T\n", g)
}
