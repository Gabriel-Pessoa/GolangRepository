package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

//método implementa interface I sem palavra reservada "implements"
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"} // semelhante ao operador new class e o parâmetro dentro dos parênteses
	i.M()
}
