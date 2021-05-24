package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Duas vezes %v é %v\n", v, v*2)
	case string:
		fmt.Printf("%q tem %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Eu não sei sobre esse tipo %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
