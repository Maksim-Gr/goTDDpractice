package pointers

import (
	"fmt"
)

type Coin int
type Wallet struct {
	balance Coin
}

func (c Coin) String() string {
	return fmt.Sprintf("%d COIN", c)
}

func (w *Wallet) Deposit(amount Coin) {
	w.balance += amount
}

func (w *Wallet) Balance() Coin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Coin) error {
	if amount > w.balance {
		return fmt.Errorf("cannot withdraw more than the current balance")
	}
	w.balance -= amount
	return nil
}
