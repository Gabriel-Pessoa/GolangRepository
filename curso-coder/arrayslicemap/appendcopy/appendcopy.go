package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	fmt.Println(array1, len(array1), cap(array1))

	var slice1 []int

	// array1 = append(array1, 1, 4, 5, 6) erro: array é uma estrutura de tamanho estático
	slice1 = append(slice1, 4, 5, 6) // slice1 é uma estrutura de tamanho dinâmica

	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	copy(slice2, slice1) // copy faz a cópia de slice não de array

	fmt.Println(slice2)
}
