package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// reparar que o receiver é um ponteiro, uma referência para o struct ("classe") .
// Então, alterações internas nos atributos da classe dentro desse método, refletem no valor original dos atributos da classe
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Gabriel", "Júlio"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Gabriel Pessoa")
	fmt.Println(p1.getNomeCompleto())
}
