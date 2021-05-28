package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) // não pode concatenar string com outro tipo

	xs := fmt.Sprint(x) // convertendo em string
	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x) // concatena e converte internamente

	fmt.Printf("O valor de x é %f\n", x)   // print formatado do tipo float: %f
	fmt.Printf("O valor de x é %.2f\n", x) // print formatado do tipo float com apenas duas casas decimais

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("%d %f %.1f %t %s", a, b, b, c, d) // %d: inteiro; %f: float; %t: bool; %s: string
	fmt.Printf("%v %v %v %v", a, b, c, d)         // %v: valor genérico
}
