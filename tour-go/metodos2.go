// métodos sem struct

package main

import (
	"fmt"
	"math"
)

type Myfloat float64

// segue o mesmo princípio, método é uma função.
func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	fmt.Println(float64(-math.Sqrt2))
}
