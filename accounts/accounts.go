package accounts

import "errors"

// Account struct
type Account struct{
	owner string
	balance int
}

//NewAccount creates Account
func NewAccount(owner string) *Account{
	account := Account{owner:owner, balance: 0}
	return &account
}

//Deposit x amount on your acount
func (a *Account) Deposit(amount int){
	a.balance += amount
}

//return only balance value
func (a Account) Balace() int{
	return a.balance
}

//Widthraw x amout on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("can't withdraw")
	}
	a.balance -=amount
	return nil
}