package game

import "fmt"

func EndPlay(balance *int, win bool, bet int) {
	if win {
		*balance += bet // Win: add bet to balance
		fmt.Printf("You won! New Balance: $%d\n", *balance)
	} else {
		*balance -= bet // Loss: subtract bet from balance
		fmt.Printf("You lost! New Balance: $%d\n", *balance)
	}
}
