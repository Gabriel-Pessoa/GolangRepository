package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // seção a de matriz subjacente names
	b := names[1:3] // seção b de matriz subjacente names
	fmt.Println(a, b)

	// uma alteração numa seção que compartilha a mesma matriz subjacente, reflete nas outras seções e na própria matriz subjacente
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
