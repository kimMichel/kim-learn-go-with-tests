package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(value Bitcoin) {
	fmt.Printf("The address of the balance in deposit is %v \n", &w.balance)
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(value Bitcoin) {
	w.balance -= value
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
