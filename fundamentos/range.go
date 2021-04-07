package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// o range do "for" itera sobre uma slice ou map.
	// ao variar sobre uma slice, dois valores são retornados de cada iteração:
	// o primeiro é o índice, o segundo uma cópia do elemento desse índice
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}
