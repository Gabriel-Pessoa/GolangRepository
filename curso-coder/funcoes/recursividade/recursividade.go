package main

import (
	"fmt"
)

func fatorial(num int) (int, error) {
	switch {
	case num < 0:
		return -1, fmt.Errorf("valor inválido: %d", num)
	case num == 1 || num == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(num - 1)
		return num * fatorialAnterior, nil
	}
}

func fatorial2(num uint) uint {
	if num == 1 || num == 0 {
		return 1
	}

	return num * fatorial2(num-1)
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	resultado2 := fatorial2(5)
	fmt.Println(resultado2)
}
