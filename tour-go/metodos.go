// go não tem classe
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

// criando método
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

/*
o método é apenas uma função com um argumento recebido
func Abs(v Vertex) float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}
*/

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
