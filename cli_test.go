package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLIGameStart(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "winning game",
			input: "0 0\n1 1\n0 1\n2 2\n0 2\n",
			contains: []string{
				"Starting a new game",
				"Player",
				"Game over:",
				"wins!",
			},
		},
		{
			name:  "invalid input",
			input: "invalid\n0 0\n1 1\n2 2\n0 1\n1 0\n0 2\n",
			contains: []string{
				"Invalid input",
				"row col",
			},
		},
		{
			name:  "out of bounds",
			input: "5 5\n0 0\n1 1\n2 2\n0 1\n1 0\n0 2\n",
			contains: []string{
				"Error:",
				"out of bounds",
			},
		},
		{
			name:  "occupied cell",
			input: "0 0\n0 0\n1 1\n2 2\n0 1\n1 0\n0 2\n",
			contains: []string{
				"Error:",
				"already contains",
			},
		},
		{
			name:  "draw game",
			input: "0 0\n0 1\n0 2\n1 0\n1 2\n1 1\n2 0\n2 2\n2 1\n",
			contains: []string{
				"Starting a new game",
				"Player",
				"Game over: It's a draw!",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdin := strings.NewReader(tt.input)
			var stdout bytes.Buffer

			cli := NewCLIGameWithIO(stdin, &stdout)
			cli.Start()

			result := stdout.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q", expected)
				}
			}
		})
	}
}
