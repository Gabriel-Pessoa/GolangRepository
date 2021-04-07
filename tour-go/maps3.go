package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m)

	m["Answer"] = 48
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m)

	v, ok := m["Answer"] //
	fmt.Println("O valor:", v, "Presente?", ok)
}
