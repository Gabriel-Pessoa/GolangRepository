package main

import "fmt"

func fibonacci() func(int) int {
	fibNMinus2 := 0
	fibNMinus1 := 1
	var fibN int

	return func(x int) int {
		if x < 1 {
			return 0
		}
		if x <= 2 {
			return 1
		}
		fibN = x
		for i := 2; i <= x; i++ { // n >= 2
			fibN = fibNMinus1 + fibNMinus2 // f(n-1) + f(n-2)
			fibNMinus2 = fibNMinus1
			fibNMinus1 = fibN
		}
		return fibN
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(fibonacci()(i))
	}
}
