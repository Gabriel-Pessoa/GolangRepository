package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) // type assertation retorna dois valores: valor, boolean. Se i detém string, então string será o valor e boolean serã true
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //panic
	fmt.Println(f)
}
