package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var e1 = errors.New("Error 1")

func x() error {
	// fmt.Errorf: formata de acordo com um especificador de formato e retorna uma string como VALOR que satisfaz o erro
	return fmt.Errorf("adding more context: %w", e1) // o erro retornado com o verbo "%w" implementa o método unwrap
}

func main() {
	e := x()

	if errors.Is(e, e1) { //Se erro "e" envolver erro "e1" retorne true, senão false.
		fmt.Println("WORKS")
		fmt.Println(x())
	}

	fmt.Println(errors.Unwrap(e1))

	test := errors.Unwrap(fmt.Errorf("%w", e1)) // retorna o erro
	fmt.Println(test)

	fmt.Printf("1º parâmetro: %v, 2º parâmetro: %v\n", e, e1)

	if _, err := os.Open("non-existing"); err != nil { // o variável "err" é do tipo "*PathError"
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Falha no caminho:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	_, err := os.Open("non-existing")

	// ok retorna true se a variável "err" é definida pela interface (*fs.PathError), caso não retorna false.
	if e, ok := err.(*fs.PathError); ok {
		fmt.Println(e.Path)
	}
}
