package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 21

	n, err := fmt.Fprint(os.Stdout, name, " tem ", age, " anos de idade.\n") // função Fprint retorna int e error, e atribue às variav. n e err

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}

	fmt.Print(n, " bytes escrito")
}
