package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CLIGame struct {
	game   *Game
	input  io.Reader
	output io.Writer
}

func NewCLIGameWithIO(input io.Reader, output io.Writer) *CLIGame {
	return &CLIGame{
		game:   NewGame(),
		input:  input,
		output: output,
	}
}

func (cli *CLIGame) Start() {
	fmt.Fprintln(cli.output, "Starting a new game of Tic Tac Toe...")
	fmt.Fprintln(cli.output, cli.game)

	scanner := bufio.NewScanner(cli.input)

	for !cli.game.IsOver() {
		fmt.Fprintf(cli.output, "Player %c, enter your move (row col): ", cli.game.Next)

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		var row, col int

		_, err := fmt.Sscanf(input, "%d %d", &row, &col)
		if err != nil {
			fmt.Fprintln(cli.output, "Invalid input. Please enter your move as 'row col'.")
			continue
		}

		err = cli.game.Move(row, col)
		if err != nil {
			fmt.Fprintf(cli.output, "Error: %v\n", err)
			continue
		}

		fmt.Fprintln(cli.output, cli.game)
	}

	if winner := cli.game.Winner(); winner != BLANK {
		fmt.Fprintf(cli.output, "Game over: Player %c wins!\n", winner)
	} else {
		fmt.Fprintln(cli.output, "Game over: It's a draw!")
	}
}
