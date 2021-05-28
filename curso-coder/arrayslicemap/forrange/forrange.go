package main

import "fmt"

func main() {
	// com o operador spread o compilador infere o tamanho do array
	numeros := [...]int{1, 2, 3, 4, 5} // sem o operador spread a variável se transforma num slice, e não num array.

	// range, semelhante ao foreach
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	for index := range numeros {
		fmt.Println(index)
	}

}
