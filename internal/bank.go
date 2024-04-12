package main

import "fmt"

type Account struct {
	ID      string
	Balance float64
}

func (a *Account) Deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}
	a.Balance += amount
	return true
}

func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}
	a.Balance -= amount
	return nil
}

func CreateAccount(id string, initialDeposit float64) *Account {
	return &Account{
		ID:      id,
		Balance: initialDeposit,
	}
}
