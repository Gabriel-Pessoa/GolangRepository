package main

import "fmt"

type nota float64 // declarei um tipo nota que se baseia no float64. Posso criar "métodos" com receiver nota

// eu não consigo usar o receiver, por exemplo: n float64. Lança um erro
func (n nota) entre(inico, fim float64) bool {
	return float64(n) >= inico && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
