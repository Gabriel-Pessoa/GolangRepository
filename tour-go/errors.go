package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"isto não funcionou",
	}
}

func main() {
	/*
		Funções frequentemente retornam um valor error e a chamada do código deve lidar com erros testando se o erro é igual a nil.
		i, err := strconv.Atoi("42")
		if err != nil {
		    fmt.Printf("Não conseguiu converter número: %v\n", err)
		}
		fmt.Println("Inteiro convertido:", i)
		Um error nil indica sucesso; um error não-nil indica fracasso.
	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
