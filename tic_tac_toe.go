package tic_tac_toe

type Board [3][3]rune

type Game struct {
	Board Board
}

func NewGame() *Game {
	const BLANK = ' '

	return &Game{
		Board: Board{
			{BLANK, BLANK, BLANK},
			{BLANK, BLANK, BLANK},
			{BLANK, BLANK, BLANK},
		},
	}
}
