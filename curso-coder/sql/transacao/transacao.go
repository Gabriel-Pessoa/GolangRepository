package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:ails self many@/cursogo") // acessa o driver definido no import
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()                                                  // inicia uma transação
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)") // cria um query base para, nesse caso, receber values

	// insert que usa a query base criada acima
	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Tiago") // dará erro, já existe uma chave com esse id
	if err != nil {
		tx.Rollback() // desfaz as alterações, nesse caso, os inserts
		log.Fatal(err)
	}

	// em caso de sucesso, aplica todas as alterações da transação
	tx.Commit()
}
