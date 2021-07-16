package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:ails self many@/cursogo") // acessa o driver definido no import
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)") // cria um query base para, nesse caso, receber values

	// insert que usa a query base criada acima
	// execuções sem se interessar com o retorno
	stmt.Exec("Maria")
	stmt.Exec("João")

	// execução interessada no retorno
	result, _ := stmt.Exec("Pedro")
	id, _ := result.LastInsertId() // obtendo o último id da inserção acima
	fmt.Println(id)
}
