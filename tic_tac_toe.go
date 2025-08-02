package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Row [3]rune

type Board [3]Row

type Game struct {
	Board Board
	Next  rune
}

const (
	BLANK = ' '
	X     = 'X'
	O     = 'O'
)

func NewGame() *Game {
	players := [2]rune{X, O}

	firstPlayer := players[rand.Intn(2)]

	return &Game{
		Board: Board{
			{BLANK, BLANK, BLANK},
			{BLANK, BLANK, BLANK},
			{BLANK, BLANK, BLANK},
		},
		Next: firstPlayer,
	}
}

func (game *Game) changePlayer() {
	if game.Next == X {
		game.Next = O
	} else {
		game.Next = X
	}
}

func (game *Game) Move(row, col int) error {
	if row < 0 || row >= len(game.Board) || col < 0 || col >= len(game.Board[row]) {
		return fmt.Errorf("Move out of bounds: %d, %d", row, col)
	}

	if game.Board[row][col] != BLANK {
		return fmt.Errorf("Illegal move: Cell (%d,%d) already contains %q", row, col, game.Board[row][col])
	}

	game.Board[row][col] = game.Next

	game.changePlayer()

	return nil
}

func printCell(c rune) string {
	switch c {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return "_"
	}
}

func (game *Game) String() string {
	var result strings.Builder
	for _, row := range game.Board {
		for i, cell := range row {
			result.WriteString(printCell(cell))

			if i != len(row)-1 {
				result.WriteString(" ")
			}
		}

		result.WriteString("\n")
	}

	return result.String()
}

func (game *Game) IsOver() bool {
	if game.Winner() != BLANK {
		return true
	}

	for _, row := range game.Board {
		for _, cell := range row {
			if cell == BLANK {
				return false
			}
		}
	}

	return true
}

func checkRow(row Row) rune {
	if row[0] == row[1] && row[1] == row[2] && row[0] != BLANK {
		return row[0]
	}

	return BLANK
}

func (game *Game) Winner() rune {
	for _, row := range game.Board {
		winner := checkRow(row)
		if winner != BLANK {
			return winner
		}
	}

	for col := range game.Board[0] {
		winner := checkRow(Row{game.Board[0][col], game.Board[1][col], game.Board[2][col]})
		if winner != BLANK {
			return winner
		}
	}

	winner := checkRow(Row{game.Board[0][0], game.Board[1][1], game.Board[2][2]})
	if winner != BLANK {
		return winner
	}

	winner = checkRow(Row{game.Board[0][2], game.Board[1][1], game.Board[2][0]})
	if winner != BLANK {
		return winner
	}

	return BLANK
}
