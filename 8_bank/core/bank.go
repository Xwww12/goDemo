package bank

import (
	"errors"
	"fmt"
)

// Customer 用户
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account 账户
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// Bank ...
type Bank interface {
	Statement() string
}

// Statement ...
func Statement(b Bank) string {
	return b.Statement()
}

// Deposit 存款
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += amount
	return nil
}

// Withdraw 取款
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement 对账
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// Transfer 转账
func (a *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to transfer should be greater than the account's balance")
	}

	a.Withdraw(amount)
	dest.Deposit(amount)
	return nil
}
