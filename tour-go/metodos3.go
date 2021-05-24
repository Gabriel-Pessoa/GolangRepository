// ponteiro receptor
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

// com ponteiro. É mais comum usar ponteiro
func (v *Vertex) Scale(f float64) {
	v.Lat = v.Lat * f
	v.Long = v.Long * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
