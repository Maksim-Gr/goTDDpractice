package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Coin(10))
		got := wallet.Balance()
		want := Coin(10)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Withdraw(Coin(10))
		got := wallet.Balance()
		want := Coin(10)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
