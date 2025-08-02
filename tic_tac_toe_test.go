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
	tests := []struct {
		name     string
		board    Board
		expected rune
	}{
		{
			name: "top row X wins",
			board: Board{
				{'X', 'X', 'X'},
				{'O', 'O', ' '},
				{' ', ' ', ' '},
			},
			expected: 'X',
		},
		{
			name: "middle row O wins",
			board: Board{
				{' ', ' ', ' '},
				{'O', 'O', 'O'},
				{'X', ' ', 'X'},
			},
			expected: 'O',
		},
		{
			name: "bottom row X wins",
			board: Board{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{'X', 'X', 'X'},
			},
			expected: 'X',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.Board = tt.board

			if !game.IsOver() {
				t.Error("Expected game to be over, but it is not")
			}

			if game.Winner() != tt.expected {
				t.Errorf("Expected winner %q, but got %q", tt.expected, game.Winner())
			}
		})
	}
}

func TestDetectsColumnWin(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected rune
	}{
		{
			name: "first column X wins",
			board: Board{
				{'X', 'O', ' '},
				{'X', 'O', ' '},
				{'X', ' ', ' '},
			},
			expected: 'X',
		},
		{
			name: "third column X wins",
			board: Board{
				{'O', ' ', 'X'},
				{'O', ' ', 'X'},
				{' ', ' ', 'X'},
			},
			expected: 'X',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.Board = tt.board

			if !game.IsOver() {
				t.Error("Expected game to be over, but it is not")
			}

			if game.Winner() != tt.expected {
				t.Errorf("Expected winner %q, but got %q", tt.expected, game.Winner())
			}
		})
	}
}

func TestDiagonalWin(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected rune
	}{
		{
			name: "top-left to bottom-right X wins",
			board: Board{
				{'X', 'O', ' '},
				{'O', 'X', ' '},
				{' ', ' ', 'X'},
			},
			expected: 'X',
		},
		{
			name: "top-right to bottom-left O wins",
			board: Board{
				{' ', ' ', 'O'},
				{'X', 'O', 'X'},
				{'O', 'X', ' '},
			},
			expected: 'O',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.Board = tt.board

			if !game.IsOver() {
				t.Error("Expected game to be over, but it is not")
			}

			if game.Winner() != tt.expected {
				t.Errorf("Expected winner %q, but got %q", tt.expected, game.Winner())
			}
		})
	}
}
