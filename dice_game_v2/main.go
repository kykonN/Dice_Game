package main

import (
	"dice_game_v2/game"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

func diceGameHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Initialize the balance
	balance := game.StartWallet()

	for {
		// Step 1: Receive the player's choice (Odd/Even)
		_, choiceBytes, err := conn.ReadMessage() // ReadMessage returns: (messageType, messageContent, error)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		choice := string(choiceBytes) // Convert the byte slice to a string

		if choice != "O" && choice != "E" {
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid choice. Please choose 'O' for Odd or 'E' for Even."))
			continue
		}

		// Step 2: Receive the bet amount from the client
		_, betBytes, err := conn.ReadMessage() // ReadMessage returns: (messageType, messageContent, error)
		if err != nil {
			fmt.Println("Error reading bet amount:", err)
			break
		}
		var bet int
		// Convert betBytes (byte slice) to an integer. You'll need to parse it correctly.
		_, err = fmt.Sscanf(string(betBytes), "%d", &bet)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid bet amount format."))
			continue
		}

		if bet <= 0 || bet > balance {
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid bet amount. Make sure it’s a positive number and doesn’t exceed your balance."))
			continue
		}

		// Step 3: Play the game (roll the dice)
		win, diceRoll := game.Play(bet, choice)

		// Step 4: Display the result to the player
		if win {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("You won! You rolled %d. New balance: $%d", diceRoll, balance+bet)))
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("You lost! You rolled %d. New balance: $%d", diceRoll, balance-bet)))
		}

		// Step 5: Update balance
		game.EndPlay(&balance, win, bet)

		// Step 6: Check if the player has run out of balance
		if balance <= 0 {
			conn.WriteMessage(websocket.TextMessage, []byte("You're out of balance. Game over!"))
			break
		}

		// Step 7: Ask if the player wants to continue
		conn.WriteMessage(websocket.TextMessage, []byte("Do you want to play again? (Y/N):"))
	}
}

func main() {
	
	rand.Seed(time.Now().UnixNano())

	// Start the WebSocket server
	http.HandleFunc("/game", diceGameHandler)
	fmt.Println("WebSocket server started on ws://localhost:8080/game")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
