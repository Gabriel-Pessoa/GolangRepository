package main

import "fmt"

func main() {
	// map: estrutura com par: chave/valor, add a mesma chave irá sobrescrever o seu valor

	// var aprovados map[int]string // não funciona, pois os maps devem ser inicializados primeiro
	// mapas devem ser inicializados

	// inicializado map
	aprovados := make(map[int]string)
	aprovados[123456] = "Maria"
	aprovados[456789] = "Pedro"
	aprovados[789123] = "Ana"

	fmt.Println(aprovados)

	// percorrendo map com for
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// acessando valor de map
	fmt.Println(aprovados[789123])

	// deletando valor de um map
	delete(aprovados, 789123)

	//tentando acessar valor excluído
	fmt.Println(aprovados[789123])

}
