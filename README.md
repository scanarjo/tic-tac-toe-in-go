# Tic Tac Toe

A simple command-line implementation of Tic Tac Toe in Go.

## Features

- Two-player gameplay with alternating turns
- Random starting player selection
- Input validation and error handling
- Win detection for rows, columns, and diagonals
- Draw detection when board is full
- Clean, testable code architecture

## Installation

```bash
git clone <repository-url>
cd tic-tac-toe-go
go mod tidy
```

## Usage

```bash
go run .
```

The game will prompt each player to enter their move as row and column coordinates (0-2):

```
Starting a new game of Tic Tac Toe...
_ _ _
_ _ _
_ _ _

Player X, enter your move (row col): 1 1
_ _ _
_ X _
_ _ _

Player O, enter your move (row col): 0 0
O _ _
_ X _
_ _ _
```
