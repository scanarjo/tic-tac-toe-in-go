package tic_tac_toe

import "math/rand"

type Board [3][3]rune

type Game struct {
	Board  Board
	Next   rune
	IsOver bool
}

func NewGame() *Game {
	const BLANK = ' '
	const X = 'X'
	const O = 'O'

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
