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
		t.Errorf("Expected %+v, but got %+v", freshBoard, game.Board)
	}
}
