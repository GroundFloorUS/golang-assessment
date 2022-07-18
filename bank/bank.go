package bank

import "net/http"

type Account struct {
	Balance int
}

//ConnectToGroundFloor makes a http GET request to Groundfloor's homepage to provide latency
func ConnectToGroundFloor() error {
	_, err := http.Get("https://groundfloor.us/")
	return err
}

func (a *Account) Deposit(amount int) int {
	a.Balance = a.Balance + amount
	return a.Balance
}

func (a *Account) Withdraw(amount int) int {
	if a.Balance >= amount {
		a.Balance = a.Balance - amount
	}
	return a.Balance
}
