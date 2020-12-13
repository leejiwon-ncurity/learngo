package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// error를 만들거라면 변수이름을 err000이라고 지정해주어야함.
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// 메소드
// receiver는 항상 struct의 맨 앞글자 lowercase로 해준다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string)  {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string  {
	return a.owner
}

// struct의 toString() 같은 것
func (a Account) String() string  {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
	
}