package main

import (
	"errors"
	"fmt"
)

// create new types from existing ones.
type Bitcoin int
// In Go, errors are values, so we can refactor it out into a variable and have a single source of truth for it.
// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
type Wallet struct {
	balance Bitcoin
}

// This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints.
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// take a pointer to that wallet so that we can change the original values within it.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}