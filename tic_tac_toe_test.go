package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	freshBoard := Board{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}

	if !reflect.DeepEqual(game.Board, freshBoard) {
		t.Errorf("Expected %q, but got %q", freshBoard, game.Board)
	}

	if game.IsOver() {
		t.Error("Expected game to not be over, but it is")
	}

	if game.Next != 'X' && game.Next != 'O' {
		t.Errorf("Game should start with 'X' or 'O' as next player, but got %q", game.Next)
	}
}

func TestNewGameStartsWithRandomPlayer(t *testing.T) {
	game := NewGame()

	firstChosen := game.Next

	for i := 0; i < 100; i++ {
		game = NewGame()
		if game.Next != firstChosen {
			return
		}
	}

	t.Errorf("Game should start with a random player, but always started with %q", firstChosen)
}

func TestMove(t *testing.T) {
	game := NewGame()

	firstPlayer := game.Next

	game.Move(0, 0)

	if game.Board[0][0] != firstPlayer {
		t.Errorf("Expected first move to be %q, but got %q", firstPlayer, game.Board[0][0])
	}

	if game.Next == firstPlayer {
		t.Errorf("Expected next player to be different, but it is still %q", game.Next)
	}

	game.Move(1, 1)

	if game.Board[1][1] == firstPlayer {
		t.Errorf("Expected 2nd move to be the other player, but got %q", firstPlayer)
	}
}

func TestMoveOutOfBounds(t *testing.T) {
	game := NewGame()

	err := game.Move(3, 3)

	if err == nil {
		t.Error("Expected an error when moving out of bounds, but got none")
	}

	err = game.Move(-1, -1)

	if err == nil {
		t.Error("Expected an error when moving out of bounds, but got none")
	}
}

func TestIllegalMove(t *testing.T) {
	game := NewGame()

	firstPlayer := game.Next

	game.Move(0, 0)

	err := game.Move(0, 0)

	if err == nil {
		t.Error("Expected an error when making an already played move, but got none")
	}

	if game.Board[0][0] != firstPlayer {
		t.Errorf("Expected cell not to be overwritten, should be %q but got %q", firstPlayer, game.Board[0][0])
	}
}

func TestString(t *testing.T) {
	game := NewGame()

	output := fmt.Sprintf("%s", game)

	expected := `_ _ _
_ _ _
_ _ _
`

	if output != expected {
		t.Errorf("Expected output %q but got %q", expected, output)
	}

	if game.Next == 'X' {
		game.Move(0, 0)
		game.Move(1, 1)
	} else {
		game.Move(1, 1)
		game.Move(0, 0)
	}

	output = fmt.Sprintf("%s", game)

	expected = `X _ _
_ O _
_ _ _
`

	if output != expected {
		t.Errorf("Expected output %q but got %q", expected, output)
	}
}

func TestIsOverWhenFull(t *testing.T) {
	game := NewGame()

	fullBoard := Board{
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
		{'O', 'X', 'O'},
	}

	game.Board = fullBoard

	if !game.IsOver() {
		t.Error("Expected game to be over when board is full, but it is not")
	}
}

func TestDetectsRowWin(t *testing.T) {
	cases := []struct {
		board  Board
		winner rune
	}{
		{
			board: Board{
				{'X', 'X', 'X'},
				{'O', 'O', ' '},
				{' ', ' ', ' '},
			},
			winner: 'X',
		},
		{
			board: Board{
				{' ', ' ', ' '},
				{'O', 'O', 'O'},
				{'X', ' ', 'X'},
			},
			winner: 'O',
		},
		{
			board: Board{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{'X', 'X', 'X'},
			},
			winner: 'X',
		},
	}

	for _, c := range cases {
		game := NewGame()
		game.Board = c.board

		game.Next = 'X' // Set next player to X for testing

		if !game.IsOver() {
			t.Error("Expected game to be over, but it is not")
		}

		if game.Winner() != c.winner {
			t.Errorf("Expected winner %q, but got %q", c.winner, game.Winner())
		}
	}
}

func TestDetectsColumnWin(t *testing.T) {
	cases := []struct {
		board  Board
		winner rune
	}{
		{
			board: Board{
				{'X', 'O', ' '},
				{'X', 'O', ' '},
				{'X', ' ', ' '},
			},
			winner: 'X',
		},
		{
			board: Board{
				{'O', ' ', 'X'},
				{'O', ' ', 'X'},
				{' ', ' ', 'X'},
			},
			winner: 'X',
		},
	}

	for _, c := range cases {
		game := NewGame()
		game.Board = c.board

		game.Next = 'X' // Set next player to X for testing

		if !game.IsOver() {
			t.Error("Expected game to be over, but it is not")
		}

		if game.Winner() != c.winner {
			t.Errorf("Expected winner %q, but got %q", c.winner, game.Winner())
		}
	}
}
