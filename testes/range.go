package main

import "fmt"

type Account struct {
	ID   int
	User string
}

func encodeAccounts(accounts []*Account) []Account {
	var accountsEncoded []Account
	for _, account := range accounts {
		accountsEncoded = append(accountsEncoded, Account{
			ID:   account.ID,
			User: account.User,
		})
	}
	return accountsEncoded
}

// numeros := [...]int{1, 2, 3, 4, 5} // sem o operador spread a variável se transforma num slice, e não num array.

func main() {
	accountSlice := []*Account{
		{ID: 1, User: "Gabriel"},
		{ID: 2, User: "João"},
		{ID: 3, User: "Maria"},
		{ID: 4, User: "Chris"},
		{ID: 5, User: "Lucas"},
		{ID: 6, User: "Lucas"},
		{ID: 7, User: "Gabriel"},
		{ID: 9, User: "Júlio"},
		{ID: 10, User: "João"},
	}
	encode := encodeAccounts(accountSlice)
	fmt.Println(encode, len(encode))

}
