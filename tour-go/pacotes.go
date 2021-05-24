//Programas começando rodando pelo pacote main
package main

import ( //Importação consignada, se utiliza dos parenteses para dois um mais imports
	"fmt"
	"math/rand" //Por convenção o nome do pacote é o mesmo que o último elemento da importação.
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Um número randômico", rand.Intn(10))
}
