package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode add mais métodos, agora tenho que implementá-los abaixo, caso queira que a struct implemente essa interface tbm
}

type bmw7 struct{} //uma struct implementa as interfaces:esportivo, luxuoso e esportivoLuxuoso. De forma implícita implementa as três interfaces

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
