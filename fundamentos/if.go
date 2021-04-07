package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// a expressão não precisa ser cercada de ( ) mas os chaves { } são obrigatórios.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//posso declarar uma variável acessada somente no escopo do if (v, parecido com o "i" do for)
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
