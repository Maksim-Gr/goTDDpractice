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

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("wanted an error to be returned")
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

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Coin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Coin(100))
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("expected an error, got none")
		}
	})
}
