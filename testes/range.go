package main

import "fmt"

type Account struct {
	ID int
}

func encodeAccounts(accounts []Account) []*Account {
	var accountsEncoded []*Account
	for _, account := range accounts {
		accountsEncoded = append(accountsEncoded, &Account{
			ID: account.ID,
		})
	}
	return accountsEncoded
}

// numeros := [...]int{1, 2, 3, 4, 5} // sem o operador spread a variável se transforma num slice, e não num array.

func main() {

	accountArray1 := []Account{}

	fmt.Println(accountArray1)

	fmt.Println(encodeAccounts(accountArray1))
}
