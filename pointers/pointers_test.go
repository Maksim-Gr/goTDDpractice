package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Coin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Coin(10))
		assertBalance(t, wallet, Coin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Coin(20)}
		wallet.Withdraw(Coin(10))
		assertBalance(t, wallet, Coin(10))
	})
}
