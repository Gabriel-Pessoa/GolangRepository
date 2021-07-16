package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente leitura
// encapsula, simula uma chamada de uma goroutine, já resolve tudo. Quem tá camando não precisa criar goroutines, criar canal; está tudo encapsulado dentro da função
// Podia passar 20 urls, que ele iria tratar todas
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { // função anônima auto-invocada
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {

	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Printf("Primeiros colocados: %v | %v \n", <-t1, <-t2)
	fmt.Printf("Segundos colocados: %v | %v \n", <-t1, <-t2)
}
