package game

import "fmt"

func StartWallet() int {
	var balance int

	for {
		fmt.Println("Insert balance: $")
		fmt.Scan(&balance)

		if balance > 0 {
			break // Exit the loop if balance is valid
		} else {
			fmt.Println("Please insert a positive balance to play.")
		}
	}

	return balance // Return the valid balance
}
