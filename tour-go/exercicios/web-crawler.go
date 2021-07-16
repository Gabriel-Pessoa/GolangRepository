package main

import (
	"fmt"
)

type Fetcher interface {

	// Fetch retorna o corpo do URL e uma parte dos URLs encontrados nessa página.
	Fetch(url string) (body string, urls []string, err error)
}

var cacheUrl []string

// O rastreamento usa o fetcher para rastrear páginas recursivamente, começando com url, até a profundidade máxima.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Buscar URLs em paralelo.

	if depth <= 0 {
		cacheUrl = nil
		return
	}
	if !hasSomeUrl(cacheUrl, url) {
		cacheUrl = append(cacheUrl, url)

	} else {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}

}

func hasSomeUrl(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
