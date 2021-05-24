package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append, função para add novos elementos em uma slice
	s = append(s, 0)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
