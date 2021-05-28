package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	i := 1.0

	for i <= 100 { // não tem while em go

		number := fmt.Sprintf("%.2f", i)
		removePoint := strings.Replace(number, ".", "", -1)
		limitStringWithThreeNumbers := removePoint[0:3]

		fmt.Printf("02%v\n", limitStringWithThreeNumbers)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturado...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é impar")
		}
	}

	// laço infinito
	for {
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second) // a cada segundo
		//time.Sleep(time.Second * 5) // a cada 5 segundos
	}

}
