package main

import (
	"fmt"
)

type DollarAccount struct {
	balance int
}

func NewDollarAccount() *DollarAccount {
	return &DollarAccount{
		balance: 0,
	}
}

func (acc *DollarAccount) Deposit(amount int) {
	acc.balance += amount
}

func (acc *DollarAccount) Withdraw(amount int) error {
	if (acc.balance - amount) < 0 {
		return fmt.Errorf("failed transaction for insufficient funds:%d", acc.balance)
	}
	acc.balance -= amount
	return nil
}

func (acc *DollarAccount) GetBalance() int {
	return acc.balance
}
