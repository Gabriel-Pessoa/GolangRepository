package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] // define um intevalo para cópia de um array: inclui o primeiro 1 e exclui o último 4 (1 à 3)
	fmt.Println(s)
}
