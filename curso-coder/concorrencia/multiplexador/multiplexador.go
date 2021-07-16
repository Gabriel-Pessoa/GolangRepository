package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Multiplexar: é pego duas fontes de dados, dois canais, convertendo em um único canal, junta a informação.

func encaminhar(origem <-chan string, destino chan string) { // recebo um dado de um canal de origem e encaminho para outro canal
	for {
		destino <- <-origem // a leitura do canal de origem encaminha para o canal destino
	}
}

// multiplexar: misturar(mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string { // recebo duas entrada (origens) e transformo em um
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func pageTitle(urls ...string) <-chan string {
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
	c := juntar(
		pageTitle("https://www.cod3r.com.br", "https://www.google.com"),
		pageTitle("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
