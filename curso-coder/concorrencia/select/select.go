package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

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

func oMaisRapido(url1, url2, url3 string) string {
	c1 := pageTitle(url1)
	c2 := pageTitle(url2)
	c3 := pageTitle(url3)

	// estrutura de controle específica para concorrência
	// o primeiro dado que chegar, ele executa o case específico
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default: // sempre caía no cenário de default, porque os outros não tinham respostas ainda
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	fmt.Println(campeao)
}
