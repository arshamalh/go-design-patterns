package main

import "fmt"

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}