package wallet

import (
	"errors"
	"fmt"
)

type Dogecoin int

func (d Dogecoin) String() string {
	return fmt.Sprintf("%d DOGE", d)
}

type Wallet struct {
	balance Dogecoin
}

func (w *Wallet) Deposit(amount Dogecoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Dogecoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withraw, insufficient funds")

func (w *Wallet) Withdraw(amount Dogecoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
