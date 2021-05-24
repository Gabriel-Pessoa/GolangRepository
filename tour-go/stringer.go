package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v anos)", p.Name, p.Age)
}

func main() {
	a := Person{"Artuhr Dent", 42}
	z := Person{"Zaphod Bumble", 9001}

	fmt.Println(a, z)
}
