/*
Há duas razões para usar um receptor de ponteiro.

O primeiro é para que o método possa modificar o valor que os seus pontos de receptor referenciam.

A segunda é para evitar copiar o valor de cada chamada de método. Isto pode ser mais eficaz se o receptor for uma estrutura grande, por exemplo.
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

func Scale(v *Vertex, f float64) {
	v.Lat = v.Lat * f
	v.Long = v.Long * f
}

func SumFunc(v *Vertex) float64 {
	return v.Lat + v.Long
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))

	p := &Vertex{10, 20}
	fmt.Println(SumFunc(p)) // p é interpretado como (*p)
}
