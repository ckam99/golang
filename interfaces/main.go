package main

import "fmt"

type IBankAccount interface {
	Deposit(amount int)
	Withdraw(amount int) error
	GetBalance() int
}

func main() {
	accounts := []IBankAccount{
		NewBitcoinAccount(),
		NewDollarAccount(),
	}
	for _, account := range accounts {
		fmt.Printf("Initial balance=%d\n", account.GetBalance())
		account.Deposit(500)
		fmt.Printf("After Deposit balance=%d\n", account.GetBalance())
		if err := account.Withdraw(100); err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Printf("After Withdraw balance=%d\n", account.GetBalance())
	}

}
