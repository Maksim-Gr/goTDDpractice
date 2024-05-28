package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Coin(10))

	got := wallet.Balance()
	want := Coin(10)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
