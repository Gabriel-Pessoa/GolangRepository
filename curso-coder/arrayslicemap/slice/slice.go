package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// slice é uma parte. Podemos ter um slice referenciando uma parte de um array
	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array. Slice define um pedaço de um array
	s2 := a2[1:3] // s2(slice2) é um pedaço do a2(array2)
	// slice não cria outro array, ele é um pedaço contínuo do array, uma cópia limitada por tamanho especificado
	fmt.Println(a2, s2)

	s3 := a2[:2] // posso ocultar a pos. inicial da cópia do array se for 0
	fmt.Println(a2, s3)

	// podemos imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // slice de slice

	fmt.Println(s2, s4) // apontam para o mesmpo array, a2
}
