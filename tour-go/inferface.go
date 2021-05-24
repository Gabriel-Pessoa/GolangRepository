package main

import (
	"fmt"
	"math"
)

// Interface define o conjunto de métodos
type Abser interface {
	Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	Lat, Long float64
}

// implementa método do tipo ponteiro
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

func main() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}
