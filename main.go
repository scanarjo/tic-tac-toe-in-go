package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	game := NewGame()
	fmt.Println("Starting a new game of Tic Tac Toe...")
	fmt.Println(game)

	scanner := bufio.NewScanner(os.Stdin)

	for !game.IsOver() {
		fmt.Printf("Player %c, enter your move (row col): ", game.Next)

		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())

			var row, col int

			_, err := fmt.Sscanf(input, "%d %d", &row, &col)
			if err != nil {
				fmt.Println("Invalid input. Please enter your move as 'row col'.")
				continue
			}

			err = game.Move(row, col)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println(game)
		}
	}
}
