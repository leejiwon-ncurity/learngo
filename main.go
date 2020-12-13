package main

import (
	"fmt"

	"github.com/leejiwon-ncurity/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jiwon")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account)
}
