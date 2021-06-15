package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // "herdei" os atributos de carro; também chamado de campos anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40" // pelo fato de criar um campo anônimo, eu consigo acessar o atributo da struct("classe") pai
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f) // visualizando a composição
}
