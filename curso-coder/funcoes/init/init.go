package main

import "fmt"

// convenção em go, quando um arquivo é lido, executa primeiramente a função init
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
