package main

import "fmt"

/* Algumas regras para o uso de constantes:
1) Declarada com a palavra reserva "const"
2) Não podem ter sequências de caracteres, booleanos, valores númericos
3) Não podem ser declaradas com a sintaxe ":="
*/
const PI = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello, World")
	fmt.Println("Feliz", "Dia", PI)

	const Truth = true
	fmt.Println("Regras GO", Truth)
}
