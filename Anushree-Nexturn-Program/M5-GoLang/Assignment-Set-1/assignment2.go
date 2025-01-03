//Exercise 2: Bank Operation System

package main

import (
	"errors"
	"fmt"
)

// Account struct to store account details
type CustomerAccount struct {
	CustomerID      int
	CustomerName    string
	AccountBalance  float64
	TransactionLogs []string
}

// Global slice to store all accounts
var customerAccounts []CustomerAccount

// Deposit method to add funds to an account
func (acc *CustomerAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc.AccountBalance += amount
	acc.TransactionLogs = append(acc.TransactionLogs, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw method to deduct funds from an account
func (acc *CustomerAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if acc.AccountBalance < amount {
		return errors.New("insufficient funds")
	}
	acc.AccountBalance -= amount
	acc.TransactionLogs = append(acc.TransactionLogs, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

// ViewTransactions method to display account transaction history
func (acc *CustomerAccount) ViewTransactions() {
	fmt.Println("Transaction History:")
	if len(acc.TransactionLogs) == 0 {
		fmt.Println("No transactions recorded.")
		return
	}
	for index, transaction := range acc.TransactionLogs {
		fmt.Printf("%d. %s\n", index+1, transaction)
	}
}

// LocateAccount searches for an account by its ID
func LocateAccount(id int) (*CustomerAccount, error) {
	for i := range customerAccounts {
		if customerAccounts[i].CustomerID == id {
			return &customerAccounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// DisplayOptions function shows the available options
func DisplayOptions() {
	fmt.Println("\n--- Welcome to Bank Operation System ---")
	fmt.Println("1. Open New Account")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Check Account Balance")
	fmt.Println("5. View Account Transactions")
	fmt.Println("6. Exit")
	fmt.Print("Please select an option: ")
}

func main() {
	// Preloaded accounts
	customerAccounts = append(customerAccounts, CustomerAccount{CustomerID: 201, CustomerName: "John Doe", AccountBalance: 4000})
	customerAccounts = append(customerAccounts, CustomerAccount{CustomerID: 202, CustomerName: "Jane Smith", AccountBalance: 6000})
	customerAccounts = append(customerAccounts, CustomerAccount{CustomerID: 203, CustomerName: "Alice Johnson", AccountBalance: 3500})

	for {
		DisplayOptions()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Customer ID: ")
			var id int
			fmt.Scan(&id)

			fmt.Print("Enter Customer Name: ")
			var name string
			fmt.Scan(&name)

			fmt.Print("Enter Initial Deposit Amount: ")
			var initialDeposit float64
			fmt.Scan(&initialDeposit)

			customerAccounts = append(customerAccounts, CustomerAccount{CustomerID: id, CustomerName: name, AccountBalance: initialDeposit})
			fmt.Println("New account successfully created.")

		case 2:
			fmt.Print("Enter Customer ID: ")
			var id int
			fmt.Scan(&id)

			account, err := LocateAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("Enter Deposit Amount: ")
			var amount float64
			fmt.Scan(&amount)

			if err := account.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case 3:
			fmt.Print("Enter Customer ID: ")
			var id int
			fmt.Scan(&id)

			account, err := LocateAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("Enter Withdrawal Amount: ")
			var amount float64
			fmt.Scan(&amount)

			if err := account.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case 4:
			fmt.Print("Enter Customer ID: ")
			var id int
			fmt.Scan(&id)

			account, err := LocateAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("Current Balance for %s: %.2f\n", account.CustomerName, account.AccountBalance)

		case 5:
			fmt.Print("Enter Customer ID: ")
			var id int
			fmt.Scan(&id)

			account, err := LocateAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			account.ViewTransactions()

		case 6:
			fmt.Println("Thank you for using the Bank Operation System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}