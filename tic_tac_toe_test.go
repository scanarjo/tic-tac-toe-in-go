package tic_tac_toe

import (
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

	if game.IsOver {
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
