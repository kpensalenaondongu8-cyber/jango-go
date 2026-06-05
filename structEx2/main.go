package  main

import (
	"fmt"
)
type BankAcc struct {
	AccountNumber int
	Owner string
	Balance  int
}
func (a *BankAcc) deposit(amount int) {
  a.Balance += amount
  fmt.Printf("You Deposited: %v\n", amount)
}
func (a *BankAcc) Withdrawal(amount int) {
	if amount > a.Balance {
		fmt.Println("Insufficient funds")
		return
	}
	a.Balance -= amount
	fmt.Printf("You withdrew: %v\n", amount)
}
func (a * BankAcc) DisplayBalance() {
   fmt.Printf("| Acct Number:%v |  Owner: %v | Balance: %v\n", a.AccountNumber, a.Owner, a.Balance)
}
func main() {
	w := BankAcc {
		AccountNumber: 8164431360,
		Owner: "Thomas",
		Balance: 1000000,
	}
	var transaction int
	var Cash int
	for {
	fmt.Println("Choose Transaction:\n",1,"Withdraw",2,"Deposit",3,"Check Balance", 4,"Exit")
	fmt.Scan(&transaction)
	if transaction == 1  {
      fmt.Println("Enter Amount to withdraw")
	  fmt.Scan(&Cash)
	  w.Withdrawal(Cash)
	  continue
	}
	if transaction == 2 {
		fmt.Println("Enter amount to Deposit")
		fmt.Scan(&Cash)
		w.deposit(Cash)
		continue
	}
	if transaction == 3 {
		w.DisplayBalance()
		continue
	}
	if transaction == 4 {
		fmt.Println("Nice banking with us")
		return
	}
}
}