package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://example.com.br"

	jsonStr := `{
		"key": "123456teste",
		"secret": "12345secret"
	}`

	req, _ := http.NewRequest("POST", url, strings.NewReader(jsonStr))

	body, _ := ioutil.ReadAll(req.Body)

	fmt.Println(string(body))

}
