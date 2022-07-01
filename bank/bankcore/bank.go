package bankcore

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(money float64) error {
	if money <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += money
	return nil
}

func (a *Account) Withdraw(money float64) error {
	if money <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}
	if a.Balance-money < 0 {
		return errors.New("the amount to withdraw is greater than the balance")
	}

	a.Balance -= money
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%+v - %+v - %+v", a.Number, a.Customer.Name, a.Balance)
}
