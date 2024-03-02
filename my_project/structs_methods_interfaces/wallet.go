package metrics

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var InsufficientFundsError = errors.New("Insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	//w is a pointer to wallet
	//in Go struct fields are automatically dereferenced
	//so we can write w.balance instead of (*w).balance
	//this is true for pointers to any type
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
	//no need to accept a pointer to a wallet
	//because we don't need to modify the wallet
	//but to avoid copying the wallet we can use a pointer
	//and to maintain consistency we use a pointer
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
