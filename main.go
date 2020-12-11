package main

import (
	"fmt"

	"github.com/leejiwon-ncurity/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jiwon")
	fmt.Println(account)
}
