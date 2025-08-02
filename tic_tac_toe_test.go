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
