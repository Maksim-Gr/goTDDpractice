package pointers

type Coin int
type Wallet struct {
	balance Coin
}

func (w *Wallet) Deposit(amount Coin) {
	w.balance += amount
}

func (w *Wallet) Balance() Coin {
	return w.balance
}
