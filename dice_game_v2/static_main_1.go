package main

import (
	"dice_game_v2/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\n--- Welcome to the Dice Game ---")

	// Step 1: Start Balance  - Player must introduce the amount he wants to play
	balance := game.StartWallet()

	for {
		// Step 2: Options Odd or Even
		var choice string
		fmt.Print("Choose Odd or Even (O/E): ")
		fmt.Scan(&choice)
		if choice != "O" && choice != "E" {
			fmt.Println("Invalid choice. Please choose 'O' for Odd or 'E' for Even.")
			continue
		}

		// Step 3: Bet amount (can't be bigger then available balance)
		//Safety and Protections
		var bet int
		fmt.Print("Enter the amount to bet: $")
		fmt.Scan(&bet)
		if bet <= 0 || bet > balance {
			fmt.Println("Invalid bet amount. Make sure it’s a positive number and doesn’t exceed your balance.")
			continue
		}

		// Step 4: Play the game
		win, _ := game.Play(bet, choice)

		// Step 5: Update balance and display outcome
		game.EndPlay(&balance, win, bet)

		// Step 6: Check if player wants to continue or if balance is depleted
		if balance <= 0 {
			fmt.Println("You're out of balance. Game over!")
			break
		}
		// Step 7: Start loop again or break it by ending the game
		var playAgain string
		fmt.Print("Do you want to play again? (Y/N): ")
		fmt.Scan(&playAgain)
		if playAgain != "Y" && playAgain != "y" {
			fmt.Printf("Thanks for playing! Final Balance: $%d\n", balance)
			break
		}
	}
}
