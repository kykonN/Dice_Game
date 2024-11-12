package game

import (
	"fmt"
	"math/rand"
)

func Play(bet int, choice string) (bool, int) {

	diceRoll := rand.Intn(6) + 1
	fmt.Printf("You rolled: %d\n", diceRoll)

	isEven := diceRoll%2 == 0
	win := (choice == "E" && isEven) || (choice == "O" && !isEven)

	if win {
		fmt.Println("Result: Win")
	} else {
		fmt.Println(" Result: Lose")
	}

	return win, diceRoll
}
