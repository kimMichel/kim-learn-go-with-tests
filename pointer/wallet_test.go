package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		result := wallet.Balance()

		fmt.Printf("The address of the balance in test is %v \n", &wallet.balance)

		expect := Bitcoin(10)

		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		result := wallet.Balance()

		fmt.Printf("The address of the balance in test is %v \n", &wallet.balance)

		expect := Bitcoin(10)

		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}
	})
}
