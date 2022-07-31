package main

import (
	"fmt"
	"sync"
	"time"
)

type TransactionType int64

var balance int = 0

const (
	Debit TransactionType = iota
	Credit
)

var wait sync.WaitGroup
var mutex sync.Mutex

func (t TransactionType) String() string {
	switch t {
	case Debit:
		return "Debit"
	case Credit:
		return "Credit"
	default:
		return fmt.Sprintf("Unknown(%d)", t)
	}
}

func transaction(amount int, transactionType TransactionType, name string) {
	if transactionType == Debit {
		if (balance - amount) < 0 {
			fmt.Printf("We are unable to process the transaction due to insufficient funds: %v [%v]\n", balance, amount)
			return
		}
		mutex.Lock()
		balance -= amount
		fmt.Printf("$%d has been debited by %v\n", amount, name)
		mutex.Unlock()
	}
	if transactionType == Credit {
		mutex.Lock()
		balance += amount
		fmt.Printf("$%d has been credited by %v\n", amount, name)
		mutex.Unlock()
	}
}

func transac(amount int, transactionType TransactionType, name string) {
	defer wait.Done()
	for i := 0; i < 5; i++ {
		transaction(amount, transactionType, name)
		time.Sleep(500 * time.Millisecond)
	}
}

func MutexExample() {
	wait.Add(3)
	go transac(10, Credit, "Marco")
	go transac(10, Debit, "Rachel")
	go transac(1, Credit, "Bamba")

	wait.Wait()

	fmt.Printf("-- your current balance is $%d\n", balance)
}
