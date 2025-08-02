package tic_tac_toe

import (
	"fmt"
	"math/rand"
	"strings"
)

type Board [3][3]rune

type Game struct {
	Board  Board
	Next   rune
	IsOver bool
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
		Next: firstPlayer,
		Board: Board{
			{BLANK, BLANK, BLANK},
			{BLANK, BLANK, BLANK},
			{BLANK, BLANK, BLANK},
		},
	}
}

func (game *Game) changePlayer() {
	switch game.Next {
	case X:
		game.Next = O
	case O:
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
	var result string
	for _, row := range game.Board {
		for _, cell := range row {
			result += fmt.Sprintf("%s ", printCell(cell))
		}
		result = strings.Trim(result, " ") + "\n"
	}
	return result
}
