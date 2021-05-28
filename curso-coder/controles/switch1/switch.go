package main

import "fmt"

func notaPraConceito(n float64) string {
	/*var nota = int(n)

	switch nota {
	case 10:
		fallthrough //semelhante ao inverso do break
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}*/

	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaPraConceito(9.5))
	fmt.Println(notaPraConceito(6.9))
	fmt.Println(notaPraConceito(2.1))
	fmt.Println(notaPraConceito(10))
	fmt.Println(notaPraConceito(11))

}
