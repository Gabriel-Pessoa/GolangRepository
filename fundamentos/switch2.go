package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quando é Sábado?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoje")
	case today + 1:
		fmt.Println("Amanhã")
	case today + 2:
		fmt.Println("Em dois dias")
	default:
		fmt.Println("Distante")
	}
}
