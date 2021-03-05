package pointers_and_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %p\n", &(wallet.balance))

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
