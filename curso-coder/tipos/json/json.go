package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	// mapeamento de como quero que esse ID apareça no JSON
	// por conversão, atributo que iniciam com letra maiúsculas são públicas (exportável) e atributos iniciados com letra minúscula são privadas (não-exportável).
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
