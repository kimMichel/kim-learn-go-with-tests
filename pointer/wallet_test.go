package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		error := wallet.Withdraw(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
		confirmExistentError(t, error)
	})

	t.Run("trying to take more than have", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		error := wallet.Withdraw(Bitcoin(100))

		confirmBalance(t, wallet, initialBalance)
		confirmError(t, error, ErrorInsufficientFunds.Error())
	})
}

func confirmBalance(t *testing.T, wallet Wallet, expect Bitcoin) {
	t.Helper()
	result := wallet.Balance()

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}

}

func confirmError(t *testing.T, value error, expect string) {
	t.Helper()
	if value == nil {
		t.Fatal("expect an error, but not occured")
	}

	result := value.Error()

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}
}

func confirmExistentError(t *testing.T, result error) {
	t.Helper()
	if result != nil {
		t.Fatal("unexpected error")
	}
}
