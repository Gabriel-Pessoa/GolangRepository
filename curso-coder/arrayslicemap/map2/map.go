package main

import "fmt"

func main() {
	// inicializando e preenchendo map
	funcionESalario := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Júnior":   1200.0, // última linha termina com vírgula ou com chave de fechamento encostada
	}

	//adicionando valor fora da declaração
	funcionESalario["Rafael Filho"] = 1350.0

	//tentativa de exclusão de chave inexistente no map não lança erro
	delete(funcionESalario, "Inexistente")

	// iterando pelo map
	for nome, salario := range funcionESalario {
		fmt.Printf("%v, %0.2f\n", nome, salario)
	}

}
