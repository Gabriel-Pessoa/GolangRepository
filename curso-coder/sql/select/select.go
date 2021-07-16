package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:ails self many@/cursogo") // acessa o driver definido no import
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, nome from usuarios where id > ?", 5) // executa a query passando 5 como parâmetro, que irá substituir a interrogação
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // fechando resultado das linhas

	// enquanto tiver próxima linha para executar
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome) // mapea o retorno da consulta e atribui a variável u
		fmt.Println(u)            // imprime no console
	}
}
