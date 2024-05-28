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
