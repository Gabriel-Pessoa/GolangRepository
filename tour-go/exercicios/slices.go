package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var matriz [][]uint8

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			matriz[x][y] = uint8(x * y)
		}
	}

	return matriz
}

func main() {
	fmt.Println(Pic(2, 2))
}
