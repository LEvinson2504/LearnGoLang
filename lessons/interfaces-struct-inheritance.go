// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type transactions interface {
	withdraw(amount float64) float64
	deposit(amount float64) float64
	balance() float64
}

type BankAccount struct {
	owner             string
	available_balance float64
}

// implement methods of interface transactions to struct BankAccount

func (account *BankAccount) balance() float64 {
	return account.available_balance
}

func (account *BankAccount) deposit(amount float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic :", r)
		}
	}()
	if amount < 0 {
		panic("Amount less than zero. Cannot deposit negative amount")
	} else {
		account.available_balance += amount
		return account.available_balance
	}
}

func (account *BankAccount) withdraw(amount float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic :", r)
		}
	}()
	if amount > account.available_balance {
		panic("Cannot withdraw amount greater than available balance.")
	} else {
		account.available_balance -= amount
		return account.available_balance
	}
}

func CreateBankAccount(owner string) *BankAccount {
	return &BankAccount{owner, 0}
}

func CheckAccountBalance(account transactions) {
	fmt.Println(account.balance())
}

func main() {
	var account1 = CreateBankAccount("Levinson")
	fmt.Println("Created a Bank account and initialised balance as zero")
	// fmt.Println(account1)
	CheckAccountBalance(account1)
	account1.deposit(1000)
	CheckAccountBalance(account1)
	account1.deposit(1000)
	CheckAccountBalance(account1)
	account1.withdraw(2000)
	fmt.Println(account1.balance())
	account1.withdraw(1000)
	account1.deposit(-1000)
	fmt.Println(account1.available_balance)
}

// Output :

// Created a Bank account and initialised balance as zero
// 0
// 1000
// 2000
// 0
// recovered from panic : Cannot withdraw amount greater than available balance.
// recovered from panic : Amount less than zero. Cannot deposit negative amount
// 0
// 
// Program exited.
