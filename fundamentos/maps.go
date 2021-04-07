package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // retorna um map com um tipo determinado, inicializado e pronto para uso
	// criei uma chave "Bell Labs" mapeando Vertex para ela
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	//Acessando valor da chave
	fmt.Println(m["Bell Labs"])
}
