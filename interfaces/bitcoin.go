package main

import (
	"fmt"
)

type BitcoinAccount struct {
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     400,
	}
}

func (acc *BitcoinAccount) Deposit(amount int) {
	acc.balance += amount
}

func (acc *BitcoinAccount) Withdraw(amount int) error {
	if (acc.balance - amount) < 0 {
		return fmt.Errorf("failed transaction for insufficient funds:%d", acc.balance)
	}
	acc.balance -= (amount + acc.fee)
	return nil
}

func (acc *BitcoinAccount) GetBalance() int {
	return acc.balance
}
