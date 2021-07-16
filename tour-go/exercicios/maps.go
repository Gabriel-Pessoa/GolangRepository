package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	str := strings.Split(s, " ")
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}

	return m
}

func main() {
	fmt.Println(WordCount("Go! I am learning. The quick brown fox jumped over the lazy dog. I ate a donut. Then I ate another donut."))
}
