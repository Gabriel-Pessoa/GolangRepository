package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go rodando em ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// windows, freebsb, openbsd, plan9
		fmt.Printf("%s. \n", os)
	}
}
