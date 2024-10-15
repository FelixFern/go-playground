package wallet

import "fmt"

func (w *Wallet) Deposit(amount float64) {
	fmt.Println(fmt.Sprintf("Added %f to wallet with id: %d", amount, w.Id))
	w.Balance += amount
}

func (w *Wallet) Deduct(amount float64) {
	fmt.Println(fmt.Sprintf("Deducted %f to wallet with id: %d", amount, w.Id))
	w.Balance -= amount
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}
