//switch sem condição
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 17:
		fmt.Println("Bom tarde")
	default:
		fmt.Println("Bom noite")
	}
}
