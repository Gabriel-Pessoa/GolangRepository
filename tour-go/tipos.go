package main

import (
	"fmt"
	"math/cmplx"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // pseudônimo para uint8

rune // pseudônimo para int32
     // representa um ponto de código Unicode

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Tipo: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo: %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo: %T Valor: %v\n", z, z)
}
