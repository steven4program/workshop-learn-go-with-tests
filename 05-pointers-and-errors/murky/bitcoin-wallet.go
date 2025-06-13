package murky

import (
	"errors"
	"fmt"
)

// https://github.com/uber-go/guide/blob/master/style.md#receivers-and-interfaces

var ErrInsufficientBalance = errors.New("you withdraw too much")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}


type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientBalance
	}

	w.balance -= amount
	return nil
}
