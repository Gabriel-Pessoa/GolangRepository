package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)...unsigned integers: uint8
	/*
			Size/Type	 Range
		1 byte unsigned	 0 to 255
		2 byte unsigned	 0 to 65,535
		4 byte unsigned	 0 to 4,294,967,295
		8 byte unsigned	 0 to 18,446,744,073,709,551,615
	*/
	var b byte = 255 // byte é um alias para unit8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // valor da tabela Unicode (int32). rune é um alias para int32
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2) // valor da tabela unicode

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("Por padrão o tipo de 49.00 é", reflect.TypeOf(49.99)) // float 64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá, mundo!"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é:", len(s1))

	// strings com multiplas linhas
	s2 := `Olá,
	mundo
	!`

	fmt.Println("O tamanho da string é:", len(s2))

	// char?? Go não tem tipo char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char) // valor da tabela unicode
}
