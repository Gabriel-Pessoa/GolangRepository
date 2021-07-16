package main

import "math"

// Inicializando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)
// Inicializando com letra minúsucla é PRIVADO (visível apenas dentro do pacote)
// obs: não posso ter funções, structs, interfaces ou variáveis com o mesmo nome dentro do pacote

// Ex:
// Ponto --> gera algo público
// ponto ou _Ponto --> gera algo privado

// Em algumas versões do Go, ele pede comentário quando se declara algo público:
// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// responsável por calcular a distânce linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
